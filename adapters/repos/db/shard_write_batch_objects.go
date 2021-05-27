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

package db

import (
	"context"
	"sync"
	"time"

	"github.com/go-openapi/strfmt"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"github.com/semi-technologies/weaviate/adapters/repos/db/helpers"
	"github.com/semi-technologies/weaviate/adapters/repos/db/inverted"
	"github.com/semi-technologies/weaviate/adapters/repos/db/storobj"
	"github.com/semi-technologies/weaviate/entities/schema"
)

// return value map[int]error gives the error for the index as it received it
func (s *Shard) putObjectBatch(ctx context.Context,
	objects []*storobj.Object) map[int]error {
	return newObjectsBatcher(s).Objects(ctx, objects)
}

// objectsBatcher is a helper type wrapping around an underlying shard that can
// execute objects batch operations on a shard (as opposed to references batch
// operations)
type objectsBatcher struct {
	sync.Mutex
	shard      *Shard
	statuses   map[strfmt.UUID]objectInsertStatus
	errs       map[int]error
	duplicates map[int]struct{}
	objects    []*storobj.Object
}

func newObjectsBatcher(s *Shard) *objectsBatcher {
	return &objectsBatcher{shard: s}
}

// Objects imports the specified objects in parallel in a batch-fashion
func (b *objectsBatcher) Objects(ctx context.Context,
	objects []*storobj.Object) map[int]error {
	beforeBatch := time.Now()
	defer b.shard.metrics.BatchObject(beforeBatch, len(objects))

	b.init(objects)
	b.storeInObjectStore(ctx)
	b.storeAdditionalStorage(ctx)
	return b.errs
}

func (b *objectsBatcher) init(objects []*storobj.Object) {
	b.objects = objects
	b.statuses = map[strfmt.UUID]objectInsertStatus{}
	b.errs = map[int]error{} // int represents original index
	b.duplicates = findDuplicatesInBatchObjects(objects)
}

// storeInObjectStore performs all storage operations on the underlying
// key/value store, this is they object-by-id store, the docID-lookup tables,
// as well as all inverted indices.
func (b *objectsBatcher) storeInObjectStore(ctx context.Context) {
	// maxPerTransaction := 30
	beforeObjectStore := time.Now()
	// wg := &sync.WaitGroup{}
	// for i := 0; i < len(b.objects); i += maxPerTransaction {
	// 	end := i + maxPerTransaction
	// 	if end > len(b.objects) {
	// 		end = len(b.objects)
	// 	}

	// 	batch := b.objects[i:end]
	// 	wg.Add(1)
	// 	go func(i int, batch []*storobj.Object) {
	// 		defer wg.Done()
	// 		var affectedIndices []int
	// 		if err := b.shard.db.Batch(func(tx *bolt.Tx) error {
	// 			var err error
	// 			affectedIndices, err = b.storeSingleBatchInTx(ctx, tx, i, batch)
	// 			return err
	// 		}); err != nil {
	// 			b.setErrorsForIndices(err, affectedIndices)
	// 		}
	// 	}(i, batch)
	// }
	// wg.Wait()

	errs := b.storeSingleBatchInLSM(ctx, b.objects)
	for i, err := range errs {
		if err != nil {
			b.setErrorAtIndex(err, i)
		}
	}

	b.shard.metrics.ObjectStore(beforeObjectStore)
}

// func (b *objectsBatcher) storeSingleBatchInTx(ctx context.Context, tx *bolt.Tx,
// 	batchId int, batch []*storobj.Object) ([]int, error) {
// 	var affectedIndices []int

// 	for j := range batch {
// 		// so we can reference potential errors
// 		affectedIndices = append(affectedIndices, batchId+j)
// 	}

// 	// only check context after assigning affected indices, otherwise a context
// 	// error can never be assigned to the correct items, see
// 	// https://github.com/semi-technologies/weaviate/issues/1363
// 	if err := ctx.Err(); err != nil {
// 		return affectedIndices, errors.Wrapf(err, "begin transaction %d of batch",
// 			batchId)
// 	}

// 	invertedMerger := inverted.NewDeltaMerger()
// 	// cache schema for the duration of the transaction
// 	classSchema := b.shard.index.getSchema.GetSchemaSkipAuth()

// 	for j, object := range batch {
// 		if err := b.storeObjectOfBatchInTx(ctx, tx, batchId, j, object); err != nil {
// 			return affectedIndices, errors.Wrapf(err, "store object %d", j)
// 		}

// 		if err := b.analyzeObjectForInvertedIndex(invertedMerger, classSchema,
// 			object); err != nil {
// 			return affectedIndices, errors.Wrapf(err, "analyze object %d", j)
// 		}
// 	}

// 	before := time.Now()
// 	if err := b.writeInvertedAdditions(tx,
// 		invertedMerger.Merge().Additions); err != nil {
// 		return affectedIndices, errors.Wrap(err, "updated inverted index")
// 	}
// 	b.shard.metrics.PutObjectUpdateInverted(before)

