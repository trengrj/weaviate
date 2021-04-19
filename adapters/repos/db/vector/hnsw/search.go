//                           _       _
// __      _____  __ ___   ___  __ _| |_ ___
// \ \ /\ / / _ \/ _` \ \ / / |/ _` | __/ _ \
//  \ V  V /  __/ (_| |\ V /| | (_| | ||  __/
//   \_/\_/ \___|\__,_| \_/ |_|\__,_|\__\___|
//
//  Copyright © 2016 - 2021 SeMI Technologies B.V. All rights reserved.
//
//  CONTACT: hello@semi.technology
//

package hnsw

import (
	"context"
	"fmt"
	"math"

	"github.com/pkg/errors"
	"github.com/semi-technologies/weaviate/adapters/repos/db/helpers"
	"github.com/semi-technologies/weaviate/adapters/repos/db/storobj"
	"github.com/semi-technologies/weaviate/adapters/repos/db/vector/hnsw/distancer"
)

func reasonableEfFromK(k int) int {
	ef := k * 8
	if ef > 100 {
		ef = 100
	}
	if k > ef {
		ef = k // otherwise results will get cut off early
	}

	return ef
}

func (h *hnsw) SearchByID(id uint64, k int) ([]uint64, error) {
	return h.knnSearch(id, k, reasonableEfFromK(k))
}

func (h *hnsw) SearchByVector(vector []float32, k int, allowList helpers.AllowList) ([]uint64, error) {
	return h.knnSearchByVector(vector, k, reasonableEfFromK(k), allowList)
}

func (h *hnsw) knnSearch(queryNodeID uint64, k int, ef int) ([]uint64, error) {
	entryPointID := h.entryPointID
	entryPointDistance, ok, err := h.distBetweenNodes(entryPointID, queryNodeID)
	if err != nil || !ok {
		return nil, errors.Wrap(err, "knn search: distance between entrypint and query node")
	}

	queryVector, err := h.vectorForID(context.Background(), queryNodeID)
	if err != nil {
		return nil, errors.Wrapf(err, "could not get vector of object at docID %d", queryNodeID)
	}

	for level := h.currentMaximumLayer; level >= 1; level-- { // stop at layer 1, not 0!
		eps := &binarySearchTreeGeneric{}
		eps.insert(entryPointID, entryPointDistance)

		res, err := h.searchLayerByVector(queryVector, *eps, 1, level, nil)
		if err != nil {
			return nil, errors.Wrapf(err, "knn search: search layer at level %d", level)
		}
		best := res.minimum()
		entryPointID = best.index
		entryPointDistance = best.dist
	}

	eps := &binarySearchTreeGeneric{}
	eps.insert(entryPointID, entryPointDistance)
	res, err := h.searchLayerByVector(queryVector, *eps, ef, 0, nil)
	if err != nil {
		return nil, errors.Wrapf(err, "knn search: search layer at level %d", 0)
	}

	flat := res.flattenInOrder()
	if len(flat) > k {
		fmt.Printf("just wasted %d allocations\n", len(flat)-k)
	}
	size := min(len(flat), k)
	out := make([]uint64, size)
	for i, elem := range flat {
		if i >= size {
			break
		}
		out[i] = elem.index
	}

	return out, nil
}

func (h *hnsw) searchLayerByVector(queryVector []float32,
	entrypoints binarySearchTreeGeneric, ef int, level int,
	allowList helpers.AllowList) (*binarySearchTreeGeneric, error) {
	visited := h.newVisitedList(entrypoints)
	candidates := &binarySearchTreeGeneric{}
	results := &binarySearchTreeGeneric{}
	distancer := h.distancerProvider.New(queryVector)

	h.insertViableEntrypointsAsCandidatesAndResults(entrypoints, candidates,
		results, level, allowList)

	for candidates.root != nil { // efficient way to see if the len is > 0
		candidate := candidates.minimum()
		// fmt.Println(candidates.flattenInOrder())
		// h.Dump()
		candidates.delete(candidate.index, candidate.dist)

		worstResultDistance, err := h.currentWorstResultDistance(results, distancer)
		if err != nil {
			return nil, errors.Wrapf(err, "calculate distance of current last result")
		}

		dist, ok, err := h.distanceToNode(distancer, candidate.index)
		if err != nil {
			return nil, errors.Wrap(err, "calculate distance between candidate and query")
		}

		if !ok {
			continue
		}

		if dist > worstResultDistance {
			break
		}

		// before := time.Now()
		// h.Lock()
		// m.addBuildingReadLocking(before)
		candidateNode := h.nodes[candidate.index]
		// h.Unlock()

		if candidateNode == nil {
			// could have been a node that already had a tombstone attached and was
			// just cleaned up while we were waiting for a read lock
			continue
		}

		// before = time.Now()
		candidateNode.Lock()
		// m.addBuildingItemLocking(before)
		connections := candidateNode.connections[level]
		candidateNode.Unlock()

		if err := h.extendCandidatesAndResultsFromNeighbors(candidates, results,
			connections, visited, distancer, ef, level, allowList,
			worstResultDistance); err != nil {
			return nil, errors.Wrap(err, "extend candidates and results from neighbors")
		}
	}

	return results, nil
}

