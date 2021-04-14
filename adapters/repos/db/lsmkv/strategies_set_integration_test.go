// +build integrationTest

package lsmkv

import (
	"context"
	"fmt"
	"math/rand"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSetCollectionStrategy_InsertAndAppend(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	dirName := fmt.Sprintf("./testdata/%d", rand.Intn(10000000))
	os.MkdirAll(dirName, 0o777)
	defer func() {
		err := os.RemoveAll(dirName)
		fmt.Println(err)
	}()

	t.Run("memtable-only", func(t *testing.T) {
		b, err := NewBucketWithStrategy(dirName, StrategySetCollection)
		require.Nil(t, err)

		// so big it effectively never triggers as part of this test
		b.SetMemtableThreshold(1e9)

		key1 := []byte("test1-key-1")
		key2 := []byte("test1-key-2")
		key3 := []byte("test1-key-3")

		t.Run("set original values and verify", func(t *testing.T) {
			orig1 := [][]byte{[]byte("value 1.1"), []byte("value 1.2")}
			orig2 := [][]byte{[]byte("value 2.1"), []byte("value 2.2")}
			orig3 := [][]byte{[]byte("value 3.1"), []byte("value 3.2")}

			err = b.Append(key1, orig1)
			require.Nil(t, err)
			err = b.Append(key2, orig2)
			require.Nil(t, err)
			err = b.Append(key3, orig3)
			require.Nil(t, err)

			res, err := b.GetCollection(key1)
			require.Nil(t, err)
			assert.Equal(t, res, orig1)
			res, err = b.GetCollection(key2)
			require.Nil(t, err)
			assert.Equal(t, res, orig2)
			res, err = b.GetCollection(key3)
			require.Nil(t, err)
			assert.Equal(t, res, orig3)
		})

		t.Run("replace some, keep one", func(t *testing.T) {
			orig1 := [][]byte{[]byte("value 1.1"), []byte("value 1.2")}
			orig2 := [][]byte{[]byte("value 2.1"), []byte("value 2.2")}
			orig3 := [][]byte{[]byte("value 3.1"), []byte("value 3.2")}
			append2 := [][]byte{[]byte("value 2.3")}
			append3 := [][]byte{[]byte("value 3.3")}

			err = b.Append(key2, append2)
			require.Nil(t, err)
			err = b.Append(key3, append3)
			require.Nil(t, err)

			res, err := b.GetCollection(key1)
			require.Nil(t, err)
			assert.Equal(t, orig1, res)
			res, err = b.GetCollection(key2)
			require.Nil(t, err)
			assert.Equal(t, append(orig2, append2...), res)
			res, err = b.GetCollection(key3)
			require.Nil(t, err)
			assert.Equal(t, append(orig3, append3...), res)
		})
	})

	t.Run("with a single flush between updates", func(t *testing.T) {
		b, err := NewBucketWithStrategy(dirName, StrategySetCollection)
		require.Nil(t, err)

		// so big it effectively never triggers as part of this test
		b.SetMemtableThreshold(1e9)

		key1 := []byte("test2-key-1")
		key2 := []byte("test2-key-2")
		key3 := []byte("test2-key-3")

		t.Run("set original values and verify", func(t *testing.T) {
			orig1 := [][]byte{[]byte("value 1.1"), []byte("value 1.2")}
			orig2 := [][]byte{[]byte("value 2.1"), []byte("value 2.2")}
			orig3 := [][]byte{[]byte("value 3.1"), []byte("value 3.2")}

			err = b.Append(key1, orig1)
			require.Nil(t, err)
			err = b.Append(key2, orig2)
			require.Nil(t, err)
			err = b.Append(key3, orig3)
			require.Nil(t, err)

			res, err := b.GetCollection(key1)
			require.Nil(t, err)
			assert.Equal(t, res, orig1)
			res, err = b.GetCollection(key2)
			require.Nil(t, err)
			assert.Equal(t, res, orig2)
			res, err = b.GetCollection(key3)
			require.Nil(t, err)
			assert.Equal(t, res, orig3)
		})

		t.Run("flush to disk", func(t *testing.T) {
			require.Nil(t, b.FlushAndSwitch())
		})

		t.Run("replace some, keep one", func(t *testing.T) {
			orig1 := [][]byte{[]byte("value 1.1"), []byte("value 1.2")}
			orig2 := [][]byte{[]byte("value 2.1"), []byte("value 2.2")}
			orig3 := [][]byte{[]byte("value 3.1"), []byte("value 3.2")}
			append2 := [][]byte{[]byte("value 2.3")}
			append3 := [][]byte{[]byte("value 3.3")}

			err = b.Append(key2, append2)
			require.Nil(t, err)
			err = b.Append(key3, append3)
			require.Nil(t, err)

			res, err := b.GetCollection(key1)
			require.Nil(t, err)
			assert.Equal(t, orig1, res)
			res, err = b.GetCollection(key2)
			require.Nil(t, err)
			assert.Equal(t, append(orig2, append2...), res)
			res, err = b.GetCollection(key3)
			require.Nil(t, err)
			assert.Equal(t, append(orig3, append3...), res)
		})
	})

	t.Run("with flushes after initial and update", func(t *testing.T) {
		b, err := NewBucketWithStrategy(dirName, StrategySetCollection)
		require.Nil(t, err)

		// so big it effectively never triggers as part of this test
		b.SetMemtableThreshold(1e9)
		key1 := []byte("test-3-key-1")
		key2 := []byte("test-3-key-2")
		key3 := []byte("test-3-key-3")

		t.Run("set original values and verify", func(t *testing.T) {
			orig1 := [][]byte{[]byte("value 1.1"), []byte("value 1.2")}
			orig2 := [][]byte{[]byte("value 2.1"), []byte("value 2.2")}
			orig3 := [][]byte{[]byte("value 3.1"), []byte("value 3.2")}

			err = b.Append(key1, orig1)
			require.Nil(t, err)
			err = b.Append(key2, orig2)
			require.Nil(t, err)
			err = b.Append(key3, orig3)
			require.Nil(t, err)

			res, err := b.GetCollection(key1)
			require.Nil(t, err)
			assert.Equal(t, res, orig1)
			res, err = b.GetCollection(key2)
			require.Nil(t, err)
			assert.Equal(t, res, orig2)
			res, err = b.GetCollection(key3)
			require.Nil(t, err)
			assert.Equal(t, res, orig3)
		})

		t.Run("flush to disk", func(t *testing.T) {
			require.Nil(t, b.FlushAndSwitch())
		})

		t.Run("replace some, keep one", func(t *testing.T) {
			orig1 := [][]byte{[]byte("value 1.1"), []byte("value 1.2")}
			orig2 := [][]byte{[]byte("value 2.1"), []byte("value 2.2")}
			orig3 := [][]byte{[]byte("value 3.1"), []byte("value 3.2")}
			append2 := [][]byte{[]byte("value 2.3")}
			append3 := [][]byte{[]byte("value 3.3")}

			err = b.Append(key2, append2)
			require.Nil(t, err)
			err = b.Append(key3, append3)
			require.Nil(t, err)

			// Flush again!
			require.Nil(t, b.FlushAndSwitch())

			res, err := b.GetCollection(key1)
			require.Nil(t, err)
			assert.Equal(t, orig1, res)
			res, err = b.GetCollection(key2)
			require.Nil(t, err)
			assert.Equal(t, append(orig2, append2...), res)
			res, err = b.GetCollection(key3)
			require.Nil(t, err)
			assert.Equal(t, append(orig3, append3...), res)
		})
	})

	t.Run("update in memtable, then do an orderly shutdown, and re-init", func(t *testing.T) {
		b, err := NewBucketWithStrategy(dirName, StrategySetCollection)
		require.Nil(t, err)

		// so big it effectively never triggers as part of this test
		b.SetMemtableThreshold(1e9)

		key1 := []byte("test4-key-1")
		key2 := []byte("test4-key-2")
		key3 := []byte("test4-key-3")

		t.Run("set original values and verify", func(t *testing.T) {
			orig1 := [][]byte{[]byte("value 1.1"), []byte("value 1.2")}
			orig2 := [][]byte{[]byte("value 2.1"), []byte("value 2.2")}
			orig3 := [][]byte{[]byte("value 3.1"), []byte("value 3.2")}

			err = b.Append(key1, orig1)
			require.Nil(t, err)
			err = b.Append(key2, orig2)
			require.Nil(t, err)
			err = b.Append(key3, orig3)
			require.Nil(t, err)

			res, err := b.GetCollection(key1)
			require.Nil(t, err)
			assert.Equal(t, res, orig1)
			res, err = b.GetCollection(key2)
			require.Nil(t, err)
			assert.Equal(t, res, orig2)
			res, err = b.GetCollection(key3)
			require.Nil(t, err)
			assert.Equal(t, res, orig3)
		})

		t.Run("replace some, keep one", func(t *testing.T) {
			orig1 := [][]byte{[]byte("value 1.1"), []byte("value 1.2")}
			orig2 := [][]byte{[]byte("value 2.1"), []byte("value 2.2")}
			orig3 := [][]byte{[]byte("value 3.1"), []byte("value 3.2")}
			append2 := [][]byte{[]byte("value 2.3")}
			append3 := [][]byte{[]byte("value 3.3")}

			err = b.Append(key2, append2)
			require.Nil(t, err)
			err = b.Append(key3, append3)
			require.Nil(t, err)

			res, err := b.GetCollection(key1)
			require.Nil(t, err)
			assert.Equal(t, orig1, res)
			res, err = b.GetCollection(key2)
			require.Nil(t, err)
			assert.Equal(t, append(orig2, append2...), res)
			res, err = b.GetCollection(key3)
			require.Nil(t, err)
			assert.Equal(t, append(orig3, append3...), res)
		})

		t.Run("orderly shutdown", func(t *testing.T) {
			b.Shutdown(context.Background())
		})

		t.Run("init another bucket on the same files", func(t *testing.T) {
			b2, err := NewBucketWithStrategy(dirName, StrategySetCollection)
			require.Nil(t, err)

			orig1 := [][]byte{[]byte("value 1.1"), []byte("value 1.2")}
			orig2 := [][]byte{[]byte("value 2.1"), []byte("value 2.2")}
			orig3 := [][]byte{[]byte("value 3.1"), []byte("value 3.2")}
			append2 := [][]byte{[]byte("value 2.3")}
			append3 := [][]byte{[]byte("value 3.3")}

			res, err := b2.GetCollection(key1)
			require.Nil(t, err)
			assert.Equal(t, orig1, res)
			res, err = b2.GetCollection(key2)
			require.Nil(t, err)
			assert.Equal(t, append(orig2, append2...), res)
			res, err = b2.GetCollection(key3)
			require.Nil(t, err)
			assert.Equal(t, append(orig3, append3...), res)
		})
	})
}

func TestSetCollectionStrategy_InsertAndDelete(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	dirName := fmt.Sprintf("./testdata/%d", rand.Intn(10000000))
	os.MkdirAll(dirName, 0o777)
	defer func() {
		err := os.RemoveAll(dirName)
		fmt.Println(err)
	}()

	t.Run("memtable-only", func(t *testing.T) {
		b, err := NewBucketWithStrategy(dirName, StrategySetCollection)
		require.Nil(t, err)

		// so big it effectively never triggers as part of this test
		b.SetMemtableThreshold(1e9)

		key1 := []byte("test1-key-1")
		key2 := []byte("test1-key-2")
		key3 := []byte("test1-key-3")

		t.Run("set original values and verify", func(t *testing.T) {
			orig1 := [][]byte{[]byte("value 1.1"), []byte("value 1.2")}
			orig2 := [][]byte{[]byte("value 2.1"), []byte("value 2.2")}
			orig3 := [][]byte{[]byte("value 3.1"), []byte("value 3.2")}

			err = b.Append(key1, orig1)
			require.Nil(t, err)
			err = b.Append(key2, orig2)
			require.Nil(t, err)
			err = b.Append(key3, orig3)
			require.Nil(t, err)

			res, err := b.GetCollection(key1)
			require.Nil(t, err)
			assert.Equal(t, orig1, res)
			res, err = b.GetCollection(key2)
			require.Nil(t, err)
			assert.Equal(t, orig2, res)
			res, err = b.GetCollection(key3)
			require.Nil(t, err)
			assert.Equal(t, orig3, res)
		})

		t.Run("delete individual keys", func(t *testing.T) {
			delete2 := []byte("value 2.1")
			delete3 := []byte("value 3.2")

			err = b.DeleteFromCollection(key2, delete2)
			require.Nil(t, err)
			err = b.DeleteFromCollection(key3, delete3)
			require.Nil(t, err)
		})

		t.Run("validate the results", func(t *testing.T) {
			expected1 := [][]byte{[]byte("value 1.1"), []byte("value 1.2")} // unchanged
			expected2 := [][]byte{[]byte("value 2.2")}                      // value1 deleted
			expected3 := [][]byte{[]byte("value 3.1")}                      // value2 deleted

			res, err := b.GetCollection(key1)
			require.Nil(t, err)
			assert.Equal(t, expected1, res)
			res, err = b.GetCollection(key2)
			require.Nil(t, err)
			assert.Equal(t, expected2, res)
			res, err = b.GetCollection(key3)
			require.Nil(t, err)
			assert.Equal(t, expected3, res)
		})

		t.Run("re-add keys which were previously deleted and new ones", func(t *testing.T) {
			readd2 := [][]byte{[]byte("value 2.1"), []byte("value 2.3")}
			readd3 := [][]byte{[]byte("value 3.2"), []byte("value 3.3")}

			err = b.Append(key2, readd2)
			require.Nil(t, err)
			err = b.Append(key3, readd3)
			require.Nil(t, err)
		})

		t.Run("validate the results again", func(t *testing.T) {
			expected1 := [][]byte{[]byte("value 1.1"), []byte("value 1.2")} // unchanged
			expected2 := [][]byte{
				[]byte("value 2.2"), // from original import
				[]byte("value 2.1"), // added again after initial deletion
				[]byte("value 2.3"), // newly added
			}
			expected3 := [][]byte{
				[]byte("value 3.1"), // form original import
				[]byte("value 3.2"), // added again after initial deletion
				[]byte("value 3.3"), // newly added
			} // value2 deleted

			res, err := b.GetCollection(key1)
			require.Nil(t, err)
			assert.Equal(t, expected1, res)
			res, err = b.GetCollection(key2)
			require.Nil(t, err)
			assert.Equal(t, expected2, res)
			res, err = b.GetCollection(key3)
			require.Nil(t, err)
			assert.Equal(t, expected3, res)
		})
	})

	t.Run("with a single flush between updates", func(t *testing.T) {
		b, err := NewBucketWithStrategy(dirName, StrategySetCollection)
		require.Nil(t, err)

		// so big it effectively never triggers as part of this test
		b.SetMemtableThreshold(1e9)

		key1 := []byte("test2-key-1")
		key2 := []byte("test2-key-2")
		key3 := []byte("test2-key-3")

		t.Run("set original values and verify", func(t *testing.T) {
			orig1 := [][]byte{[]byte("value 1.1"), []byte("value 1.2")}
			orig2 := [][]byte{[]byte("value 2.1"), []byte("value 2.2")}
			orig3 := [][]byte{[]byte("value 3.1"), []byte("value 3.2")}

			err = b.Append(key1, orig1)
			require.Nil(t, err)
			err = b.Append(key2, orig2)
			require.Nil(t, err)
			err = b.Append(key3, orig3)
			require.Nil(t, err)

			res, err := b.GetCollection(key1)
			require.Nil(t, err)
			assert.Equal(t, res, orig1)
			res, err = b.GetCollection(key2)
			require.Nil(t, err)
			assert.Equal(t, res, orig2)
			res, err = b.GetCollection(key3)
			require.Nil(t, err)
			assert.Equal(t, res, orig3)
		})

		t.Run("flush to disk", func(t *testing.T) {
			require.Nil(t, b.FlushAndSwitch())
		})

		t.Run("delete individual keys", func(t *testing.T) {
			delete2 := []byte("value 2.1")
			delete3 := []byte("value 3.2")

			err = b.DeleteFromCollection(key2, delete2)
			require.Nil(t, err)
			err = b.DeleteFromCollection(key3, delete3)
			require.Nil(t, err)
		})

		t.Run("validate the results", func(t *testing.T) {
			expected1 := [][]byte{[]byte("value 1.1"), []byte("value 1.2")} // unchanged
			expected2 := [][]byte{[]byte("value 2.2")}                      // value1 deleted
			expected3 := [][]byte{[]byte("value 3.1")}                      // value2 deleted

			res, err := b.GetCollection(key1)
			require.Nil(t, err)
			assert.Equal(t, expected1, res)
			res, err = b.GetCollection(key2)
			require.Nil(t, err)
			assert.Equal(t, expected2, res)
			res, err = b.GetCollection(key3)
			require.Nil(t, err)
			assert.Equal(t, expected3, res)
		})

		t.Run("re-add keys which were previously deleted and new ones", func(t *testing.T) {
			readd2 := [][]byte{[]byte("value 2.1"), []byte("value 2.3")}
			readd3 := [][]byte{[]byte("value 3.2"), []byte("value 3.3")}

			err = b.Append(key2, readd2)
			require.Nil(t, err)
			err = b.Append(key3, readd3)
			require.Nil(t, err)
		})

		t.Run("validate the results again", func(t *testing.T) {
			expected1 := [][]byte{[]byte("value 1.1"), []byte("value 1.2")} // unchanged
			expected2 := [][]byte{
				[]byte("value 2.2"), // from original import
				[]byte("value 2.1"), // added again after initial deletion
				[]byte("value 2.3"), // newly added
			}
			expected3 := [][]byte{
				[]byte("value 3.1"), // form original import
				[]byte("value 3.2"), // added again after initial deletion
				[]byte("value 3.3"), // newly added
			} // value2 deleted

			res, err := b.GetCollection(key1)
			require.Nil(t, err)
			assert.Equal(t, expected1, res)
			res, err = b.GetCollection(key2)
			require.Nil(t, err)
			assert.Equal(t, expected2, res)
			res, err = b.GetCollection(key3)
			require.Nil(t, err)
			assert.Equal(t, expected3, res)
		})
	})

	t.Run("with flushes in between and after the update", func(t *testing.T) {
		b, err := NewBucketWithStrategy(dirName, StrategySetCollection)
		require.Nil(t, err)

		// so big it effectively never triggers as part of this test
		b.SetMemtableThreshold(1e9)

		key1 := []byte("test3-key-1")
		key2 := []byte("test3-key-2")
		key3 := []byte("test3-key-3")

		t.Run("set original values and verify", func(t *testing.T) {
			orig1 := [][]byte{[]byte("value 1.1"), []byte("value 1.2")}
			orig2 := [][]byte{[]byte("value 2.1"), []byte("value 2.2")}
			orig3 := [][]byte{[]byte("value 3.1"), []byte("value 3.2")}

			err = b.Append(key1, orig1)
			require.Nil(t, err)
			err = b.Append(key2, orig2)
			require.Nil(t, err)
			err = b.Append(key3, orig3)
			require.Nil(t, err)

			res, err := b.GetCollection(key1)
			require.Nil(t, err)
			assert.Equal(t, res, orig1)
			res, err = b.GetCollection(key2)
			require.Nil(t, err)
			assert.Equal(t, res, orig2)
			res, err = b.GetCollection(key3)
			require.Nil(t, err)
			assert.Equal(t, res, orig3)
		})

		t.Run("flush to disk", func(t *testing.T) {
			require.Nil(t, b.FlushAndSwitch())
		})

		t.Run("delete individual keys", func(t *testing.T) {
			delete2 := []byte("value 2.1")
			delete3 := []byte("value 3.2")

			err = b.DeleteFromCollection(key2, delete2)
			require.Nil(t, err)
			err = b.DeleteFromCollection(key3, delete3)
			require.Nil(t, err)
		})

		t.Run("flush to disk - again!", func(t *testing.T) {
			require.Nil(t, b.FlushAndSwitch())
		})

		t.Run("validate", func(t *testing.T) {
			expected1 := [][]byte{[]byte("value 1.1"), []byte("value 1.2")} // unchanged
			expected2 := [][]byte{[]byte("value 2.2")}                      // value1 deleted
			expected3 := [][]byte{[]byte("value 3.1")}                      // value2 deleted

			res, err := b.GetCollection(key1)
			require.Nil(t, err)
			assert.Equal(t, expected1, res)
			res, err = b.GetCollection(key2)
			require.Nil(t, err)
			assert.Equal(t, expected2, res)
			res, err = b.GetCollection(key3)
			require.Nil(t, err)
			assert.Equal(t, expected3, res)
		})

		t.Run("re-add keys which were previously deleted and new ones", func(t *testing.T) {
			readd2 := [][]byte{[]byte("value 2.1"), []byte("value 2.3")}
			readd3 := [][]byte{[]byte("value 3.2"), []byte("value 3.3")}

			err = b.Append(key2, readd2)
			require.Nil(t, err)
			err = b.Append(key3, readd3)
			require.Nil(t, err)
		})

		t.Run("flush to disk - yet again!", func(t *testing.T) {
			require.Nil(t, b.FlushAndSwitch())
		})

		t.Run("validate the results again", func(t *testing.T) {
			expected1 := [][]byte{[]byte("value 1.1"), []byte("value 1.2")} // unchanged
			expected2 := [][]byte{
				[]byte("value 2.2"), // from original import
				[]byte("value 2.1"), // added again after initial deletion
				[]byte("value 2.3"), // newly added
			}
			expected3 := [][]byte{
				[]byte("value 3.1"), // form original import
				[]byte("value 3.2"), // added again after initial deletion
				[]byte("value 3.3"), // newly added
			} // value2 deleted

			res, err := b.GetCollection(key1)
			require.Nil(t, err)
			assert.Equal(t, expected1, res)
			res, err = b.GetCollection(key2)
			require.Nil(t, err)
			assert.Equal(t, expected2, res)
			res, err = b.GetCollection(key3)
			require.Nil(t, err)
			assert.Equal(t, expected3, res)
		})
	})

	t.Run("update in memtable, make orderly shutdown, then create a new bucket from disk",
		func(t *testing.T) {
			b, err := NewBucketWithStrategy(dirName, StrategySetCollection)
			require.Nil(t, err)

			// so big it effectively never triggers as part of this test
			b.SetMemtableThreshold(1e9)

			key1 := []byte("test4-key-1")
			key2 := []byte("test4-key-2")
			key3 := []byte("test4-key-3")

			t.Run("set original values and verify", func(t *testing.T) {
				orig1 := [][]byte{[]byte("value 1.1"), []byte("value 1.2")}
				orig2 := [][]byte{[]byte("value 2.1"), []byte("value 2.2")}
				orig3 := [][]byte{[]byte("value 3.1"), []byte("value 3.2")}

				err = b.Append(key1, orig1)
				require.Nil(t, err)
				err = b.Append(key2, orig2)
				require.Nil(t, err)
				err = b.Append(key3, orig3)
				require.Nil(t, err)

				res, err := b.GetCollection(key1)
				require.Nil(t, err)
				assert.Equal(t, res, orig1)
				res, err = b.GetCollection(key2)
				require.Nil(t, err)
				assert.Equal(t, res, orig2)
				res, err = b.GetCollection(key3)
				require.Nil(t, err)
				assert.Equal(t, res, orig3)
			})

			t.Run("delete individual keys", func(t *testing.T) {
				delete2 := []byte("value 2.1")
				delete3 := []byte("value 3.2")

				err = b.DeleteFromCollection(key2, delete2)
				require.Nil(t, err)
				err = b.DeleteFromCollection(key3, delete3)
				require.Nil(t, err)
			})

			t.Run("orderly shutdown", func(t *testing.T) {
				b.Shutdown(context.Background())
			})

			t.Run("init another bucket on the same files", func(t *testing.T) {
				b2, err := NewBucketWithStrategy(dirName, StrategySetCollection)
				require.Nil(t, err)

				expected1 := [][]byte{[]byte("value 1.1"), []byte("value 1.2")} // unchanged
				expected2 := [][]byte{[]byte("value 2.2")}                      // value1 deleted
				expected3 := [][]byte{[]byte("value 3.1")}                      // value2 deleted

				res, err := b2.GetCollection(key1)
				require.Nil(t, err)
				assert.Equal(t, expected1, res)
				res, err = b2.GetCollection(key2)
				require.Nil(t, err)
				assert.Equal(t, expected2, res)
				res, err = b2.GetCollection(key3)
				require.Nil(t, err)
				assert.Equal(t, expected3, res)
			})
		})
}