// 	return affectedIndices, nil
// }

func (b *objectsBatcher) storeSingleBatchInLSM(ctx context.Context,
	batch []*storobj.Object) []error {
	errs := make([]error, len(batch))
	errLock := &sync.Mutex{}

	// if the context is expired fail all
	if err := ctx.Err(); err != nil {
		for i := range errs {
			errs[i] = errors.Wrap(err, "begin batch")
		}
		return errs
	}

	// invertedMerger := inverted.NewDeltaMerger()
	// cache schema for the duration of the transaction
	// classSchema := b.shard.index.getSchema.GetSchemaSkipAuth()

	wg := &sync.WaitGroup{}
	for j, object := range batch {
		wg.Add(1)
		go func(index int, object *storobj.Object) {
			defer wg.Done()

			if err := b.storeObjectOfBatchInLSM(ctx, index, object); err != nil {
				errLock.Lock()
				errs[index] = err
				errLock.Unlock()
			}
		}(j, object)
	}
	wg.Wait()

	return errs
}

// nolint // TODO
func (b *objectsBatcher) analyzeObjectForInvertedIndex(merger *inverted.DeltaMerger,
	classSchema schema.Schema, obj *storobj.Object) error {
	propValues, ok := obj.Properties().(map[string]interface{})
	if !ok || propValues == nil {
		return nil
	}

	schemaClass := classSchema.FindClassByName(obj.Class())
	if schemaClass == nil {
		return errors.Errorf("class %q not present in schema", obj.Class())
	}

	analyzed, err := inverted.NewAnalyzer().Object(propValues, schemaClass.Properties,
		obj.ID())
	if err != nil {
		return err
	}

	merger.AddAdditions(analyzed, obj.DocID())
	return nil
}

// func (b *objectsBatcher) analyzeObjectAndStoreInvertedIndex(classSchema schema.Schema,
// 	obj *storobj.Object) error {
// 	propValues, ok := obj.Properties().(map[string]interface{})
// 	if !ok || propValues == nil {
// 		return nil
// 	}

// 	schemaClass := classSchema.FindClassByName(obj.Class())
// 	if schemaClass == nil {
// 		return errors.Errorf("class %q not present in schema", obj.Class())
// 	}

// 	analyzed, err := inverted.NewAnalyzer().Object(propValues, schemaClass.Properties,
// 		obj.ID())
// 	if err != nil {
// 		return err
// 	}

// 	for _, prop := range in {
// 		if prop.Has
// 		bucket := b.shard.store.Bucket(helpers.BucketFromPropNameLSM(prop.Name))
// 		if bucket == nil {
// 			return errors.Errorf("no bucket for prop '%s' found", prop.Name)
// 		}

// 		hashBucket := b.shard.store.Bucket(helpers.HashBucketFromPropNameLSM(prop.Name))
// 		if hashBucket == nil {
// 			return errors.Errorf("no hash bucket for prop '%s' found", prop.Name)
// 		}

// 		err := b.shard.extendInvertedIndexItemsLSM(bucket, hashBucket,
// 			item, prop.HasFrequency)
// 		if err != nil {
// 			return err
// 		}
// 	}

// 	return nil
// }

