//                           _       _
// __      _____  __ ___   ___  __ _| |_ ___
// \ \ /\ / / _ \/ _` \ \ / / |/ _` | __/ _ \
//  \ V  V /  __/ (_| |\ V /| | (_| | ||  __/
//   \_/\_/ \___|\__,_| \_/ |_|\__,_|\__\___|
//
//  Copyright © 2016 - 2019 Weaviate. All rights reserved.
//  LICENSE WEAVIATE OPEN SOURCE: https://www.semi.technology/playbook/playbook/contract-weaviate-OSS.html
//  LICENSE WEAVIATE ENTERPRISE: https://www.semi.technology/playbook/contract-weaviate-enterprise.html
//  CONCEPT: Bob van Luijt (@bobvanluijt)
//  CONTACT: hello@semi.technology
//

// Code generated by go-swagger; DO NOT EDIT.

package operations

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

// NewGenesisPeersLeaveParams creates a new GenesisPeersLeaveParams object
// with the default values initialized.
func NewGenesisPeersLeaveParams() *GenesisPeersLeaveParams {
	var ()
	return &GenesisPeersLeaveParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGenesisPeersLeaveParamsWithTimeout creates a new GenesisPeersLeaveParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGenesisPeersLeaveParamsWithTimeout(timeout time.Duration) *GenesisPeersLeaveParams {
	var ()
	return &GenesisPeersLeaveParams{

		timeout: timeout,
	}
}

// NewGenesisPeersLeaveParamsWithContext creates a new GenesisPeersLeaveParams object
// with the default values initialized, and the ability to set a context for a request
func NewGenesisPeersLeaveParamsWithContext(ctx context.Context) *GenesisPeersLeaveParams {
	var ()
	return &GenesisPeersLeaveParams{

		Context: ctx,
	}
}

// NewGenesisPeersLeaveParamsWithHTTPClient creates a new GenesisPeersLeaveParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGenesisPeersLeaveParamsWithHTTPClient(client *http.Client) *GenesisPeersLeaveParams {
	var ()
	return &GenesisPeersLeaveParams{
		HTTPClient: client,
	}
}

/*GenesisPeersLeaveParams contains all the parameters to send to the API endpoint
for the genesis peers leave operation typically these are written to a http.Request
*/
type GenesisPeersLeaveParams struct {

	/*PeerID
	  Name of the Weaviate peer

	*/
	PeerID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the genesis peers leave params
func (o *GenesisPeersLeaveParams) WithTimeout(timeout time.Duration) *GenesisPeersLeaveParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the genesis peers leave params
func (o *GenesisPeersLeaveParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the genesis peers leave params
func (o *GenesisPeersLeaveParams) WithContext(ctx context.Context) *GenesisPeersLeaveParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the genesis peers leave params
func (o *GenesisPeersLeaveParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the genesis peers leave params
func (o *GenesisPeersLeaveParams) WithHTTPClient(client *http.Client) *GenesisPeersLeaveParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the genesis peers leave params
func (o *GenesisPeersLeaveParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithPeerID adds the peerID to the genesis peers leave params
func (o *GenesisPeersLeaveParams) WithPeerID(peerID strfmt.UUID) *GenesisPeersLeaveParams {
	o.SetPeerID(peerID)
	return o
}

// SetPeerID adds the peerId to the genesis peers leave params
func (o *GenesisPeersLeaveParams) SetPeerID(peerID strfmt.UUID) {
	o.PeerID = peerID
}

// WriteToRequest writes these params to a swagger request
func (o *GenesisPeersLeaveParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param peerId
	if err := r.SetPathParam("peerId", o.PeerID.String()); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
