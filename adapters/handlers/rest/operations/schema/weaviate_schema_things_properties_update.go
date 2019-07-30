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

package schema

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"

	models "github.com/semi-technologies/weaviate/entities/models"
)

// WeaviateSchemaThingsPropertiesUpdateHandlerFunc turns a function with the right signature into a weaviate schema things properties update handler
type WeaviateSchemaThingsPropertiesUpdateHandlerFunc func(WeaviateSchemaThingsPropertiesUpdateParams, *models.Principal) middleware.Responder

// Handle executing the request and returning a response
func (fn WeaviateSchemaThingsPropertiesUpdateHandlerFunc) Handle(params WeaviateSchemaThingsPropertiesUpdateParams, principal *models.Principal) middleware.Responder {
	return fn(params, principal)
}

// WeaviateSchemaThingsPropertiesUpdateHandler interface for that can handle valid weaviate schema things properties update params
type WeaviateSchemaThingsPropertiesUpdateHandler interface {
	Handle(WeaviateSchemaThingsPropertiesUpdateParams, *models.Principal) middleware.Responder
}

// NewWeaviateSchemaThingsPropertiesUpdate creates a new http.Handler for the weaviate schema things properties update operation
func NewWeaviateSchemaThingsPropertiesUpdate(ctx *middleware.Context, handler WeaviateSchemaThingsPropertiesUpdateHandler) *WeaviateSchemaThingsPropertiesUpdate {
	return &WeaviateSchemaThingsPropertiesUpdate{Context: ctx, Handler: handler}
}

/*WeaviateSchemaThingsPropertiesUpdate swagger:route PUT /schema/things/{className}/properties/{propertyName} schema weaviateSchemaThingsPropertiesUpdate

Rename, or replace the keywords of the property.

*/
type WeaviateSchemaThingsPropertiesUpdate struct {
	Context *middleware.Context
	Handler WeaviateSchemaThingsPropertiesUpdateHandler
}

func (o *WeaviateSchemaThingsPropertiesUpdate) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewWeaviateSchemaThingsPropertiesUpdateParams()

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
