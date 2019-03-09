/*                          _       _
 *__      _____  __ ___   ___  __ _| |_ ___
 *\ \ /\ / / _ \/ _` \ \ / / |/ _` | __/ _ \
 * \ V  V /  __/ (_| |\ V /| | (_| | ||  __/
 *  \_/\_/ \___|\__,_| \_/ |_|\__,_|\__\___|
 *
 * Copyright © 2016 - 2019 Weaviate. All rights reserved.
 * LICENSE: https://github.com/creativesoftwarefdn/weaviate/blob/develop/LICENSE.md
 * DESIGN & CONCEPT: Bob van Luijt (@bobvanluijt)
 * CONTACT: hello@creativesoftwarefdn.org
 */ // Code generated by go-swagger; DO NOT EDIT.

package things

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// WeaviateThingHistoryGetHandlerFunc turns a function with the right signature into a weaviate thing history get handler
type WeaviateThingHistoryGetHandlerFunc func(WeaviateThingHistoryGetParams) middleware.Responder

// Handle executing the request and returning a response
func (fn WeaviateThingHistoryGetHandlerFunc) Handle(params WeaviateThingHistoryGetParams) middleware.Responder {
	return fn(params)
}

// WeaviateThingHistoryGetHandler interface for that can handle valid weaviate thing history get params
type WeaviateThingHistoryGetHandler interface {
	Handle(WeaviateThingHistoryGetParams) middleware.Responder
}

// NewWeaviateThingHistoryGet creates a new http.Handler for the weaviate thing history get operation
func NewWeaviateThingHistoryGet(ctx *middleware.Context, handler WeaviateThingHistoryGetHandler) *WeaviateThingHistoryGet {
	return &WeaviateThingHistoryGet{Context: ctx, Handler: handler}
}

/*WeaviateThingHistoryGet swagger:route GET /things/{thingId}/history things weaviateThingHistoryGet

Get a Thing's history based on its UUID.

Returns a particular Thing's history.

*/
type WeaviateThingHistoryGet struct {
	Context *middleware.Context
	Handler WeaviateThingHistoryGetHandler
}

func (o *WeaviateThingHistoryGet) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewWeaviateThingHistoryGetParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
