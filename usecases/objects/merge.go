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

package objects

import (
	"context"
	"fmt"

	"github.com/go-openapi/strfmt"
	"github.com/semi-technologies/weaviate/entities/models"
	"github.com/semi-technologies/weaviate/entities/schema"
	"github.com/semi-technologies/weaviate/entities/schema/crossref"
	"github.com/semi-technologies/weaviate/entities/search"
	"github.com/semi-technologies/weaviate/usecases/traverser"
)

type MergeDocument struct {
	Class                string
	ID                   strfmt.UUID
	PrimitiveSchema      map[string]interface{}
	References           BatchReferences
	Vector               []float32
	UpdateTime           int64
	AdditionalProperties models.AdditionalProperties
}

func (m *Manager) MergeObject(ctx context.Context, principal *models.Principal,
	id strfmt.UUID, updated *models.Object) error {
	err := m.authorizer.Authorize(principal, "update", fmt.Sprintf("objects/%s", id.String()))
	if err != nil {
		return err
	}

	previous, err := m.retrievePreviousAndValidateMergeObject(ctx, principal, id, updated)
	if err != nil {
		return NewErrInvalidUserInput("invalid merge: %v", err)
	}
	primitive, refs := m.splitPrimitiveAndRefs(updated.Properties.(map[string]interface{}),
		updated.Class, id)

	objWithVec, err := m.mergeObjectSchemaAndVectorize(ctx, previous.ClassName, previous.Schema,
		primitive, principal)
	if err != nil {
		return NewErrInternal("vectorize merged: %v", err)
	}
	mergeDoc := MergeDocument{
		Class:           updated.Class,
		ID:              id,
		PrimitiveSchema: primitive,
		References:      refs,
		Vector:          objWithVec.Vector,
		UpdateTime:      m.timeSource.Now(),
	}

	if objWithVec.Additional != nil {
		mergeDoc.AdditionalProperties = objWithVec.Additional
	}

	err = m.vectorRepo.Merge(ctx, mergeDoc)
	if err != nil {
		return NewErrInternal("repo: %v", err)
	}

	return nil
}

func (m *Manager) retrievePreviousAndValidateMergeObject(ctx context.Context, principal *models.Principal,
	id strfmt.UUID, updated *models.Object) (*search.Result, error) {
	if updated.Class == "" {
		return nil, fmt.Errorf("class is a required (and immutable) field")
	}

	object, err := m.vectorRepo.ObjectByID(ctx, id, nil, traverser.AdditionalProperties{})
	if err != nil {
		return nil, err
	}

	if object == nil {
		return nil, fmt.Errorf("object with id '%s' does not exist", id)
	}

	if object.ClassName != updated.Class {
		return nil, fmt.Errorf("class is immutable, but got '%s' for previous class '%s'",
			updated.Class, object.ClassName)
	}

	updated.ID = id
	err = m.validateObject(ctx, principal, updated)
	if err != nil {
		return nil, err
	}

	return object, nil
}

func (m *Manager) mergeObjectSchemaAndVectorize(ctx context.Context, className string,
	old interface{}, new map[string]interface{},
	principal *models.Principal) (*models.Object, error) {
	var merged map[string]interface{}
	if old == nil {
		merged = new
	} else {
		oldMap, ok := old.(map[string]interface{})
		if !ok {
			return nil, fmt.Errorf("expected previous schema to be map, but got %#v", old)
		}

		for key, value := range new {
			oldMap[key] = value
		}

		merged = oldMap
	}

	obj := &models.Object{Class: className, Properties: merged}
	if err := m.obtainVector(ctx, obj, principal); err != nil {
		return nil, err
	}

	return obj, nil
}

func (m *Manager) splitPrimitiveAndRefs(in map[string]interface{}, sourceClass string,
	sourceID strfmt.UUID) (map[string]interface{}, BatchReferences) {
	primitive := map[string]interface{}{}
	var outRefs BatchReferences

	for prop, value := range in {
		refs, ok := value.(models.MultipleRef)

		if !ok {
			// this must be a primitive filed
			primitive[prop] = value
			continue
		}

		for _, ref := range refs {
			target, _ := crossref.Parse(ref.Beacon.String())
			// safe to ignore error as validation has already been passed

			source := &crossref.RefSource{
				Local:    true,
				PeerName: "localhost",
				Property: schema.PropertyName(prop),
				Class:    schema.ClassName(sourceClass),
				TargetID: sourceID,
			}

			outRefs = append(outRefs, BatchReference{From: source, To: target})
		}
	}

	return primitive, outRefs
}
