/*                          _       _
 *__      _____  __ ___   ___  __ _| |_ ___
 *\ \ /\ / / _ \/ _` \ \ / / |/ _` | __/ _ \
 * \ V  V /  __/ (_| |\ V /| | (_| | ||  __/
 *  \_/\_/ \___|\__,_| \_/ |_|\__,_|\__\___|
 *
 * Copyright © 2016 Weaviate. All rights reserved.
 * LICENSE: https://github.com/creativesoftwarefdn/weaviate/blob/develop/LICENSE.md
 * AUTHOR: Bob van Luijt (bob@weaviate.com)
 * See www.weaviate.com for details
 * Contact: @CreativeSofwFdn / yourfriends@weaviate.com
 */

package models

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// Meta Contains meta information of the current Weaviate instance.
// swagger:model Meta

type Meta struct {

	// actions schema
	ActionsSchema *SemanticSchema `json:"actionsSchema,omitempty"`

	// The url of the host
	Hostname string `json:"hostname,omitempty"`

	// things schema
	ThingsSchema *SemanticSchema `json:"thingsSchema,omitempty"`
}

/* polymorph Meta actionsSchema false */

/* polymorph Meta hostname false */

/* polymorph Meta thingsSchema false */

// Validate validates this meta
func (m *Meta) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateActionsSchema(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateThingsSchema(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Meta) validateActionsSchema(formats strfmt.Registry) error {

	if swag.IsZero(m.ActionsSchema) { // not required
		return nil
	}

	if m.ActionsSchema != nil {

		if err := m.ActionsSchema.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("actionsSchema")
			}
			return err
		}
	}

	return nil
}

func (m *Meta) validateThingsSchema(formats strfmt.Registry) error {

	if swag.IsZero(m.ThingsSchema) { // not required
		return nil
	}

	if m.ThingsSchema != nil {

		if err := m.ThingsSchema.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("thingsSchema")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Meta) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Meta) UnmarshalBinary(b []byte) error {
	var res Meta
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
