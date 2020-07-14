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
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// SemanticPathElement On link on the semantic path chain
//
// swagger:model SemanticPathElement
type SemanticPathElement struct {

	// concept
	Concept string `json:"concept,omitempty"`

	// distance to next
	DistanceToNext *float32 `json:"distanceToNext,omitempty"`

	// distance to previous
	DistanceToPrevious *float32 `json:"distanceToPrevious,omitempty"`

	// distance to query
	DistanceToQuery float32 `json:"distanceToQuery,omitempty"`

	// distance to result
	DistanceToResult float32 `json:"distanceToResult,omitempty"`
}

// Validate validates this semantic path element
func (m *SemanticPathElement) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *SemanticPathElement) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SemanticPathElement) UnmarshalBinary(b []byte) error {
	var res SemanticPathElement
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
