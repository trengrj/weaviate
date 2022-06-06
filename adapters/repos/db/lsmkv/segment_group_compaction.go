//                           _       _
// __      _____  __ ___   ___  __ _| |_ ___
// \ \ /\ / / _ \/ _` \ \ / / |/ _` | __/ _ \
//  \ V  V /  __/ (_| |\ V /| | (_| | ||  __/
//   \_/\_/ \___|\__,_| \_/ |_|\__,_|\__\___|
//
//  Copyright © 2016 - 2022 SeMI Technologies B.V. All rights reserved.
//
//  CONTACT: hello@semi.technology
//

package lsmkv

import (
	"fmt"
	"math"
	"os"
	"path/filepath"
	"time"

	"github.com/pkg/errors"
	"github.com/prometheus/client_golang/prometheus"
)

func (ig *SegmentGroup) eligbleForCompaction() bool {
	ig.maintenanceLock.RLock()
	defer ig.maintenanceLock.RUnlock()

	// if true, the parent shard has indicated that it has
	// entered an immutable state. During this time, the
	// SegmentGroup should refrain from flushing until its
	// shard indicates otherwise
	if ig.isReadyOnly() {
		return false
	}

	// if there are at least two segments of the same level a regular compaction
	// can be performed

	levels := map[uint16]int{}

	for _, segment := range ig.segments {
		levels[segment.level]++
		if levels[segment.level] > 1 {
			return true
		}
	}

	return false
}

func (ig *SegmentGroup) bestCompactionCandidatePair() []int {
	ig.maintenanceLock.RLock()
	defer ig.maintenanceLock.RUnlock()

	// first determine the lowest level with candidates
	levels := map[uint16]int{}

	for _, segment := range ig.segments {
		levels[segment.level]++
	}

	currLowestLevel := uint16(math.MaxUint16)
	found := false
	for level, count := range levels {
		if count < 2 {
			continue
		}

		if level < currLowestLevel {
			currLowestLevel = level
			found = true
		}
	}

	if !found {
		return nil
	}

	// now pick any two segements which match the level
	var res []int

	for i, segment := range ig.segments {
		if len(res) >= 2 {
			break
		}

		if segment.level == currLowestLevel {
			res = append(res, i)
		}
	}

	return res
}

// segmentAtPos retrieves the segment for the given position using a read-lock
func (ig *SegmentGroup) segmentAtPos(pos int) *segment {
	ig.maintenanceLock.RLock()
	defer ig.maintenanceLock.RUnlock()

	return ig.segments[pos]
}

