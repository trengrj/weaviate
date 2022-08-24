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
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// NewObjectsListParams creates a new ObjectsListParams object
// with the default values initialized.
func NewObjectsListParams() ObjectsListParams {
	// initialize parameters with default values

	offsetDefault := int64(0)

	return ObjectsListParams{
		Offset: &offsetDefault,
	}
}

// ObjectsListParams contains all the bound params for the objects list operation
// typically these are obtained from a http.Request
//
// swagger:parameters objects.list
type ObjectsListParams struct {
	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*Class parameter specifies the class from which to query objects
	  In: query
	*/
	Class *string
	/*Include additional information, such as classification infos. Allowed values include: classification, vector, interpretation
	  In: query
	*/
	Include *string
	/*The maximum number of items to be returned per page. Default value is set in Weaviate config.
	  In: query
	*/
	Limit *int64
	/*The starting index of the result window. Default value is 0.
	  In: query
	  Default: 0
	*/
	Offset *int64
	/*Order parameter to tell how to order (asc or desc) data within given field
	  In: query
	*/
	Order *string
	/*Sort parameter to pass an information about the names of the sort fields
	  In: query
	*/
	Sort *string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewObjectsListParams() beforehand.
func (o *ObjectsListParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	qs := runtime.Values(r.URL.Query())

	qClass, qhkClass, _ := qs.GetOK("class")
	if err := o.bindClass(qClass, qhkClass, route.Formats); err != nil {
		res = append(res, err)
	}

	qInclude, qhkInclude, _ := qs.GetOK("include")
	if err := o.bindInclude(qInclude, qhkInclude, route.Formats); err != nil {
		res = append(res, err)
	}

	qLimit, qhkLimit, _ := qs.GetOK("limit")
	if err := o.bindLimit(qLimit, qhkLimit, route.Formats); err != nil {
		res = append(res, err)
	}

	qOffset, qhkOffset, _ := qs.GetOK("offset")
	if err := o.bindOffset(qOffset, qhkOffset, route.Formats); err != nil {
		res = append(res, err)
	}

	qOrder, qhkOrder, _ := qs.GetOK("order")
	if err := o.bindOrder(qOrder, qhkOrder, route.Formats); err != nil {
		res = append(res, err)
	}

	qSort, qhkSort, _ := qs.GetOK("sort")
	if err := o.bindSort(qSort, qhkSort, route.Formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindClass binds and validates parameter Class from query.
func (o *ObjectsListParams) bindClass(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false
	if raw == "" { // empty values pass all other validations
		return nil
	}

	o.Class = &raw

	return nil
}

// bindInclude binds and validates parameter Include from query.
func (o *ObjectsListParams) bindInclude(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false
	if raw == "" { // empty values pass all other validations
		return nil
	}

	o.Include = &raw

	return nil
}

// bindLimit binds and validates parameter Limit from query.
func (o *ObjectsListParams) bindLimit(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false
	if raw == "" { // empty values pass all other validations
		return nil
	}

	value, err := swag.ConvertInt64(raw)
	if err != nil {
		return errors.InvalidType("limit", "query", "int64", raw)
	}
	o.Limit = &value

	return nil
}

// bindOffset binds and validates parameter Offset from query.
func (o *ObjectsListParams) bindOffset(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false
	if raw == "" { // empty values pass all other validations
		// Default values have been previously initialized by NewObjectsListParams()
		return nil
	}

	value, err := swag.ConvertInt64(raw)
	if err != nil {
		return errors.InvalidType("offset", "query", "int64", raw)
	}
	o.Offset = &value

	return nil
}

// bindOrder binds and validates parameter Order from query.
func (o *ObjectsListParams) bindOrder(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false
	if raw == "" { // empty values pass all other validations
		return nil
	}

	o.Order = &raw

	return nil
}

// bindSort binds and validates parameter Sort from query.
func (o *ObjectsListParams) bindSort(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false
	if raw == "" { // empty values pass all other validations
		return nil
	}

	o.Sort = &raw

	return nil
}
