/*                          _       _
 *__      _____  __ ___   ___  __ _| |_ ___
 *\ \ /\ / / _ \/ _` \ \ / / |/ _` | __/ _ \
 * \ V  V /  __/ (_| |\ V /| | (_| | ||  __/
 *  \_/\_/ \___|\__,_| \_/ |_|\__,_|\__\___|
 *
 * Copyright © 2016 - 2018 Weaviate. All rights reserved.
 * LICENSE: https://github.com/creativesoftwarefdn/weaviate/blob/develop/LICENSE.md
 * AUTHOR: Bob van Luijt (bob@kub.design)
 * See www.creativesoftwarefdn.org for details
 * Contact: @CreativeSofwFdn / bob@kub.design
 */
// Code generated by go-swagger; DO NOT EDIT.

package actions

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	context "golang.org/x/net/context"

	middleware "github.com/go-openapi/runtime/middleware"
)

// WeaviateActionsValidateHandlerFunc turns a function with the right signature into a weaviate actions validate handler
type WeaviateActionsValidateHandlerFunc func(context.Context, WeaviateActionsValidateParams) middleware.Responder

// Handle executing the request and returning a response
func (fn WeaviateActionsValidateHandlerFunc) Handle(ctx context.Context, params WeaviateActionsValidateParams) middleware.Responder {
	return fn(ctx, params)
}

// WeaviateActionsValidateHandler interface for that can handle valid weaviate actions validate params
type WeaviateActionsValidateHandler interface {
	Handle(context.Context, WeaviateActionsValidateParams) middleware.Responder
}

// NewWeaviateActionsValidate creates a new http.Handler for the weaviate actions validate operation
func NewWeaviateActionsValidate(ctx *middleware.Context, handler WeaviateActionsValidateHandler) *WeaviateActionsValidate {
	return &WeaviateActionsValidate{Context: ctx, Handler: handler}
}

/*WeaviateActionsValidate swagger:route POST /actions/validate actions weaviateActionsValidate

Validate an Action based on a schema.

Validate an Action's schema and meta-data. It has to be based on a schema, which is related to the given Action to be accepted by this validation.

*/
type WeaviateActionsValidate struct {
	Context *middleware.Context
	Handler WeaviateActionsValidateHandler
}

func (o *WeaviateActionsValidate) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewWeaviateActionsValidateParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(r.Context(), Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
