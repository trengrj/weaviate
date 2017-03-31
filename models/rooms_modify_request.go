package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/validate"
)

// RoomsModifyRequest rooms modify request
// swagger:model RoomsModifyRequest
type RoomsModifyRequest struct {

	// name
	Name string `json:"name,omitempty"`

	// type
	Type string `json:"type,omitempty"`
}

// Validate validates this rooms modify request
func (m *RoomsModifyRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateType(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var roomsModifyRequestTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["attic","backyard","basement","bathroom","bedroom","default","den","diningRoom","entryway","familyRoom","frontyard","garage","hallway","kitchen","livingRoom","masterBedroom","office","other","shed","unknownRoomType"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		roomsModifyRequestTypeTypePropEnum = append(roomsModifyRequestTypeTypePropEnum, v)
	}
}

const (
	roomsModifyRequestTypeAttic           string = "attic"
	roomsModifyRequestTypeBackyard        string = "backyard"
	roomsModifyRequestTypeBasement        string = "basement"
	roomsModifyRequestTypeBathroom        string = "bathroom"
	roomsModifyRequestTypeBedroom         string = "bedroom"
	roomsModifyRequestTypeDefault         string = "default"
	roomsModifyRequestTypeDen             string = "den"
	roomsModifyRequestTypeDiningRoom      string = "diningRoom"
	roomsModifyRequestTypeEntryway        string = "entryway"
	roomsModifyRequestTypeFamilyRoom      string = "familyRoom"
	roomsModifyRequestTypeFrontyard       string = "frontyard"
	roomsModifyRequestTypeGarage          string = "garage"
	roomsModifyRequestTypeHallway         string = "hallway"
	roomsModifyRequestTypeKitchen         string = "kitchen"
	roomsModifyRequestTypeLivingRoom      string = "livingRoom"
	roomsModifyRequestTypeMasterBedroom   string = "masterBedroom"
	roomsModifyRequestTypeOffice          string = "office"
	roomsModifyRequestTypeOther           string = "other"
	roomsModifyRequestTypeShed            string = "shed"
	roomsModifyRequestTypeUnknownRoomType string = "unknownRoomType"
)

// prop value enum
func (m *RoomsModifyRequest) validateTypeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, roomsModifyRequestTypeTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *RoomsModifyRequest) validateType(formats strfmt.Registry) error {

	if swag.IsZero(m.Type) { // not required
		return nil
	}

	// value enum
	if err := m.validateTypeEnum("type", "body", m.Type); err != nil {
		return err
	}

	return nil
}