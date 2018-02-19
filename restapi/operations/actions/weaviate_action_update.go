/*                          _       _
 *__      _____  __ ___   ___  __ _| |_ ___
 *\ \ /\ / / _ \/ _` \ \ / / |/ _` | __/ _ \
 * \ V  V /  __/ (_| |\ V /| | (_| | ||  __/
 *  \_/\_/ \___|\__,_| \_/ |_|\__,_|\__\___|
 *
 * Copyright © 2016 - 2018 - 2018 Weaviate. All rights reserved.
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

// WeaviateActionUpdateHandlerFunc turns a function with the right signature into a weaviate action update handler
type WeaviateActionUpdateHandlerFunc func(WeaviateActionUpdateParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn WeaviateActionUpdateHandlerFunc) Handle(params WeaviateActionUpdateParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// WeaviateActionUpdateHandler interface for that can handle valid weaviate action update params
type WeaviateActionUpdateHandler interface {
	Handle(WeaviateActionUpdateParams, interface{}) middleware.Responder
}

// NewWeaviateActionUpdate creates a new http.Handler for the weaviate action update operation
func NewWeaviateActionUpdate(ctx *middleware.Context, handler WeaviateActionUpdateHandler) *WeaviateActionUpdate {
	return &WeaviateActionUpdate{Context: ctx, Handler: handler}
}

/*WeaviateActionUpdate swagger:route PUT /actions/{actionId} actions weaviateActionUpdate

Update an action based on its uuid related to this key.

Updates an action's data.

*/
type WeaviateActionUpdate struct {
	Context *middleware.Context
	Handler WeaviateActionUpdateHandler
}

func (o *WeaviateActionUpdate) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewWeaviateActionUpdateParams()

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
