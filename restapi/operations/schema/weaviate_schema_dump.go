/*                          _       _
 *__      _____  __ ___   ___  __ _| |_ ___
 *\ \ /\ / / _ \/ _` \ \ / / |/ _` | __/ _ \
 * \ V  V /  __/ (_| |\ V /| | (_| | ||  __/
 *  \_/\_/ \___|\__,_| \_/ |_|\__,_|\__\___|
 *
 * Copyright © 2016 - 2019 Weaviate. All rights reserved.
 * LICENSE: https://github.com/creativesoftwarefdn/weaviate/blob/develop/LICENSE.md
 * DESIGN & CONCEPT: Bob van Luijt (@bobvanluijt)
 * CONTACT: hello@creativesoftwarefdn.org
 */ // Code generated by go-swagger; DO NOT EDIT.

package schema

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	context "golang.org/x/net/context"

	errors "github.com/go-openapi/errors"
	middleware "github.com/go-openapi/runtime/middleware"
	strfmt "github.com/go-openapi/strfmt"
	swag "github.com/go-openapi/swag"

	models "github.com/creativesoftwarefdn/weaviate/models"
)

// WeaviateSchemaDumpHandlerFunc turns a function with the right signature into a weaviate schema dump handler
type WeaviateSchemaDumpHandlerFunc func(context.Context, WeaviateSchemaDumpParams) middleware.Responder

// Handle executing the request and returning a response
func (fn WeaviateSchemaDumpHandlerFunc) Handle(ctx context.Context, params WeaviateSchemaDumpParams) middleware.Responder {
	return fn(ctx, params)
}

// WeaviateSchemaDumpHandler interface for that can handle valid weaviate schema dump params
type WeaviateSchemaDumpHandler interface {
	Handle(context.Context, WeaviateSchemaDumpParams) middleware.Responder
}

// NewWeaviateSchemaDump creates a new http.Handler for the weaviate schema dump operation
func NewWeaviateSchemaDump(ctx *middleware.Context, handler WeaviateSchemaDumpHandler) *WeaviateSchemaDump {
	return &WeaviateSchemaDump{Context: ctx, Handler: handler}
}

/*WeaviateSchemaDump swagger:route GET /schema schema weaviateSchemaDump

Dump the current the database schema.

*/
type WeaviateSchemaDump struct {
	Context *middleware.Context
	Handler WeaviateSchemaDumpHandler
}

func (o *WeaviateSchemaDump) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewWeaviateSchemaDumpParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(r.Context(), Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}

// WeaviateSchemaDumpOKBody weaviate schema dump o k body
// swagger:model WeaviateSchemaDumpOKBody
type WeaviateSchemaDumpOKBody struct {

	// actions
	Actions *models.SemanticSchema `json:"actions,omitempty"`

	// things
	Things *models.SemanticSchema `json:"things,omitempty"`
}

// Validate validates this weaviate schema dump o k body
func (o *WeaviateSchemaDumpOKBody) Validate(formats strfmt.Registry) error {
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

func (o *WeaviateSchemaDumpOKBody) validateActions(formats strfmt.Registry) error {

	if swag.IsZero(o.Actions) { // not required
		return nil
	}

	if o.Actions != nil {
		if err := o.Actions.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("weaviateSchemaDumpOK" + "." + "actions")
			}
			return err
		}
	}

	return nil
}

func (o *WeaviateSchemaDumpOKBody) validateThings(formats strfmt.Registry) error {

	if swag.IsZero(o.Things) { // not required
		return nil
	}

	if o.Things != nil {
		if err := o.Things.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("weaviateSchemaDumpOK" + "." + "things")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *WeaviateSchemaDumpOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *WeaviateSchemaDumpOKBody) UnmarshalBinary(b []byte) error {
	var res WeaviateSchemaDumpOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
