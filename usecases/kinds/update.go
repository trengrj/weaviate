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

package kinds

import (
	"context"
	"fmt"

	"github.com/go-openapi/strfmt"
	"github.com/semi-technologies/weaviate/entities/models"
	"github.com/semi-technologies/weaviate/entities/schema/kind"
	"github.com/semi-technologies/weaviate/usecases/traverser"
)

// UpdateAction Class Instance to the connected DB. If the class contains a network
// ref, it has a side-effect on the schema: The schema will be updated to
// include this particular network ref class.
func (m *Manager) UpdateAction(ctx context.Context, principal *models.Principal, id strfmt.UUID,
	class *models.Action) (*models.Action, error) {
	err := m.authorizer.Authorize(principal, "update", fmt.Sprintf("actions/%s", id.String()))
	if err != nil {
		return nil, err
	}

	unlock, err := m.locks.LockSchema()
	if err != nil {
		return nil, NewErrInternal("could not acquire lock: %v", err)
	}
	defer unlock()

	return m.updateActionToConnectorAndSchema(ctx, principal, id, class)
}

func (m *Manager) updateActionToConnectorAndSchema(ctx context.Context, principal *models.Principal,
	id strfmt.UUID, class *models.Action) (*models.Action, error) {
	if id != class.ID {
		return nil, NewErrInvalidUserInput("invalid update: field 'id' is immutable")
	}

	originalAction, err := m.getActionFromRepo(ctx, id, traverser.UnderscoreProperties{})
	if err != nil {
		return nil, err
	}

	m.logger.
		WithField("action", "kinds_update_requested").
		WithField("kind", kind.Action).
		WithField("original", originalAction).
		WithField("updated", class).
		WithField("id", id).
		Debug("received update kind request")

	err = m.validateAction(ctx, principal, class)
	if err != nil {
		return nil, NewErrInvalidUserInput("invalid action: %v", err)
	}

	class.LastUpdateTimeUnix = m.timeSource.Now()

	err = m.vectorizeAndPutAction(ctx, class)
	if err != nil {
		return nil, NewErrInternal("update action: %v", err)
	}

	return class, nil
}

// UpdateThing Class Instance to the connected DB. If the class contains a network
// ref, it has a side-effect on the schema: The schema will be updated to
// include this particular network ref class.
func (m *Manager) UpdateThing(ctx context.Context, principal *models.Principal,
	id strfmt.UUID, class *models.Thing) (*models.Thing, error) {
	err := m.authorizer.Authorize(principal, "update", fmt.Sprintf("things/%s", id.String()))
	if err != nil {
		return nil, err
	}

	unlock, err := m.locks.LockSchema()
	if err != nil {
		return nil, NewErrInternal("could not acquire lock: %v", err)
	}
	defer unlock()

	return m.updateThingToConnectorAndSchema(ctx, principal, id, class)
}

func (m *Manager) updateThingToConnectorAndSchema(ctx context.Context, principal *models.Principal,
	id strfmt.UUID, class *models.Thing) (*models.Thing, error) {
	if id != class.ID {
		return nil, NewErrInvalidUserInput("invalid update: field 'id' is immutable")
	}

	originalThing, err := m.getThingFromRepo(ctx, id, traverser.UnderscoreProperties{})
	if err != nil {
		return nil, err
	}

	m.logger.
		WithField("action", "kinds_update_requested").
		WithField("kind", kind.Thing).
		WithField("original", originalThing).
		WithField("updated", class).
		WithField("id", id).
		Debug("received update kind request")

	err = m.validateThing(ctx, principal, class)
	if err != nil {
		return nil, NewErrInvalidUserInput("invalid thing: %v", err)
	}

	class.LastUpdateTimeUnix = m.timeSource.Now()

	err = m.vectorizeAndPutThing(ctx, class)
	if err != nil {
		return nil, NewErrInternal("update thing: %v", err)
	}

	return class, nil
}
