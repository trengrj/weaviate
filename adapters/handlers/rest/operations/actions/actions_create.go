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

package actions

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"

	"github.com/semi-technologies/weaviate/entities/models"
)

// ActionsCreateHandlerFunc turns a function with the right signature into a actions create handler
type ActionsCreateHandlerFunc func(ActionsCreateParams, *models.Principal) middleware.Responder

// Handle executing the request and returning a response
func (fn ActionsCreateHandlerFunc) Handle(params ActionsCreateParams, principal *models.Principal) middleware.Responder {
	return fn(params, principal)
}

// ActionsCreateHandler interface for that can handle valid actions create params
type ActionsCreateHandler interface {
	Handle(ActionsCreateParams, *models.Principal) middleware.Responder
}

// NewActionsCreate creates a new http.Handler for the actions create operation
func NewActionsCreate(ctx *middleware.Context, handler ActionsCreateHandler) *ActionsCreate {
	return &ActionsCreate{Context: ctx, Handler: handler}
}

/*ActionsCreate swagger:route POST /actions actions actionsCreate

Create Actions between two Things (object and subject).

Registers a new Action. Provided meta-data and schema values are validated.

*/
type ActionsCreate struct {
	Context *middleware.Context
	Handler ActionsCreateHandler
}

func (o *ActionsCreate) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewActionsCreateParams()

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
