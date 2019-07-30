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

package contextionary_api

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new contextionary api API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for contextionary api API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
WeaviateC11yCorpusGet checks if a word or word string is part of the contextionary

Analyzes a sentence based on the contextionary
*/
func (a *Client) WeaviateC11yCorpusGet(params *WeaviateC11yCorpusGetParams, authInfo runtime.ClientAuthInfoWriter) error {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewWeaviateC11yCorpusGetParams()
	}

	_, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "weaviate.c11y.corpus.get",
		Method:             "POST",
		PathPattern:        "/c11y/corpus",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json", "application/yaml"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &WeaviateC11yCorpusGetReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return err
	}
	return nil

}

/*
WeaviateC11yWords checks if a word or word string is part of the contextionary

Checks if a word or wordString is part of the contextionary. Words should be concatenated as described here: https://github.com/semi-technologies/weaviate/blob/master/docs/en/use/ontology-schema.md#camelcase
*/
func (a *Client) WeaviateC11yWords(params *WeaviateC11yWordsParams, authInfo runtime.ClientAuthInfoWriter) (*WeaviateC11yWordsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewWeaviateC11yWordsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "weaviate.c11y.words",
		Method:             "GET",
		PathPattern:        "/c11y/words/{words}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json", "application/yaml"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &WeaviateC11yWordsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*WeaviateC11yWordsOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
