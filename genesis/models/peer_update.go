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

// PeerUpdate peer update
// swagger:model PeerUpdate
type PeerUpdate struct {

	// Name of the peer, must be valid DNS name
	PeerName string `json:"peerName,omitempty"`

	// Host or IP of the peer, defaults to peerName
	// Format: uri
	PeerURI strfmt.URI `json:"peerUri,omitempty"`
}

// Validate validates this peer update
func (m *PeerUpdate) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePeerURI(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PeerUpdate) validatePeerURI(formats strfmt.Registry) error {

	if swag.IsZero(m.PeerURI) { // not required
		return nil
	}

	if err := validate.FormatOf("peerUri", "body", "uri", m.PeerURI.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PeerUpdate) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PeerUpdate) UnmarshalBinary(b []byte) error {
	var res PeerUpdate
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
