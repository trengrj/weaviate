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

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/loads"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/runtime/security"
	"github.com/go-openapi/runtime/yamlpc"
	"github.com/go-openapi/spec"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/semi-technologies/weaviate/adapters/handlers/rest/operations/batch"
	"github.com/semi-technologies/weaviate/adapters/handlers/rest/operations/classifications"
	"github.com/semi-technologies/weaviate/adapters/handlers/rest/operations/graphql"
	"github.com/semi-technologies/weaviate/adapters/handlers/rest/operations/meta"
	"github.com/semi-technologies/weaviate/adapters/handlers/rest/operations/objects"
	"github.com/semi-technologies/weaviate/adapters/handlers/rest/operations/schema"
	"github.com/semi-technologies/weaviate/adapters/handlers/rest/operations/well_known"
	"github.com/semi-technologies/weaviate/entities/models"
)

// NewWeaviateAPI creates a new Weaviate instance
func NewWeaviateAPI(spec *loads.Document) *WeaviateAPI {
	return &WeaviateAPI{
		handlers:            make(map[string]map[string]http.Handler),
		formats:             strfmt.Default,
		defaultConsumes:     "application/json",
		defaultProduces:     "application/json",
		customConsumers:     make(map[string]runtime.Consumer),
		customProducers:     make(map[string]runtime.Producer),
		PreServerShutdown:   func() {},
		ServerShutdown:      func() {},
		spec:                spec,
		ServeError:          errors.ServeError,
		BasicAuthenticator:  security.BasicAuth,
		APIKeyAuthenticator: security.APIKeyAuth,
		BearerAuthenticator: security.BearerAuth,

		JSONConsumer: runtime.JSONConsumer(),
		YamlConsumer: yamlpc.YAMLConsumer(),

		JSONProducer: runtime.JSONProducer(),

		WellKnownGetWellKnownOpenidConfigurationHandler: well_known.GetWellKnownOpenidConfigurationHandlerFunc(func(params well_known.GetWellKnownOpenidConfigurationParams, principal *models.Principal) middleware.Responder {
			return middleware.NotImplemented("operation well_known.GetWellKnownOpenidConfiguration has not yet been implemented")
		}),
		BatchBatchObjectsCreateHandler: batch.BatchObjectsCreateHandlerFunc(func(params batch.BatchObjectsCreateParams, principal *models.Principal) middleware.Responder {
			return middleware.NotImplemented("operation batch.BatchObjectsCreate has not yet been implemented")
		}),
		BatchBatchObjectsDeleteHandler: batch.BatchObjectsDeleteHandlerFunc(func(params batch.BatchObjectsDeleteParams, principal *models.Principal) middleware.Responder {
			return middleware.NotImplemented("operation batch.BatchObjectsDelete has not yet been implemented")
		}),
		BatchBatchReferencesCreateHandler: batch.BatchReferencesCreateHandlerFunc(func(params batch.BatchReferencesCreateParams, principal *models.Principal) middleware.Responder {
			return middleware.NotImplemented("operation batch.BatchReferencesCreate has not yet been implemented")
		}),
		ClassificationsClassificationsGetHandler: classifications.ClassificationsGetHandlerFunc(func(params classifications.ClassificationsGetParams, principal *models.Principal) middleware.Responder {
			return middleware.NotImplemented("operation classifications.ClassificationsGet has not yet been implemented")
		}),
		ClassificationsClassificationsPostHandler: classifications.ClassificationsPostHandlerFunc(func(params classifications.ClassificationsPostParams, principal *models.Principal) middleware.Responder {
			return middleware.NotImplemented("operation classifications.ClassificationsPost has not yet been implemented")
		}),
		GraphqlGraphqlBatchHandler: graphql.GraphqlBatchHandlerFunc(func(params graphql.GraphqlBatchParams, principal *models.Principal) middleware.Responder {
			return middleware.NotImplemented("operation graphql.GraphqlBatch has not yet been implemented")
		}),
		GraphqlGraphqlPostHandler: graphql.GraphqlPostHandlerFunc(func(params graphql.GraphqlPostParams, principal *models.Principal) middleware.Responder {
			return middleware.NotImplemented("operation graphql.GraphqlPost has not yet been implemented")
		}),
		MetaMetaGetHandler: meta.MetaGetHandlerFunc(func(params meta.MetaGetParams, principal *models.Principal) middleware.Responder {
			return middleware.NotImplemented("operation meta.MetaGet has not yet been implemented")
		}),
		ObjectsObjectsClassDeleteHandler: objects.ObjectsClassDeleteHandlerFunc(func(params objects.ObjectsClassDeleteParams, principal *models.Principal) middleware.Responder {
			return middleware.NotImplemented("operation objects.ObjectsClassDelete has not yet been implemented")
		}),
		ObjectsObjectsClassGetHandler: objects.ObjectsClassGetHandlerFunc(func(params objects.ObjectsClassGetParams, principal *models.Principal) middleware.Responder {
			return middleware.NotImplemented("operation objects.ObjectsClassGet has not yet been implemented")
		}),
		ObjectsObjectsClassPutHandler: objects.ObjectsClassPutHandlerFunc(func(params objects.ObjectsClassPutParams, principal *models.Principal) middleware.Responder {
			return middleware.NotImplemented("operation objects.ObjectsClassPut has not yet been implemented")
		}),
		ObjectsObjectsCreateHandler: objects.ObjectsCreateHandlerFunc(func(params objects.ObjectsCreateParams, principal *models.Principal) middleware.Responder {
			return middleware.NotImplemented("operation objects.ObjectsCreate has not yet been implemented")
		}),
		ObjectsObjectsDeleteHandler: objects.ObjectsDeleteHandlerFunc(func(params objects.ObjectsDeleteParams, principal *models.Principal) middleware.Responder {
			return middleware.NotImplemented("operation objects.ObjectsDelete has not yet been implemented")
		}),
		ObjectsObjectsGetHandler: objects.ObjectsGetHandlerFunc(func(params objects.ObjectsGetParams, principal *models.Principal) middleware.Responder {
			return middleware.NotImplemented("operation objects.ObjectsGet has not yet been implemented")
		}),
		ObjectsObjectsHeadHandler: objects.ObjectsHeadHandlerFunc(func(params objects.ObjectsHeadParams, principal *models.Principal) middleware.Responder {
			return middleware.NotImplemented("operation objects.ObjectsHead has not yet been implemented")
		}),
		ObjectsObjectsListHandler: objects.ObjectsListHandlerFunc(func(params objects.ObjectsListParams, principal *models.Principal) middleware.Responder {
			return middleware.NotImplemented("operation objects.ObjectsList has not yet been implemented")
		}),
		ObjectsObjectsPatchHandler: objects.ObjectsPatchHandlerFunc(func(params objects.ObjectsPatchParams, principal *models.Principal) middleware.Responder {
			return middleware.NotImplemented("operation objects.ObjectsPatch has not yet been implemented")
		}),
		ObjectsObjectsReferencesCreateHandler: objects.ObjectsReferencesCreateHandlerFunc(func(params objects.ObjectsReferencesCreateParams, principal *models.Principal) middleware.Responder {
			return middleware.NotImplemented("operation objects.ObjectsReferencesCreate has not yet been implemented")
		}),
		ObjectsObjectsReferencesDeleteHandler: objects.ObjectsReferencesDeleteHandlerFunc(func(params objects.ObjectsReferencesDeleteParams, principal *models.Principal) middleware.Responder {
			return middleware.NotImplemented("operation objects.ObjectsReferencesDelete has not yet been implemented")
		}),
		ObjectsObjectsReferencesUpdateHandler: objects.ObjectsReferencesUpdateHandlerFunc(func(params objects.ObjectsReferencesUpdateParams, principal *models.Principal) middleware.Responder {
			return middleware.NotImplemented("operation objects.ObjectsReferencesUpdate has not yet been implemented")
		}),
		ObjectsObjectsUpdateHandler: objects.ObjectsUpdateHandlerFunc(func(params objects.ObjectsUpdateParams, principal *models.Principal) middleware.Responder {
			return middleware.NotImplemented("operation objects.ObjectsUpdate has not yet been implemented")
		}),
		ObjectsObjectsValidateHandler: objects.ObjectsValidateHandlerFunc(func(params objects.ObjectsValidateParams, principal *models.Principal) middleware.Responder {
			return middleware.NotImplemented("operation objects.ObjectsValidate has not yet been implemented")
		}),
		SchemaSchemaDumpHandler: schema.SchemaDumpHandlerFunc(func(params schema.SchemaDumpParams, principal *models.Principal) middleware.Responder {
			return middleware.NotImplemented("operation schema.SchemaDump has not yet been implemented")
		}),
		SchemaSchemaObjectsCreateHandler: schema.SchemaObjectsCreateHandlerFunc(func(params schema.SchemaObjectsCreateParams, principal *models.Principal) middleware.Responder {
			return middleware.NotImplemented("operation schema.SchemaObjectsCreate has not yet been implemented")
		}),
		SchemaSchemaObjectsDeleteHandler: schema.SchemaObjectsDeleteHandlerFunc(func(params schema.SchemaObjectsDeleteParams, principal *models.Principal) middleware.Responder {
			return middleware.NotImplemented("operation schema.SchemaObjectsDelete has not yet been implemented")
		}),
		SchemaSchemaObjectsGetHandler: schema.SchemaObjectsGetHandlerFunc(func(params schema.SchemaObjectsGetParams, principal *models.Principal) middleware.Responder {
			return middleware.NotImplemented("operation schema.SchemaObjectsGet has not yet been implemented")
		}),
		SchemaSchemaObjectsPropertiesAddHandler: schema.SchemaObjectsPropertiesAddHandlerFunc(func(params schema.SchemaObjectsPropertiesAddParams, principal *models.Principal) middleware.Responder {
			return middleware.NotImplemented("operation schema.SchemaObjectsPropertiesAdd has not yet been implemented")
		}),
		SchemaSchemaObjectsShardsGetHandler: schema.SchemaObjectsShardsGetHandlerFunc(func(params schema.SchemaObjectsShardsGetParams, principal *models.Principal) middleware.Responder {
			return middleware.NotImplemented("operation schema.SchemaObjectsShardsGet has not yet been implemented")
		}),
		SchemaSchemaObjectsShardsUpdateHandler: schema.SchemaObjectsShardsUpdateHandlerFunc(func(params schema.SchemaObjectsShardsUpdateParams, principal *models.Principal) middleware.Responder {
			return middleware.NotImplemented("operation schema.SchemaObjectsShardsUpdate has not yet been implemented")
		}),
		SchemaSchemaObjectsUpdateHandler: schema.SchemaObjectsUpdateHandlerFunc(func(params schema.SchemaObjectsUpdateParams, principal *models.Principal) middleware.Responder {
			return middleware.NotImplemented("operation schema.SchemaObjectsUpdate has not yet been implemented")
		}),
		WeaviateRootHandler: WeaviateRootHandlerFunc(func(params WeaviateRootParams, principal *models.Principal) middleware.Responder {
			return middleware.NotImplemented("operation WeaviateRoot has not yet been implemented")
		}),
		WeaviateWellknownLivenessHandler: WeaviateWellknownLivenessHandlerFunc(func(params WeaviateWellknownLivenessParams, principal *models.Principal) middleware.Responder {
			return middleware.NotImplemented("operation WeaviateWellknownLiveness has not yet been implemented")
		}),
		WeaviateWellknownReadinessHandler: WeaviateWellknownReadinessHandlerFunc(func(params WeaviateWellknownReadinessParams, principal *models.Principal) middleware.Responder {
			return middleware.NotImplemented("operation WeaviateWellknownReadiness has not yet been implemented")
		}),

		OidcAuth: func(token string, scopes []string) (*models.Principal, error) {
			return nil, errors.NotImplemented("oauth2 bearer auth (oidc) has not yet been implemented")
		},
		// default authorizer is authorized meaning no requests are blocked
		APIAuthorizer: security.Authorized(),
	}
}

