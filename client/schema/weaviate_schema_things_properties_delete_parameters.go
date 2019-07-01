/*                          _       _
 *__      _____  __ ___   ___  __ _| |_ ___
 *\ \ /\ / / _ \/ _` \ \ / / |/ _` | __/ _ \
 * \ V  V /  __/ (_| |\ V /| | (_| | ||  __/
 *  \_/\_/ \___|\__,_| \_/ |_|\__,_|\__\___|
 *
 * Copyright © 2016 - 2019 Weaviate. All rights reserved.
 * LICENSE WEAVIATE OPEN SOURCE: https://www.semi.technology/playbook/playbook/contract-weaviate-OSS.html
 * LICENSE WEAVIATE ENTERPRISE: https://www.semi.technology/playbook/contract-weaviate-enterprise.html
 * CONCEPT: Bob van Luijt (@bobvanluijt)
 * CONTACT: hello@semi.technology
 */ // Code generated by go-swagger; DO NOT EDIT.

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

// NewWeaviateSchemaThingsPropertiesDeleteParams creates a new WeaviateSchemaThingsPropertiesDeleteParams object
// with the default values initialized.
func NewWeaviateSchemaThingsPropertiesDeleteParams() *WeaviateSchemaThingsPropertiesDeleteParams {
	var ()
	return &WeaviateSchemaThingsPropertiesDeleteParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewWeaviateSchemaThingsPropertiesDeleteParamsWithTimeout creates a new WeaviateSchemaThingsPropertiesDeleteParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewWeaviateSchemaThingsPropertiesDeleteParamsWithTimeout(timeout time.Duration) *WeaviateSchemaThingsPropertiesDeleteParams {
	var ()
	return &WeaviateSchemaThingsPropertiesDeleteParams{

		timeout: timeout,
	}
}

// NewWeaviateSchemaThingsPropertiesDeleteParamsWithContext creates a new WeaviateSchemaThingsPropertiesDeleteParams object
// with the default values initialized, and the ability to set a context for a request
func NewWeaviateSchemaThingsPropertiesDeleteParamsWithContext(ctx context.Context) *WeaviateSchemaThingsPropertiesDeleteParams {
	var ()
	return &WeaviateSchemaThingsPropertiesDeleteParams{

		Context: ctx,
	}
}

// NewWeaviateSchemaThingsPropertiesDeleteParamsWithHTTPClient creates a new WeaviateSchemaThingsPropertiesDeleteParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewWeaviateSchemaThingsPropertiesDeleteParamsWithHTTPClient(client *http.Client) *WeaviateSchemaThingsPropertiesDeleteParams {
	var ()
	return &WeaviateSchemaThingsPropertiesDeleteParams{
		HTTPClient: client,
	}
}

/*WeaviateSchemaThingsPropertiesDeleteParams contains all the parameters to send to the API endpoint
for the weaviate schema things properties delete operation typically these are written to a http.Request
*/
type WeaviateSchemaThingsPropertiesDeleteParams struct {

	/*ClassName*/
	ClassName string
	/*PropertyName*/
	PropertyName string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the weaviate schema things properties delete params
func (o *WeaviateSchemaThingsPropertiesDeleteParams) WithTimeout(timeout time.Duration) *WeaviateSchemaThingsPropertiesDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the weaviate schema things properties delete params
func (o *WeaviateSchemaThingsPropertiesDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the weaviate schema things properties delete params
func (o *WeaviateSchemaThingsPropertiesDeleteParams) WithContext(ctx context.Context) *WeaviateSchemaThingsPropertiesDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the weaviate schema things properties delete params
func (o *WeaviateSchemaThingsPropertiesDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the weaviate schema things properties delete params
func (o *WeaviateSchemaThingsPropertiesDeleteParams) WithHTTPClient(client *http.Client) *WeaviateSchemaThingsPropertiesDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the weaviate schema things properties delete params
func (o *WeaviateSchemaThingsPropertiesDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithClassName adds the className to the weaviate schema things properties delete params
func (o *WeaviateSchemaThingsPropertiesDeleteParams) WithClassName(className string) *WeaviateSchemaThingsPropertiesDeleteParams {
	o.SetClassName(className)
	return o
}

// SetClassName adds the className to the weaviate schema things properties delete params
func (o *WeaviateSchemaThingsPropertiesDeleteParams) SetClassName(className string) {
	o.ClassName = className
}

// WithPropertyName adds the propertyName to the weaviate schema things properties delete params
func (o *WeaviateSchemaThingsPropertiesDeleteParams) WithPropertyName(propertyName string) *WeaviateSchemaThingsPropertiesDeleteParams {
	o.SetPropertyName(propertyName)
	return o
}

// SetPropertyName adds the propertyName to the weaviate schema things properties delete params
func (o *WeaviateSchemaThingsPropertiesDeleteParams) SetPropertyName(propertyName string) {
	o.PropertyName = propertyName
}

// WriteToRequest writes these params to a swagger request
func (o *WeaviateSchemaThingsPropertiesDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param className
	if err := r.SetPathParam("className", o.ClassName); err != nil {
		return err
	}

	// path param propertyName
	if err := r.SetPathParam("propertyName", o.PropertyName); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
