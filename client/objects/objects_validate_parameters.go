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

// NewObjectsValidateParams creates a new ObjectsValidateParams object
// with the default values initialized.
func NewObjectsValidateParams() *ObjectsValidateParams {
	var ()
	return &ObjectsValidateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewObjectsValidateParamsWithTimeout creates a new ObjectsValidateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewObjectsValidateParamsWithTimeout(timeout time.Duration) *ObjectsValidateParams {
	var ()
	return &ObjectsValidateParams{

		timeout: timeout,
	}
}

// NewObjectsValidateParamsWithContext creates a new ObjectsValidateParams object
// with the default values initialized, and the ability to set a context for a request
func NewObjectsValidateParamsWithContext(ctx context.Context) *ObjectsValidateParams {
	var ()
	return &ObjectsValidateParams{

		Context: ctx,
	}
}

// NewObjectsValidateParamsWithHTTPClient creates a new ObjectsValidateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewObjectsValidateParamsWithHTTPClient(client *http.Client) *ObjectsValidateParams {
	var ()
	return &ObjectsValidateParams{
		HTTPClient: client,
	}
}

/*ObjectsValidateParams contains all the parameters to send to the API endpoint
for the objects validate operation typically these are written to a http.Request
*/
type ObjectsValidateParams struct {

	/*Body*/
	Body *models.Object

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the objects validate params
func (o *ObjectsValidateParams) WithTimeout(timeout time.Duration) *ObjectsValidateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the objects validate params
func (o *ObjectsValidateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the objects validate params
func (o *ObjectsValidateParams) WithContext(ctx context.Context) *ObjectsValidateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the objects validate params
func (o *ObjectsValidateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the objects validate params
func (o *ObjectsValidateParams) WithHTTPClient(client *http.Client) *ObjectsValidateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the objects validate params
func (o *ObjectsValidateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the objects validate params
func (o *ObjectsValidateParams) WithBody(body *models.Object) *ObjectsValidateParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the objects validate params
func (o *ObjectsValidateParams) SetBody(body *models.Object) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *ObjectsValidateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
