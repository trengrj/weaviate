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

package lsmkv

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"io"
	"os"
	"sort"

	"github.com/pkg/errors"
	"github.com/semi-technologies/weaviate/adapters/repos/db/lsmkv/segmentindex"
)

func (l *Memtable) flush() error {
	// close the commit log first, this also forces it to be fsynced. If
	// something fails there, don't proceed with flushing. The commit log will
	// only be deleted at the very end, if the flush was successful
	// (indicated by a successful close of the flush file - which indicates a
	// successful fsync)

	if err := l.commitlog.close(); err != nil {
		return errors.Wrap(err, "close commit log file")
	}

	if l.Size() == 0 {
		// this is an empty memtable, nothing to do
		return nil
	}

	f, err := os.Create(l.path + ".db")
	if err != nil {
		return err
	}

	w := bufio.NewWriterSize(f, int(float64(l.size)*1.3)) // calculate 30% overhead for disk representation

	var keys []keyIndex
	switch l.strategy {
	case StrategyReplace:
		if keys, err = l.flushDataReplace(w); err != nil {
			return err
		}

	case StrategySetCollection, StrategyMapCollection:
		if keys, err = l.flushDataCollection(w); err != nil {
			return err
		}

	}

	currentOffset := uint64(keys[len(keys)-1].valueEnd)
	indexBytes, err := buildAndMarshalPrimaryIndex(keys)
	if err != nil {
		return err
	}

	// pretend that primary index was already written
	currentOffset = currentOffset + uint64(len(indexBytes)) + uint64(l.secondaryIndices)*8

	secondaryIndicesBytes := bytes.NewBuffer(nil)

	if l.secondaryIndices > 0 {
		offsets := make([]uint64, l.secondaryIndices)
		for pos := range offsets {
			secondaryBytes, err := buildAndMarshalSecondaryIndex(pos, keys)
			if err != nil {
				return err
			}

			if _, err := secondaryIndicesBytes.Write(secondaryBytes); err != nil {
				return err
			}

			offsets[pos] = currentOffset
			currentOffset = offsets[pos] + uint64(len(secondaryBytes))
		}

		if err := binary.Write(w, binary.LittleEndian, &offsets); err != nil {
			return err
		}
	}

	// write primary index
	if _, err := w.Write(indexBytes); err != nil {
		return err
	}

	// write secondary indices
	if _, err := secondaryIndicesBytes.WriteTo(w); err != nil {
		return err
	}

	if err := w.Flush(); err != nil {
		return err
	}

	if err := f.Close(); err != nil {
		return err
	}

	// only now that the file has been flushed is it safe to delete the commit log
	// TODO: there might be an interest in keeping the commit logs around for
	// longer as they might come in handy for replication
	return l.commitlog.delete()
}

// pos indicates the position of a secondary index, assumes unsorted keys and
// sorts them
func buildAndMarshalSecondaryIndex(pos int, keys []keyIndex) ([]byte, error) {
	keyNodes := make([]segmentindex.Node, len(keys))
	i := 0
	for _, key := range keys {
		if pos >= len(key.secondaryKeys) {
			// a secondary key is not guaranteed to be present. For example, a delete
			// operation could pe performed using only the primary key
			continue
		}

		keyNodes[i] = segmentindex.Node{
			Key:   key.secondaryKeys[pos],
			Start: uint64(key.valueStart),
			End:   uint64(key.valueEnd),
		}
		i++
	}

	keyNodes = keyNodes[:i]

	sort.Slice(keyNodes, func(a, b int) bool {
		return bytes.Compare(keyNodes[a].Key, keyNodes[b].Key) < 0
	})

	index := segmentindex.NewBalanced(keyNodes)
	indexBytes, err := index.MarshalBinary()
	if err != nil {
		return nil, err
	}

	return indexBytes, nil
}

// assumes sorted keys and does NOT sort them again
func buildAndMarshalPrimaryIndex(keys []keyIndex) ([]byte, error) {
	keyNodes := make([]segmentindex.Node, len(keys))
	for i, key := range keys {
		keyNodes[i] = segmentindex.Node{
			Key:   key.key,
			Start: uint64(key.valueStart),
			End:   uint64(key.valueEnd),
		}
	}
	index := segmentindex.NewBalanced(keyNodes)

	indexBytes, err := index.MarshalBinary()
	if err != nil {
		return nil, err
	}

	return indexBytes, nil
}

// SegmentOffset describes the general offset in a segment until the data
// starts, it is comprised of 2 bytes for level, 2 bytes for version,
// 2 bytes for secondary index count, 2 bytes for strategy, 8 bytes
// for the pointer to the index part
const SegmentHeaderSize = 16

