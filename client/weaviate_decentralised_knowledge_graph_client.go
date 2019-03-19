/*                          _       _
 *__      _____  __ ___   ___  __ _| |_ ___
 *\ \ /\ / / _ \/ _` \ \ / / |/ _` | __/ _ \
 * \ V  V /  __/ (_| |\ V /| | (_| | ||  __/
 *  \_/\_/ \___|\__,_| \_/ |_|\__,_|\__\___|
 *
 * Copyright © 2016 - 2019 Weaviate. All rights reserved.
 * LICENSE: https://github.com/creativesoftwarefdn/weaviate/blob/develop/LICENSE.md
 * DESIGN & CONCEPT: Bob van Luijt (@bobvanluijt)
 * CONTACT: hello@creativesoftwarefdn.org
 */ // Code generated by go-swagger; DO NOT EDIT.

package client

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/creativesoftwarefdn/weaviate/client/actions"
	"github.com/creativesoftwarefdn/weaviate/client/contextionary_api"
	"github.com/creativesoftwarefdn/weaviate/client/graphql"
	"github.com/creativesoftwarefdn/weaviate/client/meta"
	"github.com/creativesoftwarefdn/weaviate/client/operations"
	"github.com/creativesoftwarefdn/weaviate/client/p2_p"
	"github.com/creativesoftwarefdn/weaviate/client/schema"
	"github.com/creativesoftwarefdn/weaviate/client/things"
)

// Default weaviate decentralised knowledge graph HTTP client.
var Default = NewHTTPClient(nil)

const (
	// DefaultHost is the default Host
	// found in Meta (info) section of spec file
	DefaultHost string = "localhost"
	// DefaultBasePath is the default BasePath
	// found in Meta (info) section of spec file
	DefaultBasePath string = "/weaviate/v1"
)

// DefaultSchemes are the default schemes found in Meta (info) section of spec file
var DefaultSchemes = []string{"https"}

// NewHTTPClient creates a new weaviate decentralised knowledge graph HTTP client.
func NewHTTPClient(formats strfmt.Registry) *WeaviateDecentralisedKnowledgeGraph {
	return NewHTTPClientWithConfig(formats, nil)
}

// NewHTTPClientWithConfig creates a new weaviate decentralised knowledge graph HTTP client,
// using a customizable transport config.
func NewHTTPClientWithConfig(formats strfmt.Registry, cfg *TransportConfig) *WeaviateDecentralisedKnowledgeGraph {
	// ensure nullable parameters have default
	if cfg == nil {
		cfg = DefaultTransportConfig()
	}

	// create transport and client
	transport := httptransport.New(cfg.Host, cfg.BasePath, cfg.Schemes)
	return New(transport, formats)
}

// New creates a new weaviate decentralised knowledge graph client
func New(transport runtime.ClientTransport, formats strfmt.Registry) *WeaviateDecentralisedKnowledgeGraph {
	// ensure nullable parameters have default
	if formats == nil {
		formats = strfmt.Default
	}

	cli := new(WeaviateDecentralisedKnowledgeGraph)
	cli.Transport = transport

	cli.Actions = actions.New(transport, formats)

	cli.ContextionaryAPI = contextionary_api.New(transport, formats)

	cli.Graphql = graphql.New(transport, formats)

	cli.Meta = meta.New(transport, formats)

	cli.Operations = operations.New(transport, formats)

	cli.P2P = p2_p.New(transport, formats)

	cli.Schema = schema.New(transport, formats)

	cli.Things = things.New(transport, formats)

	return cli
}

// DefaultTransportConfig creates a TransportConfig with the
// default settings taken from the meta section of the spec file.
func DefaultTransportConfig() *TransportConfig {
	return &TransportConfig{
		Host:     DefaultHost,
		BasePath: DefaultBasePath,
		Schemes:  DefaultSchemes,
	}
}

// TransportConfig contains the transport related info,
// found in the meta section of the spec file.
type TransportConfig struct {
	Host     string
	BasePath string
	Schemes  []string
}

// WithHost overrides the default host,
// provided by the meta section of the spec file.
func (cfg *TransportConfig) WithHost(host string) *TransportConfig {
	cfg.Host = host
	return cfg
}

// WithBasePath overrides the default basePath,
// provided by the meta section of the spec file.
func (cfg *TransportConfig) WithBasePath(basePath string) *TransportConfig {
	cfg.BasePath = basePath
	return cfg
}

// WithSchemes overrides the default schemes,
// provided by the meta section of the spec file.
func (cfg *TransportConfig) WithSchemes(schemes []string) *TransportConfig {
	cfg.Schemes = schemes
	return cfg
}

// WeaviateDecentralisedKnowledgeGraph is a client for weaviate decentralised knowledge graph
type WeaviateDecentralisedKnowledgeGraph struct {
	Actions *actions.Client

	ContextionaryAPI *contextionary_api.Client

	Graphql *graphql.Client

	Meta *meta.Client

	Operations *operations.Client

	P2P *p2_p.Client

	Schema *schema.Client

	Things *things.Client

	Transport runtime.ClientTransport
}

// SetTransport changes the transport on the client and all its subresources
func (c *WeaviateDecentralisedKnowledgeGraph) SetTransport(transport runtime.ClientTransport) {
	c.Transport = transport

	c.Actions.SetTransport(transport)

	c.ContextionaryAPI.SetTransport(transport)

	c.Graphql.SetTransport(transport)

	c.Meta.SetTransport(transport)

	c.Operations.SetTransport(transport)

	c.P2P.SetTransport(transport)

	c.Schema.SetTransport(transport)

	c.Things.SetTransport(transport)

}
