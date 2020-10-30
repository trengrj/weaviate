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

package contextionary_api

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
)

// NewC11yWordsParams creates a new C11yWordsParams object
// no default values defined in spec.
func NewC11yWordsParams() C11yWordsParams {
	return C11yWordsParams{}
}

// C11yWordsParams contains all the bound params for the c11y words operation
// typically these are obtained from a http.Request
//
// swagger:parameters c11y.words
type C11yWordsParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*CamelCase list of words to validate.
	  Required: true
	  In: path
	*/
	Words string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewC11yWordsParams() beforehand.
func (o *C11yWordsParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	rWords, rhkWords, _ := route.Params.GetOK("words")
	if err := o.bindWords(rWords, rhkWords, route.Formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindWords binds and validates parameter Words from path.
func (o *C11yWordsParams) bindWords(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route

	o.Words = raw

	return nil
}