func (ig *SegmentGroup) compactOnce() error {
	// Is it safe to only occasionally lock instead of the entire duration? Yes,
	// because other than compaction the only change to the segments array could
	// be an append because of a new flush cycle, so we do not need to guarantee
	// that the array contents stay stable over the duration of an entire
	// compaction. We do however need to protect against a read-while-write (race
	// condition) on the array. Thus any read from ig.segments need to protected
	pair := ig.bestCompactionCandidatePair()
	if pair == nil {
		// nothing to do
		return nil
	}

	path := fmt.Sprintf("%s.tmp", ig.segmentAtPos(pair[1]).path)
	f, err := os.Create(path)
	if err != nil {
		return err
	}

	scratchSpacePath := ig.segmentAtPos(pair[1]).path + "compaction.scratch.d"

	// the assumption is that both pairs are of the same level, so we can just
	// take either value. If we want to support asymmetric compaction, then we
	// might have to choose this value more intelligently
	level := ig.segmentAtPos(pair[0]).level
	secondaryIndices := ig.segmentAtPos(pair[0]).secondaryIndexCount

	strategy := ig.segmentAtPos(pair[0]).strategy

	switch strategy {

	// TODO: call metrics just once with variable strategy label

	case SegmentStrategyReplace:
		c := newCompactorReplace(f, ig.segmentAtPos(pair[0]).newCursor(),
			ig.segmentAtPos(pair[1]).newCursor(), level, secondaryIndices, scratchSpacePath)

		if ig.metrics != nil {
			ig.metrics.CompactionReplace.With(prometheus.Labels{"path": ig.dir}).Set(1)
			defer ig.metrics.CompactionReplace.With(prometheus.Labels{"path": ig.dir}).Set(0)
		}

		if err := c.do(); err != nil {
			return err
		}
	case SegmentStrategySetCollection:
		c := newCompactorSetCollection(f, ig.segmentAtPos(pair[0]).newCollectionCursor(),
			ig.segmentAtPos(pair[1]).newCollectionCursor(), level, secondaryIndices,
			scratchSpacePath)

		if ig.metrics != nil {
			ig.metrics.CompactionSet.With(prometheus.Labels{"path": ig.dir}).Set(1)
			defer ig.metrics.CompactionSet.With(prometheus.Labels{"path": ig.dir}).Set(0)
		}

		if err := c.do(); err != nil {
			return err
		}
	case SegmentStrategyMapCollection:
		c := newCompactorMapCollection(f,
			ig.segmentAtPos(pair[0]).newCollectionCursorReusable(),
			ig.segmentAtPos(pair[1]).newCollectionCursorReusable(),
			level, secondaryIndices, scratchSpacePath, ig.mapRequiresSorting)

		if ig.metrics != nil {
			ig.metrics.CompactionMap.With(prometheus.Labels{"path": ig.dir}).Set(1)
			defer ig.metrics.CompactionMap.With(prometheus.Labels{"path": ig.dir}).Set(0)
		}

		if err := c.do(); err != nil {
			return err
		}

	default:
		return errors.Errorf("unrecognized strategy %v", strategy)
	}

	if err := f.Close(); err != nil {
		return errors.Wrap(err, "close compacted segment file")
	}

	if err := ig.replaceCompactedSegments(pair[0], pair[1], path); err != nil {
		return errors.Wrap(err, "replace compacted segments")
	}

	return nil
}

func (ig *SegmentGroup) replaceCompactedSegments(old1, old2 int,
	newPathTmp string) error {
	ig.maintenanceLock.Lock()
	defer ig.maintenanceLock.Unlock()

	if err := ig.segments[old1].close(); err != nil {
		return errors.Wrap(err, "close disk segment")
	}

	if err := ig.segments[old2].close(); err != nil {
		return errors.Wrap(err, "close disk segment")
	}

	if err := ig.segments[old1].drop(); err != nil {
		return errors.Wrap(err, "drop disk segment")
	}

	if err := ig.segments[old2].drop(); err != nil {
		return errors.Wrap(err, "drop disk segment")
	}

	ig.segments[old1] = nil
	ig.segments[old2] = nil

	// the old segments have been deletd, we can now safely remove the .tmp
	// extension from the new segment which carried the name of the second old
	// segment
	newPath, err := ig.stripTmpExtension(newPathTmp)
	if err != nil {
		return errors.Wrap(err, "strip .tmp extension of new segment")
	}

	exists := ig.makeExistsOnLower(old1)
	seg, err := newSegment(newPath, ig.logger, ig.metrics, exists)
	if err != nil {
		return errors.Wrap(err, "create new segment")
	}

	ig.segments[old2] = seg

	ig.segments = append(ig.segments[:old1], ig.segments[old1+1:]...)

	return nil
}

func (ig *SegmentGroup) stripTmpExtension(oldPath string) (string, error) {
	ext := filepath.Ext(oldPath)
	if ext != ".tmp" {
		return "", errors.Errorf("segment %q did not have .tmp extension", oldPath)
	}
	newPath := oldPath[:len(oldPath)-len(ext)]

	if err := os.Rename(oldPath, newPath); err != nil {
		return "", errors.Wrapf(err, "rename %q -> %q", oldPath, newPath)
	}

	return newPath, nil
}

