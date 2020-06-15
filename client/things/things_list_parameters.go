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

package things

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
	"github.com/go-openapi/swag"
)

// NewThingsListParams creates a new ThingsListParams object
// with the default values initialized.
func NewThingsListParams() *ThingsListParams {
	var ()
	return &ThingsListParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewThingsListParamsWithTimeout creates a new ThingsListParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewThingsListParamsWithTimeout(timeout time.Duration) *ThingsListParams {
	var ()
	return &ThingsListParams{

		timeout: timeout,
	}
}

// NewThingsListParamsWithContext creates a new ThingsListParams object
// with the default values initialized, and the ability to set a context for a request
func NewThingsListParamsWithContext(ctx context.Context) *ThingsListParams {
	var ()
	return &ThingsListParams{

		Context: ctx,
	}
}

// NewThingsListParamsWithHTTPClient creates a new ThingsListParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewThingsListParamsWithHTTPClient(client *http.Client) *ThingsListParams {
	var ()
	return &ThingsListParams{
		HTTPClient: client,
	}
}

/*ThingsListParams contains all the parameters to send to the API endpoint
for the things list operation typically these are written to a http.Request
*/
type ThingsListParams struct {

	/*Limit
	  The maximum number of items to be returned per page. Default value is set in Weaviate config.

	*/
	Limit *int64
	/*Meta
	  Should additional meta information (e.g. about classified properties) be included? Defaults to false.

	*/
	Meta *bool

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the things list params
func (o *ThingsListParams) WithTimeout(timeout time.Duration) *ThingsListParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the things list params
func (o *ThingsListParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the things list params
func (o *ThingsListParams) WithContext(ctx context.Context) *ThingsListParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the things list params
func (o *ThingsListParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the things list params
func (o *ThingsListParams) WithHTTPClient(client *http.Client) *ThingsListParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the things list params
func (o *ThingsListParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithLimit adds the limit to the things list params
func (o *ThingsListParams) WithLimit(limit *int64) *ThingsListParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the things list params
func (o *ThingsListParams) SetLimit(limit *int64) {
	o.Limit = limit
}

// WithMeta adds the meta to the things list params
func (o *ThingsListParams) WithMeta(meta *bool) *ThingsListParams {
	o.SetMeta(meta)
	return o
}

// SetMeta adds the meta to the things list params
func (o *ThingsListParams) SetMeta(meta *bool) {
	o.Meta = meta
}

// WriteToRequest writes these params to a swagger request
func (o *ThingsListParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Limit != nil {

		// query param limit
		var qrLimit int64
		if o.Limit != nil {
			qrLimit = *o.Limit
		}
		qLimit := swag.FormatInt64(qrLimit)
		if qLimit != "" {
			if err := r.SetQueryParam("limit", qLimit); err != nil {
				return err
			}
		}

	}

	if o.Meta != nil {

		// query param meta
		var qrMeta bool
		if o.Meta != nil {
			qrMeta = *o.Meta
		}
		qMeta := swag.FormatBool(qrMeta)
		if qMeta != "" {
			if err := r.SetQueryParam("meta", qMeta); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