func (l *Memtable) flushDataReplace(f io.Writer) ([]keyIndex, error) {
	flat := l.key.flattenInOrder()

	totalDataLength := totalKeyAndValueSize(flat)
	perObjectAdditions := len(flat) * (1 + 8 + 4 + int(l.secondaryIndices)*4) // 1 byte for the tombstone, 8 bytes value length encoding, 4 bytes key length encoding, + 4 bytes key encoding for every secondary index
	headerSize := SegmentHeaderSize
	header := segmentHeader{
		indexStart:       uint64(totalDataLength + perObjectAdditions + headerSize),
		level:            0, // always level zero on a new one
		version:          0, // always version 0 for now
		secondaryIndices: l.secondaryIndices,
		strategy:         SegmentStrategyFromString(l.strategy),
	}

	n, err := header.WriteTo(f)
	if err != nil {
		return nil, err
	}
	headerSize = int(n)
	keys := make([]keyIndex, len(flat))

	totalWritten := headerSize
	for i, node := range flat {
		writtenForNode := 0
		if err := binary.Write(f, binary.LittleEndian, node.tombstone); err != nil {
			return nil, errors.Wrapf(err, "write tombstone for node %d", i)
		}
		writtenForNode += 1

		valueLength := uint64(len(node.value))
		if err := binary.Write(f, binary.LittleEndian, &valueLength); err != nil {
			return nil, errors.Wrapf(err, "write value length encoding for node %d", i)
		}
		writtenForNode += 8

		n, err := f.Write(node.value)
		if err != nil {
			return nil, errors.Wrapf(err, "write node %d", i)
		}
		writtenForNode += n

		keyLength := uint32(len(node.key))
		if err := binary.Write(f, binary.LittleEndian, &keyLength); err != nil {
			return nil, errors.Wrapf(err, "write key length encoding for node %d", i)
		}
		writtenForNode += 4

		n, err = f.Write(node.key)
		if err != nil {
			return nil, errors.Wrapf(err, "write node %d", i)
		}
		writtenForNode += n

		for j := 0; j < int(l.secondaryIndices); j++ {
			var secondaryKeyLength uint32
			if j < len(node.secondaryKeys) {
				secondaryKeyLength = uint32(len(node.secondaryKeys[j]))
			}

			// write the key length in any case
			if err := binary.Write(f, binary.LittleEndian, &secondaryKeyLength); err != nil {
				return nil, errors.Wrapf(err, "write secondary key length encoding for node %d", i)
			}
			writtenForNode += 4

			if secondaryKeyLength == 0 {
				// we're done here
				continue
			}

			// only write the key if it exists
			n, err = f.Write(node.secondaryKeys[j])
			if err != nil {
				return nil, errors.Wrapf(err, "write secondary key %d node %d", j, i)
			}
			writtenForNode += n
		}

		keys[i] = keyIndex{
			valueStart:    totalWritten,
			valueEnd:      totalWritten + writtenForNode,
			key:           node.key,
			secondaryKeys: node.secondaryKeys,
		}

		totalWritten += writtenForNode
	}

	return keys, nil
}

func (l *Memtable) flushDataCollection(f io.Writer) ([]keyIndex, error) {
	flat := l.keyMulti.flattenInOrder()

	totalDataLength := totalValueSizeCollection(flat)
	header := segmentHeader{
		indexStart:       uint64(totalDataLength + SegmentHeaderSize),
		level:            0, // always level zero on a new one
		version:          0, // always version 0 for now
		secondaryIndices: l.secondaryIndices,
		strategy:         SegmentStrategyFromString(l.strategy),
	}

	n, err := header.WriteTo(f)
	if err != nil {
		return nil, err
	}
	headerSize := int(n)
	keys := make([]keyIndex, len(flat))

	totalWritten := headerSize
	for i, node := range flat {
		writtenForNode := 0

		valueLen := uint64(len(node.values))
		if err := binary.Write(f, binary.LittleEndian, &valueLen); err != nil {
			return nil, errors.Wrapf(err, "write values len for node %d", i)
		}
		writtenForNode += 8

		for _, value := range node.values {
			if err := binary.Write(f, binary.LittleEndian, value.tombstone); err != nil {
				return nil, errors.Wrapf(err, "write tombstone for value on node %d", i)
			}
			writtenForNode += 1

			valueLen := uint64(len(value.value))
			if err := binary.Write(f, binary.LittleEndian, valueLen); err != nil {
				return nil, errors.Wrapf(err, "write len of value on node %d", i)
			}
			writtenForNode += 8

			n, err := f.Write(value.value)
			if err != nil {
				return nil, errors.Wrapf(err, "write value on node %d", i)
			}
			writtenForNode += n
		}

		keyLen := uint32(len(node.key))
		if err := binary.Write(f, binary.LittleEndian, &keyLen); err != nil {
			return nil, errors.Wrapf(err, "write key len for node %d", i)
		}
		writtenForNode += 4

		if n, err := f.Write(node.key); err != nil {
			return nil, errors.Wrapf(err, "write key on node %d", i)
		} else {
			writtenForNode += n
		}

		keys[i] = keyIndex{
			valueStart: totalWritten,
			valueEnd:   totalWritten + writtenForNode,
			key:        node.key,
		}

		totalWritten += writtenForNode
	}

	return keys, nil
}

func totalKeyAndValueSize(in []*binarySearchNode) int {
	var sum int
	for _, n := range in {
		sum += len(n.value)
		sum += len(n.key)
		for _, sec := range n.secondaryKeys {
			sum += len(sec)
		}
	}

	return sum
}

func totalValueSizeCollection(in []*binarySearchNodeMulti) int {
	var sum int
	for _, n := range in {
		sum += 8 // uint64 to indicate array length
		for _, v := range n.values {
			sum += 1 // bool to indicate value tombstone
			sum += 8 // uint64 to indicate value length
			sum += len(v.value)
		}

		sum += 4 // uint32 to indicate key size
		sum += len(n.key)
	}

	return sum
}
