//                           _       _
// __      _____  __ ___   ___  __ _| |_ ___
// \ \ /\ / / _ \/ _` \ \ / / |/ _` | __/ _ \
//  \ V  V /  __/ (_| |\ V /| | (_| | ||  __/
//   \_/\_/ \___|\__,_| \_/ |_|\__,_|\__\___|
//
//  Copyright © 2016 - 2019 SeMI Holding B.V. (registered @ Dutch Chamber of Commerce no 75221632). All rights reserved.
//  LICENSE WEAVIATE OPEN SOURCE: https://www.semi.technology/playbook/playbook/contract-weaviate-OSS.html
//  LICENSE WEAVIATE ENTERPRISE: https://www.semi.technology/playbook/contract-weaviate-enterprise.html
//  CONCEPT: Bob van Luijt (@bobvanluijt)
//  CONTACT: hello@semi.technology
//

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

	models "github.com/semi-technologies/weaviate/entities/models"
)

// NewSchemaThingsPropertiesAddParams creates a new SchemaThingsPropertiesAddParams object
// with the default values initialized.
func NewSchemaThingsPropertiesAddParams() *SchemaThingsPropertiesAddParams {
	var ()
	return &SchemaThingsPropertiesAddParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewSchemaThingsPropertiesAddParamsWithTimeout creates a new SchemaThingsPropertiesAddParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewSchemaThingsPropertiesAddParamsWithTimeout(timeout time.Duration) *SchemaThingsPropertiesAddParams {
	var ()
	return &SchemaThingsPropertiesAddParams{

		timeout: timeout,
	}
}

// NewSchemaThingsPropertiesAddParamsWithContext creates a new SchemaThingsPropertiesAddParams object
// with the default values initialized, and the ability to set a context for a request
func NewSchemaThingsPropertiesAddParamsWithContext(ctx context.Context) *SchemaThingsPropertiesAddParams {
	var ()
	return &SchemaThingsPropertiesAddParams{

		Context: ctx,
	}
}

// NewSchemaThingsPropertiesAddParamsWithHTTPClient creates a new SchemaThingsPropertiesAddParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewSchemaThingsPropertiesAddParamsWithHTTPClient(client *http.Client) *SchemaThingsPropertiesAddParams {
	var ()
	return &SchemaThingsPropertiesAddParams{
		HTTPClient: client,
	}
}

/*SchemaThingsPropertiesAddParams contains all the parameters to send to the API endpoint
for the schema things properties add operation typically these are written to a http.Request
*/
type SchemaThingsPropertiesAddParams struct {

	/*Body*/
	Body *models.Property
	/*ClassName*/
	ClassName string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the schema things properties add params
func (o *SchemaThingsPropertiesAddParams) WithTimeout(timeout time.Duration) *SchemaThingsPropertiesAddParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the schema things properties add params
func (o *SchemaThingsPropertiesAddParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the schema things properties add params
func (o *SchemaThingsPropertiesAddParams) WithContext(ctx context.Context) *SchemaThingsPropertiesAddParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the schema things properties add params
func (o *SchemaThingsPropertiesAddParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the schema things properties add params
func (o *SchemaThingsPropertiesAddParams) WithHTTPClient(client *http.Client) *SchemaThingsPropertiesAddParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the schema things properties add params
func (o *SchemaThingsPropertiesAddParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the schema things properties add params
func (o *SchemaThingsPropertiesAddParams) WithBody(body *models.Property) *SchemaThingsPropertiesAddParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the schema things properties add params
func (o *SchemaThingsPropertiesAddParams) SetBody(body *models.Property) {
	o.Body = body
}

// WithClassName adds the className to the schema things properties add params
func (o *SchemaThingsPropertiesAddParams) WithClassName(className string) *SchemaThingsPropertiesAddParams {
	o.SetClassName(className)
	return o
}

// SetClassName adds the className to the schema things properties add params
func (o *SchemaThingsPropertiesAddParams) SetClassName(className string) {
	o.ClassName = className
}

// WriteToRequest writes these params to a swagger request
func (o *SchemaThingsPropertiesAddParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param className
	if err := r.SetPathParam("className", o.ClassName); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