/*WeaviateAPI Cloud-native, modular vector search engine */
type WeaviateAPI struct {
	spec            *loads.Document
	context         *middleware.Context
	handlers        map[string]map[string]http.Handler
	formats         strfmt.Registry
	customConsumers map[string]runtime.Consumer
	customProducers map[string]runtime.Producer
	defaultConsumes string
	defaultProduces string
	Middleware      func(middleware.Builder) http.Handler

	// BasicAuthenticator generates a runtime.Authenticator from the supplied basic auth function.
	// It has a default implementation in the security package, however you can replace it for your particular usage.
	BasicAuthenticator func(security.UserPassAuthentication) runtime.Authenticator
	// APIKeyAuthenticator generates a runtime.Authenticator from the supplied token auth function.
	// It has a default implementation in the security package, however you can replace it for your particular usage.
	APIKeyAuthenticator func(string, string, security.TokenAuthentication) runtime.Authenticator
	// BearerAuthenticator generates a runtime.Authenticator from the supplied bearer token auth function.
	// It has a default implementation in the security package, however you can replace it for your particular usage.
	BearerAuthenticator func(string, security.ScopedTokenAuthentication) runtime.Authenticator

	// JSONConsumer registers a consumer for the following mime types:
	//   - application/json
	JSONConsumer runtime.Consumer
	// YamlConsumer registers a consumer for the following mime types:
	//   - application/yaml
	YamlConsumer runtime.Consumer

	// JSONProducer registers a producer for the following mime types:
	//   - application/json
	JSONProducer runtime.Producer

	// OidcAuth registers a function that takes an access token and a collection of required scopes and returns a principal
	// it performs authentication based on an oauth2 bearer token provided in the request
	OidcAuth func(string, []string) (*models.Principal, error)

	// APIAuthorizer provides access control (ACL/RBAC/ABAC) by providing access to the request and authenticated principal
	APIAuthorizer runtime.Authorizer

	// WellKnownGetWellKnownOpenidConfigurationHandler sets the operation handler for the get well known openid configuration operation
	WellKnownGetWellKnownOpenidConfigurationHandler well_known.GetWellKnownOpenidConfigurationHandler
	// BatchBatchObjectsCreateHandler sets the operation handler for the batch objects create operation
	BatchBatchObjectsCreateHandler batch.BatchObjectsCreateHandler
	// BatchBatchObjectsDeleteHandler sets the operation handler for the batch objects delete operation
	BatchBatchObjectsDeleteHandler batch.BatchObjectsDeleteHandler
	// BatchBatchReferencesCreateHandler sets the operation handler for the batch references create operation
	BatchBatchReferencesCreateHandler batch.BatchReferencesCreateHandler
	// ClassificationsClassificationsGetHandler sets the operation handler for the classifications get operation
	ClassificationsClassificationsGetHandler classifications.ClassificationsGetHandler
	// ClassificationsClassificationsPostHandler sets the operation handler for the classifications post operation
	ClassificationsClassificationsPostHandler classifications.ClassificationsPostHandler
	// GraphqlGraphqlBatchHandler sets the operation handler for the graphql batch operation
	GraphqlGraphqlBatchHandler graphql.GraphqlBatchHandler
	// GraphqlGraphqlPostHandler sets the operation handler for the graphql post operation
	GraphqlGraphqlPostHandler graphql.GraphqlPostHandler
	// MetaMetaGetHandler sets the operation handler for the meta get operation
	MetaMetaGetHandler meta.MetaGetHandler
	// ObjectsObjectsClassDeleteHandler sets the operation handler for the objects class delete operation
	ObjectsObjectsClassDeleteHandler objects.ObjectsClassDeleteHandler
	// ObjectsObjectsClassGetHandler sets the operation handler for the objects class get operation
	ObjectsObjectsClassGetHandler objects.ObjectsClassGetHandler
	// ObjectsObjectsClassPutHandler sets the operation handler for the objects class put operation
	ObjectsObjectsClassPutHandler objects.ObjectsClassPutHandler
	// ObjectsObjectsCreateHandler sets the operation handler for the objects create operation
	ObjectsObjectsCreateHandler objects.ObjectsCreateHandler
	// ObjectsObjectsDeleteHandler sets the operation handler for the objects delete operation
	ObjectsObjectsDeleteHandler objects.ObjectsDeleteHandler
	// ObjectsObjectsGetHandler sets the operation handler for the objects get operation
	ObjectsObjectsGetHandler objects.ObjectsGetHandler
	// ObjectsObjectsHeadHandler sets the operation handler for the objects head operation
	ObjectsObjectsHeadHandler objects.ObjectsHeadHandler
	// ObjectsObjectsListHandler sets the operation handler for the objects list operation
	ObjectsObjectsListHandler objects.ObjectsListHandler
	// ObjectsObjectsPatchHandler sets the operation handler for the objects patch operation
	ObjectsObjectsPatchHandler objects.ObjectsPatchHandler
	// ObjectsObjectsReferencesCreateHandler sets the operation handler for the objects references create operation
	ObjectsObjectsReferencesCreateHandler objects.ObjectsReferencesCreateHandler
	// ObjectsObjectsReferencesDeleteHandler sets the operation handler for the objects references delete operation
	ObjectsObjectsReferencesDeleteHandler objects.ObjectsReferencesDeleteHandler
	// ObjectsObjectsReferencesUpdateHandler sets the operation handler for the objects references update operation
	ObjectsObjectsReferencesUpdateHandler objects.ObjectsReferencesUpdateHandler
	// ObjectsObjectsUpdateHandler sets the operation handler for the objects update operation
	ObjectsObjectsUpdateHandler objects.ObjectsUpdateHandler
	// ObjectsObjectsValidateHandler sets the operation handler for the objects validate operation
	ObjectsObjectsValidateHandler objects.ObjectsValidateHandler
	// SchemaSchemaDumpHandler sets the operation handler for the schema dump operation
	SchemaSchemaDumpHandler schema.SchemaDumpHandler
	// SchemaSchemaObjectsCreateHandler sets the operation handler for the schema objects create operation
	SchemaSchemaObjectsCreateHandler schema.SchemaObjectsCreateHandler
	// SchemaSchemaObjectsDeleteHandler sets the operation handler for the schema objects delete operation
	SchemaSchemaObjectsDeleteHandler schema.SchemaObjectsDeleteHandler
	// SchemaSchemaObjectsGetHandler sets the operation handler for the schema objects get operation
	SchemaSchemaObjectsGetHandler schema.SchemaObjectsGetHandler
	// SchemaSchemaObjectsPropertiesAddHandler sets the operation handler for the schema objects properties add operation
	SchemaSchemaObjectsPropertiesAddHandler schema.SchemaObjectsPropertiesAddHandler
	// SchemaSchemaObjectsShardsGetHandler sets the operation handler for the schema objects shards get operation
	SchemaSchemaObjectsShardsGetHandler schema.SchemaObjectsShardsGetHandler
	// SchemaSchemaObjectsShardsUpdateHandler sets the operation handler for the schema objects shards update operation
	SchemaSchemaObjectsShardsUpdateHandler schema.SchemaObjectsShardsUpdateHandler
	// SchemaSchemaObjectsUpdateHandler sets the operation handler for the schema objects update operation
	SchemaSchemaObjectsUpdateHandler schema.SchemaObjectsUpdateHandler
	// WeaviateRootHandler sets the operation handler for the weaviate root operation
	WeaviateRootHandler WeaviateRootHandler
	// WeaviateWellknownLivenessHandler sets the operation handler for the weaviate wellknown liveness operation
	WeaviateWellknownLivenessHandler WeaviateWellknownLivenessHandler
	// WeaviateWellknownReadinessHandler sets the operation handler for the weaviate wellknown readiness operation
	WeaviateWellknownReadinessHandler WeaviateWellknownReadinessHandler
	// ServeError is called when an error is received, there is a default handler
	// but you can set your own with this
	ServeError func(http.ResponseWriter, *http.Request, error)

	// PreServerShutdown is called before the HTTP(S) server is shutdown
	// This allows for custom functions to get executed before the HTTP(S) server stops accepting traffic
	PreServerShutdown func()

	// ServerShutdown is called when the HTTP(S) server is shut down and done
	// handling all active connections and does not accept connections any more
	ServerShutdown func()

	// Custom command line argument groups with their descriptions
	CommandLineOptionsGroups []swag.CommandLineOptionsGroup

	// User defined logger function.
	Logger func(string, ...interface{})
}

