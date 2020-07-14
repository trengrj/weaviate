//                           _       _
// __      _____  __ ___   ___  __ _| |_ ___
// \ \ /\ / / _ \/ _` \ \ / / |/ _` | __/ _ \
//  \ V  V /  __/ (_| |\ V /| | (_| | ||  __/
//   \_/\_/ \___|\__,_| \_/ |_|\__,_|\__\___|
//
//  Copyright © 2016 - 2020 SeMI Technologies B.V. All rights reserved.
//
//  CONTACT: hello@semi.technology
//

// Code generated by go-swagger; DO NOT EDIT.

package batching

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new batching API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for batching API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	BatchingActionsCreate(params *BatchingActionsCreateParams, authInfo runtime.ClientAuthInfoWriter) (*BatchingActionsCreateOK, error)

	BatchingReferencesCreate(params *BatchingReferencesCreateParams, authInfo runtime.ClientAuthInfoWriter) (*BatchingReferencesCreateOK, error)

	BatchingThingsCreate(params *BatchingThingsCreateParams, authInfo runtime.ClientAuthInfoWriter) (*BatchingThingsCreateOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  BatchingActionsCreate creates new actions based on an action template as a batch

  Register new Actions in bulk. Given meta-data and schema values are validated.
*/
func (a *Client) BatchingActionsCreate(params *BatchingActionsCreateParams, authInfo runtime.ClientAuthInfoWriter) (*BatchingActionsCreateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewBatchingActionsCreateParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "batching.actions.create",
		Method:             "POST",
		PathPattern:        "/batching/actions",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json", "application/yaml"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &BatchingActionsCreateReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*BatchingActionsCreateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for batching.actions.create: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  BatchingReferencesCreate creates new cross references between arbitrary classes in bulk

  Register cross-references between any class items (things or actions) in bulk.
*/
func (a *Client) BatchingReferencesCreate(params *BatchingReferencesCreateParams, authInfo runtime.ClientAuthInfoWriter) (*BatchingReferencesCreateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewBatchingReferencesCreateParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "batching.references.create",
		Method:             "POST",
		PathPattern:        "/batching/references",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json", "application/yaml"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &BatchingReferencesCreateReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*BatchingReferencesCreateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for batching.references.create: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  BatchingThingsCreate creates new things based on a thing template as a batch

  Register new Things in bulk. Provided meta-data and schema values are validated.
*/
func (a *Client) BatchingThingsCreate(params *BatchingThingsCreateParams, authInfo runtime.ClientAuthInfoWriter) (*BatchingThingsCreateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewBatchingThingsCreateParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "batching.things.create",
		Method:             "POST",
		PathPattern:        "/batching/things",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json", "application/yaml"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &BatchingThingsCreateReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*BatchingThingsCreateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for batching.things.create: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
