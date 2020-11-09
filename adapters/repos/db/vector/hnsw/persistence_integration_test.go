//                           _       _
// __      _____  __ ___   ___  __ _| |_ ___
// \ \ /\ / / _ \/ _` \ \ / / |/ _` | __/ _ \
//  \ V  V /  __/ (_| |\ V /| | (_| | ||  __/
//   \_/\_/ \___|\__,_| \_/ |_|\__,_|\__\___|
//
//  Copyright © 2016 - 2020 SeMI Technologies B.V. All rights reserved.
//
//  CONTACT: hello@semi.technology
//

// +build integrationTest

package hnsw

import (
	"fmt"
	"math/rand"
	"os"
	"testing"
	"time"

	"github.com/semi-technologies/weaviate/adapters/repos/db/vector/hnsw/distancer"
	"github.com/sirupsen/logrus/hooks/test"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestHnswPersistence(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	dirName := fmt.Sprintf("./testdata/%d", rand.Intn(10000000))
	os.MkdirAll(dirName, 0o777)
	indexID := "integrationtest"
	defer func() {
		err := os.RemoveAll(dirName)
		fmt.Println(err)
	}()

	logger, _ := test.NewNullLogger()
	cl, clErr := NewCommitLogger(dirName, indexID, 0, logger)
	makeCL := func() (CommitLogger, error) {
		return cl, clErr
	}
	index, err := New(Config{
		RootPath:              dirName,
		ID:                    indexID,
		MakeCommitLoggerThunk: makeCL,
		MaximumConnections:    30,
		EFConstruction:        60,
		DistanceProvider:      distancer.NewCosineProvider(),
		VectorForIDThunk:      testVectorForID,
	})
	require.Nil(t, err)

	for i, vec := range testVectors {
		err := index.Add(i, vec)
		require.Nil(t, err)
	}

	// see index_test.go for more context
	expectedResults := []int{
		3, 5, 4, // cluster 2
		7, 8, 6, // cluster 3
		2, 1, 0, // cluster 1
	}

	t.Run("verify that the results match originally", func(t *testing.T) {
		position := 3
		res, err := index.knnSearchByVector(testVectors[position], 50, 36, nil)
		require.Nil(t, err)
		assert.Equal(t, expectedResults, res)
	})

	// destroy the index
	time.Sleep(100 * time.Millisecond) // TODO: make commit log write consinstent
	index = nil

	// build a new index from the (uncondensed) commit log
	secondIndex, err := New(Config{
		RootPath:              dirName,
		ID:                    indexID,
		MakeCommitLoggerThunk: makeCL,
		MaximumConnections:    30,
		EFConstruction:        60,
		DistanceProvider:      distancer.NewCosineProvider(),
		VectorForIDThunk:      testVectorForID,
	})
	require.Nil(t, err)

	t.Run("verify that the results match after rebuiling from disk",
		func(t *testing.T) {
			position := 3
			res, err := secondIndex.knnSearchByVector(testVectors[position], 50, 36, nil)
			require.Nil(t, err)
			assert.Equal(t, expectedResults, res)
		})
}

func TestHnswPersistence_WithDeletion_WithoutTombstoneCleanup(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	dirName := fmt.Sprintf("./testdata/%d", rand.Intn(10000000))
	os.MkdirAll(dirName, 0o777)
	indexID := "integrationtest_deletion"
	defer func() {
		err := os.RemoveAll(dirName)
		fmt.Println(err)
	}()

	logger, _ := test.NewNullLogger()
	cl, clErr := NewCommitLogger(dirName, indexID, 0, logger)
	makeCL := func() (CommitLogger, error) {
		return cl, clErr
	}
	index, err := New(Config{
		RootPath:              dirName,
		ID:                    indexID,
		MakeCommitLoggerThunk: makeCL,
		MaximumConnections:    30,
		EFConstruction:        60,
		DistanceProvider:      distancer.NewCosineProvider(),
		VectorForIDThunk:      testVectorForID,
	})
	require.Nil(t, err)

	for i, vec := range testVectors {
		err := index.Add(i, vec)
		require.Nil(t, err)
	}

	t.Run("delete some elements", func(t *testing.T) {
		err := index.Delete(6)
		require.Nil(t, err)
		err = index.Delete(8)
		require.Nil(t, err)
	})

	// see index_test.go for more context
	expectedResults := []int{
		3, 5, 4, // cluster 2
		7,       // cluster 3 with element 6 and 8 deleted
		2, 1, 0, // cluster 1
	}

	t.Run("verify that the results match originally", func(t *testing.T) {
		position := 3
		res, err := index.knnSearchByVector(testVectors[position], 50, 36, nil)
		require.Nil(t, err)
		assert.Equal(t, expectedResults, res)
	})

	dumpIndex(index, "without_cleanup_original_index_before_storage")

	// destroy the index
	time.Sleep(100 * time.Millisecond) // TODO: make commit log write consinstent
	index = nil

	// build a new index from the (uncondensed) commit log
	secondIndex, err := New(Config{
		RootPath:              dirName,
		ID:                    indexID,
		MakeCommitLoggerThunk: makeCL,
		MaximumConnections:    30,
		EFConstruction:        60,
		DistanceProvider:      distancer.NewCosineProvider(),
		VectorForIDThunk:      testVectorForID,
	})
	require.Nil(t, err)

	dumpIndex(secondIndex, "without_cleanup_after_rebuild")
	t.Run("verify that the results match after rebuiling from disk",
		func(t *testing.T) {
			position := 3
			res, err := secondIndex.knnSearchByVector(testVectors[position], 50, 36, nil)
			require.Nil(t, err)
			assert.Equal(t, expectedResults, res)
		})
}

func TestHnswPersistence_WithDeletion_WithTombstoneCleanup(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	dirName := fmt.Sprintf("./testdata/%d", rand.Intn(10000000))
	os.MkdirAll(dirName, 0o777)
	indexID := "integrationtest_tombstonecleanup"
	defer func() {
		err := os.RemoveAll(dirName)
		fmt.Println(err)
	}()

	logger, _ := test.NewNullLogger()
	makeCL := func() (CommitLogger, error) {
		return NewCommitLogger(dirName, indexID, 0, logger)
	}
	index, err := New(Config{
		RootPath:              dirName,
		ID:                    indexID,
		MakeCommitLoggerThunk: makeCL,
		MaximumConnections:    30,
		EFConstruction:        60,
		DistanceProvider:      distancer.NewCosineProvider(),
		VectorForIDThunk:      testVectorForID,
	})
	require.Nil(t, err)

	for i, vec := range testVectors {
		err := index.Add(i, vec)
		require.Nil(t, err)
	}
	// dumpIndex(index, "with cleanup after import")

	t.Run("delete some elements and permanently delete tombstoned elements",
		func(t *testing.T) {
			err := index.Delete(6)
			require.Nil(t, err)
			err = index.Delete(8)
			require.Nil(t, err)

			err = index.CleanUpTombstonedNodes()
			require.Nil(t, err)
		})

	// dumpIndex(index, "with cleanup after delete")

	// see index_test.go for more context
	expectedResults := []int{
		3, 5, 4, // cluster 2
		7,       // cluster 3 with element 6 and 8 deleted
		2, 1, 0, // cluster 1
	}

	t.Run("verify that the results match originally", func(t *testing.T) {
		position := 3
		res, err := index.knnSearchByVector(testVectors[position], 50, 36, nil)
		require.Nil(t, err)
		assert.Equal(t, expectedResults, res)
	})

	// destroy the index
	time.Sleep(100 * time.Millisecond) // TODO: make commit log write consinstent
	index = nil

	// build a new index from the (uncondensed) commit log
	secondIndex, err := New(Config{
		RootPath:              dirName,
		ID:                    indexID,
		MakeCommitLoggerThunk: makeCL,
		MaximumConnections:    30,
		EFConstruction:        60,
		DistanceProvider:      distancer.NewCosineProvider(),
		VectorForIDThunk:      testVectorForID,
	})
	require.Nil(t, err)
	// dumpIndex(secondIndex, "with cleanup second index")

	t.Run("verify that the results match after rebuiling from disk",
		func(t *testing.T) {
			position := 3
			res, err := secondIndex.knnSearchByVector(testVectors[position], 50, 36, nil)
			require.Nil(t, err)
			assert.Equal(t, expectedResults, res)
		})

	t.Run("further deleting all elements and reimporting one", func(t *testing.T) {
		toDelete := []int{0, 1, 2, 3, 4, 5, 7}

		for _, id := range toDelete {
			err := secondIndex.Delete(id)
			require.Nil(t, err)
		}

		err = secondIndex.CleanUpTombstonedNodes()
		require.Nil(t, err)

		err := secondIndex.Add(3, testVectors[3])
		require.Nil(t, err)
	})

	// dumpIndex(secondIndex)

	time.Sleep(100 * time.Millisecond) // TODO: make commit log write consinstent
	secondIndex = nil
	// build a new index from the (uncondensed) commit log
	thirdIndex, err := New(Config{
		RootPath:              dirName,
		ID:                    indexID,
		MakeCommitLoggerThunk: makeCL,
		MaximumConnections:    30,
		EFConstruction:        60,
		DistanceProvider:      distancer.NewCosineProvider(),
		VectorForIDThunk:      testVectorForID,
	})
	require.Nil(t, err)

	// dumpIndex(thirdIndex)

	t.Run("verify that the results match after rebuiling from disk",
		func(t *testing.T) {
			position := 3
			res, err := thirdIndex.knnSearchByVector(testVectors[position], 50, 36, nil)
			require.Nil(t, err)
			assert.Equal(t, []int{3}, res)
		})

	t.Run("delete all elements so the commitlog ends with an empty graph", func(t *testing.T) {
		toDelete := []int{3}

		for _, id := range toDelete {
			err := thirdIndex.Delete(id)
			require.Nil(t, err)
		}

		err = thirdIndex.CleanUpTombstonedNodes()
		require.Nil(t, err)
	})

	time.Sleep(100 * time.Millisecond) // TODO: make commit log write consinstent
	thirdIndex = nil
	// build a new index from the (uncondensed) commit log
	fourthIndex, err := New(Config{
		RootPath:              dirName,
		ID:                    indexID,
		MakeCommitLoggerThunk: makeCL,
		MaximumConnections:    30,
		EFConstruction:        60,
		DistanceProvider:      distancer.NewCosineProvider(),
		VectorForIDThunk:      testVectorForID,
	})
	require.Nil(t, err)

	t.Run("load from disk and try to insert again", func(t *testing.T) {
		for i, vec := range testVectors {
			err := fourthIndex.Add(i, vec)
			require.Nil(t, err)
		}
	})

	t.Run("verify that searching works normally", func(t *testing.T) {
		expectedResults := []int{
			3, 5, 4, // cluster 2
			7, 8, 6, // cluster 3 with element 6 and 8 deleted
			2, 1, 0, // cluster 1
		}
		position := 3
		res, err := fourthIndex.knnSearchByVector(testVectors[position], 50, 36, nil)
		require.Nil(t, err)
		assert.Equal(t, expectedResults, res)
	})
}
