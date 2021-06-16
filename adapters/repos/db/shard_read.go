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
	"bytes"
	"context"
	"encoding/binary"

	"github.com/go-openapi/strfmt"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"github.com/semi-technologies/weaviate/adapters/repos/db/helpers"
	"github.com/semi-technologies/weaviate/adapters/repos/db/inverted"
	"github.com/semi-technologies/weaviate/adapters/repos/db/storobj"
	"github.com/semi-technologies/weaviate/entities/filters"
	"github.com/semi-technologies/weaviate/entities/multi"
	"github.com/semi-technologies/weaviate/usecases/traverser"
)

func (s *Shard) objectByID(ctx context.Context, id strfmt.UUID,
	props traverser.SelectProperties,
	additional traverser.AdditionalProperties) (*storobj.Object, error) {
	idBytes, err := uuid.MustParse(id.String()).MarshalBinary()
	if err != nil {
		return nil, err
	}

	bytes, err := s.store.Bucket(helpers.ObjectsBucketLSM).Get(idBytes)
	if err != nil {
		return nil, err
	}

	if bytes == nil {
		return nil, nil
	}

	obj, err := storobj.FromBinary(bytes)
	if err != nil {
		return nil, errors.Wrap(err, "unmarshal object")
	}

	return obj, nil
}

func (s *Shard) multiObjectByID(ctx context.Context,
	query []multi.Identifier) ([]*storobj.Object, error) {
	objects := make([]*storobj.Object, len(query))

	ids := make([][]byte, len(query))
	for i, q := range query {
		idBytes, err := uuid.MustParse(q.ID).MarshalBinary()
		if err != nil {
			return nil, err
		}

		ids[i] = idBytes
	}

	bucket := s.store.Bucket(helpers.ObjectsBucketLSM)
	for i, id := range ids {
		bytes, err := bucket.Get(id)
		if err != nil {
			return nil, err
		}

		if bytes == nil {
			continue
		}

		obj, err := storobj.FromBinary(bytes)
		if err != nil {
			return nil, errors.Wrap(err, "unmarshal kind object")
		}
		objects[i] = obj
	}

	return objects, nil
}

// TODO: This does an actual read which is not really needed, if we see this
// come up in profiling, we could optimize this by adding an explicit Exists()
// on the LSMKV which only checks the bloom filters, which at least in the case
// of a true negative would be considerably faster. For a (false) positive,
// we'd still need to check, though.
func (s *Shard) exists(ctx context.Context, id strfmt.UUID) (bool, error) {
	idBytes, err := uuid.MustParse(id.String()).MarshalBinary()
	if err != nil {
		return false, err
	}

	bytes, err := s.store.Bucket(helpers.ObjectsBucketLSM).Get(idBytes)
	if err != nil {
		return false, errors.Wrap(err, "read request")
	}

	if bytes == nil {
		return false, nil
	}

	return true, nil
}

func (s *Shard) objectByIndexID(ctx context.Context,
	indexID uint64, acceptDeleted bool) (*storobj.Object, error) {
	keyBuf := bytes.NewBuffer(nil)
	binary.Write(keyBuf, binary.LittleEndian, &indexID)
	key := keyBuf.Bytes()

	// uuidLookup, err := s.store.Bucket(helpers.ObjectsBucketLSM).GetBySecondary(0, key)
	// if err != nil {
	// 	return nil, err
	// }

	// if uuidLookup == nil {
	// 	return nil, storobj.NewErrNotFoundf(indexID,
	// 		"doc id inverted resolved to a nil object, i.e. no uuid found")
	// }

	// lookup, err := docid.LookupFromBinary(uuidLookup)
	// if err != nil {
	// 	return nil, errors.Wrap(err, "unmarshal docID lookup")
	// }

	// if lookup.Deleted && !acceptDeleted {
	// 	return nil, storobj.NewErrNotFoundf(indexID,
	// 		"doc id is marked as deleted at %s",
	// 		lookup.DeletionTime.String())
	// }

	bytes, err := s.store.Bucket(helpers.ObjectsBucketLSM).
		GetBySecondary(0, key)
	if err != nil {
		return nil, err
	}

	if bytes == nil {
		return nil, storobj.NewErrNotFoundf(indexID,
			"uuid found for docID, but object is nil")
	}

	obj, err := storobj.FromBinary(bytes)
	if err != nil {
		return nil, errors.Wrap(err, "unmarshal kind object")
	}

	return obj, nil
}

