//                           _       _
// __      _____  __ ___   ___  __ _| |_ ___
// \ \ /\ / / _ \/ _` \ \ / / |/ _` | __/ _ \
//  \ V  V /  __/ (_| |\ V /| | (_| | ||  __/
//   \_/\_/ \___|\__,_| \_/ |_|\__,_|\__\___|
//
//  Copyright © 2016 - 2021 SeMI Technologies B.V. All rights reserved.
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

	"github.com/semi-technologies/weaviate/entities/models"
)

// NewObjectsCreateParams creates a new ObjectsCreateParams object
// with the default values initialized.
func NewObjectsCreateParams() *ObjectsCreateParams {
	var ()
	return &ObjectsCreateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewObjectsCreateParamsWithTimeout creates a new ObjectsCreateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewObjectsCreateParamsWithTimeout(timeout time.Duration) *ObjectsCreateParams {
	var ()
	return &ObjectsCreateParams{

		timeout: timeout,
	}
}

// NewObjectsCreateParamsWithContext creates a new ObjectsCreateParams object
// with the default values initialized, and the ability to set a context for a request
func NewObjectsCreateParamsWithContext(ctx context.Context) *ObjectsCreateParams {
	var ()
	return &ObjectsCreateParams{

		Context: ctx,
	}
}

// NewObjectsCreateParamsWithHTTPClient creates a new ObjectsCreateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewObjectsCreateParamsWithHTTPClient(client *http.Client) *ObjectsCreateParams {
	var ()
	return &ObjectsCreateParams{
		HTTPClient: client,
	}
}

/*ObjectsCreateParams contains all the parameters to send to the API endpoint
for the objects create operation typically these are written to a http.Request
*/
type ObjectsCreateParams struct {

	/*Body*/
	Body *models.Object

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the objects create params
func (o *ObjectsCreateParams) WithTimeout(timeout time.Duration) *ObjectsCreateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the objects create params
func (o *ObjectsCreateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the objects create params
func (o *ObjectsCreateParams) WithContext(ctx context.Context) *ObjectsCreateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the objects create params
func (o *ObjectsCreateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the objects create params
func (o *ObjectsCreateParams) WithHTTPClient(client *http.Client) *ObjectsCreateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the objects create params
func (o *ObjectsCreateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the objects create params
func (o *ObjectsCreateParams) WithBody(body *models.Object) *ObjectsCreateParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the objects create params
func (o *ObjectsCreateParams) SetBody(body *models.Object) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *ObjectsCreateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
