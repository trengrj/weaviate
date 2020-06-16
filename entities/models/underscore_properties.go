//                           _       _
// __      _____  __ ___   ___  __ _| |_ ___
// \ \ /\ / / _ \/ _` \ \ / / |/ _` | __/ _ \
//  \ V  V /  __/ (_| |\ V /| | (_| | ||  __/
//   \_/\_/ \___|\__,_| \_/ |_|\__,_|\__\___|
//
//  Copyright © 2016 - 2020 SeMI Holding B.V. (registered @ Dutch Chamber of Commerce no 75221632). All rights reserved.
//  LICENSE WEAVIATE OPEN SOURCE: https://www.semi.technology/playbook/playbook/contract-weaviate-OSS.html
//  LICENSE WEAVIATE ENTERPRISE: https://www.semi.technology/playbook/contract-weaviate-enterprise.html
//  CONCEPT: Bob van Luijt (@bobvanluijt)
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

// UnderscoreProperties Additional Meta information about a single thing/action object.
//
// swagger:model UnderscoreProperties
type UnderscoreProperties struct {

	// If this object was subject of a classificiation, additional meta info about this classification is available here
	Classification *UnderscorePropertiesClassification `json:"classification,omitempty"`

	// This object's position in the Contextionary vector space
	Vector C11yVector `json:"vector,omitempty"`

	// Additional information about how the object was vectorized
	VectorizationMeta *VectorizationMeta `json:"vectorizationMeta,omitempty"`
}

// Validate validates this underscore properties
func (m *UnderscoreProperties) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateClassification(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVector(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVectorizationMeta(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *UnderscoreProperties) validateClassification(formats strfmt.Registry) error {

	if swag.IsZero(m.Classification) { // not required
		return nil
	}

	if m.Classification != nil {
		if err := m.Classification.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("classification")
			}
			return err
		}
	}

	return nil
}

func (m *UnderscoreProperties) validateVector(formats strfmt.Registry) error {

	if swag.IsZero(m.Vector) { // not required
		return nil
	}

	if err := m.Vector.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("vector")
		}
		return err
	}

	return nil
}

func (m *UnderscoreProperties) validateVectorizationMeta(formats strfmt.Registry) error {

	if swag.IsZero(m.VectorizationMeta) { // not required
		return nil
	}

	if m.VectorizationMeta != nil {
		if err := m.VectorizationMeta.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("vectorizationMeta")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *UnderscoreProperties) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UnderscoreProperties) UnmarshalBinary(b []byte) error {
	var res UnderscoreProperties
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