func (s *Shard) vectorByIndexID(ctx context.Context, indexID uint64) ([]float32, error) {
	obj, err := s.objectByIndexID(ctx, indexID, true)
	if err != nil {
		return nil, err
	}

	return obj.Vector, nil
}

func (s *Shard) objectSearch(ctx context.Context, limit int,
	filters *filters.LocalFilter, additional traverser.AdditionalProperties) ([]*storobj.Object, error) {
	if filters == nil {
		return s.objectList(ctx, limit, additional)
	}

	return inverted.NewSearcher(s.store, s.index.getSchema.GetSchemaSkipAuth(),
		s.invertedRowCache, s.propertyIndices, s.index.classSearcher,
		s.deletedDocIDs).
		Object(ctx, limit, filters, additional, s.index.Config.ClassName)
}

func (s *Shard) objectVectorSearch(ctx context.Context, searchVector []float32,
	limit int, filters *filters.LocalFilter, additional traverser.AdditionalProperties) ([]*storobj.Object, error) {
	var allowList helpers.AllowList
	if filters != nil {
		list, err := inverted.NewSearcher(s.store, s.index.getSchema.GetSchemaSkipAuth(),
			s.invertedRowCache, s.propertyIndices, s.index.classSearcher,
			s.deletedDocIDs).
			DocIDs(ctx, filters, additional, s.index.Config.ClassName)
		if err != nil {
			return nil, errors.Wrap(err, "build inverted filter allow list")
		}

		allowList = list
	}
	ids, err := s.vectorIndex.SearchByVector(searchVector, limit, allowList)
	if err != nil {
		return nil, errors.Wrap(err, "vector search")
	}

	if len(ids) == 0 {
		return nil, nil
	}

	return s.objectsByDocID(ids)
}

func (s *Shard) objectsByDocID(ids []uint64) ([]*storobj.Object, error) {
	out := make([]*storobj.Object, len(ids))

	bucket := s.store.Bucket(helpers.ObjectsBucketLSM)
	if bucket == nil {
		return nil, errors.Errorf("objects bucket not found")
	}

	i := 0

	for _, id := range ids {
		keyBuf := bytes.NewBuffer(nil)
		binary.Write(keyBuf, binary.LittleEndian, &id)
		docIDBytes := keyBuf.Bytes()
		res, err := bucket.GetBySecondary(0, docIDBytes)
		if err != nil {
			return nil, err
		}

		if res == nil {
			continue
		}

		unmarshalled, err := storobj.FromBinary(res)
		if err != nil {
			return nil, errors.Wrapf(err, "unmarshal data object at position %d", i)
		}

		out[i] = unmarshalled
		i++
	}

	return out[:i], nil
}

func (s *Shard) objectList(ctx context.Context, limit int,
	additional traverser.AdditionalProperties) ([]*storobj.Object, error) {
	out := make([]*storobj.Object, limit)
	i := 0
	cursor := s.store.Bucket(helpers.ObjectsBucketLSM).Cursor()
	defer cursor.Close()

	for k, v := cursor.First(); k != nil && i < limit; k, v = cursor.Next() {
		obj, err := storobj.FromBinary(v)
		if err != nil {
			return nil, errors.Wrapf(err, "unmarhsal item %d", i)
		}

		out[i] = obj
		i++
	}

	return out[:i], nil
}
