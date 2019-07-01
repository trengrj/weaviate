/*                          _       _
 *__      _____  __ ___   ___  __ _| |_ ___
 *\ \ /\ / / _ \/ _` \ \ / / |/ _` | __/ _ \
 * \ V  V /  __/ (_| |\ V /| | (_| | ||  __/
 *  \_/\_/ \___|\__,_| \_/ |_|\__,_|\__\___|
 *
 * Copyright © 2016 - 2019 Weaviate. All rights reserved.
 * LICENSE WEAVIATE OPEN SOURCE: https://www.semi.technology/playbook/playbook/contract-weaviate-OSS.html
 * LICENSE WEAVIATE ENTERPRISE: https://www.semi.technology/playbook/contract-weaviate-enterprise.html
 * CONCEPT: Bob van Luijt (@bobvanluijt)
 * CONTACT: hello@semi.technology
 */ // Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// C11yWordsResponse An array of available words and contexts.
// swagger:model C11yWordsResponse
type C11yWordsResponse struct {

	// concatenated word
	ConcatenatedWord *C11yWordsResponseConcatenatedWord `json:"concatenatedWord,omitempty"`

	// Weighted results for per individual word
	IndividualWords []*C11yWordsResponseIndividualWordsItems0 `json:"individualWords"`
}

// Validate validates this c11y words response
func (m *C11yWordsResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateConcatenatedWord(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIndividualWords(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *C11yWordsResponse) validateConcatenatedWord(formats strfmt.Registry) error {

	if swag.IsZero(m.ConcatenatedWord) { // not required
		return nil
	}

	if m.ConcatenatedWord != nil {
		if err := m.ConcatenatedWord.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("concatenatedWord")
			}
			return err
		}
	}

	return nil
}

func (m *C11yWordsResponse) validateIndividualWords(formats strfmt.Registry) error {

	if swag.IsZero(m.IndividualWords) { // not required
		return nil
	}

	for i := 0; i < len(m.IndividualWords); i++ {
		if swag.IsZero(m.IndividualWords[i]) { // not required
			continue
		}

		if m.IndividualWords[i] != nil {
			if err := m.IndividualWords[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("individualWords" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *C11yWordsResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *C11yWordsResponse) UnmarshalBinary(b []byte) error {
	var res C11yWordsResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// C11yWordsResponseConcatenatedWord Weighted results for all words
// swagger:model C11yWordsResponseConcatenatedWord
type C11yWordsResponseConcatenatedWord struct {

	// concatenated nearest neighbors
	ConcatenatedNearestNeighbors C11yNearestNeighbors `json:"concatenatedNearestNeighbors,omitempty"`

	// concatenated vector
	ConcatenatedVector C11yVector `json:"concatenatedVector,omitempty"`

	// concatenated word
	ConcatenatedWord string `json:"concatenatedWord,omitempty"`

	// single words
	SingleWords []string `json:"singleWords"`
}

// Validate validates this c11y words response concatenated word
func (m *C11yWordsResponseConcatenatedWord) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateConcatenatedNearestNeighbors(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateConcatenatedVector(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *C11yWordsResponseConcatenatedWord) validateConcatenatedNearestNeighbors(formats strfmt.Registry) error {

	if swag.IsZero(m.ConcatenatedNearestNeighbors) { // not required
		return nil
	}

	if err := m.ConcatenatedNearestNeighbors.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("concatenatedWord" + "." + "concatenatedNearestNeighbors")
		}
		return err
	}

	return nil
}

func (m *C11yWordsResponseConcatenatedWord) validateConcatenatedVector(formats strfmt.Registry) error {

	if swag.IsZero(m.ConcatenatedVector) { // not required
		return nil
	}

	if err := m.ConcatenatedVector.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("concatenatedWord" + "." + "concatenatedVector")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *C11yWordsResponseConcatenatedWord) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *C11yWordsResponseConcatenatedWord) UnmarshalBinary(b []byte) error {
	var res C11yWordsResponseConcatenatedWord
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// C11yWordsResponseIndividualWordsItems0 c11y words response individual words items0
// swagger:model C11yWordsResponseIndividualWordsItems0
type C11yWordsResponseIndividualWordsItems0 struct {

	// in c11y
	InC11y bool `json:"inC11y,omitempty"`

	// info
	Info *C11yWordsResponseIndividualWordsItems0Info `json:"info,omitempty"`

	// word
	Word string `json:"word,omitempty"`
}

// Validate validates this c11y words response individual words items0
func (m *C11yWordsResponseIndividualWordsItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateInfo(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *C11yWordsResponseIndividualWordsItems0) validateInfo(formats strfmt.Registry) error {

	if swag.IsZero(m.Info) { // not required
		return nil
	}

	if m.Info != nil {
		if err := m.Info.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("info")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *C11yWordsResponseIndividualWordsItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *C11yWordsResponseIndividualWordsItems0) UnmarshalBinary(b []byte) error {
	var res C11yWordsResponseIndividualWordsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// C11yWordsResponseIndividualWordsItems0Info c11y words response individual words items0 info
// swagger:model C11yWordsResponseIndividualWordsItems0Info
type C11yWordsResponseIndividualWordsItems0Info struct {

	// nearest neighbors
	NearestNeighbors C11yNearestNeighbors `json:"nearestNeighbors,omitempty"`

	// vector
	Vector C11yVector `json:"vector,omitempty"`
}

// Validate validates this c11y words response individual words items0 info
func (m *C11yWordsResponseIndividualWordsItems0Info) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateNearestNeighbors(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVector(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *C11yWordsResponseIndividualWordsItems0Info) validateNearestNeighbors(formats strfmt.Registry) error {

	if swag.IsZero(m.NearestNeighbors) { // not required
		return nil
	}

	if err := m.NearestNeighbors.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("info" + "." + "nearestNeighbors")
		}
		return err
	}

	return nil
}

func (m *C11yWordsResponseIndividualWordsItems0Info) validateVector(formats strfmt.Registry) error {

	if swag.IsZero(m.Vector) { // not required
		return nil
	}

	if err := m.Vector.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("info" + "." + "vector")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *C11yWordsResponseIndividualWordsItems0Info) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *C11yWordsResponseIndividualWordsItems0Info) UnmarshalBinary(b []byte) error {
	var res C11yWordsResponseIndividualWordsItems0Info
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
