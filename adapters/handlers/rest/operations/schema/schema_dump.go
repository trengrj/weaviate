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

package schema

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	errors "github.com/go-openapi/errors"
	middleware "github.com/go-openapi/runtime/middleware"
	strfmt "github.com/go-openapi/strfmt"
	swag "github.com/go-openapi/swag"

	models "github.com/semi-technologies/weaviate/entities/models"
)

// SchemaDumpHandlerFunc turns a function with the right signature into a schema dump handler
type SchemaDumpHandlerFunc func(SchemaDumpParams, *models.Principal) middleware.Responder

// Handle executing the request and returning a response
func (fn SchemaDumpHandlerFunc) Handle(params SchemaDumpParams, principal *models.Principal) middleware.Responder {
	return fn(params, principal)
}

// SchemaDumpHandler interface for that can handle valid schema dump params
type SchemaDumpHandler interface {
	Handle(SchemaDumpParams, *models.Principal) middleware.Responder
}

// NewSchemaDump creates a new http.Handler for the schema dump operation
func NewSchemaDump(ctx *middleware.Context, handler SchemaDumpHandler) *SchemaDump {
	return &SchemaDump{Context: ctx, Handler: handler}
}

/*SchemaDump swagger:route GET /schema schema schemaDump

Dump the current the database schema.

*/
type SchemaDump struct {
	Context *middleware.Context
	Handler SchemaDumpHandler
}

func (o *SchemaDump) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewSchemaDumpParams()

	uprinc, aCtx, err := o.Context.Authorize(r, route)
	if err != nil {
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}
	if aCtx != nil {
		r = aCtx
	}
	var principal *models.Principal
	if uprinc != nil {
		principal = uprinc.(*models.Principal) // this is really a models.Principal, I promise
	}

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params, principal) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}

// SchemaDumpOKBody schema dump o k body
// swagger:model SchemaDumpOKBody
type SchemaDumpOKBody struct {

	// actions
	Actions *models.Schema `json:"actions,omitempty"`

	// things
	Things *models.Schema `json:"things,omitempty"`
}

// Validate validates this schema dump o k body
func (o *SchemaDumpOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateActions(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateThings(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *SchemaDumpOKBody) validateActions(formats strfmt.Registry) error {

	if swag.IsZero(o.Actions) { // not required
		return nil
	}

	if o.Actions != nil {
		if err := o.Actions.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("schemaDumpOK" + "." + "actions")
			}
			return err
		}
	}

	return nil
}

func (o *SchemaDumpOKBody) validateThings(formats strfmt.Registry) error {

	if swag.IsZero(o.Things) { // not required
		return nil
	}

	if o.Things != nil {
		if err := o.Things.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("schemaDumpOK" + "." + "things")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *SchemaDumpOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *SchemaDumpOKBody) UnmarshalBinary(b []byte) error {
	var res SchemaDumpOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
