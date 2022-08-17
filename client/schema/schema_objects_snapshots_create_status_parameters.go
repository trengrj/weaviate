//                           _       _
// __      _____  __ ___   ___  __ _| |_ ___
// \ \ /\ / / _ \/ _` \ \ / / |/ _` | __/ _ \
//  \ V  V /  __/ (_| |\ V /| | (_| | ||  __/
//   \_/\_/ \___|\__,_| \_/ |_|\__,_|\__\___|
//
//  Copyright © 2016 - 2022 SeMI Technologies B.V. All rights reserved.
//
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
	"github.com/go-openapi/strfmt"
)

// NewSchemaObjectsSnapshotsCreateStatusParams creates a new SchemaObjectsSnapshotsCreateStatusParams object
// with the default values initialized.
func NewSchemaObjectsSnapshotsCreateStatusParams() *SchemaObjectsSnapshotsCreateStatusParams {
	var ()
	return &SchemaObjectsSnapshotsCreateStatusParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewSchemaObjectsSnapshotsCreateStatusParamsWithTimeout creates a new SchemaObjectsSnapshotsCreateStatusParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewSchemaObjectsSnapshotsCreateStatusParamsWithTimeout(timeout time.Duration) *SchemaObjectsSnapshotsCreateStatusParams {
	var ()
	return &SchemaObjectsSnapshotsCreateStatusParams{

		timeout: timeout,
	}
}

// NewSchemaObjectsSnapshotsCreateStatusParamsWithContext creates a new SchemaObjectsSnapshotsCreateStatusParams object
// with the default values initialized, and the ability to set a context for a request
func NewSchemaObjectsSnapshotsCreateStatusParamsWithContext(ctx context.Context) *SchemaObjectsSnapshotsCreateStatusParams {
	var ()
	return &SchemaObjectsSnapshotsCreateStatusParams{

		Context: ctx,
	}
}

// NewSchemaObjectsSnapshotsCreateStatusParamsWithHTTPClient creates a new SchemaObjectsSnapshotsCreateStatusParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewSchemaObjectsSnapshotsCreateStatusParamsWithHTTPClient(client *http.Client) *SchemaObjectsSnapshotsCreateStatusParams {
	var ()
	return &SchemaObjectsSnapshotsCreateStatusParams{
		HTTPClient: client,
	}
}

/*SchemaObjectsSnapshotsCreateStatusParams contains all the parameters to send to the API endpoint
for the schema objects snapshots create status operation typically these are written to a http.Request
*/
type SchemaObjectsSnapshotsCreateStatusParams struct {

	/*ClassName
	  The name of the class

	*/
	ClassName string
	/*ID
	  The ID of a snapshot. Must be URL-safe and work as a filesystem path, only lowercase, numbers, underscore, minus characters allowed.

	*/
	ID string
	/*StorageName
	  Storage name e.g. filesystem, gcs, s3.

	*/
	StorageName string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the schema objects snapshots create status params
func (o *SchemaObjectsSnapshotsCreateStatusParams) WithTimeout(timeout time.Duration) *SchemaObjectsSnapshotsCreateStatusParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the schema objects snapshots create status params
func (o *SchemaObjectsSnapshotsCreateStatusParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the schema objects snapshots create status params
func (o *SchemaObjectsSnapshotsCreateStatusParams) WithContext(ctx context.Context) *SchemaObjectsSnapshotsCreateStatusParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the schema objects snapshots create status params
func (o *SchemaObjectsSnapshotsCreateStatusParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the schema objects snapshots create status params
func (o *SchemaObjectsSnapshotsCreateStatusParams) WithHTTPClient(client *http.Client) *SchemaObjectsSnapshotsCreateStatusParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the schema objects snapshots create status params
func (o *SchemaObjectsSnapshotsCreateStatusParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithClassName adds the className to the schema objects snapshots create status params
func (o *SchemaObjectsSnapshotsCreateStatusParams) WithClassName(className string) *SchemaObjectsSnapshotsCreateStatusParams {
	o.SetClassName(className)
	return o
}

// SetClassName adds the className to the schema objects snapshots create status params
func (o *SchemaObjectsSnapshotsCreateStatusParams) SetClassName(className string) {
	o.ClassName = className
}

// WithID adds the id to the schema objects snapshots create status params
func (o *SchemaObjectsSnapshotsCreateStatusParams) WithID(id string) *SchemaObjectsSnapshotsCreateStatusParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the schema objects snapshots create status params
func (o *SchemaObjectsSnapshotsCreateStatusParams) SetID(id string) {
	o.ID = id
}

// WithStorageName adds the storageName to the schema objects snapshots create status params
func (o *SchemaObjectsSnapshotsCreateStatusParams) WithStorageName(storageName string) *SchemaObjectsSnapshotsCreateStatusParams {
	o.SetStorageName(storageName)
	return o
}

// SetStorageName adds the storageName to the schema objects snapshots create status params
func (o *SchemaObjectsSnapshotsCreateStatusParams) SetStorageName(storageName string) {
	o.StorageName = storageName
}

// WriteToRequest writes these params to a swagger request
func (o *SchemaObjectsSnapshotsCreateStatusParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param className
	if err := r.SetPathParam("className", o.ClassName); err != nil {
		return err
	}

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	// path param storageName
	if err := r.SetPathParam("storageName", o.StorageName); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
