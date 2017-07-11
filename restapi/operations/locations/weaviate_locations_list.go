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
  package locations

 
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// WeaviateLocationsListHandlerFunc turns a function with the right signature into a weaviate locations list handler
type WeaviateLocationsListHandlerFunc func(WeaviateLocationsListParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn WeaviateLocationsListHandlerFunc) Handle(params WeaviateLocationsListParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// WeaviateLocationsListHandler interface for that can handle valid weaviate locations list params
type WeaviateLocationsListHandler interface {
	Handle(WeaviateLocationsListParams, interface{}) middleware.Responder
}

// NewWeaviateLocationsList creates a new http.Handler for the weaviate locations list operation
func NewWeaviateLocationsList(ctx *middleware.Context, handler WeaviateLocationsListHandler) *WeaviateLocationsList {
	return &WeaviateLocationsList{Context: ctx, Handler: handler}
}

/*WeaviateLocationsList swagger:route GET /locations locations weaviateLocationsList

Get a list of locations related to this key.

Lists all locations in reverse order of creation, owned by the user that belongs to the used token.

*/
type WeaviateLocationsList struct {
	Context *middleware.Context
	Handler WeaviateLocationsListHandler
}

func (o *WeaviateLocationsList) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, _ := o.Context.RouteInfo(r)
	var Params = NewWeaviateLocationsListParams()

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