func (h *hnsw) newVisitedList(entrypoints binarySearchTreeGeneric) []bool {
	h.Lock()
	size := len(h.nodes) + defaultIndexGrowthDelta // add delta to be add some
	// buffer if a visited list is created shortly before a growth operation
	h.Unlock()

	visited := make([]bool, size)
	for _, elem := range entrypoints.flattenInOrder() {
		visited[elem.index] = true
	}
	return visited
}

func (h *hnsw) insertViableEntrypointsAsCandidatesAndResults(
	entrypoints binarySearchTreeGeneric, candidates,
	results *binarySearchTreeGeneric, level int, allowList helpers.AllowList) {
	for _, ep := range entrypoints.flattenInOrder() {
		candidates.insert(ep.index, ep.dist)
		if level == 0 && allowList != nil {
			// we are on the lowest level containing the actual candidates and we
			// have an allow list (i.e. the user has probably set some sort of a
			// filter restricting this search further. As a result we have to
			// ignore items not on the list
			if !allowList.Contains(ep.index) {
				continue
			}
		}

		if h.hasTombstone(ep.index) {
			continue
		}

		results.insert(ep.index, ep.dist)
	}
}

func (h *hnsw) currentWorstResultDistance(results *binarySearchTreeGeneric,
	distancer distancer.Distancer) (float32, error) {
	if results.root != nil {
		id := results.maximum().index
		d, ok, err := h.distanceToNode(distancer, id)
		if err != nil {
			return 0, errors.Wrap(err,
				"calculated distance between worst result and query")
		}

		if !ok {
			return math.MaxFloat32, nil
		}
		return d, nil
	} else {
		// if the entrypoint (which we received from a higher layer doesn't match
		// the allow List the result list is empty. In this case we can just set
		// the worstDistance to an arbitrarily large number, so that any
		// (allowed) candidate will have a lower distance in comparison
		return math.MaxFloat32, nil
	}
}

func (h *hnsw) extendCandidatesAndResultsFromNeighbors(candidates,
	results *binarySearchTreeGeneric, connections []uint64,
	visited []bool, distancer distancer.Distancer, ef int,
	level int, allowList helpers.AllowList, worstResultDistance float32,
) error {
	for _, neighborID := range connections {
		if ok := visited[neighborID]; ok {
			// skip if we've already visited this neighbor
			continue
		}

		// make sure we never visit this neighbor again
		visited[neighborID] = true

		distance, ok, err := h.distanceToNode(distancer, neighborID)
		if err != nil {
			return errors.Wrap(err, "calculate distance between candidate and query")
		}

		if !ok {
			// node was deleted in the underlying object store
			continue
		}

		resLenBefore := results.len() // calculating just once saves a bit of time
		if distance < worstResultDistance || resLenBefore < ef {
			candidates.insert(neighborID, distance)
			if level == 0 && allowList != nil {
				// we are on the lowest level containing the actual candidates and we
				// have an allow list (i.e. the user has probably set some sort of a
				// filter restricting this search further. As a result we have to
				// ignore items not on the list
				if !allowList.Contains(neighborID) {
					continue
				}
			}

			if h.hasTombstone(neighborID) {
				continue
			}

			results.insert(neighborID, distance)

			// +1 because we have added one node size calculating the len
			if resLenBefore+1 > ef {
				max := results.maximum()
				results.delete(max.index, max.dist)
			}
		}
	}

	return nil
}

func (h *hnsw) distanceToNode(distancer distancer.Distancer,
	nodeID uint64) (float32, bool, error) {
	candidateVec, err := h.vectorForID(context.Background(), nodeID)
	if err != nil {
		var e storobj.ErrNotFound
		if errors.As(err, &e) {
			h.handleDeletedNode(e.DocID)
			return 0, false, nil
		} else {
			// not a typed error, we can recover from, return with err
			return 0, false, errors.Wrapf(err, "get vector of docID %d", nodeID)
		}
	}

	dist, _, err := distancer.Distance(candidateVec)
	if err != nil {
		return 0, false, errors.Wrap(err, "calculate distance between candidate and query")
	}

	return dist, true, nil
}

