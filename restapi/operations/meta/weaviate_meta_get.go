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

package meta

// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// WeaviateMetaGetHandlerFunc turns a function with the right signature into a weaviate meta get handler
type WeaviateMetaGetHandlerFunc func(WeaviateMetaGetParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn WeaviateMetaGetHandlerFunc) Handle(params WeaviateMetaGetParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// WeaviateMetaGetHandler interface for that can handle valid weaviate meta get params
type WeaviateMetaGetHandler interface {
	Handle(WeaviateMetaGetParams, interface{}) middleware.Responder
}

// NewWeaviateMetaGet creates a new http.Handler for the weaviate meta get operation
func NewWeaviateMetaGet(ctx *middleware.Context, handler WeaviateMetaGetHandler) *WeaviateMetaGet {
	return &WeaviateMetaGet{Context: ctx, Handler: handler}
}

/*WeaviateMetaGet swagger:route GET /meta meta weaviateMetaGet

Returns meta information of the current Weaviate instance.

Gives meta information about the server and can be used to provide information to another Weaviate instance that wants to interact with the current instance.

*/
type WeaviateMetaGet struct {
	Context *middleware.Context
	Handler WeaviateMetaGetHandler
}

func (o *WeaviateMetaGet) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewWeaviateMetaGetParams()

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
