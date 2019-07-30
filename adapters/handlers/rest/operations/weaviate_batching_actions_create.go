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

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"encoding/json"
	"net/http"
	"strconv"

	errors "github.com/go-openapi/errors"
	middleware "github.com/go-openapi/runtime/middleware"
	strfmt "github.com/go-openapi/strfmt"
	swag "github.com/go-openapi/swag"
	validate "github.com/go-openapi/validate"

	models "github.com/semi-technologies/weaviate/entities/models"
)

// WeaviateBatchingActionsCreateHandlerFunc turns a function with the right signature into a weaviate batching actions create handler
type WeaviateBatchingActionsCreateHandlerFunc func(WeaviateBatchingActionsCreateParams, *models.Principal) middleware.Responder

// Handle executing the request and returning a response
func (fn WeaviateBatchingActionsCreateHandlerFunc) Handle(params WeaviateBatchingActionsCreateParams, principal *models.Principal) middleware.Responder {
	return fn(params, principal)
}

// WeaviateBatchingActionsCreateHandler interface for that can handle valid weaviate batching actions create params
type WeaviateBatchingActionsCreateHandler interface {
	Handle(WeaviateBatchingActionsCreateParams, *models.Principal) middleware.Responder
}

// NewWeaviateBatchingActionsCreate creates a new http.Handler for the weaviate batching actions create operation
func NewWeaviateBatchingActionsCreate(ctx *middleware.Context, handler WeaviateBatchingActionsCreateHandler) *WeaviateBatchingActionsCreate {
	return &WeaviateBatchingActionsCreate{Context: ctx, Handler: handler}
}

/*WeaviateBatchingActionsCreate swagger:route POST /batching/actions batching actions weaviateBatchingActionsCreate

Creates new Actions based on an Action template as a batch.

Register new Actions in bulk. Given meta-data and schema values are validated.

*/
type WeaviateBatchingActionsCreate struct {
	Context *middleware.Context
	Handler WeaviateBatchingActionsCreateHandler
}

func (o *WeaviateBatchingActionsCreate) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewWeaviateBatchingActionsCreateParams()

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

// WeaviateBatchingActionsCreateBody weaviate batching actions create body
// swagger:model WeaviateBatchingActionsCreateBody
type WeaviateBatchingActionsCreateBody struct {

	// actions
	Actions []*models.Action `json:"actions"`

	// Define which fields need to be returned. Default value is ALL
	Fields []*string `json:"fields"`
}

// Validate validates this weaviate batching actions create body
func (o *WeaviateBatchingActionsCreateBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateActions(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateFields(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *WeaviateBatchingActionsCreateBody) validateActions(formats strfmt.Registry) error {

	if swag.IsZero(o.Actions) { // not required
		return nil
	}

	for i := 0; i < len(o.Actions); i++ {
		if swag.IsZero(o.Actions[i]) { // not required
			continue
		}

		if o.Actions[i] != nil {
			if err := o.Actions[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("body" + "." + "actions" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

var weaviateBatchingActionsCreateBodyFieldsItemsEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["ALL","class","schema","id","creationTimeUnix"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		weaviateBatchingActionsCreateBodyFieldsItemsEnum = append(weaviateBatchingActionsCreateBodyFieldsItemsEnum, v)
	}
}

func (o *WeaviateBatchingActionsCreateBody) validateFieldsItemsEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, weaviateBatchingActionsCreateBodyFieldsItemsEnum); err != nil {
		return err
	}
	return nil
}

func (o *WeaviateBatchingActionsCreateBody) validateFields(formats strfmt.Registry) error {

	if swag.IsZero(o.Fields) { // not required
		return nil
	}

	for i := 0; i < len(o.Fields); i++ {
		if swag.IsZero(o.Fields[i]) { // not required
			continue
		}

		// value enum
		if err := o.validateFieldsItemsEnum("body"+"."+"fields"+"."+strconv.Itoa(i), "body", *o.Fields[i]); err != nil {
			return err
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *WeaviateBatchingActionsCreateBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *WeaviateBatchingActionsCreateBody) UnmarshalBinary(b []byte) error {
	var res WeaviateBatchingActionsCreateBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