// the underlying object seems to have been deleted, to recover from
// this situation let's add a tombstone to the deleted object, so it
// will be cleaned up and skip this candidate in the current search
func (h *hnsw) handleDeletedNode(docID uint64) {
	if h.hasTombstone(docID) {
		// nothing to do, this node already has a tombstone, it will be cleaned up
		// in the next deletion cycle
		return
	}

	h.addTombstone(docID)
	h.logger.WithField("action", "attach_tombstone_to_deleted_node").
		WithField("node_id", docID).
		Info("found a deleted node (%d) without a tombstone, "+
			"tombstone was added", docID)
}

func (h *hnsw) knnSearchByVector(searchVec []float32, k int,
	ef int, allowList helpers.AllowList) ([]uint64, error) {
	if h.isEmpty() {
		return nil, nil
	}

	entryPointID := h.entryPointID
	entryPointDistance, ok, err := h.distBetweenNodeAndVec(entryPointID, searchVec)
	if err != nil {
		return nil, errors.Wrap(err, "knn search: distance between entrypint and query node")
	}

	if !ok {
		return nil, fmt.Errorf("entrypoint was deleted in the object store, " +
			"it has been flagged for cleanup and should be fixed in the next cleanup cycle")
	}

	// stop at layer 1, not 0!
	for level := h.currentMaximumLayer; level >= 1; level-- {
		eps := &binarySearchTreeGeneric{}
		eps.insert(entryPointID, entryPointDistance)
		res, err := h.searchLayerByVector(searchVec, *eps, 1, level, nil)
		if err != nil {
			return nil, errors.Wrapf(err, "knn search: search layer at level %d", level)
		}

		// There might be situations where we did not find a better entrypoint at
		// that particular level, so instead we're keeping whatever entrypoint we
		// had before (i.e. either from a previous level or even the main
		// entrypoint)
		if res.root != nil {
			best := res.flattenInOrder()

			for _, cand := range best {
				if !h.nodeByID(cand.index).isUnderMaintenance() {
					entryPointID = cand.index
					entryPointDistance = cand.dist
					break
				}

				// if we managed to go through the loop without finding a single
				// suitable node, we simply stick with the original, i.e. the global
				// entrypoint
			}
		}
	}

	eps := &binarySearchTreeGeneric{}
	eps.insert(entryPointID, entryPointDistance)
	res, err := h.searchLayerByVector(searchVec, *eps, ef, 0, allowList)
	if err != nil {
		return nil, errors.Wrapf(err, "knn search: search layer at level %d", 0)
	}

	flat := res.flattenInOrder()
	if len(flat) > k {
		fmt.Printf("search by vec: just wasted %d allocations\n", len(flat)-k)
	}
	size := min(len(flat), k)
	out := make([]uint64, size)
	for i, elem := range flat {
		if i >= size {
			break
		}
		out[i] = elem.index
	}

	return out, nil
}

func (h *hnsw) selectNeighborsSimple(input binarySearchTreeGeneric,
	max int, denyList helpers.AllowList) []uint64 {
	// flat := input.flattenInOrder()
	// if len(flat) > max {
	// 	fmt.Printf("select neighbors: just wasted %d allocations\n", len(flat)-max)
	// }

	// maxSize := min(len(flat), max)
	out := make([]uint64, max)
	actualSize := 0
	for input.len() > 0 && actualSize < max {
		elem := input.minimum()
		input.delete(elem.index, elem.dist)
		if denyList != nil && denyList.Contains(elem.index) {
			continue
		}

		out[actualSize] = elem.index
		actualSize++
	}

	return out[:actualSize]
}

func (h *hnsw) selectNeighborsSimpleFromId(nodeId uint64, ids []uint64,
	max int, denyList helpers.AllowList) ([]uint64, error) {
	vec, err := h.vectorForID(context.Background(), nodeId)
	if err != nil {
		return nil, err
	}

	distancer := h.distancerProvider.New(vec)

	bst := &binarySearchTreeGeneric{}
	for _, id := range ids {

		vecA, err := h.vectorForID(context.Background(), id)
		if err != nil {
			var e storobj.ErrNotFound
			if errors.As(err, &e) {
				h.handleDeletedNode(e.DocID)
				continue
			} else {
				// not a typed error, we can recover from, return with err
				return nil, errors.Wrapf(err,
					"could not get vector of object at docID %d", id)
			}
		}
		dist, _, err := distancer.Distance(vecA)
		if err != nil {
			return nil, errors.Wrap(err, "select neighbors simple from id")
		}

		bst.insert(id, dist)
	}

	return h.selectNeighborsSimple(*bst, max, denyList), nil
}