func (ig *SegmentGroup) initCompactionCycle(interval time.Duration) {
	if interval == 0 {
		return
	}

	go func() {
		t := time.Tick(interval)
		for {
			select {
			case <-ig.stopCompactionCycle:
				ig.logger.WithField("action", "lsm_compaction_stop_cycle").
					WithField("path", ig.dir).
					Debug("stop compaction cycle")
				return
			case <-t:
				ig.monitorSegments()

				if ig.eligbleForCompaction() {
					if err := ig.compactOnce(); err != nil {
						ig.logger.WithField("action", "lsm_compaction").
							WithField("path", ig.dir).
							WithError(err).
							Errorf("compaction failed")
					}
				} else {
					ig.logger.WithField("action", "lsm_compaction").
						WithField("path", ig.dir).
						Trace("no segment eligble for compaction")
				}
			}
		}
	}()
}

func (ig *SegmentGroup) Len() int {
	ig.maintenanceLock.RLock()
	defer ig.maintenanceLock.RUnlock()

	return len(ig.segments)
}

func (ig *SegmentGroup) monitorSegments() {
	if ig.metrics == nil {
		return
	}

	ig.metrics.ActiveSegments.With(prometheus.Labels{
		"strategy": ig.strategy,
		"path":     ig.dir,
	}).Set(float64(ig.Len()))

	stats := ig.segmentLevelStats()
	stats.fillMissingLevels()
	stats.report(ig.metrics, ig.strategy, ig.dir)
}

type segmentLevelStats struct {
	indexes  map[uint16]int
	payloads map[uint16]int
	count    map[uint16]int
}

func newSegmentLevelStats() segmentLevelStats {
	return segmentLevelStats{
		indexes:  map[uint16]int{},
		payloads: map[uint16]int{},
		count:    map[uint16]int{},
	}
}

func (ig *SegmentGroup) segmentLevelStats() segmentLevelStats {
	ig.maintenanceLock.RLock()
	defer ig.maintenanceLock.RUnlock()

	stats := newSegmentLevelStats()

	for _, seg := range ig.segments {
		stats.count[seg.level]++

		cur := stats.indexes[seg.level]
		cur += seg.index.Size()
		stats.indexes[seg.level] = cur

		cur = stats.payloads[seg.level]
		cur += seg.PayloadSize()
		stats.payloads[seg.level] = cur
	}

	return stats
}

// fill missing levels
//
// Imagine we had exactly two segments of level 4 before, and there were just
// compacted to single segment of level 5. As a result, there should be no
// more segments of level 4. However, our current logic only loops over
// existing segments. As a result, we need to check what the highest level
// is, then for every level lower than the highest check if we are missing
// data. If yes, we need to explicitly set the gauges to 0.
func (s *segmentLevelStats) fillMissingLevels() {
	maxLevel := uint16(0)
	for level := range s.count {
		if level > maxLevel {
			maxLevel = level
		}
	}

	if maxLevel > 0 {
		for level := uint16(0); level < maxLevel; level++ {
			if _, ok := s.count[level]; ok {
				continue
			}

			// there is no entry for this level, we must explicitly set it to 0
			s.count[level] = 0
			s.indexes[level] = 0
			s.payloads[level] = 0
		}
	}
}

func (s *segmentLevelStats) report(metrics *Metrics,
	strategy, dir string) {
	for level, size := range s.indexes {
		metrics.SegmentSize.With(prometheus.Labels{
			"strategy": strategy,
			"unit":     "index",
			"level":    fmt.Sprint(level),
			"path":     dir,
		}).Set(float64(size))
	}

	for level, size := range s.payloads {
		metrics.SegmentSize.With(prometheus.Labels{
			"strategy": strategy,
			"unit":     "payload",
			"level":    fmt.Sprint(level),
			"path":     dir,
		}).Set(float64(size))
	}

	for level, count := range s.count {
		metrics.SegmentCount.With(prometheus.Labels{
			"strategy": strategy,
			"level":    fmt.Sprint(level),
			"path":     dir,
		}).Set(float64(count))
	}
}
