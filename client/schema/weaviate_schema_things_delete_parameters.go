/*                          _       _
 *__      _____  __ ___   ___  __ _| |_ ___
 *\ \ /\ / / _ \/ _` \ \ / / |/ _` | __/ _ \
 * \ V  V /  __/ (_| |\ V /| | (_| | ||  __/
 *  \_/\_/ \___|\__,_| \_/ |_|\__,_|\__\___|
 *
 * Copyright © 2016 - 2018 Weaviate. All rights reserved.
 * LICENSE: https://github.com/creativesoftwarefdn/weaviate/blob/develop/LICENSE.md
 * AUTHOR: Bob van Luijt (bob@kub.design)
 * See www.creativesoftwarefdn.org for details
 * Contact: @CreativeSofwFdn / bob@kub.design
 */
// Code generated by go-swagger; DO NOT EDIT.

package schema

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewWeaviateSchemaThingsDeleteParams creates a new WeaviateSchemaThingsDeleteParams object
// with the default values initialized.
func NewWeaviateSchemaThingsDeleteParams() *WeaviateSchemaThingsDeleteParams {
	var ()
	return &WeaviateSchemaThingsDeleteParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewWeaviateSchemaThingsDeleteParamsWithTimeout creates a new WeaviateSchemaThingsDeleteParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewWeaviateSchemaThingsDeleteParamsWithTimeout(timeout time.Duration) *WeaviateSchemaThingsDeleteParams {
	var ()
	return &WeaviateSchemaThingsDeleteParams{

		timeout: timeout,
	}
}

// NewWeaviateSchemaThingsDeleteParamsWithContext creates a new WeaviateSchemaThingsDeleteParams object
// with the default values initialized, and the ability to set a context for a request
func NewWeaviateSchemaThingsDeleteParamsWithContext(ctx context.Context) *WeaviateSchemaThingsDeleteParams {
	var ()
	return &WeaviateSchemaThingsDeleteParams{

		Context: ctx,
	}
}

// NewWeaviateSchemaThingsDeleteParamsWithHTTPClient creates a new WeaviateSchemaThingsDeleteParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewWeaviateSchemaThingsDeleteParamsWithHTTPClient(client *http.Client) *WeaviateSchemaThingsDeleteParams {
	var ()
	return &WeaviateSchemaThingsDeleteParams{
		HTTPClient: client,
	}
}

/*WeaviateSchemaThingsDeleteParams contains all the parameters to send to the API endpoint
for the weaviate schema things delete operation typically these are written to a http.Request
*/
type WeaviateSchemaThingsDeleteParams struct {

	/*ClassName*/
	ClassName string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the weaviate schema things delete params
func (o *WeaviateSchemaThingsDeleteParams) WithTimeout(timeout time.Duration) *WeaviateSchemaThingsDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the weaviate schema things delete params
func (o *WeaviateSchemaThingsDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the weaviate schema things delete params
func (o *WeaviateSchemaThingsDeleteParams) WithContext(ctx context.Context) *WeaviateSchemaThingsDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the weaviate schema things delete params
func (o *WeaviateSchemaThingsDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the weaviate schema things delete params
func (o *WeaviateSchemaThingsDeleteParams) WithHTTPClient(client *http.Client) *WeaviateSchemaThingsDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the weaviate schema things delete params
func (o *WeaviateSchemaThingsDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithClassName adds the className to the weaviate schema things delete params
func (o *WeaviateSchemaThingsDeleteParams) WithClassName(className string) *WeaviateSchemaThingsDeleteParams {
	o.SetClassName(className)
	return o
}

// SetClassName adds the className to the weaviate schema things delete params
func (o *WeaviateSchemaThingsDeleteParams) SetClassName(className string) {
	o.ClassName = className
}

// WriteToRequest writes these params to a swagger request
func (o *WeaviateSchemaThingsDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param className
	if err := r.SetPathParam("className", o.ClassName); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
