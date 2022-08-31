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
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// NewObjectsClassGetParams creates a new ObjectsClassGetParams object
// with the default values initialized.
func NewObjectsClassGetParams() *ObjectsClassGetParams {
	var ()
	return &ObjectsClassGetParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewObjectsClassGetParamsWithTimeout creates a new ObjectsClassGetParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewObjectsClassGetParamsWithTimeout(timeout time.Duration) *ObjectsClassGetParams {
	var ()
	return &ObjectsClassGetParams{

		timeout: timeout,
	}
}

// NewObjectsClassGetParamsWithContext creates a new ObjectsClassGetParams object
// with the default values initialized, and the ability to set a context for a request
func NewObjectsClassGetParamsWithContext(ctx context.Context) *ObjectsClassGetParams {
	var ()
	return &ObjectsClassGetParams{

		Context: ctx,
	}
}

// NewObjectsClassGetParamsWithHTTPClient creates a new ObjectsClassGetParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewObjectsClassGetParamsWithHTTPClient(client *http.Client) *ObjectsClassGetParams {
	var ()
	return &ObjectsClassGetParams{
		HTTPClient: client,
	}
}

/*ObjectsClassGetParams contains all the parameters to send to the API endpoint
for the objects class get operation typically these are written to a http.Request
*/
type ObjectsClassGetParams struct {

	/*ClassName*/
	ClassName string
	/*ID
	  Unique ID of the Object.

	*/
	ID strfmt.UUID
	/*Include
	  Include additional information, such as classification infos. Allowed values include: classification, vector, interpretation

	*/
	Include *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the objects class get params
func (o *ObjectsClassGetParams) WithTimeout(timeout time.Duration) *ObjectsClassGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the objects class get params
func (o *ObjectsClassGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the objects class get params
func (o *ObjectsClassGetParams) WithContext(ctx context.Context) *ObjectsClassGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the objects class get params
func (o *ObjectsClassGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the objects class get params
func (o *ObjectsClassGetParams) WithHTTPClient(client *http.Client) *ObjectsClassGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the objects class get params
func (o *ObjectsClassGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithClassName adds the className to the objects class get params
func (o *ObjectsClassGetParams) WithClassName(className string) *ObjectsClassGetParams {
	o.SetClassName(className)
	return o
}

// SetClassName adds the className to the objects class get params
func (o *ObjectsClassGetParams) SetClassName(className string) {
	o.ClassName = className
}

// WithID adds the id to the objects class get params
func (o *ObjectsClassGetParams) WithID(id strfmt.UUID) *ObjectsClassGetParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the objects class get params
func (o *ObjectsClassGetParams) SetID(id strfmt.UUID) {
	o.ID = id
}

// WithInclude adds the include to the objects class get params
func (o *ObjectsClassGetParams) WithInclude(include *string) *ObjectsClassGetParams {
	o.SetInclude(include)
	return o
}

// SetInclude adds the include to the objects class get params
func (o *ObjectsClassGetParams) SetInclude(include *string) {
	o.Include = include
}

// WriteToRequest writes these params to a swagger request
func (o *ObjectsClassGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param className
	if err := r.SetPathParam("className", o.ClassName); err != nil {
		return err
	}

	// path param id
	if err := r.SetPathParam("id", o.ID.String()); err != nil {
		return err
	}

	if o.Include != nil {

		// query param include
		var qrInclude string
		if o.Include != nil {
			qrInclude = *o.Include
		}
		qInclude := qrInclude
		if qInclude != "" {
			if err := r.SetQueryParam("include", qInclude); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
