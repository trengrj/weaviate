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
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// GraphQLQuery GraphQL query based on: http://facebook.github.io/graphql/
// swagger:model GraphQLQuery

type GraphQLQuery struct {

	// Name of the operation if multiple exist in query.
	OperationName string `json:"operationName,omitempty"`

	// Query based on GraphQL syntax
	Query string `json:"query,omitempty"`

	// Additional variables for the query.
	Variables interface{} `json:"variables,omitempty"`
}

/* polymorph GraphQLQuery operationName false */

/* polymorph GraphQLQuery query false */

/* polymorph GraphQLQuery variables false */

// Validate validates this graph q l query
func (m *GraphQLQuery) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *GraphQLQuery) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GraphQLQuery) UnmarshalBinary(b []byte) error {
	var res GraphQLQuery
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
