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

package batch

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new batch API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for batch API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	BatchObjectsCreate(params *BatchObjectsCreateParams, authInfo runtime.ClientAuthInfoWriter) (*BatchObjectsCreateOK, error)

	BatchReferencesCreate(params *BatchReferencesCreateParams, authInfo runtime.ClientAuthInfoWriter) (*BatchReferencesCreateOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  BatchObjectsCreate creates new objects based on a object template as a batch

  Register new Objects in bulk. Provided meta-data and schema values are validated.
*/
func (a *Client) BatchObjectsCreate(params *BatchObjectsCreateParams, authInfo runtime.ClientAuthInfoWriter) (*BatchObjectsCreateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewBatchObjectsCreateParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "batch.objects.create",
		Method:             "POST",
		PathPattern:        "/batch/objects",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json", "application/yaml"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &BatchObjectsCreateReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*BatchObjectsCreateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for batch.objects.create: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  BatchReferencesCreate creates new cross references between arbitrary classes in bulk

  Register cross-references between any class items (objects or objects) in bulk.
*/
func (a *Client) BatchReferencesCreate(params *BatchReferencesCreateParams, authInfo runtime.ClientAuthInfoWriter) (*BatchReferencesCreateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewBatchReferencesCreateParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "batch.references.create",
		Method:             "POST",
		PathPattern:        "/batch/references",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json", "application/yaml"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &BatchReferencesCreateReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*BatchReferencesCreateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for batch.references.create: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
