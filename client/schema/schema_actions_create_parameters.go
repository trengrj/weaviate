//                           _       _
// __      _____  __ ___   ___  __ _| |_ ___
// \ \ /\ / / _ \/ _` \ \ / / |/ _` | __/ _ \
//  \ V  V /  __/ (_| |\ V /| | (_| | ||  __/
//   \_/\_/ \___|\__,_| \_/ |_|\__,_|\__\___|
//
//  Copyright © 2016 - 2019 Weaviate. All rights reserved.
//  LICENSE: https://github.com/semi-technologies/weaviate/blob/develop/LICENSE.md
//  DESIGN & CONCEPT: Bob van Luijt (@bobvanluijt)
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

// NewSchemaActionsCreateParams creates a new SchemaActionsCreateParams object
// with the default values initialized.
func NewSchemaActionsCreateParams() *SchemaActionsCreateParams {
	var ()
	return &SchemaActionsCreateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewSchemaActionsCreateParamsWithTimeout creates a new SchemaActionsCreateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewSchemaActionsCreateParamsWithTimeout(timeout time.Duration) *SchemaActionsCreateParams {
	var ()
	return &SchemaActionsCreateParams{

		timeout: timeout,
	}
}

// NewSchemaActionsCreateParamsWithContext creates a new SchemaActionsCreateParams object
// with the default values initialized, and the ability to set a context for a request
func NewSchemaActionsCreateParamsWithContext(ctx context.Context) *SchemaActionsCreateParams {
	var ()
	return &SchemaActionsCreateParams{

		Context: ctx,
	}
}

// NewSchemaActionsCreateParamsWithHTTPClient creates a new SchemaActionsCreateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewSchemaActionsCreateParamsWithHTTPClient(client *http.Client) *SchemaActionsCreateParams {
	var ()
	return &SchemaActionsCreateParams{
		HTTPClient: client,
	}
}

/*SchemaActionsCreateParams contains all the parameters to send to the API endpoint
for the schema actions create operation typically these are written to a http.Request
*/
type SchemaActionsCreateParams struct {

	/*ActionClass*/
	ActionClass *models.Class

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the schema actions create params
func (o *SchemaActionsCreateParams) WithTimeout(timeout time.Duration) *SchemaActionsCreateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the schema actions create params
func (o *SchemaActionsCreateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the schema actions create params
func (o *SchemaActionsCreateParams) WithContext(ctx context.Context) *SchemaActionsCreateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the schema actions create params
func (o *SchemaActionsCreateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the schema actions create params
func (o *SchemaActionsCreateParams) WithHTTPClient(client *http.Client) *SchemaActionsCreateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the schema actions create params
func (o *SchemaActionsCreateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithActionClass adds the actionClass to the schema actions create params
func (o *SchemaActionsCreateParams) WithActionClass(actionClass *models.Class) *SchemaActionsCreateParams {
	o.SetActionClass(actionClass)
	return o
}

// SetActionClass adds the actionClass to the schema actions create params
func (o *SchemaActionsCreateParams) SetActionClass(actionClass *models.Class) {
	o.ActionClass = actionClass
}

// WriteToRequest writes these params to a swagger request
func (o *SchemaActionsCreateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.ActionClass != nil {
		if err := r.SetBodyParam(o.ActionClass); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
