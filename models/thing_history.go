/*                          _       _
 *__      _____  __ ___   ___  __ _| |_ ___
 *\ \ /\ / / _ \/ _` \ \ / / |/ _` | __/ _ \
 * \ V  V /  __/ (_| |\ V /| | (_| | ||  __/
 *  \_/\_/ \___|\__,_| \_/ |_|\__,_|\__\___|
 *
 * Copyright © 2016 - 2018 Weaviate. All rights reserved.
 * LICENSE: https://github.com/creativesoftwarefdn/weaviate/blob/develop/LICENSE.md
 * AUTHOR: Bob van Luijt (bob@kub.design)
 * See www.creativesoftwarefdn.org for details
 * Contact: @CreativeSofwFdn / bob@kub.design
 */

package models

import (
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// ThingHistory thing history
// swagger:model ThingHistory

type ThingHistory struct {

	// Indication whether the action is deleted
	Deleted bool `json:"deleted,omitempty"`

	// key
	Key *SingleRef `json:"key,omitempty"`

	// An array with the history of the things.
	PropertyHistory []*ThingHistoryObject `json:"propertyHistory"`
}

/* polymorph ThingHistory deleted false */

/* polymorph ThingHistory key false */

/* polymorph ThingHistory propertyHistory false */

// Validate validates this thing history
func (m *ThingHistory) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateKey(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validatePropertyHistory(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ThingHistory) validateKey(formats strfmt.Registry) error {

	if swag.IsZero(m.Key) { // not required
		return nil
	}

	if m.Key != nil {

		if err := m.Key.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("key")
			}
			return err
		}
	}

	return nil
}

func (m *ThingHistory) validatePropertyHistory(formats strfmt.Registry) error {

	if swag.IsZero(m.PropertyHistory) { // not required
		return nil
	}

	for i := 0; i < len(m.PropertyHistory); i++ {

		if swag.IsZero(m.PropertyHistory[i]) { // not required
			continue
		}

		if m.PropertyHistory[i] != nil {

			if err := m.PropertyHistory[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("propertyHistory" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *ThingHistory) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ThingHistory) UnmarshalBinary(b []byte) error {
	var res ThingHistory
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
