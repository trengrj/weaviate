/*                          _       _
 *__      _____  __ ___   ___  __ _| |_ ___
 *\ \ /\ / / _ \/ _` \ \ / / |/ _` | __/ _ \
 * \ V  V /  __/ (_| |\ V /| | (_| | ||  __/
 *  \_/\_/ \___|\__,_| \_/ |_|\__,_|\__\___|
 *
 * Copyright © 2016 - 2019 Weaviate. All rights reserved.
 * LICENSE: https://github.com/semi-technologies/weaviate/blob/develop/LICENSE.md
 * DESIGN & CONCEPT: Bob van Luijt (@bobvanluijt)
 * CONTACT: hello@semi.technology
 */ // Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"

	models "github.com/semi-technologies/weaviate/entities/models"
)

// GetWellKnownOpenidConfigurationHandlerFunc turns a function with the right signature into a get well known openid configuration handler
type GetWellKnownOpenidConfigurationHandlerFunc func(GetWellKnownOpenidConfigurationParams, *models.Principal) middleware.Responder

// Handle executing the request and returning a response
func (fn GetWellKnownOpenidConfigurationHandlerFunc) Handle(params GetWellKnownOpenidConfigurationParams, principal *models.Principal) middleware.Responder {
	return fn(params, principal)
}

// GetWellKnownOpenidConfigurationHandler interface for that can handle valid get well known openid configuration params
type GetWellKnownOpenidConfigurationHandler interface {
	Handle(GetWellKnownOpenidConfigurationParams, *models.Principal) middleware.Responder
}

// NewGetWellKnownOpenidConfiguration creates a new http.Handler for the get well known openid configuration operation
func NewGetWellKnownOpenidConfiguration(ctx *middleware.Context, handler GetWellKnownOpenidConfigurationHandler) *GetWellKnownOpenidConfiguration {
	return &GetWellKnownOpenidConfiguration{Context: ctx, Handler: handler}
}

/*GetWellKnownOpenidConfiguration swagger:route GET /.well-known/openid-configuration well-known oidc discovery getWellKnownOpenidConfiguration

OIDC discovery information if OIDC auth is enabled

OIDC Discovery page, redirects to the token issuer if one is configured

*/
type GetWellKnownOpenidConfiguration struct {
	Context *middleware.Context
	Handler GetWellKnownOpenidConfigurationHandler
}

func (o *GetWellKnownOpenidConfiguration) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetWellKnownOpenidConfigurationParams()

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
