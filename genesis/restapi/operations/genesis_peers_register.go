// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// GenesisPeersRegisterHandlerFunc turns a function with the right signature into a genesis peers register handler
type GenesisPeersRegisterHandlerFunc func(GenesisPeersRegisterParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GenesisPeersRegisterHandlerFunc) Handle(params GenesisPeersRegisterParams) middleware.Responder {
	return fn(params)
}

// GenesisPeersRegisterHandler interface for that can handle valid genesis peers register params
type GenesisPeersRegisterHandler interface {
	Handle(GenesisPeersRegisterParams) middleware.Responder
}

// NewGenesisPeersRegister creates a new http.Handler for the genesis peers register operation
func NewGenesisPeersRegister(ctx *middleware.Context, handler GenesisPeersRegisterHandler) *GenesisPeersRegister {
	return &GenesisPeersRegister{Context: ctx, Handler: handler}
}

/*GenesisPeersRegister swagger:route POST /peers/register genesisPeersRegister

Register a new Weaviate peer in the network

*/
type GenesisPeersRegister struct {
	Context *middleware.Context
	Handler GenesisPeersRegisterHandler
}

func (o *GenesisPeersRegister) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGenesisPeersRegisterParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
