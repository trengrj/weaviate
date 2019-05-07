/*                          _       _
 *__      _____  __ ___   ___  __ _| |_ ___
 *\ \ /\ / / _ \/ _` \ \ / / |/ _` | __/ _ \
 * \ V  V /  __/ (_| |\ V /| | (_| | ||  __/
 *  \_/\_/ \___|\__,_| \_/ |_|\__,_|\__\___|
 *
 * Copyright © 2016 - 2019 Weaviate. All rights reserved.
 * LICENSE: https://github.com/semi-technologies/weaviate/blob/develop/LICENSE.md
 * DESIGN & CONCEPT: Bob van Luijt (@bobvanluijt)
 * CONTACT: hello@semi.technology
 */ // Code generated by go-swagger; DO NOT EDIT.

package contextionary_api

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

// NewWeaviateC11yWordsParams creates a new WeaviateC11yWordsParams object
// with the default values initialized.
func NewWeaviateC11yWordsParams() *WeaviateC11yWordsParams {
	var ()
	return &WeaviateC11yWordsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewWeaviateC11yWordsParamsWithTimeout creates a new WeaviateC11yWordsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewWeaviateC11yWordsParamsWithTimeout(timeout time.Duration) *WeaviateC11yWordsParams {
	var ()
	return &WeaviateC11yWordsParams{

		timeout: timeout,
	}
}

// NewWeaviateC11yWordsParamsWithContext creates a new WeaviateC11yWordsParams object
// with the default values initialized, and the ability to set a context for a request
func NewWeaviateC11yWordsParamsWithContext(ctx context.Context) *WeaviateC11yWordsParams {
	var ()
	return &WeaviateC11yWordsParams{

		Context: ctx,
	}
}

// NewWeaviateC11yWordsParamsWithHTTPClient creates a new WeaviateC11yWordsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewWeaviateC11yWordsParamsWithHTTPClient(client *http.Client) *WeaviateC11yWordsParams {
	var ()
	return &WeaviateC11yWordsParams{
		HTTPClient: client,
	}
}

/*WeaviateC11yWordsParams contains all the parameters to send to the API endpoint
for the weaviate c11y words operation typically these are written to a http.Request
*/
type WeaviateC11yWordsParams struct {

	/*Words
	  CamelCase list of words to validate.

	*/
	Words string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the weaviate c11y words params
func (o *WeaviateC11yWordsParams) WithTimeout(timeout time.Duration) *WeaviateC11yWordsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the weaviate c11y words params
func (o *WeaviateC11yWordsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the weaviate c11y words params
func (o *WeaviateC11yWordsParams) WithContext(ctx context.Context) *WeaviateC11yWordsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the weaviate c11y words params
func (o *WeaviateC11yWordsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the weaviate c11y words params
func (o *WeaviateC11yWordsParams) WithHTTPClient(client *http.Client) *WeaviateC11yWordsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the weaviate c11y words params
func (o *WeaviateC11yWordsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithWords adds the words to the weaviate c11y words params
func (o *WeaviateC11yWordsParams) WithWords(words string) *WeaviateC11yWordsParams {
	o.SetWords(words)
	return o
}

// SetWords adds the words to the weaviate c11y words params
func (o *WeaviateC11yWordsParams) SetWords(words string) {
	o.Words = words
}

// WriteToRequest writes these params to a swagger request
func (o *WeaviateC11yWordsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param words
	if err := r.SetPathParam("words", o.Words); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
