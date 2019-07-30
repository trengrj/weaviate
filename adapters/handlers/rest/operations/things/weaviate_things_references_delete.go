//                           _       _
// __      _____  __ ___   ___  __ _| |_ ___
// \ \ /\ / / _ \/ _` \ \ / / |/ _` | __/ _ \
//  \ V  V /  __/ (_| |\ V /| | (_| | ||  __/
//   \_/\_/ \___|\__,_| \_/ |_|\__,_|\__\___|
//
//  Copyright © 2016 - 2019 Weaviate. All rights reserved.
//  LICENSE WEAVIATE OPEN SOURCE: https://www.semi.technology/playbook/playbook/contract-weaviate-OSS.html
//  LICENSE WEAVIATE ENTERPRISE: https://www.semi.technology/playbook/contract-weaviate-enterprise.html
//  CONCEPT: Bob van Luijt (@bobvanluijt)
//  CONTACT: hello@semi.technology
//

// Code generated by go-swagger; DO NOT EDIT.

package things

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"

	models "github.com/semi-technologies/weaviate/entities/models"
)

// WeaviateThingsReferencesDeleteHandlerFunc turns a function with the right signature into a weaviate things references delete handler
type WeaviateThingsReferencesDeleteHandlerFunc func(WeaviateThingsReferencesDeleteParams, *models.Principal) middleware.Responder

// Handle executing the request and returning a response
func (fn WeaviateThingsReferencesDeleteHandlerFunc) Handle(params WeaviateThingsReferencesDeleteParams, principal *models.Principal) middleware.Responder {
	return fn(params, principal)
}

// WeaviateThingsReferencesDeleteHandler interface for that can handle valid weaviate things references delete params
type WeaviateThingsReferencesDeleteHandler interface {
	Handle(WeaviateThingsReferencesDeleteParams, *models.Principal) middleware.Responder
}

// NewWeaviateThingsReferencesDelete creates a new http.Handler for the weaviate things references delete operation
func NewWeaviateThingsReferencesDelete(ctx *middleware.Context, handler WeaviateThingsReferencesDeleteHandler) *WeaviateThingsReferencesDelete {
	return &WeaviateThingsReferencesDelete{Context: ctx, Handler: handler}
}

/*WeaviateThingsReferencesDelete swagger:route DELETE /things/{id}/references/{propertyName} things weaviateThingsReferencesDelete

Delete the single reference that is given in the body from the list of references that this property has.

Delete the single reference that is given in the body from the list of references that this property has.

*/
type WeaviateThingsReferencesDelete struct {
	Context *middleware.Context
	Handler WeaviateThingsReferencesDeleteHandler
}

func (o *WeaviateThingsReferencesDelete) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewWeaviateThingsReferencesDeleteParams()

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
