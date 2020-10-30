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

// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// Property property
//
// swagger:model Property
type Property struct {

	// DEPRECATED - do not use anymore.
	Cardinality string `json:"cardinality,omitempty"`

	// Can be a reference to another type when it starts with a capital (for example Person), otherwise "string" or "int".
	DataType []string `json:"dataType"`

	// Description of the property.
	Description string `json:"description,omitempty"`

	// Optional. By default each property is fully indexed both for full-text, as well as vector-search. You can ignore properties in searches by explicitly setting index to false. Not set is the same as true
	Index *bool `json:"index,omitempty"`

	// keywords
	Keywords Keywords `json:"keywords,omitempty"`

	// Name of the property as URI relative to the schema URL.
	Name string `json:"name,omitempty"`

	// Set this to true if the object vector should include this property's name in calculating the overall vector position. If set to false (default), only the property value will be used.
	VectorizePropertyName bool `json:"vectorizePropertyName,omitempty"`
}

// Validate validates this property
func (m *Property) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateKeywords(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Property) validateKeywords(formats strfmt.Registry) error {
	if swag.IsZero(m.Keywords) { // not required
		return nil
	}

	if err := m.Keywords.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("keywords")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Property) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Property) UnmarshalBinary(b []byte) error {
	var res Property
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
