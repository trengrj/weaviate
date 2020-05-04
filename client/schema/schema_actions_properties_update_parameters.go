//                           _       _
// __      _____  __ ___   ___  __ _| |_ ___
// \ \ /\ / / _ \/ _` \ \ / / |/ _` | __/ _ \
//  \ V  V /  __/ (_| |\ V /| | (_| | ||  __/
//   \_/\_/ \___|\__,_| \_/ |_|\__,_|\__\___|
//
//  Copyright © 2016 - 2020 SeMI Holding B.V. (registered @ Dutch Chamber of Commerce no 75221632). All rights reserved.
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

// NewSchemaActionsPropertiesUpdateParams creates a new SchemaActionsPropertiesUpdateParams object
// with the default values initialized.
func NewSchemaActionsPropertiesUpdateParams() *SchemaActionsPropertiesUpdateParams {
	var ()
	return &SchemaActionsPropertiesUpdateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewSchemaActionsPropertiesUpdateParamsWithTimeout creates a new SchemaActionsPropertiesUpdateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewSchemaActionsPropertiesUpdateParamsWithTimeout(timeout time.Duration) *SchemaActionsPropertiesUpdateParams {
	var ()
	return &SchemaActionsPropertiesUpdateParams{

		timeout: timeout,
	}
}

// NewSchemaActionsPropertiesUpdateParamsWithContext creates a new SchemaActionsPropertiesUpdateParams object
// with the default values initialized, and the ability to set a context for a request
func NewSchemaActionsPropertiesUpdateParamsWithContext(ctx context.Context) *SchemaActionsPropertiesUpdateParams {
	var ()
	return &SchemaActionsPropertiesUpdateParams{

		Context: ctx,
	}
}

// NewSchemaActionsPropertiesUpdateParamsWithHTTPClient creates a new SchemaActionsPropertiesUpdateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewSchemaActionsPropertiesUpdateParamsWithHTTPClient(client *http.Client) *SchemaActionsPropertiesUpdateParams {
	var ()
	return &SchemaActionsPropertiesUpdateParams{
		HTTPClient: client,
	}
}

/*SchemaActionsPropertiesUpdateParams contains all the parameters to send to the API endpoint
for the schema actions properties update operation typically these are written to a http.Request
*/
type SchemaActionsPropertiesUpdateParams struct {

	/*Body*/
	Body *models.Property
	/*ClassName*/
	ClassName string
	/*PropertyName*/
	PropertyName string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the schema actions properties update params
func (o *SchemaActionsPropertiesUpdateParams) WithTimeout(timeout time.Duration) *SchemaActionsPropertiesUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the schema actions properties update params
func (o *SchemaActionsPropertiesUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the schema actions properties update params
func (o *SchemaActionsPropertiesUpdateParams) WithContext(ctx context.Context) *SchemaActionsPropertiesUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the schema actions properties update params
func (o *SchemaActionsPropertiesUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the schema actions properties update params
func (o *SchemaActionsPropertiesUpdateParams) WithHTTPClient(client *http.Client) *SchemaActionsPropertiesUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the schema actions properties update params
func (o *SchemaActionsPropertiesUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the schema actions properties update params
func (o *SchemaActionsPropertiesUpdateParams) WithBody(body *models.Property) *SchemaActionsPropertiesUpdateParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the schema actions properties update params
func (o *SchemaActionsPropertiesUpdateParams) SetBody(body *models.Property) {
	o.Body = body
}

// WithClassName adds the className to the schema actions properties update params
func (o *SchemaActionsPropertiesUpdateParams) WithClassName(className string) *SchemaActionsPropertiesUpdateParams {
	o.SetClassName(className)
	return o
}

// SetClassName adds the className to the schema actions properties update params
func (o *SchemaActionsPropertiesUpdateParams) SetClassName(className string) {
	o.ClassName = className
}

// WithPropertyName adds the propertyName to the schema actions properties update params
func (o *SchemaActionsPropertiesUpdateParams) WithPropertyName(propertyName string) *SchemaActionsPropertiesUpdateParams {
	o.SetPropertyName(propertyName)
	return o
}

// SetPropertyName adds the propertyName to the schema actions properties update params
func (o *SchemaActionsPropertiesUpdateParams) SetPropertyName(propertyName string) {
	o.PropertyName = propertyName
}

// WriteToRequest writes these params to a swagger request
func (o *SchemaActionsPropertiesUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param propertyName
	if err := r.SetPathParam("propertyName", o.PropertyName); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