// SetDefaultProduces sets the default produces media type
func (o *WeaviateAPI) SetDefaultProduces(mediaType string) {
	o.defaultProduces = mediaType
}

// SetDefaultConsumes returns the default consumes media type
func (o *WeaviateAPI) SetDefaultConsumes(mediaType string) {
	o.defaultConsumes = mediaType
}

// SetSpec sets a spec that will be served for the clients.
func (o *WeaviateAPI) SetSpec(spec *loads.Document) {
	o.spec = spec
}

// DefaultProduces returns the default produces media type
func (o *WeaviateAPI) DefaultProduces() string {
	return o.defaultProduces
}

// DefaultConsumes returns the default consumes media type
func (o *WeaviateAPI) DefaultConsumes() string {
	return o.defaultConsumes
}

// Formats returns the registered string formats
func (o *WeaviateAPI) Formats() strfmt.Registry {
	return o.formats
}

// RegisterFormat registers a custom format validator
func (o *WeaviateAPI) RegisterFormat(name string, format strfmt.Format, validator strfmt.Validator) {
	o.formats.Add(name, format, validator)
}

// Validate validates the registrations in the WeaviateAPI
func (o *WeaviateAPI) Validate() error {
	var unregistered []string

	if o.JSONConsumer == nil {
		unregistered = append(unregistered, "JSONConsumer")
	}
	if o.YamlConsumer == nil {
		unregistered = append(unregistered, "YamlConsumer")
	}

	if o.JSONProducer == nil {
		unregistered = append(unregistered, "JSONProducer")
	}

	if o.OidcAuth == nil {
		unregistered = append(unregistered, "OidcAuth")
	}

	if o.WellKnownGetWellKnownOpenidConfigurationHandler == nil {
		unregistered = append(unregistered, "well_known.GetWellKnownOpenidConfigurationHandler")
	}
	if o.BatchBatchObjectsCreateHandler == nil {
		unregistered = append(unregistered, "batch.BatchObjectsCreateHandler")
	}
	if o.BatchBatchObjectsDeleteHandler == nil {
		unregistered = append(unregistered, "batch.BatchObjectsDeleteHandler")
	}
	if o.BatchBatchReferencesCreateHandler == nil {
		unregistered = append(unregistered, "batch.BatchReferencesCreateHandler")
	}
	if o.ClassificationsClassificationsGetHandler == nil {
		unregistered = append(unregistered, "classifications.ClassificationsGetHandler")
	}
	if o.ClassificationsClassificationsPostHandler == nil {
		unregistered = append(unregistered, "classifications.ClassificationsPostHandler")
	}
	if o.GraphqlGraphqlBatchHandler == nil {
		unregistered = append(unregistered, "graphql.GraphqlBatchHandler")
	}
	if o.GraphqlGraphqlPostHandler == nil {
		unregistered = append(unregistered, "graphql.GraphqlPostHandler")
	}
	if o.MetaMetaGetHandler == nil {
		unregistered = append(unregistered, "meta.MetaGetHandler")
	}
	if o.ObjectsObjectsClassDeleteHandler == nil {
		unregistered = append(unregistered, "objects.ObjectsClassDeleteHandler")
	}
	if o.ObjectsObjectsClassGetHandler == nil {
		unregistered = append(unregistered, "objects.ObjectsClassGetHandler")
	}
	if o.ObjectsObjectsClassPutHandler == nil {
		unregistered = append(unregistered, "objects.ObjectsClassPutHandler")
	}
	if o.ObjectsObjectsCreateHandler == nil {
		unregistered = append(unregistered, "objects.ObjectsCreateHandler")
	}
	if o.ObjectsObjectsDeleteHandler == nil {
		unregistered = append(unregistered, "objects.ObjectsDeleteHandler")
	}
	if o.ObjectsObjectsGetHandler == nil {
		unregistered = append(unregistered, "objects.ObjectsGetHandler")
	}
	if o.ObjectsObjectsHeadHandler == nil {
		unregistered = append(unregistered, "objects.ObjectsHeadHandler")
	}
	if o.ObjectsObjectsListHandler == nil {
		unregistered = append(unregistered, "objects.ObjectsListHandler")
	}
	if o.ObjectsObjectsPatchHandler == nil {
		unregistered = append(unregistered, "objects.ObjectsPatchHandler")
	}
	if o.ObjectsObjectsReferencesCreateHandler == nil {
		unregistered = append(unregistered, "objects.ObjectsReferencesCreateHandler")
	}
	if o.ObjectsObjectsReferencesDeleteHandler == nil {
		unregistered = append(unregistered, "objects.ObjectsReferencesDeleteHandler")
	}
	if o.ObjectsObjectsReferencesUpdateHandler == nil {
		unregistered = append(unregistered, "objects.ObjectsReferencesUpdateHandler")
	}
	if o.ObjectsObjectsUpdateHandler == nil {
		unregistered = append(unregistered, "objects.ObjectsUpdateHandler")
	}
	if o.ObjectsObjectsValidateHandler == nil {
		unregistered = append(unregistered, "objects.ObjectsValidateHandler")
	}
	if o.SchemaSchemaDumpHandler == nil {
		unregistered = append(unregistered, "schema.SchemaDumpHandler")
	}
	if o.SchemaSchemaObjectsCreateHandler == nil {
		unregistered = append(unregistered, "schema.SchemaObjectsCreateHandler")
	}
	if o.SchemaSchemaObjectsDeleteHandler == nil {
		unregistered = append(unregistered, "schema.SchemaObjectsDeleteHandler")
	}
	if o.SchemaSchemaObjectsGetHandler == nil {
		unregistered = append(unregistered, "schema.SchemaObjectsGetHandler")
	}
	if o.SchemaSchemaObjectsPropertiesAddHandler == nil {
		unregistered = append(unregistered, "schema.SchemaObjectsPropertiesAddHandler")
	}
	if o.SchemaSchemaObjectsShardsGetHandler == nil {
		unregistered = append(unregistered, "schema.SchemaObjectsShardsGetHandler")
	}
	if o.SchemaSchemaObjectsShardsUpdateHandler == nil {
		unregistered = append(unregistered, "schema.SchemaObjectsShardsUpdateHandler")
	}
	if o.SchemaSchemaObjectsUpdateHandler == nil {
		unregistered = append(unregistered, "schema.SchemaObjectsUpdateHandler")
	}
	if o.WeaviateRootHandler == nil {
		unregistered = append(unregistered, "WeaviateRootHandler")
	}
	if o.WeaviateWellknownLivenessHandler == nil {
		unregistered = append(unregistered, "WeaviateWellknownLivenessHandler")
	}
	if o.WeaviateWellknownReadinessHandler == nil {
		unregistered = append(unregistered, "WeaviateWellknownReadinessHandler")
	}

	if len(unregistered) > 0 {
		return fmt.Errorf("missing registration: %s", strings.Join(unregistered, ", "))
	}

	return nil
}

