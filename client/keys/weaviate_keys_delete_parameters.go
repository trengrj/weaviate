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
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewWeaviateKeysDeleteParams creates a new WeaviateKeysDeleteParams object
// with the default values initialized.
func NewWeaviateKeysDeleteParams() *WeaviateKeysDeleteParams {
	var ()
	return &WeaviateKeysDeleteParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewWeaviateKeysDeleteParamsWithTimeout creates a new WeaviateKeysDeleteParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewWeaviateKeysDeleteParamsWithTimeout(timeout time.Duration) *WeaviateKeysDeleteParams {
	var ()
	return &WeaviateKeysDeleteParams{

		timeout: timeout,
	}
}

// NewWeaviateKeysDeleteParamsWithContext creates a new WeaviateKeysDeleteParams object
// with the default values initialized, and the ability to set a context for a request
func NewWeaviateKeysDeleteParamsWithContext(ctx context.Context) *WeaviateKeysDeleteParams {
	var ()
	return &WeaviateKeysDeleteParams{

		Context: ctx,
	}
}

// NewWeaviateKeysDeleteParamsWithHTTPClient creates a new WeaviateKeysDeleteParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewWeaviateKeysDeleteParamsWithHTTPClient(client *http.Client) *WeaviateKeysDeleteParams {
	var ()
	return &WeaviateKeysDeleteParams{
		HTTPClient: client,
	}
}

/*WeaviateKeysDeleteParams contains all the parameters to send to the API endpoint
for the weaviate keys delete operation typically these are written to a http.Request
*/
type WeaviateKeysDeleteParams struct {

	/*KeyID
	  Unique ID of the key.

	*/
	KeyID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the weaviate keys delete params
func (o *WeaviateKeysDeleteParams) WithTimeout(timeout time.Duration) *WeaviateKeysDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the weaviate keys delete params
func (o *WeaviateKeysDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the weaviate keys delete params
func (o *WeaviateKeysDeleteParams) WithContext(ctx context.Context) *WeaviateKeysDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the weaviate keys delete params
func (o *WeaviateKeysDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the weaviate keys delete params
func (o *WeaviateKeysDeleteParams) WithHTTPClient(client *http.Client) *WeaviateKeysDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the weaviate keys delete params
func (o *WeaviateKeysDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithKeyID adds the keyID to the weaviate keys delete params
func (o *WeaviateKeysDeleteParams) WithKeyID(keyID strfmt.UUID) *WeaviateKeysDeleteParams {
	o.SetKeyID(keyID)
	return o
}

// SetKeyID adds the keyId to the weaviate keys delete params
func (o *WeaviateKeysDeleteParams) SetKeyID(keyID strfmt.UUID) {
	o.KeyID = keyID
}

// WriteToRequest writes these params to a swagger request
func (o *WeaviateKeysDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
