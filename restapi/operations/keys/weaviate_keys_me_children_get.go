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

package keys

// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// WeaviateKeysMeChildrenGetHandlerFunc turns a function with the right signature into a weaviate keys me children get handler
type WeaviateKeysMeChildrenGetHandlerFunc func(WeaviateKeysMeChildrenGetParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn WeaviateKeysMeChildrenGetHandlerFunc) Handle(params WeaviateKeysMeChildrenGetParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// WeaviateKeysMeChildrenGetHandler interface for that can handle valid weaviate keys me children get params
type WeaviateKeysMeChildrenGetHandler interface {
	Handle(WeaviateKeysMeChildrenGetParams, interface{}) middleware.Responder
}

// NewWeaviateKeysMeChildrenGet creates a new http.Handler for the weaviate keys me children get operation
func NewWeaviateKeysMeChildrenGet(ctx *middleware.Context, handler WeaviateKeysMeChildrenGetHandler) *WeaviateKeysMeChildrenGet {
	return &WeaviateKeysMeChildrenGet{Context: ctx, Handler: handler}
}

/*WeaviateKeysMeChildrenGet swagger:route GET /keys/me/children keys weaviateKeysMeChildrenGet

Get an object of this keys' children related to the key used for request.

Get children of used key, only one step deep. A child can have children of its own.

*/
type WeaviateKeysMeChildrenGet struct {
	Context *middleware.Context
	Handler WeaviateKeysMeChildrenGetHandler
}

func (o *WeaviateKeysMeChildrenGet) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewWeaviateKeysMeChildrenGetParams()

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
