/*                          _       _
 *__      _____  __ ___   ___  __ _| |_ ___
 *\ \ /\ / / _ \/ _` \ \ / / |/ _` | __/ _ \
 * \ V  V /  __/ (_| |\ V /| | (_| | ||  __/
 *  \_/\_/ \___|\__,_| \_/ |_|\__,_|\__\___|
 *
 * Copyright © 2016 Weaviate. All rights reserved.
 * LICENSE: https://github.com/creativesoftwarefdn/weaviate/blob/develop/LICENSE.md
 * AUTHOR: Bob van Luijt (bob@weaviate.com)
 * See www.weaviate.com for details
 * Contact: @CreativeSofwFdn / yourfriends@weaviate.com
 */

package models

import (
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// PatchDocument A JSONPatch document as defined by RFC 6902.
// swagger:model PatchDocument

type PatchDocument struct {

	// A string containing a JSON Pointer value.
	From string `json:"from,omitempty"`

	// The operation to be performed.
	// Required: true
	Op *string `json:"op"`

	// A JSON-Pointer.
	// Required: true
	Path *string `json:"path"`

	// The value to be used within the operations.
	Value interface{} `json:"value,omitempty"`
}

/* polymorph PatchDocument from false */

/* polymorph PatchDocument op false */

/* polymorph PatchDocument path false */

/* polymorph PatchDocument value false */

// Validate validates this patch document
func (m *PatchDocument) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateOp(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validatePath(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var patchDocumentTypeOpPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["add","remove","replace","move","copy","test"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		patchDocumentTypeOpPropEnum = append(patchDocumentTypeOpPropEnum, v)
	}
}

const (
	// PatchDocumentOpAdd captures enum value "add"
	PatchDocumentOpAdd string = "add"
	// PatchDocumentOpRemove captures enum value "remove"
	PatchDocumentOpRemove string = "remove"
	// PatchDocumentOpReplace captures enum value "replace"
	PatchDocumentOpReplace string = "replace"
	// PatchDocumentOpMove captures enum value "move"
	PatchDocumentOpMove string = "move"
	// PatchDocumentOpCopy captures enum value "copy"
	PatchDocumentOpCopy string = "copy"
	// PatchDocumentOpTest captures enum value "test"
	PatchDocumentOpTest string = "test"
)

// prop value enum
func (m *PatchDocument) validateOpEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, patchDocumentTypeOpPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *PatchDocument) validateOp(formats strfmt.Registry) error {

	if err := validate.Required("op", "body", m.Op); err != nil {
		return err
	}

	// value enum
	if err := m.validateOpEnum("op", "body", *m.Op); err != nil {
		return err
	}

	return nil
}

func (m *PatchDocument) validatePath(formats strfmt.Registry) error {

	if err := validate.Required("path", "body", m.Path); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PatchDocument) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PatchDocument) UnmarshalBinary(b []byte) error {
	var res PatchDocument
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
