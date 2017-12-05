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

package actions

// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// WeaviateActionsGetHandlerFunc turns a function with the right signature into a weaviate actions get handler
type WeaviateActionsGetHandlerFunc func(WeaviateActionsGetParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn WeaviateActionsGetHandlerFunc) Handle(params WeaviateActionsGetParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// WeaviateActionsGetHandler interface for that can handle valid weaviate actions get params
type WeaviateActionsGetHandler interface {
	Handle(WeaviateActionsGetParams, interface{}) middleware.Responder
}

// NewWeaviateActionsGet creates a new http.Handler for the weaviate actions get operation
func NewWeaviateActionsGet(ctx *middleware.Context, handler WeaviateActionsGetHandler) *WeaviateActionsGet {
	return &WeaviateActionsGet{Context: ctx, Handler: handler}
}

/*WeaviateActionsGet swagger:route GET /actions/{actionId} actions weaviateActionsGet

Get a specific action based on its uuid and a thing uuid related to this key. Also available as Websocket bus.

Lists actions.

*/
type WeaviateActionsGet struct {
	Context *middleware.Context
	Handler WeaviateActionsGetHandler
}

func (o *WeaviateActionsGet) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewWeaviateActionsGetParams()

	uprinc, aCtx, err := o.Context.Authorize(r, route)
	if err != nil {
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}
	if aCtx != nil {
		r = aCtx
	}
	var principal interface{}
	if uprinc != nil {
		principal = uprinc
	}

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params, principal) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
