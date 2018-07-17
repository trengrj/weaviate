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
// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// SingleRef single ref
// swagger:model SingleRef
type SingleRef struct {

	// Location of the cross reference.
	// Format: uuid
	NrDollarCref strfmt.UUID `json:"$cref,omitempty"`

	// url of location. http://localhost means this database. This option can be used to refer to other databases.
	LocationURL *string `json:"locationUrl,omitempty"`

	// Type should be Thing, Action or Key
	// Enum: [Thing Action Key]
	Type string `json:"type,omitempty"`
}

// Validate validates this single ref
func (m *SingleRef) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateNrDollarCref(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SingleRef) validateNrDollarCref(formats strfmt.Registry) error {

	if swag.IsZero(m.NrDollarCref) { // not required
		return nil
	}

	if err := validate.FormatOf("$cref", "body", "uuid", m.NrDollarCref.String(), formats); err != nil {
		return err
	}

	return nil
}

var singleRefTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Thing","Action","Key"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		singleRefTypeTypePropEnum = append(singleRefTypeTypePropEnum, v)
	}
}

const (

	// SingleRefTypeThing captures enum value "Thing"
	SingleRefTypeThing string = "Thing"

	// SingleRefTypeAction captures enum value "Action"
	SingleRefTypeAction string = "Action"

	// SingleRefTypeKey captures enum value "Key"
	SingleRefTypeKey string = "Key"
)

// prop value enum
func (m *SingleRef) validateTypeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, singleRefTypeTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *SingleRef) validateType(formats strfmt.Registry) error {

	if swag.IsZero(m.Type) { // not required
		return nil
	}

	// value enum
	if err := m.validateTypeEnum("type", "body", m.Type); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SingleRef) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SingleRef) UnmarshalBinary(b []byte) error {
	var res SingleRef
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
