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

package graphql

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

// NewGraphqlBatchParams creates a new GraphqlBatchParams object
// with the default values initialized.
func NewGraphqlBatchParams() *GraphqlBatchParams {
	var ()
	return &GraphqlBatchParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGraphqlBatchParamsWithTimeout creates a new GraphqlBatchParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGraphqlBatchParamsWithTimeout(timeout time.Duration) *GraphqlBatchParams {
	var ()
	return &GraphqlBatchParams{

		timeout: timeout,
	}
}

// NewGraphqlBatchParamsWithContext creates a new GraphqlBatchParams object
// with the default values initialized, and the ability to set a context for a request
func NewGraphqlBatchParamsWithContext(ctx context.Context) *GraphqlBatchParams {
	var ()
	return &GraphqlBatchParams{

		Context: ctx,
	}
}

// NewGraphqlBatchParamsWithHTTPClient creates a new GraphqlBatchParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGraphqlBatchParamsWithHTTPClient(client *http.Client) *GraphqlBatchParams {
	var ()
	return &GraphqlBatchParams{
		HTTPClient: client,
	}
}

/*GraphqlBatchParams contains all the parameters to send to the API endpoint
for the graphql batch operation typically these are written to a http.Request
*/
type GraphqlBatchParams struct {

	/*Body
	  The GraphQL queries.

	*/
	Body models.GraphQLQueries

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the graphql batch params
func (o *GraphqlBatchParams) WithTimeout(timeout time.Duration) *GraphqlBatchParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the graphql batch params
func (o *GraphqlBatchParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the graphql batch params
func (o *GraphqlBatchParams) WithContext(ctx context.Context) *GraphqlBatchParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the graphql batch params
func (o *GraphqlBatchParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the graphql batch params
func (o *GraphqlBatchParams) WithHTTPClient(client *http.Client) *GraphqlBatchParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the graphql batch params
func (o *GraphqlBatchParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the graphql batch params
func (o *GraphqlBatchParams) WithBody(body models.GraphQLQueries) *GraphqlBatchParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the graphql batch params
func (o *GraphqlBatchParams) SetBody(body models.GraphQLQueries) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *GraphqlBatchParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
