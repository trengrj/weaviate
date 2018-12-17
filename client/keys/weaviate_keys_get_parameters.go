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

package keys

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
)

// NewWeaviateKeysGetParams creates a new WeaviateKeysGetParams object
// with the default values initialized.
func NewWeaviateKeysGetParams() *WeaviateKeysGetParams {
	var ()
	return &WeaviateKeysGetParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewWeaviateKeysGetParamsWithTimeout creates a new WeaviateKeysGetParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewWeaviateKeysGetParamsWithTimeout(timeout time.Duration) *WeaviateKeysGetParams {
	var ()
	return &WeaviateKeysGetParams{

		timeout: timeout,
	}
}

// NewWeaviateKeysGetParamsWithContext creates a new WeaviateKeysGetParams object
// with the default values initialized, and the ability to set a context for a request
func NewWeaviateKeysGetParamsWithContext(ctx context.Context) *WeaviateKeysGetParams {
	var ()
	return &WeaviateKeysGetParams{

		Context: ctx,
	}
}

// NewWeaviateKeysGetParamsWithHTTPClient creates a new WeaviateKeysGetParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewWeaviateKeysGetParamsWithHTTPClient(client *http.Client) *WeaviateKeysGetParams {
	var ()
	return &WeaviateKeysGetParams{
		HTTPClient: client,
	}
}

/*WeaviateKeysGetParams contains all the parameters to send to the API endpoint
for the weaviate keys get operation typically these are written to a http.Request
*/
type WeaviateKeysGetParams struct {

	/*KeyID
	  Unique ID of the key.

	*/
	KeyID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the weaviate keys get params
func (o *WeaviateKeysGetParams) WithTimeout(timeout time.Duration) *WeaviateKeysGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the weaviate keys get params
func (o *WeaviateKeysGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the weaviate keys get params
func (o *WeaviateKeysGetParams) WithContext(ctx context.Context) *WeaviateKeysGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the weaviate keys get params
func (o *WeaviateKeysGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the weaviate keys get params
func (o *WeaviateKeysGetParams) WithHTTPClient(client *http.Client) *WeaviateKeysGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the weaviate keys get params
func (o *WeaviateKeysGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithKeyID adds the keyID to the weaviate keys get params
func (o *WeaviateKeysGetParams) WithKeyID(keyID strfmt.UUID) *WeaviateKeysGetParams {
	o.SetKeyID(keyID)
	return o
}

// SetKeyID adds the keyId to the weaviate keys get params
func (o *WeaviateKeysGetParams) SetKeyID(keyID strfmt.UUID) {
	o.KeyID = keyID
}

// WriteToRequest writes these params to a swagger request
func (o *WeaviateKeysGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param keyId
	if err := r.SetPathParam("keyId", o.KeyID.String()); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
