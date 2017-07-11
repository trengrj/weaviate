/*                          _       _
 *__      _____  __ ___   ___  __ _| |_ ___
 *\ \ /\ / / _ \/ _` \ \ / / |/ _` | __/ _ \
 * \ V  V /  __/ (_| |\ V /| | (_| | ||  __/
 *  \_/\_/ \___|\__,_| \_/ |_|\__,_|\__\___|
 *
 * Copyright © 2016 Weaviate. All rights reserved.
 * LICENSE: https://github.com/weaviate/weaviate/blob/master/LICENSE
 * AUTHOR: Bob van Luijt (bob@weaviate.com)
 * See www.weaviate.com for details
 * Contact: @weaviate_iot / yourfriends@weaviate.com
 */
  package groups

 
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// WeaviateGroupsCreateHandlerFunc turns a function with the right signature into a weaviate groups create handler
type WeaviateGroupsCreateHandlerFunc func(WeaviateGroupsCreateParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn WeaviateGroupsCreateHandlerFunc) Handle(params WeaviateGroupsCreateParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// WeaviateGroupsCreateHandler interface for that can handle valid weaviate groups create params
type WeaviateGroupsCreateHandler interface {
	Handle(WeaviateGroupsCreateParams, interface{}) middleware.Responder
}

// NewWeaviateGroupsCreate creates a new http.Handler for the weaviate groups create operation
func NewWeaviateGroupsCreate(ctx *middleware.Context, handler WeaviateGroupsCreateHandler) *WeaviateGroupsCreate {
	return &WeaviateGroupsCreate{Context: ctx, Handler: handler}
}

/*WeaviateGroupsCreate swagger:route POST /groups groups weaviateGroupsCreate

Create a new group related to this key.

Creates group.

*/
type WeaviateGroupsCreate struct {
	Context *middleware.Context
	Handler WeaviateGroupsCreateHandler
}

func (o *WeaviateGroupsCreate) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, _ := o.Context.RouteInfo(r)
	var Params = NewWeaviateGroupsCreateParams()

	uprinc, err := o.Context.Authorize(r, route)
	if err != nil {
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
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
