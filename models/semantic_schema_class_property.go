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

// SemanticSchemaClassProperty semantic schema class property
// swagger:model SemanticSchemaClassProperty

type SemanticSchemaClassProperty struct {

	// Can be a reference ($cref) to another type when starts with a capital (for example Person) otherwise "string" or "int".
	AtDataType []string `json:"@dataType"`

	// Description of the property
	Description string `json:"description,omitempty"`

	// Describes the kind of class. For example Geolocation for the class City.
	Keywords []*SemanticSchemaClassPropertyKeywordsItems0 `json:"keywords"`

	// Name of the property as URI relative to the schema URL.
	Name string `json:"name,omitempty"`
}

/* polymorph SemanticSchemaClassProperty @dataType false */

/* polymorph SemanticSchemaClassProperty description false */

/* polymorph SemanticSchemaClassProperty keywords false */

/* polymorph SemanticSchemaClassProperty name false */

// Validate validates this semantic schema class property
func (m *SemanticSchemaClassProperty) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAtDataType(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateKeywords(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SemanticSchemaClassProperty) validateAtDataType(formats strfmt.Registry) error {

	if swag.IsZero(m.AtDataType) { // not required
		return nil
	}

	return nil
}

func (m *SemanticSchemaClassProperty) validateKeywords(formats strfmt.Registry) error {

	if swag.IsZero(m.Keywords) { // not required
		return nil
	}

	for i := 0; i < len(m.Keywords); i++ {

		if swag.IsZero(m.Keywords[i]) { // not required
			continue
		}

		if m.Keywords[i] != nil {

			if err := m.Keywords[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("keywords" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *SemanticSchemaClassProperty) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SemanticSchemaClassProperty) UnmarshalBinary(b []byte) error {
	var res SemanticSchemaClassProperty
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// SemanticSchemaClassPropertyKeywordsItems0 semantic schema class property keywords items0
// swagger:model SemanticSchemaClassPropertyKeywordsItems0

type SemanticSchemaClassPropertyKeywordsItems0 struct {

	// kind
	Kind string `json:"kind,omitempty"`

	// weight
	Weight float32 `json:"weight,omitempty"`
}

/* polymorph SemanticSchemaClassPropertyKeywordsItems0 kind false */

/* polymorph SemanticSchemaClassPropertyKeywordsItems0 weight false */

// Validate validates this semantic schema class property keywords items0
func (m *SemanticSchemaClassPropertyKeywordsItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *SemanticSchemaClassPropertyKeywordsItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SemanticSchemaClassPropertyKeywordsItems0) UnmarshalBinary(b []byte) error {
	var res SemanticSchemaClassPropertyKeywordsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
