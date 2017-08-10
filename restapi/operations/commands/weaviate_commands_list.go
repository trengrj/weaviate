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
   

package commands

 
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// WeaviateCommandsListHandlerFunc turns a function with the right signature into a weaviate commands list handler
type WeaviateCommandsListHandlerFunc func(WeaviateCommandsListParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn WeaviateCommandsListHandlerFunc) Handle(params WeaviateCommandsListParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// WeaviateCommandsListHandler interface for that can handle valid weaviate commands list params
type WeaviateCommandsListHandler interface {
	Handle(WeaviateCommandsListParams, interface{}) middleware.Responder
}

// NewWeaviateCommandsList creates a new http.Handler for the weaviate commands list operation
func NewWeaviateCommandsList(ctx *middleware.Context, handler WeaviateCommandsListHandler) *WeaviateCommandsList {
	return &WeaviateCommandsList{Context: ctx, Handler: handler}
}

/*WeaviateCommandsList swagger:route GET /commands commands weaviateCommandsList

Get a list of commands related to this key.

Lists all commands in reverse order of creation, owned by the user that belongs to the used token.

*/
type WeaviateCommandsList struct {
	Context *middleware.Context
	Handler WeaviateCommandsListHandler
}

func (o *WeaviateCommandsList) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewWeaviateCommandsListParams()

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