// nolint // TODO
func (b *objectsBatcher) writeInvertedAdditions(in []inverted.MergeProperty) error {
	for _, prop := range in {
		bucket := b.shard.store.Bucket(helpers.BucketFromPropNameLSM(prop.Name))
		if bucket == nil {
			return errors.Errorf("no bucket for prop '%s' found", prop.Name)
		}

		hashBucket := b.shard.store.Bucket(helpers.HashBucketFromPropNameLSM(prop.Name))
		if hashBucket == nil {
			return errors.Errorf("no hash bucket for prop '%s' found", prop.Name)
		}

		for _, item := range prop.MergeItems {
			err := b.shard.batchExtendInvertedIndexItemsLSM(bucket, hashBucket,
				item, prop.HasFrequency)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

// func (b *objectsBatcher) storeObjectOfBatchInTx(ctx context.Context, tx *bolt.Tx,
// 	batchId int, objectIndex int, object *storobj.Object) error {
// 	if _, ok := b.duplicates[batchId+objectIndex]; ok {
// 		return nil
// 	}
// 	uuidParsed, err := uuid.Parse(object.ID().String())
// 	if err != nil {
// 		return errors.Wrap(err, "invalid id")
// 	}

// 	idBytes, err := uuidParsed.MarshalBinary()
// 	if err != nil {
// 		return err
// 	}

// 	status, err := b.shard.putObjectInTx(tx, object, idBytes, true)
// 	if err != nil {
// 		return err
// 	}

// 	b.setStatusForID(status, object.ID())

// 	if err := ctx.Err(); err != nil {
// 		return errors.Wrapf(err, "end transaction %d of batch", batchId)
// 	}
// 	return nil
// }

func (b *objectsBatcher) storeObjectOfBatchInLSM(ctx context.Context,
	objectIndex int, object *storobj.Object) error {
	if _, ok := b.duplicates[objectIndex]; ok {
		return nil
	}
	uuidParsed, err := uuid.Parse(object.ID().String())
	if err != nil {
		return errors.Wrap(err, "invalid id")
	}

	idBytes, err := uuidParsed.MarshalBinary()
	if err != nil {
		return err
	}

	status, err := b.shard.putObjectLSM(object, idBytes, false)
	if err != nil {
		return err
	}

	b.setStatusForID(status, object.ID())

	if err := ctx.Err(); err != nil {
		return errors.Wrapf(err, "end store object %d of batch", objectIndex)
	}
	return nil
}

// setStatusForID is thread-safe as it uses the underlying mutex to lock the
// statuses map when writing into it
func (b *objectsBatcher) setStatusForID(status objectInsertStatus, id strfmt.UUID) {
	b.Lock()
	defer b.Unlock()
	b.statuses[id] = status
}

// storeAdditionalStorage stores the object in all non-key-value stores,
// such as the main vector index as well as the property-specific indices, such
// as the geo-index.
func (b *objectsBatcher) storeAdditionalStorage(ctx context.Context) {
	if ok := b.checkContext(ctx); !ok {
		// if the context is no longer OK, there's no point in continuing - abort
		// early
		return
	}

	beforeVectorIndex := time.Now()
	wg := &sync.WaitGroup{}
	for i, object := range b.objects {
		if b.shouldSkipInAdditionalStorage(i) {
			continue
		}

		wg.Add(1)
		status := b.statuses[object.ID()]
		go func(object *storobj.Object, status objectInsertStatus, index int) {
			defer wg.Done()
			b.storeSingleObjectInAdditionalStorage(ctx, object, status, index)
		}(object, status, i)
	}
	wg.Wait()
	b.shard.metrics.VectorIndex(beforeVectorIndex)
}

func (b *objectsBatcher) shouldSkipInAdditionalStorage(i int) bool {
	if ok := b.hasErrorAtIndex(i); ok {
		// had an error prior, ignore
		return true
	}

	// no need to lock the mutex for a duplicate check, as we only ever write
	// during init() in there - not concurrently
	if _, ok := b.duplicates[i]; ok {
		// is a duplicate, ignore
		return true
	}

	return false
}

func (b *objectsBatcher) storeSingleObjectInAdditionalStorage(ctx context.Context,
	object *storobj.Object, status objectInsertStatus, index int) {
	if err := ctx.Err(); err != nil {
		b.setErrorAtIndex(errors.Wrap(err, "insert to vector index"), index)
		return
	}

	if err := b.shard.updateVectorIndex(object.Vector, status); err != nil {
		b.setErrorAtIndex(errors.Wrap(err, "insert to vector index"), index)
		return
	}

	if err := b.shard.updatePropertySpecificIndices(object, status); err != nil {
		b.setErrorAtIndex(errors.Wrap(err, "update prop-specific indices"), index)
		return
	}
}

// hasErrorAtIndex is thread-safe as it uses the underlying mutex to lock
// before reading from the errs map
func (b *objectsBatcher) hasErrorAtIndex(i int) bool {
	b.Lock()
	defer b.Unlock()
	_, ok := b.errs[i]

	return ok
}

// setErrorAtIndex is thread-safe as it uses the underlying mutex to lock
// writing into the errs map
func (b *objectsBatcher) setErrorAtIndex(err error, index int) {
	b.Lock()
	defer b.Unlock()
	b.errs[index] = err
}

// // setErrorsForIndices is thread-safe as it uses the underlying mutex to lock
// // writing into the errs map
// func (b *objectsBatcher) setErrorsForIndices(err error, affectedIndices []int) {
// 	b.Lock()
// 	defer b.Unlock()

// 	err = errors.Wrap(err, "bolt batch tx")
// 	for _, affected := range affectedIndices {
// 		b.errs[affected] = err
// 	}
// }

// checkContext does nothing if the context is still active. But if the context
// has error'd, it marks all objects which have not previously error'd yet with
// the ctx error
func (s *objectsBatcher) checkContext(ctx context.Context) bool {
	if err := ctx.Err(); err != nil {
		for i, err := range s.errs {
			if err == nil {
				// already has an error, ignore
				continue
			}

			s.errs[i] = errors.Wrapf(err,
				"inverted indexing complete, about to start vector indexing")
		}

		return false
	}

	return true
}

// returns the originalIndexIDs to be ignored
func findDuplicatesInBatchObjects(in []*storobj.Object) map[int]struct{} {
	count := map[strfmt.UUID]int{}
	for _, obj := range in {
		count[obj.ID()] = count[obj.ID()] + 1
	}

	ignore := map[int]struct{}{}
	for i, obj := range in {
		if c := count[obj.ID()]; c > 1 {
			count[obj.ID()] = c - 1
			ignore[i] = struct{}{}
		}
	}

	return ignore
}
