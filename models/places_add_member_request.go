package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
)

// PlacesAddMemberRequest places add member request
// swagger:model PlacesAddMemberRequest
type PlacesAddMemberRequest struct {

	// Email of the new member of the place.
	MemberID string `json:"memberId,omitempty"`
}

// Validate validates this places add member request
func (m *PlacesAddMemberRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}