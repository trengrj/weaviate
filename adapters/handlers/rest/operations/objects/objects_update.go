//                           _       _
// __      _____  __ ___   ___  __ _| |_ ___
// \ \ /\ / / _ \/ _` \ \ / / |/ _` | __/ _ \
//  \ V  V /  __/ (_| |\ V /| | (_| | ||  __/
//   \_/\_/ \___|\__,_| \_/ |_|\__,_|\__\___|
//
//  Copyright © 2016 - 2022 SeMI Technologies B.V. All rights reserved.
//
//  CONTACT: hello@semi.technology
//

// Code generated by go-swagger; DO NOT EDIT.

package objects

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"

	"github.com/semi-technologies/weaviate/entities/models"
)

// ObjectsUpdateHandlerFunc turns a function with the right signature into a objects update handler
type ObjectsUpdateHandlerFunc func(ObjectsUpdateParams, *models.Principal) middleware.Responder

// Handle executing the request and returning a response
func (fn ObjectsUpdateHandlerFunc) Handle(params ObjectsUpdateParams, principal *models.Principal) middleware.Responder {
	return fn(params, principal)
}

// ObjectsUpdateHandler interface for that can handle valid objects update params
type ObjectsUpdateHandler interface {
	Handle(ObjectsUpdateParams, *models.Principal) middleware.Responder
}

// NewObjectsUpdate creates a new http.Handler for the objects update operation
func NewObjectsUpdate(ctx *middleware.Context, handler ObjectsUpdateHandler) *ObjectsUpdate {
	return &ObjectsUpdate{Context: ctx, Handler: handler}
}

/*
ObjectsUpdate swagger:route PUT /objects/{id} objects objectsUpdate

Update an Object based on its UUID.

Updates an Object's data. Given meta-data and schema values are validated. LastUpdateTime is set to the time this function is called.
*/
type ObjectsUpdate struct {
	Context *middleware.Context
	Handler ObjectsUpdateHandler
}

func (o *ObjectsUpdate) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewObjectsUpdateParams()

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
