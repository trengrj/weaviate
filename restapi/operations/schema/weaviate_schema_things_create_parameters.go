/*                          _       _
 *__      _____  __ ___   ___  __ _| |_ ___
 *\ \ /\ / / _ \/ _` \ \ / / |/ _` | __/ _ \
 * \ V  V /  __/ (_| |\ V /| | (_| | ||  __/
 *  \_/\_/ \___|\__,_| \_/ |_|\__,_|\__\___|
 *
 * Copyright © 2016 - 2019 Weaviate. All rights reserved.
 * LICENSE: https://github.com/creativesoftwarefdn/weaviate/blob/develop/LICENSE.md
 * DESIGN: Bob van Luijt (bob@k10y.co)
 */
// Code generated by go-swagger; DO NOT EDIT.

package schema

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"io"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"

	models "github.com/creativesoftwarefdn/weaviate/models"
)

// NewWeaviateSchemaThingsCreateParams creates a new WeaviateSchemaThingsCreateParams object
// no default values defined in spec.
func NewWeaviateSchemaThingsCreateParams() WeaviateSchemaThingsCreateParams {

	return WeaviateSchemaThingsCreateParams{}
}

// WeaviateSchemaThingsCreateParams contains all the bound params for the weaviate schema things create operation
// typically these are obtained from a http.Request
//
// swagger:parameters weaviate.schema.things.create
type WeaviateSchemaThingsCreateParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*
	  Required: true
	  In: body
	*/
	ThingClass *models.SemanticSchemaClass
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewWeaviateSchemaThingsCreateParams() beforehand.
func (o *WeaviateSchemaThingsCreateParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	if runtime.HasBody(r) {
		defer r.Body.Close()
		var body models.SemanticSchemaClass
		if err := route.Consumer.Consume(r.Body, &body); err != nil {
			if err == io.EOF {
				res = append(res, errors.Required("thingClass", "body"))
			} else {
				res = append(res, errors.NewParseError("thingClass", "body", "", err))
			}
		} else {
			// validate body object
			if err := body.Validate(route.Formats); err != nil {
				res = append(res, err)
			}

			if len(res) == 0 {
				o.ThingClass = &body
			}
		}
	} else {
		res = append(res, errors.Required("thingClass", "body"))
	}
	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
