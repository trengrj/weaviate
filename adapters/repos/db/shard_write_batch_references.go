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

package db

import (
	"context"
	"sync"
	"time"

	"github.com/google/uuid"
	"github.com/pkg/errors"
	"github.com/semi-technologies/weaviate/usecases/objects"
	bolt "go.etcd.io/bbolt"
)

// return value map[int]error gives the error for the index as it received it
func (s *Shard) addReferencesBatch(ctx context.Context,
	refs objects.BatchReferences) map[int]error {
	return newReferencesBatcher(s).References(ctx, refs)
}

// referencesBatcher is a helper type wrapping around an underlying shard that can
// execute references batch operations on a shard (as opposed to object batch
// operations)
type referencesBatcher struct {
	sync.Mutex
	shard *Shard
	errs  map[int]error
	refs  objects.BatchReferences
}

func newReferencesBatcher(s *Shard) *referencesBatcher {
	return &referencesBatcher{
		shard: s,
	}
}

func (b *referencesBatcher) References(ctx context.Context,
	refs objects.BatchReferences) map[int]error {
	b.init(refs)
	b.storeInObjectStore(ctx)
	return b.errs
}

func (b *referencesBatcher) init(refs objects.BatchReferences) {
	b.refs = refs
	b.errs = map[int]error{} // int represents original index
}

func (b *referencesBatcher) storeInObjectStore(
	ctx context.Context) {
	maxPerTransaction := 30

	wg := &sync.WaitGroup{}
	for i := 0; i < len(b.refs); i += maxPerTransaction {
		end := i + maxPerTransaction
		if end > len(b.refs) {
			end = len(b.refs)
		}

		batch := b.refs[i:end]
		wg.Add(1)
		go func(i int, batch objects.BatchReferences) {
			defer wg.Done()
			var affectedIndices []int
			if err := b.shard.db.Batch(func(tx *bolt.Tx) error {
				var err error
				affectedIndices, err = b.storeSingleBatchInTx(ctx, tx, i, batch)
				return err
			}); err != nil {
				b.setErrorsForIndices(err, affectedIndices)
			}
		}(i, batch)
	}
	wg.Wait()

	// adding references can not alter the vector position, so no need to alter
	// the vector index
}

func (b *referencesBatcher) storeSingleBatchInTx(ctx context.Context, tx *bolt.Tx,
	batchId int, batch objects.BatchReferences) ([]int, error) {
	var affectedIndices []int
	for i := range batch {
		// so we can reference potential errors
		affectedIndices = append(affectedIndices, batchId+i)
	}

	for _, ref := range batch {
		uuidParsed, err := uuid.Parse(ref.From.TargetID.String())
		if err != nil {
			return nil, errors.Wrap(err, "invalid id")
		}

		idBytes, err := uuidParsed.MarshalBinary()
		if err != nil {
			return nil, err
		}

		mergeDoc := mergeDocFromBatchReference(ref)
		_, _, err = b.shard.mutableMergeObjectInTx(tx, mergeDoc, idBytes)
		if err != nil {
			return nil, err
		}

		// since we do a mutable update, we no longer need to carry the updates
		// forward to the additional storages. Since we only add references, it is
		// also impossible that any of their content was changed
	}

	return affectedIndices, nil
}

func (b *referencesBatcher) setErrorsForIndices(err error, affectedIndices []int) {
	b.Lock()
	defer b.Unlock()

	err = errors.Wrap(err, "bolt batch tx")
	for _, affected := range affectedIndices {
		b.errs[affected] = err
	}
}

func mergeDocFromBatchReference(ref objects.BatchReference) objects.MergeDocument {
	return objects.MergeDocument{
		Class:      ref.From.Class.String(),
		ID:         ref.From.TargetID,
		UpdateTime: time.Now().UnixNano(),
		References: objects.BatchReferences{ref},
	}
}
