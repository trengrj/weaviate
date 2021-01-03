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

package objects

import (
	"context"
	"fmt"

	"github.com/go-openapi/strfmt"
	"github.com/semi-technologies/weaviate/entities/models"
	"github.com/semi-technologies/weaviate/entities/schema/kind"
	"github.com/semi-technologies/weaviate/usecases/traverser"
)

// DeleteObjectReference from connected DB
func (m *Manager) DeleteObjectReference(ctx context.Context, principal *models.Principal,
	id strfmt.UUID, propertyName string, property *models.SingleRef) error {
	err := m.authorizer.Authorize(principal, "update", fmt.Sprintf("objects/%s", id.String()))
	if err != nil {
		return err
	}

	unlock, err := m.locks.LockSchema()
	if err != nil {
		return NewErrInternal("could not acquire lock: %v", err)
	}
	defer unlock()

	return m.deleteObjectReferenceFromConnector(ctx, principal, id, propertyName, property)
}

func (m *Manager) deleteObjectReferenceFromConnector(ctx context.Context, principal *models.Principal,
	id strfmt.UUID, propertyName string, property *models.SingleRef) error {
	// get object to see if it exists
	objectRes, err := m.getObjectFromRepo(ctx, id, traverser.UnderscoreProperties{})
	if err != nil {
		return err
	}

	object := objectRes.Object()
	// NOTE: The reference itself is not being validated, to allow for deletion
	// of broken references
	err = m.validateCanModifyReference(principal, kind.Object, object.Class, propertyName)
	if err != nil {
		return err
	}

	extended, err := m.removeReferenceFromClassProps(object.Schema, propertyName, property)
	if err != nil {
		return err
	}
	object.Schema = extended
	object.LastUpdateTimeUnix = m.timeSource.Now()

	err = m.vectorRepo.PutObject(ctx, object, objectRes.Vector)
	if err != nil {
		return NewErrInternal("could not store object: %v", err)
	}

	return nil
}

// DeleteThingReference from connected DB
// func (m *Manager) DeleteThingReference(ctx context.Context, principal *models.Principal,
// 	id strfmt.UUID, propertyName string, property *models.SingleRef) error {
// 	err := m.authorizer.Authorize(principal, "update", fmt.Sprintf("things/%s", id.String()))
// 	if err != nil {
// 		return err
// 	}

// 	unlock, err := m.locks.LockSchema()
// 	if err != nil {
// 		return NewErrInternal("could not acquire lock: %v", err)
// 	}
// 	defer unlock()

// 	return m.deleteThingReferenceFromConnector(ctx, principal, id, propertyName, property)
// }

// func (m *Manager) deleteThingReferenceFromConnector(ctx context.Context, principal *models.Principal,
// 	id strfmt.UUID, propertyName string, property *models.SingleRef) error {
// 	// get thing to see if it exists
// 	thingRes, err := m.getThingFromRepo(ctx, id, traverser.UnderscoreProperties{})
// 	if err != nil {
// 		return err
// 	}

// 	thing := thingRes.Thing()
// 	// NOTE: The reference itself is not being validated, to allow for deletion
// 	// of broken references
// 	err = m.validateCanModifyReference(principal, kind.Thing, thing.Class, propertyName)
// 	if err != nil {
// 		return err
// 	}

// 	extended, err := m.removeReferenceFromClassProps(thing.Schema, propertyName, property)
// 	if err != nil {
// 		return err
// 	}
// 	thing.Schema = extended
// 	thing.LastUpdateTimeUnix = m.timeSource.Now()

// 	err = m.vectorRepo.PutThing(ctx, thing, thingRes.Vector)
// 	if err != nil {
// 		return NewErrInternal("could not store thing: %v", err)
// 	}

// 	return nil
// }

func (m *Manager) removeReferenceFromClassProps(props interface{}, propertyName string,
	property *models.SingleRef) (interface{}, error) {
	if props == nil {
		props = map[string]interface{}{}
	}

	propsMap := props.(map[string]interface{})

	_, ok := propsMap[propertyName]
	if !ok {
		propsMap[propertyName] = models.MultipleRef{}
	}

	existingRefs := propsMap[propertyName]
	existingMultipleRef, ok := existingRefs.(models.MultipleRef)
	if !ok {
		return nil, NewErrInternal("expected list for reference props, but got %T", existingRefs)
	}

	propsMap[propertyName] = removeRef(existingMultipleRef, property)
	return propsMap, nil
}

func removeRef(refs models.MultipleRef, property *models.SingleRef) models.MultipleRef {
	// Remove if this reference is found.
	for i, currentRef := range refs {
		if currentRef.Beacon != property.Beacon {
			continue
		}

		// remove this one without memory leaks, see
		// https://github.com/golang/go/wiki/SliceTricks#delete
		copy(refs[i:], refs[i+1:])
		refs[len(refs)-1] = nil // or the zero value of T
		refs = refs[:len(refs)-1]
		break // we can only remove one at the same time, so break the loop.
	}

	return refs
}
