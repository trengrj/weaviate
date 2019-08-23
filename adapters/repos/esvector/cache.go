package esvector

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"

	"github.com/elastic/go-elasticsearch/v5/esapi"
	"github.com/go-openapi/strfmt"
	"github.com/semi-technologies/weaviate/entities/models"
	"github.com/semi-technologies/weaviate/entities/schema/crossref"
	"github.com/semi-technologies/weaviate/entities/schema/kind"
	"github.com/semi-technologies/weaviate/entities/search"
)

func (r *Repo) PopulateCache(ctx context.Context, kind kind.Kind, id strfmt.UUID) error {
	manager := newCacheManager(r)
	_, err := manager.populate(ctx, kind, id)
	if err != nil {
		return fmt.Errorf("populate cache for %s with id %s: %v", kind.Name(), id, err)
	}

	return nil
}

type cacheManager struct {
	repo *Repo
}

func newCacheManager(r *Repo) *cacheManager {
	return &cacheManager{r}
}

type refClassAndSchema struct {
	class  string
	schema map[string]interface{}
}

func (c *cacheManager) populate(ctx context.Context, kind kind.Kind, id strfmt.UUID) (*search.Result, error) {
	obj, err := c.getObject(ctx, kind, id)
	if err != nil {
		return nil, err
	}

	if obj == nil {
		return nil, fmt.Errorf("%s with id '%s' not found", kind.Name(), id)
	}

	if obj.CacheHot {
		// nothing to do, cache is already hot
		// TODO: make this check dependent on when the cache was last populated,
		// otherwise we can never renew if it becomes stale
		return prepareForStoringAsCache(obj), nil
	}

	resolvedSchema := map[string]interface{}{}
	schemaMap := obj.Schema.(map[string]interface{})
	for prop, value := range schemaMap {
		if gc, ok := value.(*models.GeoCoordinates); ok {
			resolvedSchema[prop] = map[string]interface{}{
				"lat": gc.Latitude,
				"lon": gc.Longitude,
			}
			continue
		}

		if _, ok := value.(*models.MultipleRef); ok {
			return nil, fmt.Errorf("if you see this message you have found a bug in weaviate" +
				", congrutulations! please open an issue at github.com/semi-technologies/weaviate" +
				" with the following error: found *models.MultipleRef in cache population, but" +
				" expected to only ever see models.MultipleRef")
		}

		refs, ok := value.(models.MultipleRef)
		if ok {
			resolvedRefs, err := c.resolveRefs(ctx, refs)
			if err != nil {
				return nil, err
			}

			resolvedSchema[prop] = groupRefByClassType(resolvedRefs)
			continue
		}

		resolvedSchema[prop] = value
	}

	obj.Schema = resolvedSchema

	if err := c.repo.upsertCache(ctx, id.String(), obj.Kind, obj.ClassName, resolvedSchema); err != nil {
		return nil, err
	}

	return obj, nil
}

func (c *cacheManager) getObject(ctx context.Context, k kind.Kind, id strfmt.UUID) (*search.Result, error) {
	switch k {
	case kind.Thing:
		return c.repo.ThingByID(ctx, id, 0)
	case kind.Action:
		return c.repo.ActionByID(ctx, id, 0)
	default:
		return nil, fmt.Errorf("impossible kind: %v", k)
	}
}

func (c *cacheManager) resolveRefs(ctx context.Context, refs models.MultipleRef) ([]refClassAndSchema, error) {
	var resolvedRefs []refClassAndSchema

	refSlice := []*models.SingleRef(refs)
	for _, ref := range refSlice {
		details, err := crossref.Parse(ref.Beacon.String())
		if err != nil {
			return nil, fmt.Errorf("parse %s: %v", ref.Beacon, err)
		}

		innerRef, err := c.populate(ctx, details.Kind, details.TargetID)
		if err != nil {
			return nil, fmt.Errorf("populate %s: %v", ref.Beacon, err)
		}

		if innerRef.Schema == nil {
			continue
		}

		resolvedRefs = append(resolvedRefs, refClassAndSchema{
			class:  innerRef.ClassName,
			schema: innerRef.Schema.(map[string]interface{}),
		})
	}

	return resolvedRefs, nil
}

func groupRefByClassType(refs []refClassAndSchema) map[string]interface{} {
	output := map[string]interface{}{}
	for _, ref := range refs {
		if slice, ok := output[ref.class]; ok {
			output[ref.class] = append(slice.([]interface{}), ref.schema)
		} else {
			output[ref.class] = []interface{}{ref.schema}
		}
	}

	return output
}

func (r *Repo) upsertCache(ctx context.Context, id string, k kind.Kind,
	className string, cache map[string]interface{}) error {
	// copy otherwise we modify the original when adding the cacheHot field
	cacheCopy := copyMap(cache)
	cacheCopy[keyCacheHot.String()] = true

	body := map[string]interface{}{
		"doc": map[string]interface{}{
			keyCache.String(): cacheCopy,
		},
	}

	var buf bytes.Buffer
	err := json.NewEncoder(&buf).Encode(body)
	if err != nil {
		return fmt.Errorf("upsert cache: encode json: %v", err)
	}

	req := esapi.UpdateRequest{
		Index:      classIndexFromClassName(k, className),
		DocumentID: id,
		Body:       &buf,
	}

	res, err := req.Do(ctx, r.client)
	if err != nil {
		return fmt.Errorf("upsert cache: %v", err)
	}

	if err := errorResToErr(res, r.logger); err != nil {
		r.logger.WithField("action", "vector_index_upsert_cache").
			WithError(err).
			WithField("request", req).
			WithField("res", res).
			WithField("body_before_marshal", body).
			WithField("body", buf.String()).
			Errorf("upsert cache failed")

		return fmt.Errorf("upsert cache: %v", err)
	}

	return nil
}

func copyMap(in map[string]interface{}) map[string]interface{} {
	out := make(map[string]interface{})
	for key, value := range in {
		out[key] = value
	}

	return out
}

// this function is only ever called if we had a hot cache. calling it will
// prepare the output for another insertion and will itself take refs from
// cache if it has any or leave them untouched if not cached
func prepareForStoringAsCache(in *search.Result) *search.Result {
	schema := map[string]interface{}{}

	for prop, value := range in.Schema.(map[string]interface{}) {
		switch v := value.(type) {
		case *models.GeoCoordinates:
			// geoocordniates need to be prepared for ES
			schema[prop] = map[string]interface{}{
				"lat": v.Latitude,
				"lon": v.Longitude,
			}

		case models.MultipleRef:
			// refs need to be taken from cache if present or left untouched otherwise
			if ref, ok := in.CacheSchema[prop]; ok {
				schema[prop] = ref
			} else {
				schema[prop] = value
			}
		default:
			// primitive props are good to go
			schema[prop] = value
		}
	}

	in.Schema = schema
	return in
}
