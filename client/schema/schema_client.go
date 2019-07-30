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

package schema

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new schema API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for schema API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
WeaviateSchemaActionsCreate creates a new action class in the ontology
*/
func (a *Client) WeaviateSchemaActionsCreate(params *WeaviateSchemaActionsCreateParams, authInfo runtime.ClientAuthInfoWriter) (*WeaviateSchemaActionsCreateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewWeaviateSchemaActionsCreateParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "weaviate.schema.actions.create",
		Method:             "POST",
		PathPattern:        "/schema/actions",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json", "application/yaml"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &WeaviateSchemaActionsCreateReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*WeaviateSchemaActionsCreateOK), nil

}

/*
WeaviateSchemaActionsDelete removes an action class and all data in the instances from the ontology
*/
func (a *Client) WeaviateSchemaActionsDelete(params *WeaviateSchemaActionsDeleteParams, authInfo runtime.ClientAuthInfoWriter) (*WeaviateSchemaActionsDeleteOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewWeaviateSchemaActionsDeleteParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "weaviate.schema.actions.delete",
		Method:             "DELETE",
		PathPattern:        "/schema/actions/{className}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json", "application/yaml"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &WeaviateSchemaActionsDeleteReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*WeaviateSchemaActionsDeleteOK), nil

}

/*
WeaviateSchemaActionsPropertiesAdd adds a property to an action class
*/
func (a *Client) WeaviateSchemaActionsPropertiesAdd(params *WeaviateSchemaActionsPropertiesAddParams, authInfo runtime.ClientAuthInfoWriter) (*WeaviateSchemaActionsPropertiesAddOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewWeaviateSchemaActionsPropertiesAddParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "weaviate.schema.actions.properties.add",
		Method:             "POST",
		PathPattern:        "/schema/actions/{className}/properties",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json", "application/yaml"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &WeaviateSchemaActionsPropertiesAddReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*WeaviateSchemaActionsPropertiesAddOK), nil

}

/*
WeaviateSchemaActionsPropertiesDelete removes a property from an action class
*/
func (a *Client) WeaviateSchemaActionsPropertiesDelete(params *WeaviateSchemaActionsPropertiesDeleteParams, authInfo runtime.ClientAuthInfoWriter) (*WeaviateSchemaActionsPropertiesDeleteOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewWeaviateSchemaActionsPropertiesDeleteParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "weaviate.schema.actions.properties.delete",
		Method:             "DELETE",
		PathPattern:        "/schema/actions/{className}/properties/{propertyName}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json", "application/yaml"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &WeaviateSchemaActionsPropertiesDeleteReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*WeaviateSchemaActionsPropertiesDeleteOK), nil

}

/*
WeaviateSchemaActionsPropertiesUpdate renames or replace the keywords of the property
*/
func (a *Client) WeaviateSchemaActionsPropertiesUpdate(params *WeaviateSchemaActionsPropertiesUpdateParams, authInfo runtime.ClientAuthInfoWriter) (*WeaviateSchemaActionsPropertiesUpdateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewWeaviateSchemaActionsPropertiesUpdateParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "weaviate.schema.actions.properties.update",
		Method:             "PUT",
		PathPattern:        "/schema/actions/{className}/properties/{propertyName}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json", "application/yaml"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &WeaviateSchemaActionsPropertiesUpdateReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*WeaviateSchemaActionsPropertiesUpdateOK), nil

}

/*
WeaviateSchemaActionsUpdate renames or replace the keywords of the action
*/
func (a *Client) WeaviateSchemaActionsUpdate(params *WeaviateSchemaActionsUpdateParams, authInfo runtime.ClientAuthInfoWriter) (*WeaviateSchemaActionsUpdateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewWeaviateSchemaActionsUpdateParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "weaviate.schema.actions.update",
		Method:             "PUT",
		PathPattern:        "/schema/actions/{className}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json", "application/yaml"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &WeaviateSchemaActionsUpdateReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*WeaviateSchemaActionsUpdateOK), nil

}

