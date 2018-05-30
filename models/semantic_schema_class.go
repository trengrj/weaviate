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

// SemanticSchemaClass semantic schema class
// swagger:model SemanticSchemaClass

type SemanticSchemaClass struct {

	// Name of the class as URI relative to the schema URL.
	Class string `json:"class,omitempty"`

	// Description of the class
	Description string `json:"description,omitempty"`

	// Describes the kind of class. For example Geolocation for the class City.
	Keywords []*SemanticSchemaClassKeywordsItems0 `json:"keywords"`

	// The properties of the class.
	Properties []*SemanticSchemaClassProperty `json:"properties"`
}

/* polymorph SemanticSchemaClass class false */

/* polymorph SemanticSchemaClass description false */

/* polymorph SemanticSchemaClass keywords false */

/* polymorph SemanticSchemaClass properties false */

// Validate validates this semantic schema class
func (m *SemanticSchemaClass) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateKeywords(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateProperties(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SemanticSchemaClass) validateKeywords(formats strfmt.Registry) error {

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

func (m *SemanticSchemaClass) validateProperties(formats strfmt.Registry) error {

	if swag.IsZero(m.Properties) { // not required
		return nil
	}

	for i := 0; i < len(m.Properties); i++ {

		if swag.IsZero(m.Properties[i]) { // not required
			continue
		}

		if m.Properties[i] != nil {

			if err := m.Properties[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("properties" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *SemanticSchemaClass) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SemanticSchemaClass) UnmarshalBinary(b []byte) error {
	var res SemanticSchemaClass
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// SemanticSchemaClassKeywordsItems0 semantic schema class keywords items0
// swagger:model SemanticSchemaClassKeywordsItems0

type SemanticSchemaClassKeywordsItems0 struct {

	// kind
	Kind string `json:"kind,omitempty"`

	// weight
	Weight float32 `json:"weight,omitempty"`
}

/* polymorph SemanticSchemaClassKeywordsItems0 kind false */

/* polymorph SemanticSchemaClassKeywordsItems0 weight false */

// Validate validates this semantic schema class keywords items0
func (m *SemanticSchemaClassKeywordsItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *SemanticSchemaClassKeywordsItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SemanticSchemaClassKeywordsItems0) UnmarshalBinary(b []byte) error {
	var res SemanticSchemaClassKeywordsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
