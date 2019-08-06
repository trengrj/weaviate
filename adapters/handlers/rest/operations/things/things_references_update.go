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

// ThingsReferencesUpdateHandlerFunc turns a function with the right signature into a things references update handler
type ThingsReferencesUpdateHandlerFunc func(ThingsReferencesUpdateParams, *models.Principal) middleware.Responder

// Handle executing the request and returning a response
func (fn ThingsReferencesUpdateHandlerFunc) Handle(params ThingsReferencesUpdateParams, principal *models.Principal) middleware.Responder {
	return fn(params, principal)
}

// ThingsReferencesUpdateHandler interface for that can handle valid things references update params
type ThingsReferencesUpdateHandler interface {
	Handle(ThingsReferencesUpdateParams, *models.Principal) middleware.Responder
}

// NewThingsReferencesUpdate creates a new http.Handler for the things references update operation
func NewThingsReferencesUpdate(ctx *middleware.Context, handler ThingsReferencesUpdateHandler) *ThingsReferencesUpdate {
	return &ThingsReferencesUpdate{Context: ctx, Handler: handler}
}

/*ThingsReferencesUpdate swagger:route PUT /things/{id}/references/{propertyName} things thingsReferencesUpdate

Replace all references to a class-property.

Replace all references to a class-property.

*/
type ThingsReferencesUpdate struct {
	Context *middleware.Context
	Handler ThingsReferencesUpdateHandler
}

func (o *ThingsReferencesUpdate) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewThingsReferencesUpdateParams()

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
