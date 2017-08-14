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

// WeaviateLocationsGetHandlerFunc turns a function with the right signature into a weaviate locations get handler
type WeaviateLocationsGetHandlerFunc func(WeaviateLocationsGetParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn WeaviateLocationsGetHandlerFunc) Handle(params WeaviateLocationsGetParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// WeaviateLocationsGetHandler interface for that can handle valid weaviate locations get params
type WeaviateLocationsGetHandler interface {
	Handle(WeaviateLocationsGetParams, interface{}) middleware.Responder
}

// NewWeaviateLocationsGet creates a new http.Handler for the weaviate locations get operation
func NewWeaviateLocationsGet(ctx *middleware.Context, handler WeaviateLocationsGetHandler) *WeaviateLocationsGet {
	return &WeaviateLocationsGet{Context: ctx, Handler: handler}
}

/*WeaviateLocationsGet swagger:route GET /locations/{locationId} locations weaviateLocationsGet

Get a location based on its uuid related to this key.

Get a location.

*/
type WeaviateLocationsGet struct {
	Context *middleware.Context
	Handler WeaviateLocationsGetHandler
}

func (o *WeaviateLocationsGet) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewWeaviateLocationsGetParams()

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
