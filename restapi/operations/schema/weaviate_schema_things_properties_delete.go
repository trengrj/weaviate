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

package schema

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	context "golang.org/x/net/context"

	middleware "github.com/go-openapi/runtime/middleware"
)

// WeaviateSchemaThingsPropertiesDeleteHandlerFunc turns a function with the right signature into a weaviate schema things properties delete handler
type WeaviateSchemaThingsPropertiesDeleteHandlerFunc func(context.Context, WeaviateSchemaThingsPropertiesDeleteParams) middleware.Responder

// Handle executing the request and returning a response
func (fn WeaviateSchemaThingsPropertiesDeleteHandlerFunc) Handle(ctx context.Context, params WeaviateSchemaThingsPropertiesDeleteParams) middleware.Responder {
	return fn(ctx, params)
}

// WeaviateSchemaThingsPropertiesDeleteHandler interface for that can handle valid weaviate schema things properties delete params
type WeaviateSchemaThingsPropertiesDeleteHandler interface {
	Handle(context.Context, WeaviateSchemaThingsPropertiesDeleteParams) middleware.Responder
}

// NewWeaviateSchemaThingsPropertiesDelete creates a new http.Handler for the weaviate schema things properties delete operation
func NewWeaviateSchemaThingsPropertiesDelete(ctx *middleware.Context, handler WeaviateSchemaThingsPropertiesDeleteHandler) *WeaviateSchemaThingsPropertiesDelete {
	return &WeaviateSchemaThingsPropertiesDelete{Context: ctx, Handler: handler}
}

/*WeaviateSchemaThingsPropertiesDelete swagger:route DELETE /schema/things/{className}/properties/{propertyName} schema weaviateSchemaThingsPropertiesDelete

Remove a property from a Thing class.

*/
type WeaviateSchemaThingsPropertiesDelete struct {
	Context *middleware.Context
	Handler WeaviateSchemaThingsPropertiesDeleteHandler
}

func (o *WeaviateSchemaThingsPropertiesDelete) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewWeaviateSchemaThingsPropertiesDeleteParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(r.Context(), Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