/*
WeaviateSchemaDump dumps the current the database schema
*/
func (a *Client) WeaviateSchemaDump(params *WeaviateSchemaDumpParams, authInfo runtime.ClientAuthInfoWriter) (*WeaviateSchemaDumpOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewWeaviateSchemaDumpParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "weaviate.schema.dump",
		Method:             "GET",
		PathPattern:        "/schema",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json", "application/yaml"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &WeaviateSchemaDumpReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*WeaviateSchemaDumpOK), nil

}

/*
WeaviateSchemaThingsCreate creates a new thing class in the ontology
*/
func (a *Client) WeaviateSchemaThingsCreate(params *WeaviateSchemaThingsCreateParams, authInfo runtime.ClientAuthInfoWriter) (*WeaviateSchemaThingsCreateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewWeaviateSchemaThingsCreateParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "weaviate.schema.things.create",
		Method:             "POST",
		PathPattern:        "/schema/things",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json", "application/yaml"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &WeaviateSchemaThingsCreateReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*WeaviateSchemaThingsCreateOK), nil

}

/*
WeaviateSchemaThingsDelete removes a thing class and all data in the instances from the ontology
*/
func (a *Client) WeaviateSchemaThingsDelete(params *WeaviateSchemaThingsDeleteParams, authInfo runtime.ClientAuthInfoWriter) (*WeaviateSchemaThingsDeleteOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewWeaviateSchemaThingsDeleteParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "weaviate.schema.things.delete",
		Method:             "DELETE",
		PathPattern:        "/schema/things/{className}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json", "application/yaml"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &WeaviateSchemaThingsDeleteReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*WeaviateSchemaThingsDeleteOK), nil

}

/*
WeaviateSchemaThingsPropertiesAdd adds a property to a thing class
*/
func (a *Client) WeaviateSchemaThingsPropertiesAdd(params *WeaviateSchemaThingsPropertiesAddParams, authInfo runtime.ClientAuthInfoWriter) (*WeaviateSchemaThingsPropertiesAddOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewWeaviateSchemaThingsPropertiesAddParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "weaviate.schema.things.properties.add",
		Method:             "POST",
		PathPattern:        "/schema/things/{className}/properties",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json", "application/yaml"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &WeaviateSchemaThingsPropertiesAddReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*WeaviateSchemaThingsPropertiesAddOK), nil

}

/*
WeaviateSchemaThingsPropertiesDelete removes a property from a thing class
*/
func (a *Client) WeaviateSchemaThingsPropertiesDelete(params *WeaviateSchemaThingsPropertiesDeleteParams, authInfo runtime.ClientAuthInfoWriter) (*WeaviateSchemaThingsPropertiesDeleteOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewWeaviateSchemaThingsPropertiesDeleteParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "weaviate.schema.things.properties.delete",
		Method:             "DELETE",
		PathPattern:        "/schema/things/{className}/properties/{propertyName}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json", "application/yaml"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &WeaviateSchemaThingsPropertiesDeleteReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*WeaviateSchemaThingsPropertiesDeleteOK), nil

}

/*
WeaviateSchemaThingsPropertiesUpdate renames or replace the keywords of the property
*/
func (a *Client) WeaviateSchemaThingsPropertiesUpdate(params *WeaviateSchemaThingsPropertiesUpdateParams, authInfo runtime.ClientAuthInfoWriter) (*WeaviateSchemaThingsPropertiesUpdateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewWeaviateSchemaThingsPropertiesUpdateParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "weaviate.schema.things.properties.update",
		Method:             "PUT",
		PathPattern:        "/schema/things/{className}/properties/{propertyName}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json", "application/yaml"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &WeaviateSchemaThingsPropertiesUpdateReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*WeaviateSchemaThingsPropertiesUpdateOK), nil

}

/*
WeaviateSchemaThingsUpdate renames or replace the keywords of the thing
*/
func (a *Client) WeaviateSchemaThingsUpdate(params *WeaviateSchemaThingsUpdateParams, authInfo runtime.ClientAuthInfoWriter) (*WeaviateSchemaThingsUpdateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewWeaviateSchemaThingsUpdateParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "weaviate.schema.things.update",
		Method:             "PUT",
		PathPattern:        "/schema/things/{className}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json", "application/yaml"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &WeaviateSchemaThingsUpdateReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*WeaviateSchemaThingsUpdateOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
