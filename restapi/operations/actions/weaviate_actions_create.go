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

package actions

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	errors "github.com/go-openapi/errors"
	middleware "github.com/go-openapi/runtime/middleware"
	strfmt "github.com/go-openapi/strfmt"
	swag "github.com/go-openapi/swag"

	models "github.com/creativesoftwarefdn/weaviate/models"
)

// WeaviateActionsCreateHandlerFunc turns a function with the right signature into a weaviate actions create handler
type WeaviateActionsCreateHandlerFunc func(WeaviateActionsCreateParams, *models.Principal) middleware.Responder

// Handle executing the request and returning a response
func (fn WeaviateActionsCreateHandlerFunc) Handle(params WeaviateActionsCreateParams, principal *models.Principal) middleware.Responder {
	return fn(params, principal)
}

// WeaviateActionsCreateHandler interface for that can handle valid weaviate actions create params
type WeaviateActionsCreateHandler interface {
	Handle(WeaviateActionsCreateParams, *models.Principal) middleware.Responder
}

// NewWeaviateActionsCreate creates a new http.Handler for the weaviate actions create operation
func NewWeaviateActionsCreate(ctx *middleware.Context, handler WeaviateActionsCreateHandler) *WeaviateActionsCreate {
	return &WeaviateActionsCreate{Context: ctx, Handler: handler}
}

/*WeaviateActionsCreate swagger:route POST /actions actions weaviateActionsCreate

Create Actions between two Things (object and subject).

Registers a new Action. Provided meta-data and schema values are validated.

*/
type WeaviateActionsCreate struct {
	Context *middleware.Context
	Handler WeaviateActionsCreateHandler
}

func (o *WeaviateActionsCreate) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewWeaviateActionsCreateParams()

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

// WeaviateActionsCreateBody weaviate actions create body
// swagger:model WeaviateActionsCreateBody
type WeaviateActionsCreateBody struct {

	// action
	Action *models.ActionCreate `json:"action,omitempty"`

	// If `async` is true, return a 202 with the new ID of the Action. You will receive this response before the data is made persistent. If `async` is false, you will receive confirmation after the value is made persistent. The value of `async` defaults to false.
	Async bool `json:"async,omitempty"`
}

// Validate validates this weaviate actions create body
func (o *WeaviateActionsCreateBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateAction(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *WeaviateActionsCreateBody) validateAction(formats strfmt.Registry) error {

	if swag.IsZero(o.Action) { // not required
		return nil
	}

	if o.Action != nil {
		if err := o.Action.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("body" + "." + "action")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *WeaviateActionsCreateBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *WeaviateActionsCreateBody) UnmarshalBinary(b []byte) error {
	var res WeaviateActionsCreateBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
