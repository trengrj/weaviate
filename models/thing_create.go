/*                          _       _
 *__      _____  __ ___   ___  __ _| |_ ___
 *\ \ /\ / / _ \/ _` \ \ / / |/ _` | __/ _ \
 * \ V  V /  __/ (_| |\ V /| | (_| | ||  __/
 *  \_/\_/ \___|\__,_| \_/ |_|\__,_|\__\___|
 *
 * Copyright © 2016 Weaviate. All rights reserved.
 * LICENSE: https://github.com/weaviate/weaviate/blob/master/LICENSE
 * AUTHOR: Bob van Luijt (bob@weaviate.com)
 * See www.weaviate.com for details
 * Contact: @weaviate_iot / yourfriends@weaviate.com
 */
   

package models

 
 

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// ThingCreate thing create
// swagger:model ThingCreate
type ThingCreate struct {

	// commands Id
	CommandsID strfmt.UUID `json:"commandsId,omitempty"`

	// description
	Description string `json:"description,omitempty"`

	// groups
	Groups string `json:"groups,omitempty"`

	// location Id
	LocationID strfmt.UUID `json:"locationId,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// owner
	Owner string `json:"owner,omitempty"`

	// serial number
	SerialNumber string `json:"serialNumber,omitempty"`

	// tags
	Tags []interface{} `json:"tags"`

	// thing template Id
	ThingTemplateID strfmt.UUID `json:"thingTemplateId,omitempty"`
}

// Validate validates this thing create
func (m *ThingCreate) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *ThingCreate) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ThingCreate) UnmarshalBinary(b []byte) error {
	var res ThingCreate
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
