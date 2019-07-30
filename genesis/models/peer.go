//                           _       _
// __      _____  __ ___   ___  __ _| |_ ___
// \ \ /\ / / _ \/ _` \ \ / / |/ _` | __/ _ \
//  \ V  V /  __/ (_| |\ V /| | (_| | ||  __/
//   \_/\_/ \___|\__,_| \_/ |_|\__,_|\__\___|
// 
//  Copyright © 2016 - 2019 Weaviate. All rights reserved.
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
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Peer peer
// swagger:model Peer
type Peer struct {
	PeerUpdate

	// Unique ID of this peer registration, will be updated if the peer conntects again to the network.
	// Format: uuid
	ID strfmt.UUID `json:"id,omitempty"`

	// When we were received a ping from this peer from the last time
	LastContactAt int64 `json:"last_contact_at,omitempty"`

	// The latest known hash of the local schema of the peer
	SchemaHash string `json:"schema_hash,omitempty"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *Peer) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 PeerUpdate
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.PeerUpdate = aO0

	// AO1
	var dataAO1 struct {
		ID strfmt.UUID `json:"id,omitempty"`

		LastContactAt int64 `json:"last_contact_at,omitempty"`

		SchemaHash string `json:"schema_hash,omitempty"`
	}
	if err := swag.ReadJSON(raw, &dataAO1); err != nil {
		return err
	}

	m.ID = dataAO1.ID

	m.LastContactAt = dataAO1.LastContactAt

	m.SchemaHash = dataAO1.SchemaHash

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m Peer) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	aO0, err := swag.WriteJSON(m.PeerUpdate)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	var dataAO1 struct {
		ID strfmt.UUID `json:"id,omitempty"`

		LastContactAt int64 `json:"last_contact_at,omitempty"`

		SchemaHash string `json:"schema_hash,omitempty"`
	}

	dataAO1.ID = m.ID

	dataAO1.LastContactAt = m.LastContactAt

	dataAO1.SchemaHash = m.SchemaHash

	jsonDataAO1, errAO1 := swag.WriteJSON(dataAO1)
	if errAO1 != nil {
		return nil, errAO1
	}
	_parts = append(_parts, jsonDataAO1)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this peer
func (m *Peer) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with PeerUpdate
	if err := m.PeerUpdate.Validate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Peer) validateID(formats strfmt.Registry) error {

	if swag.IsZero(m.ID) { // not required
		return nil
	}

	if err := validate.FormatOf("id", "body", "uuid", m.ID.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Peer) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Peer) UnmarshalBinary(b []byte) error {
	var res Peer
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
