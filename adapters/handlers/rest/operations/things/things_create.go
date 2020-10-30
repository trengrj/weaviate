//                           _       _
// __      _____  __ ___   ___  __ _| |_ ___
// \ \ /\ / / _ \/ _` \ \ / / |/ _` | __/ _ \
//  \ V  V /  __/ (_| |\ V /| | (_| | ||  __/
//   \_/\_/ \___|\__,_| \_/ |_|\__,_|\__\___|
//
//  Copyright © 2016 - 2020 SeMI Technologies B.V. All rights reserved.
//
//  CONTACT: hello@semi.technology
//

// Code generated by go-swagger; DO NOT EDIT.

package things

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"

	"github.com/semi-technologies/weaviate/entities/models"
)

// ThingsCreateHandlerFunc turns a function with the right signature into a things create handler
type ThingsCreateHandlerFunc func(ThingsCreateParams, *models.Principal) middleware.Responder

// Handle executing the request and returning a response
func (fn ThingsCreateHandlerFunc) Handle(params ThingsCreateParams, principal *models.Principal) middleware.Responder {
	return fn(params, principal)
}

// ThingsCreateHandler interface for that can handle valid things create params
type ThingsCreateHandler interface {
	Handle(ThingsCreateParams, *models.Principal) middleware.Responder
}

// NewThingsCreate creates a new http.Handler for the things create operation
func NewThingsCreate(ctx *middleware.Context, handler ThingsCreateHandler) *ThingsCreate {
	return &ThingsCreate{Context: ctx, Handler: handler}
}

/*ThingsCreate swagger:route POST /things things thingsCreate

Create a new Thing based on a Thing template.

Registers a new Thing. Given meta-data and schema values are validated.

*/
type ThingsCreate struct {
	Context *middleware.Context
	Handler ThingsCreateHandler
}

func (o *ThingsCreate) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	Params := NewThingsCreateParams()

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
