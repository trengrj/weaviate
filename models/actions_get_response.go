/*                          _       _
 *__      _____  __ ___   ___  __ _| |_ ___
 *\ \ /\ / / _ \/ _` \ \ / / |/ _` | __/ _ \
 * \ V  V /  __/ (_| |\ V /| | (_| | ||  __/
 *  \_/\_/ \___|\__,_| \_/ |_|\__,_|\__\___|
 *
 * Copyright © 2016 - 2019 Weaviate. All rights reserved.
 * LICENSE: https://github.com/creativesoftwarefdn/weaviate/blob/develop/LICENSE.md
 * DESIGN & CONCEPT: Bob van Luijt (@bobvanluijt)
 * CONTACT: hello@creativesoftwarefdn.org
 */ // Code generated by go-swagger; DO NOT EDIT.

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

// ActionsGetResponse actions get response
// swagger:model ActionsGetResponse
type ActionsGetResponse struct {
	Action

	// result
	Result *ActionsGetResponseAO1Result `json:"result,omitempty"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *ActionsGetResponse) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 Action
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.Action = aO0

	// AO1
	var dataAO1 struct {
		Result *ActionsGetResponseAO1Result `json:"result,omitempty"`
	}
	if err := swag.ReadJSON(raw, &dataAO1); err != nil {
		return err
	}

	m.Result = dataAO1.Result

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m ActionsGetResponse) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	aO0, err := swag.WriteJSON(m.Action)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	var dataAO1 struct {
		Result *ActionsGetResponseAO1Result `json:"result,omitempty"`
	}

	dataAO1.Result = m.Result

	jsonDataAO1, errAO1 := swag.WriteJSON(dataAO1)
	if errAO1 != nil {
		return nil, errAO1
	}
	_parts = append(_parts, jsonDataAO1)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this actions get response
func (m *ActionsGetResponse) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with Action
	if err := m.Action.Validate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateResult(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ActionsGetResponse) validateResult(formats strfmt.Registry) error {

	if swag.IsZero(m.Result) { // not required
		return nil
	}

	if m.Result != nil {
		if err := m.Result.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("result")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ActionsGetResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ActionsGetResponse) UnmarshalBinary(b []byte) error {
	var res ActionsGetResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// ActionsGetResponseAO1Result Results for this specific Action.
// swagger:model ActionsGetResponseAO1Result
type ActionsGetResponseAO1Result struct {

	// errors
	Errors *ErrorResponse `json:"errors,omitempty"`

	// status
	// Enum: [SUCCESS PENDING FAILED]
	Status *string `json:"status,omitempty"`
}

// Validate validates this actions get response a o1 result
func (m *ActionsGetResponseAO1Result) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateErrors(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ActionsGetResponseAO1Result) validateErrors(formats strfmt.Registry) error {

	if swag.IsZero(m.Errors) { // not required
		return nil
	}

	if m.Errors != nil {
		if err := m.Errors.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("result" + "." + "errors")
			}
			return err
		}
	}

	return nil
}

var actionsGetResponseAO1ResultTypeStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["SUCCESS","PENDING","FAILED"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		actionsGetResponseAO1ResultTypeStatusPropEnum = append(actionsGetResponseAO1ResultTypeStatusPropEnum, v)
	}
}

const (

	// ActionsGetResponseAO1ResultStatusSUCCESS captures enum value "SUCCESS"
	ActionsGetResponseAO1ResultStatusSUCCESS string = "SUCCESS"

	// ActionsGetResponseAO1ResultStatusPENDING captures enum value "PENDING"
	ActionsGetResponseAO1ResultStatusPENDING string = "PENDING"

	// ActionsGetResponseAO1ResultStatusFAILED captures enum value "FAILED"
	ActionsGetResponseAO1ResultStatusFAILED string = "FAILED"
)

// prop value enum
func (m *ActionsGetResponseAO1Result) validateStatusEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, actionsGetResponseAO1ResultTypeStatusPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *ActionsGetResponseAO1Result) validateStatus(formats strfmt.Registry) error {

	if swag.IsZero(m.Status) { // not required
		return nil
	}

	// value enum
	if err := m.validateStatusEnum("result"+"."+"status", "body", *m.Status); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ActionsGetResponseAO1Result) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ActionsGetResponseAO1Result) UnmarshalBinary(b []byte) error {
	var res ActionsGetResponseAO1Result
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
