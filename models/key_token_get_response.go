/*                          _       _
 *__      _____  __ ___   ___  __ _| |_ ___
 *\ \ /\ / / _ \/ _` \ \ / / |/ _` | __/ _ \
 * \ V  V /  __/ (_| |\ V /| | (_| | ||  __/
 *  \_/\_/ \___|\__,_| \_/ |_|\__,_|\__\___|
 *
 * Copyright © 2016 Weaviate. All rights reserved.
 * LICENSE: https://github.com/weaviate/weaviate/blob/master/LICENSE
 * AUTHOR: Bob van Luijt (bob@weaviate.com)
 * See www.weaviate.com for details
 * Contact: @weaviate_iot / yourfriends@weaviate.com
 */
  package models

 
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// KeyTokenGetResponse key token get response
// swagger:model KeyTokenGetResponse
type KeyTokenGetResponse struct {
	KeyGetResponse

	// Key for user to use.
	Key string `json:"key,omitempty"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *KeyTokenGetResponse) UnmarshalJSON(raw []byte) error {

	var aO0 KeyGetResponse
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.KeyGetResponse = aO0

	var data struct {
		Key string `json:"key,omitempty"`
	}
	if err := swag.ReadJSON(raw, &data); err != nil {
		return err
	}

	m.Key = data.Key

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m KeyTokenGetResponse) MarshalJSON() ([]byte, error) {
	var _parts [][]byte

	aO0, err := swag.WriteJSON(m.KeyGetResponse)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	var data struct {
		Key string `json:"key,omitempty"`
	}

	data.Key = m.Key

	jsonData, err := swag.WriteJSON(data)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, jsonData)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this key token get response
func (m *KeyTokenGetResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.KeyGetResponse.Validate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
