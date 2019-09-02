//                           _       _
// __      _____  __ ___   ___  __ _| |_ ___
// \ \ /\ / / _ \/ _` \ \ / / |/ _` | __/ _ \
//  \ V  V /  __/ (_| |\ V /| | (_| | ||  __/
//   \_/\_/ \___|\__,_| \_/ |_|\__,_|\__\___|
//
//  Copyright © 2016 - 2019 SeMI Holding B.V. (registered @ Dutch Chamber of Commerce no 75221632). All rights reserved.
//  LICENSE WEAVIATE OPEN SOURCE: https://www.semi.technology/playbook/playbook/contract-weaviate-OSS.html
//  LICENSE WEAVIATE ENTERPRISE: https://www.semi.technology/playbook/contract-weaviate-enterprise.html
//  CONCEPT: Bob van Luijt (@bobvanluijt)
//  CONTACT: hello@semi.technology
//

// Code generated by go-swagger; DO NOT EDIT.

package actions

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"

	models "github.com/semi-technologies/weaviate/entities/models"
)

// ActionsGetHandlerFunc turns a function with the right signature into a actions get handler
type ActionsGetHandlerFunc func(ActionsGetParams, *models.Principal) middleware.Responder

// Handle executing the request and returning a response
func (fn ActionsGetHandlerFunc) Handle(params ActionsGetParams, principal *models.Principal) middleware.Responder {
	return fn(params, principal)
}

// ActionsGetHandler interface for that can handle valid actions get params
type ActionsGetHandler interface {
	Handle(ActionsGetParams, *models.Principal) middleware.Responder
}

// NewActionsGet creates a new http.Handler for the actions get operation
func NewActionsGet(ctx *middleware.Context, handler ActionsGetHandler) *ActionsGet {
	return &ActionsGet{Context: ctx, Handler: handler}
}

/*ActionsGet swagger:route GET /actions/{id} actions actionsGet

Get a specific Action based on its UUID and a Thing UUID. Also available as Websocket bus.

Lists Actions.

*/
type ActionsGet struct {
	Context *middleware.Context
	Handler ActionsGetHandler
}

func (o *ActionsGet) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewActionsGetParams()

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
