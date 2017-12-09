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
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// ThingsListResponse List of things.
// swagger:model ThingsListResponse

type ThingsListResponse struct {

	// The actual list of things.
	Things []*ThingGetResponse `json:"things"`

	// The total number of things for the query. The number of items in a response may be smaller due to paging.
	TotalResults int64 `json:"totalResults,omitempty"`
}

/* polymorph ThingsListResponse things false */

/* polymorph ThingsListResponse totalResults false */

// Validate validates this things list response
func (m *ThingsListResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateThings(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ThingsListResponse) validateThings(formats strfmt.Registry) error {

	if swag.IsZero(m.Things) { // not required
		return nil
	}

	for i := 0; i < len(m.Things); i++ {

		if swag.IsZero(m.Things[i]) { // not required
			continue
		}

		if m.Things[i] != nil {

			if err := m.Things[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("things" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *ThingsListResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ThingsListResponse) UnmarshalBinary(b []byte) error {
	var res ThingsListResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
