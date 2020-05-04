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

package p2_p

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

// NewP2pGenesisUpdateParams creates a new P2pGenesisUpdateParams object
// with the default values initialized.
func NewP2pGenesisUpdateParams() *P2pGenesisUpdateParams {
	var ()
	return &P2pGenesisUpdateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewP2pGenesisUpdateParamsWithTimeout creates a new P2pGenesisUpdateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewP2pGenesisUpdateParamsWithTimeout(timeout time.Duration) *P2pGenesisUpdateParams {
	var ()
	return &P2pGenesisUpdateParams{

		timeout: timeout,
	}
}

// NewP2pGenesisUpdateParamsWithContext creates a new P2pGenesisUpdateParams object
// with the default values initialized, and the ability to set a context for a request
func NewP2pGenesisUpdateParamsWithContext(ctx context.Context) *P2pGenesisUpdateParams {
	var ()
	return &P2pGenesisUpdateParams{

		Context: ctx,
	}
}

// NewP2pGenesisUpdateParamsWithHTTPClient creates a new P2pGenesisUpdateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewP2pGenesisUpdateParamsWithHTTPClient(client *http.Client) *P2pGenesisUpdateParams {
	var ()
	return &P2pGenesisUpdateParams{
		HTTPClient: client,
	}
}

/*P2pGenesisUpdateParams contains all the parameters to send to the API endpoint
for the p2p genesis update operation typically these are written to a http.Request
*/
type P2pGenesisUpdateParams struct {

	/*Peers*/
	Peers models.PeerUpdateList

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the p2p genesis update params
func (o *P2pGenesisUpdateParams) WithTimeout(timeout time.Duration) *P2pGenesisUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the p2p genesis update params
func (o *P2pGenesisUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the p2p genesis update params
func (o *P2pGenesisUpdateParams) WithContext(ctx context.Context) *P2pGenesisUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the p2p genesis update params
func (o *P2pGenesisUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the p2p genesis update params
func (o *P2pGenesisUpdateParams) WithHTTPClient(client *http.Client) *P2pGenesisUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the p2p genesis update params
func (o *P2pGenesisUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithPeers adds the peers to the p2p genesis update params
func (o *P2pGenesisUpdateParams) WithPeers(peers models.PeerUpdateList) *P2pGenesisUpdateParams {
	o.SetPeers(peers)
	return o
}

// SetPeers adds the peers to the p2p genesis update params
func (o *P2pGenesisUpdateParams) SetPeers(peers models.PeerUpdateList) {
	o.Peers = peers
}

// WriteToRequest writes these params to a swagger request
func (o *P2pGenesisUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Peers != nil {
		if err := r.SetBodyParam(o.Peers); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
