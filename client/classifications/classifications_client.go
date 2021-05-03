//                           _       _
// __      _____  __ ___   ___  __ _| |_ ___
// \ \ /\ / / _ \/ _` \ \ / / |/ _` | __/ _ \
//  \ V  V /  __/ (_| |\ V /| | (_| | ||  __/
//   \_/\_/ \___|\__,_| \_/ |_|\__,_|\__\___|
//
//  Copyright © 2016 - 2021 SeMI Technologies B.V. All rights reserved.
//
//  CONTACT: hello@semi.technology
//

// Code generated by go-swagger; DO NOT EDIT.

package classifications

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new classifications API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for classifications API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	ClassificationsGet(params *ClassificationsGetParams, authInfo runtime.ClientAuthInfoWriter) (*ClassificationsGetOK, error)

	ClassificationsPost(params *ClassificationsPostParams, authInfo runtime.ClientAuthInfoWriter) (*ClassificationsPostCreated, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  ClassificationsGet views previously created classification

  Get status, results and metadata of a previously created classification
*/
func (a *Client) ClassificationsGet(params *ClassificationsGetParams, authInfo runtime.ClientAuthInfoWriter) (*ClassificationsGetOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewClassificationsGetParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "classifications.get",
		Method:             "GET",
		PathPattern:        "/classifications/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json", "application/yaml"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ClassificationsGetReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ClassificationsGetOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for classifications.get: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ClassificationsPost starts a classification

  Trigger a classification based on the specified params. Classifications will run in the background, use GET /classifications/<id> to retrieve the status of your classification.
*/
func (a *Client) ClassificationsPost(params *ClassificationsPostParams, authInfo runtime.ClientAuthInfoWriter) (*ClassificationsPostCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewClassificationsPostParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "classifications.post",
		Method:             "POST",
		PathPattern:        "/classifications/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json", "application/yaml"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ClassificationsPostReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ClassificationsPostCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for classifications.post: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
