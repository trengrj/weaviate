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
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// C11yExtension A resource describing an extension to the contextinoary, containing both the identifier and the definition of the extension
//
// swagger:model C11yExtension
type C11yExtension struct {

	// The new concept you want to extend. Must be an all-lowercase single word, or a space delimited compound word. Examples: 'foobarium', 'my custom concept'
	Concept string `json:"concept,omitempty"`

	// A list of space-delimited words or a sentence describing what the custom concept is about. Avoid using the custom concept itself. An Example definition for the custom concept 'foobarium': would be 'a naturally occourring element which can only be seen by programmers'
	Definition string `json:"definition,omitempty"`

	// Weight of the definition of the new concept where 1='override existing definition entirely' and 0='ignore custom definition'. Note that if the custom concept is not present in the contextionary yet, the weight cannot be less than 1.
	Weight float32 `json:"weight,omitempty"`
}

// Validate validates this c11y extension
func (m *C11yExtension) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *C11yExtension) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *C11yExtension) UnmarshalBinary(b []byte) error {
	var res C11yExtension
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
