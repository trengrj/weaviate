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

	"github.com/semi-technologies/weaviate/entities/models"
)

// NewThingsPatchParams creates a new ThingsPatchParams object
// with the default values initialized.
func NewThingsPatchParams() *ThingsPatchParams {
	var ()
	return &ThingsPatchParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewThingsPatchParamsWithTimeout creates a new ThingsPatchParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewThingsPatchParamsWithTimeout(timeout time.Duration) *ThingsPatchParams {
	var ()
	return &ThingsPatchParams{

		timeout: timeout,
	}
}

// NewThingsPatchParamsWithContext creates a new ThingsPatchParams object
// with the default values initialized, and the ability to set a context for a request
func NewThingsPatchParamsWithContext(ctx context.Context) *ThingsPatchParams {
	var ()
	return &ThingsPatchParams{

		Context: ctx,
	}
}

// NewThingsPatchParamsWithHTTPClient creates a new ThingsPatchParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewThingsPatchParamsWithHTTPClient(client *http.Client) *ThingsPatchParams {
	var ()
	return &ThingsPatchParams{
		HTTPClient: client,
	}
}

/*ThingsPatchParams contains all the parameters to send to the API endpoint
for the things patch operation typically these are written to a http.Request
*/
type ThingsPatchParams struct {

	/*Body
	  RFC 7396-style patch, the body contains the thing object to merge into the existing thing object.

	*/
	Body *models.Thing
	/*ID
	  Unique ID of the Thing.

	*/
	ID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the things patch params
func (o *ThingsPatchParams) WithTimeout(timeout time.Duration) *ThingsPatchParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the things patch params
func (o *ThingsPatchParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the things patch params
func (o *ThingsPatchParams) WithContext(ctx context.Context) *ThingsPatchParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the things patch params
func (o *ThingsPatchParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the things patch params
func (o *ThingsPatchParams) WithHTTPClient(client *http.Client) *ThingsPatchParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the things patch params
func (o *ThingsPatchParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the things patch params
func (o *ThingsPatchParams) WithBody(body *models.Thing) *ThingsPatchParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the things patch params
func (o *ThingsPatchParams) SetBody(body *models.Thing) {
	o.Body = body
}

// WithID adds the id to the things patch params
func (o *ThingsPatchParams) WithID(id strfmt.UUID) *ThingsPatchParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the things patch params
func (o *ThingsPatchParams) SetID(id strfmt.UUID) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *ThingsPatchParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param id
	if err := r.SetPathParam("id", o.ID.String()); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