// ServeErrorFor gets a error handler for a given operation id
func (o *WeaviateAPI) ServeErrorFor(operationID string) func(http.ResponseWriter, *http.Request, error) {
	return o.ServeError
}

// AuthenticatorsFor gets the authenticators for the specified security schemes
func (o *WeaviateAPI) AuthenticatorsFor(schemes map[string]spec.SecurityScheme) map[string]runtime.Authenticator {
	result := make(map[string]runtime.Authenticator)
	for name := range schemes {
		switch name {
		case "oidc":
			result[name] = o.BearerAuthenticator(name, func(token string, scopes []string) (interface{}, error) {
				return o.OidcAuth(token, scopes)
			})

		}
	}
	return result
}

// Authorizer returns the registered authorizer
func (o *WeaviateAPI) Authorizer() runtime.Authorizer {
	return o.APIAuthorizer
}

// ConsumersFor gets the consumers for the specified media types.
// MIME type parameters are ignored here.
func (o *WeaviateAPI) ConsumersFor(mediaTypes []string) map[string]runtime.Consumer {
	result := make(map[string]runtime.Consumer, len(mediaTypes))
	for _, mt := range mediaTypes {
		switch mt {
		case "application/json":
			result["application/json"] = o.JSONConsumer
		case "application/yaml":
			result["application/yaml"] = o.YamlConsumer
		}

		if c, ok := o.customConsumers[mt]; ok {
			result[mt] = c
		}
	}
	return result
}

