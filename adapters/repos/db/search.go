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
	"fmt"
	"strings"
	"sync"

	"github.com/pkg/errors"
	"github.com/semi-technologies/weaviate/adapters/repos/db/refcache"
	"github.com/semi-technologies/weaviate/adapters/repos/db/storobj"
	"github.com/semi-technologies/weaviate/entities/aggregation"
	"github.com/semi-technologies/weaviate/entities/filters"
	"github.com/semi-technologies/weaviate/entities/schema"
	"github.com/semi-technologies/weaviate/entities/search"
	"github.com/semi-technologies/weaviate/usecases/traverser"
)

func (db *DB) Aggregate(ctx context.Context,
	params traverser.AggregateParams) (*aggregation.Result, error) {
	idx := db.GetIndex(schema.ClassName(params.ClassName))
	if idx == nil {
		return nil, fmt.Errorf("tried to browse non-existing index for %s", params.ClassName)
	}

	return idx.aggregate(ctx, params)
}

func (db *DB) ClassSearch(ctx context.Context,
	params traverser.GetParams) ([]search.Result, error) {
	idx := db.GetIndex(schema.ClassName(params.ClassName))
	if idx == nil {
		return nil, fmt.Errorf("tried to browse non-existing index for %s", params.ClassName)
	}

	if params.Pagination == nil {
		return nil, fmt.Errorf("invalid params, pagination object is nil")
	}

	res, err := idx.objectSearch(ctx, params.Pagination.Limit, params.Filters, params.AdditionalProperties)
	if err != nil {
		return nil, errors.Wrapf(err, "object search at index %s", idx.ID())
	}

	return db.enrichRefsForList(ctx, storobj.SearchResults(res, params.AdditionalProperties),
		params.Properties, params.AdditionalProperties)
}

func (db *DB) VectorClassSearch(ctx context.Context,
	params traverser.GetParams) ([]search.Result, error) {
	if params.SearchVector == nil {
		return db.ClassSearch(ctx, params)
	}

	idx := db.GetIndex(schema.ClassName(params.ClassName))
	if idx == nil {
		return nil, fmt.Errorf("tried to browse non-existing index for %s", params.ClassName)
	}

	res, err := idx.objectVectorSearch(ctx, params.SearchVector,
		params.Pagination.Limit, params.Filters, params.AdditionalProperties)
	if err != nil {
		return nil, errors.Wrapf(err, "object vector search at index %s", idx.ID())
	}

	return db.enrichRefsForList(ctx, storobj.SearchResults(res, params.AdditionalProperties),
		params.Properties, params.AdditionalProperties)
}

func (db *DB) VectorSearch(ctx context.Context, vector []float32, limit int,
	filters *filters.LocalFilter) ([]search.Result, error) {
	var found search.Results

	wg := &sync.WaitGroup{}
	mutex := &sync.Mutex{}
	var searchErrors []error

	emptyAdditional := traverser.AdditionalProperties{}
	for _, index := range db.indices {
		wg.Add(1)
		go func(index *Index, wg *sync.WaitGroup) {
			defer wg.Done()

			// TODO support all additional props
			res, err := index.objectVectorSearch(ctx, vector, limit, filters, emptyAdditional)
			if err != nil {
				mutex.Lock()
				searchErrors = append(searchErrors, errors.Wrapf(err, "search index %s", index.ID()))
				mutex.Unlock()
			}

			mutex.Lock()
			found = append(found, storobj.SearchResults(res, emptyAdditional)...)
			mutex.Unlock()
		}(index, wg)
	}

	wg.Wait()

	if len(searchErrors) > 0 {
		var msg strings.Builder
		for i, err := range searchErrors {
			if i != 0 {
				msg.WriteString(", ")
			}
			msg.WriteString(err.Error())
		}
		return nil, errors.New(msg.String())
	}

	found, err := found.SortByDistanceToVector(vector)
	if err != nil {
		return nil, errors.Wrapf(err, "re-sort when merging indices")
	}

	if len(found) > limit {
		found = found[:limit]
	}

	// not enriching by refs, as a vector search result cannot provide
	// SelectProperties
	return found, nil
}

func (d *DB) ObjectSearch(ctx context.Context, limit int, filters *filters.LocalFilter,
	additional traverser.AdditionalProperties) (search.Results, error) {
	return d.objectSearch(ctx, limit, filters, additional)
}

func (d *DB) objectSearch(ctx context.Context, limit int,
	filters *filters.LocalFilter,
	additional traverser.AdditionalProperties) (search.Results, error) {
	var found search.Results

	// TODO: Search in parallel, rather than sequentially or this will be
	// painfully slow on large schemas
	for _, index := range d.indices {
		// TODO support all additional props
		res, err := index.objectSearch(ctx, limit, filters, additional)
		if err != nil {
			return nil, errors.Wrapf(err, "search index %s", index.ID())
		}

		found = append(found, storobj.SearchResults(res, additional)...)
		if len(found) >= limit {
			// we are done
			break
		}
	}

	if len(found) > limit {
		found = found[:limit]
	}

	return found, nil
}

func (d *DB) enrichRefsForList(ctx context.Context, objs search.Results,
	props traverser.SelectProperties, additional traverser.AdditionalProperties) (search.Results, error) {
	res, err := refcache.NewResolver(refcache.NewCacher(d, d.logger)).
		Do(ctx, objs, props, additional)
	if err != nil {
		return nil, errors.Wrap(err, "resolve cross-refs")
	}

	return res, nil
}
