//                           _       _
// __      _____  __ ___   ___  __ _| |_ ___
// \ \ /\ / / _ \/ _` \ \ / / |/ _` | __/ _ \
//  \ V  V /  __/ (_| |\ V /| | (_| | ||  __/
//   \_/\_/ \___|\__,_| \_/ |_|\__,_|\__\___|
//
//  Copyright © 2016 - 2020 SeMI Technologies B.V. All rights reserved.
//
//  CONTACT: hello@semi.technology
//

// Code generated by go-swagger; DO NOT EDIT.

package schema

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"

	"github.com/semi-technologies/weaviate/entities/models"
)

// SchemaActionsDeleteHandlerFunc turns a function with the right signature into a schema actions delete handler
type SchemaActionsDeleteHandlerFunc func(SchemaActionsDeleteParams, *models.Principal) middleware.Responder

// Handle executing the request and returning a response
func (fn SchemaActionsDeleteHandlerFunc) Handle(params SchemaActionsDeleteParams, principal *models.Principal) middleware.Responder {
	return fn(params, principal)
}

// SchemaActionsDeleteHandler interface for that can handle valid schema actions delete params
type SchemaActionsDeleteHandler interface {
	Handle(SchemaActionsDeleteParams, *models.Principal) middleware.Responder
}

// NewSchemaActionsDelete creates a new http.Handler for the schema actions delete operation
func NewSchemaActionsDelete(ctx *middleware.Context, handler SchemaActionsDeleteHandler) *SchemaActionsDelete {
	return &SchemaActionsDelete{Context: ctx, Handler: handler}
}

/*SchemaActionsDelete swagger:route DELETE /schema/actions/{className} schema schemaActionsDelete

Remove an Action class (and all data in the instances) from the schema.

*/
type SchemaActionsDelete struct {
	Context *middleware.Context
	Handler SchemaActionsDeleteHandler
}

func (o *SchemaActionsDelete) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewSchemaActionsDeleteParams()

	uprinc, aCtx, err := o.Context.Authorize(r, route)
	if err != nil {
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}
	if aCtx != nil {
		r = aCtx
	}
	var principal *models.Principal
	if uprinc != nil {
		principal = uprinc.(*models.Principal) // this is really a models.Principal, I promise
	}

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params, principal) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
