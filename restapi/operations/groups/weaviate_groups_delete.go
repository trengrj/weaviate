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

// WeaviateGroupsDeleteHandlerFunc turns a function with the right signature into a weaviate groups delete handler
type WeaviateGroupsDeleteHandlerFunc func(WeaviateGroupsDeleteParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn WeaviateGroupsDeleteHandlerFunc) Handle(params WeaviateGroupsDeleteParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// WeaviateGroupsDeleteHandler interface for that can handle valid weaviate groups delete params
type WeaviateGroupsDeleteHandler interface {
	Handle(WeaviateGroupsDeleteParams, interface{}) middleware.Responder
}

// NewWeaviateGroupsDelete creates a new http.Handler for the weaviate groups delete operation
func NewWeaviateGroupsDelete(ctx *middleware.Context, handler WeaviateGroupsDeleteHandler) *WeaviateGroupsDelete {
	return &WeaviateGroupsDelete{Context: ctx, Handler: handler}
}

/*WeaviateGroupsDelete swagger:route DELETE /groups/{groupId} groups weaviateGroupsDelete

Delete a group based on its uuid related to this key.

Deletes an group.

*/
type WeaviateGroupsDelete struct {
	Context *middleware.Context
	Handler WeaviateGroupsDeleteHandler
}

func (o *WeaviateGroupsDelete) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewWeaviateGroupsDeleteParams()

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