// ProducersFor gets the producers for the specified media types.
// MIME type parameters are ignored here.
func (o *WeaviateAPI) ProducersFor(mediaTypes []string) map[string]runtime.Producer {
	result := make(map[string]runtime.Producer, len(mediaTypes))
	for _, mt := range mediaTypes {
		switch mt {
		case "application/json":
			result["application/json"] = o.JSONProducer
		}

		if p, ok := o.customProducers[mt]; ok {
			result[mt] = p
		}
	}
	return result
}

// HandlerFor gets a http.Handler for the provided operation method and path
func (o *WeaviateAPI) HandlerFor(method, path string) (http.Handler, bool) {
	if o.handlers == nil {
		return nil, false
	}
	um := strings.ToUpper(method)
	if _, ok := o.handlers[um]; !ok {
		return nil, false
	}
	if path == "/" {
		path = ""
	}
	h, ok := o.handlers[um][path]
	return h, ok
}

// Context returns the middleware context for the weaviate API
func (o *WeaviateAPI) Context() *middleware.Context {
	if o.context == nil {
		o.context = middleware.NewRoutableContext(o.spec, o, nil)
	}

	return o.context
}

func (o *WeaviateAPI) initHandlerCache() {
	o.Context() // don't care about the result, just that the initialization happened
	if o.handlers == nil {
		o.handlers = make(map[string]map[string]http.Handler)
	}

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/.well-known/openid-configuration"] = well_known.NewGetWellKnownOpenidConfiguration(o.context, o.WellKnownGetWellKnownOpenidConfigurationHandler)
	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/batch/objects"] = batch.NewBatchObjectsCreate(o.context, o.BatchBatchObjectsCreateHandler)
	if o.handlers["DELETE"] == nil {
		o.handlers["DELETE"] = make(map[string]http.Handler)
	}
	o.handlers["DELETE"]["/batch/objects"] = batch.NewBatchObjectsDelete(o.context, o.BatchBatchObjectsDeleteHandler)
	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/batch/references"] = batch.NewBatchReferencesCreate(o.context, o.BatchBatchReferencesCreateHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/classifications/{id}"] = classifications.NewClassificationsGet(o.context, o.ClassificationsClassificationsGetHandler)
	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/classifications"] = classifications.NewClassificationsPost(o.context, o.ClassificationsClassificationsPostHandler)
	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/graphql/batch"] = graphql.NewGraphqlBatch(o.context, o.GraphqlGraphqlBatchHandler)
	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/graphql"] = graphql.NewGraphqlPost(o.context, o.GraphqlGraphqlPostHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/meta"] = meta.NewMetaGet(o.context, o.MetaMetaGetHandler)
	if o.handlers["DELETE"] == nil {
		o.handlers["DELETE"] = make(map[string]http.Handler)
	}
	o.handlers["DELETE"]["/objects/{className}/{id}"] = objects.NewObjectsClassDelete(o.context, o.ObjectsObjectsClassDeleteHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/objects/{className}/{id}"] = objects.NewObjectsClassGet(o.context, o.ObjectsObjectsClassGetHandler)
	if o.handlers["PUT"] == nil {
		o.handlers["PUT"] = make(map[string]http.Handler)
	}
	o.handlers["PUT"]["/objects/{className}/{id}"] = objects.NewObjectsClassPut(o.context, o.ObjectsObjectsClassPutHandler)
	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/objects"] = objects.NewObjectsCreate(o.context, o.ObjectsObjectsCreateHandler)
	if o.handlers["DELETE"] == nil {
		o.handlers["DELETE"] = make(map[string]http.Handler)
	}
	o.handlers["DELETE"]["/objects/{id}"] = objects.NewObjectsDelete(o.context, o.ObjectsObjectsDeleteHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/objects/{id}"] = objects.NewObjectsGet(o.context, o.ObjectsObjectsGetHandler)
	if o.handlers["HEAD"] == nil {
		o.handlers["HEAD"] = make(map[string]http.Handler)
	}
	o.handlers["HEAD"]["/objects/{id}"] = objects.NewObjectsHead(o.context, o.ObjectsObjectsHeadHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/objects"] = objects.NewObjectsList(o.context, o.ObjectsObjectsListHandler)
	if o.handlers["PATCH"] == nil {
		o.handlers["PATCH"] = make(map[string]http.Handler)
	}
	o.handlers["PATCH"]["/objects/{id}"] = objects.NewObjectsPatch(o.context, o.ObjectsObjectsPatchHandler)
	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/objects/{id}/references/{propertyName}"] = objects.NewObjectsReferencesCreate(o.context, o.ObjectsObjectsReferencesCreateHandler)
	if o.handlers["DELETE"] == nil {
		o.handlers["DELETE"] = make(map[string]http.Handler)
	}
	o.handlers["DELETE"]["/objects/{id}/references/{propertyName}"] = objects.NewObjectsReferencesDelete(o.context, o.ObjectsObjectsReferencesDeleteHandler)
	if o.handlers["PUT"] == nil {
		o.handlers["PUT"] = make(map[string]http.Handler)
	}
	o.handlers["PUT"]["/objects/{id}/references/{propertyName}"] = objects.NewObjectsReferencesUpdate(o.context, o.ObjectsObjectsReferencesUpdateHandler)
	if o.handlers["PUT"] == nil {
		o.handlers["PUT"] = make(map[string]http.Handler)
	}
	o.handlers["PUT"]["/objects/{id}"] = objects.NewObjectsUpdate(o.context, o.ObjectsObjectsUpdateHandler)
	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/objects/validate"] = objects.NewObjectsValidate(o.context, o.ObjectsObjectsValidateHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/schema"] = schema.NewSchemaDump(o.context, o.SchemaSchemaDumpHandler)
	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/schema"] = schema.NewSchemaObjectsCreate(o.context, o.SchemaSchemaObjectsCreateHandler)
	if o.handlers["DELETE"] == nil {
		o.handlers["DELETE"] = make(map[string]http.Handler)
	}
	o.handlers["DELETE"]["/schema/{className}"] = schema.NewSchemaObjectsDelete(o.context, o.SchemaSchemaObjectsDeleteHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/schema/{className}"] = schema.NewSchemaObjectsGet(o.context, o.SchemaSchemaObjectsGetHandler)
	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/schema/{className}/properties"] = schema.NewSchemaObjectsPropertiesAdd(o.context, o.SchemaSchemaObjectsPropertiesAddHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/schema/{className}/shards"] = schema.NewSchemaObjectsShardsGet(o.context, o.SchemaSchemaObjectsShardsGetHandler)
	if o.handlers["PUT"] == nil {
		o.handlers["PUT"] = make(map[string]http.Handler)
	}
	o.handlers["PUT"]["/schema/{className}/shards/{shardName}"] = schema.NewSchemaObjectsShardsUpdate(o.context, o.SchemaSchemaObjectsShardsUpdateHandler)
	if o.handlers["PUT"] == nil {
		o.handlers["PUT"] = make(map[string]http.Handler)
	}
	o.handlers["PUT"]["/schema/{className}"] = schema.NewSchemaObjectsUpdate(o.context, o.SchemaSchemaObjectsUpdateHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"][""] = NewWeaviateRoot(o.context, o.WeaviateRootHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/.well-known/live"] = NewWeaviateWellknownLiveness(o.context, o.WeaviateWellknownLivenessHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/.well-known/ready"] = NewWeaviateWellknownReadiness(o.context, o.WeaviateWellknownReadinessHandler)
}

// Serve creates a http handler to serve the API over HTTP
// can be used directly in http.ListenAndServe(":8000", api.Serve(nil))
func (o *WeaviateAPI) Serve(builder middleware.Builder) http.Handler {
	o.Init()

	if o.Middleware != nil {
		return o.Middleware(builder)
	}
	return o.context.APIHandler(builder)
}

// Init allows you to just initialize the handler cache, you can then recompose the middleware as you see fit
func (o *WeaviateAPI) Init() {
	if len(o.handlers) == 0 {
		o.initHandlerCache()
	}
}

// RegisterConsumer allows you to add (or override) a consumer for a media type.
func (o *WeaviateAPI) RegisterConsumer(mediaType string, consumer runtime.Consumer) {
	o.customConsumers[mediaType] = consumer
}

// RegisterProducer allows you to add (or override) a producer for a media type.
func (o *WeaviateAPI) RegisterProducer(mediaType string, producer runtime.Producer) {
	o.customProducers[mediaType] = producer
}

// AddMiddlewareFor adds a http middleware to existing handler
func (o *WeaviateAPI) AddMiddlewareFor(method, path string, builder middleware.Builder) {
	um := strings.ToUpper(method)
	if path == "/" {
		path = ""
	}
	o.Init()
	if h, ok := o.handlers[um][path]; ok {
		o.handlers[method][path] = builder(h)
	}
}
