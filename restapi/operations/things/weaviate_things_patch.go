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

package things

// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// WeaviateThingsPatchHandlerFunc turns a function with the right signature into a weaviate things patch handler
type WeaviateThingsPatchHandlerFunc func(WeaviateThingsPatchParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn WeaviateThingsPatchHandlerFunc) Handle(params WeaviateThingsPatchParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// WeaviateThingsPatchHandler interface for that can handle valid weaviate things patch params
type WeaviateThingsPatchHandler interface {
	Handle(WeaviateThingsPatchParams, interface{}) middleware.Responder
}

// NewWeaviateThingsPatch creates a new http.Handler for the weaviate things patch operation
func NewWeaviateThingsPatch(ctx *middleware.Context, handler WeaviateThingsPatchHandler) *WeaviateThingsPatch {
	return &WeaviateThingsPatch{Context: ctx, Handler: handler}
}

/*WeaviateThingsPatch swagger:route PATCH /things/{thingId} things weaviateThingsPatch

Update a thing based on its uuid (using patch semantics) related to this key.

Updates a thing data. This method supports patch semantics.

*/
type WeaviateThingsPatch struct {
	Context *middleware.Context
	Handler WeaviateThingsPatchHandler
}

func (o *WeaviateThingsPatch) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewWeaviateThingsPatchParams()

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
