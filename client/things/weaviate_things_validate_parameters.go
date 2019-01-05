/*                          _       _
 *__      _____  __ ___   ___  __ _| |_ ___
 *\ \ /\ / / _ \/ _` \ \ / / |/ _` | __/ _ \
 * \ V  V /  __/ (_| |\ V /| | (_| | ||  __/
 *  \_/\_/ \___|\__,_| \_/ |_|\__,_|\__\___|
 *
 * Copyright © 2016 - 2019 Weaviate. All rights reserved.
 * LICENSE: https://github.com/creativesoftwarefdn/weaviate/blob/develop/LICENSE.md
 * DESIGN: Bob van Luijt (bob@k10y.co)
 */
// Code generated by go-swagger; DO NOT EDIT.

package things

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/creativesoftwarefdn/weaviate/models"
)

// NewWeaviateThingsValidateParams creates a new WeaviateThingsValidateParams object
// with the default values initialized.
func NewWeaviateThingsValidateParams() *WeaviateThingsValidateParams {
	var ()
	return &WeaviateThingsValidateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewWeaviateThingsValidateParamsWithTimeout creates a new WeaviateThingsValidateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewWeaviateThingsValidateParamsWithTimeout(timeout time.Duration) *WeaviateThingsValidateParams {
	var ()
	return &WeaviateThingsValidateParams{

		timeout: timeout,
	}
}

// NewWeaviateThingsValidateParamsWithContext creates a new WeaviateThingsValidateParams object
// with the default values initialized, and the ability to set a context for a request
func NewWeaviateThingsValidateParamsWithContext(ctx context.Context) *WeaviateThingsValidateParams {
	var ()
	return &WeaviateThingsValidateParams{

		Context: ctx,
	}
}

// NewWeaviateThingsValidateParamsWithHTTPClient creates a new WeaviateThingsValidateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewWeaviateThingsValidateParamsWithHTTPClient(client *http.Client) *WeaviateThingsValidateParams {
	var ()
	return &WeaviateThingsValidateParams{
		HTTPClient: client,
	}
}

/*WeaviateThingsValidateParams contains all the parameters to send to the API endpoint
for the weaviate things validate operation typically these are written to a http.Request
*/
type WeaviateThingsValidateParams struct {

	/*Body*/
	Body *models.ThingCreate

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the weaviate things validate params
func (o *WeaviateThingsValidateParams) WithTimeout(timeout time.Duration) *WeaviateThingsValidateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the weaviate things validate params
func (o *WeaviateThingsValidateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the weaviate things validate params
func (o *WeaviateThingsValidateParams) WithContext(ctx context.Context) *WeaviateThingsValidateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the weaviate things validate params
func (o *WeaviateThingsValidateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the weaviate things validate params
func (o *WeaviateThingsValidateParams) WithHTTPClient(client *http.Client) *WeaviateThingsValidateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the weaviate things validate params
func (o *WeaviateThingsValidateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the weaviate things validate params
func (o *WeaviateThingsValidateParams) WithBody(body *models.ThingCreate) *WeaviateThingsValidateParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the weaviate things validate params
func (o *WeaviateThingsValidateParams) SetBody(body *models.ThingCreate) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *WeaviateThingsValidateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
