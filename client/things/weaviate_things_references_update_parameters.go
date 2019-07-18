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

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/semi-technologies/weaviate/entities/models"
)

// NewWeaviateThingsReferencesUpdateParams creates a new WeaviateThingsReferencesUpdateParams object
// with the default values initialized.
func NewWeaviateThingsReferencesUpdateParams() *WeaviateThingsReferencesUpdateParams {
	var ()
	return &WeaviateThingsReferencesUpdateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewWeaviateThingsReferencesUpdateParamsWithTimeout creates a new WeaviateThingsReferencesUpdateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewWeaviateThingsReferencesUpdateParamsWithTimeout(timeout time.Duration) *WeaviateThingsReferencesUpdateParams {
	var ()
	return &WeaviateThingsReferencesUpdateParams{

		timeout: timeout,
	}
}

// NewWeaviateThingsReferencesUpdateParamsWithContext creates a new WeaviateThingsReferencesUpdateParams object
// with the default values initialized, and the ability to set a context for a request
func NewWeaviateThingsReferencesUpdateParamsWithContext(ctx context.Context) *WeaviateThingsReferencesUpdateParams {
	var ()
	return &WeaviateThingsReferencesUpdateParams{

		Context: ctx,
	}
}

// NewWeaviateThingsReferencesUpdateParamsWithHTTPClient creates a new WeaviateThingsReferencesUpdateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewWeaviateThingsReferencesUpdateParamsWithHTTPClient(client *http.Client) *WeaviateThingsReferencesUpdateParams {
	var ()
	return &WeaviateThingsReferencesUpdateParams{
		HTTPClient: client,
	}
}

/*WeaviateThingsReferencesUpdateParams contains all the parameters to send to the API endpoint
for the weaviate things references update operation typically these are written to a http.Request
*/
type WeaviateThingsReferencesUpdateParams struct {

	/*Body*/
	Body models.MultipleRef
	/*ID
	  Unique ID of the Thing.

	*/
	ID strfmt.UUID
	/*PropertyName
	  Unique name of the property related to the Thing.

	*/
	PropertyName string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the weaviate things references update params
func (o *WeaviateThingsReferencesUpdateParams) WithTimeout(timeout time.Duration) *WeaviateThingsReferencesUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the weaviate things references update params
func (o *WeaviateThingsReferencesUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the weaviate things references update params
func (o *WeaviateThingsReferencesUpdateParams) WithContext(ctx context.Context) *WeaviateThingsReferencesUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the weaviate things references update params
func (o *WeaviateThingsReferencesUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the weaviate things references update params
func (o *WeaviateThingsReferencesUpdateParams) WithHTTPClient(client *http.Client) *WeaviateThingsReferencesUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the weaviate things references update params
func (o *WeaviateThingsReferencesUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the weaviate things references update params
func (o *WeaviateThingsReferencesUpdateParams) WithBody(body models.MultipleRef) *WeaviateThingsReferencesUpdateParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the weaviate things references update params
func (o *WeaviateThingsReferencesUpdateParams) SetBody(body models.MultipleRef) {
	o.Body = body
}

// WithID adds the id to the weaviate things references update params
func (o *WeaviateThingsReferencesUpdateParams) WithID(id strfmt.UUID) *WeaviateThingsReferencesUpdateParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the weaviate things references update params
func (o *WeaviateThingsReferencesUpdateParams) SetID(id strfmt.UUID) {
	o.ID = id
}

// WithPropertyName adds the propertyName to the weaviate things references update params
func (o *WeaviateThingsReferencesUpdateParams) WithPropertyName(propertyName string) *WeaviateThingsReferencesUpdateParams {
	o.SetPropertyName(propertyName)
	return o
}

// SetPropertyName adds the propertyName to the weaviate things references update params
func (o *WeaviateThingsReferencesUpdateParams) SetPropertyName(propertyName string) {
	o.PropertyName = propertyName
}

// WriteToRequest writes these params to a swagger request
func (o *WeaviateThingsReferencesUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param propertyName
	if err := r.SetPathParam("propertyName", o.PropertyName); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
