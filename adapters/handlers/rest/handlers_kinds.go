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

package rest

import (
	"context"
	"fmt"
	"strings"

	middleware "github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/semi-technologies/weaviate/adapters/handlers/rest/operations"
	"github.com/semi-technologies/weaviate/adapters/handlers/rest/operations/objects"
	"github.com/semi-technologies/weaviate/entities/models"
	"github.com/semi-technologies/weaviate/entities/schema/crossref"
	"github.com/semi-technologies/weaviate/usecases/auth/authorization/errors"
	"github.com/semi-technologies/weaviate/usecases/config"
	"github.com/semi-technologies/weaviate/usecases/kinds"
	"github.com/semi-technologies/weaviate/usecases/projector"
	"github.com/semi-technologies/weaviate/usecases/traverser"
	"github.com/sirupsen/logrus"
)

type kindHandlers struct {
	manager kindsManager
	logger  logrus.FieldLogger
	config  config.Config
}

type kindsManager interface {
	AddObject(context.Context, *models.Principal, *models.Object) (*models.Object, error)
	ValidateObject(context.Context, *models.Principal, *models.Object) error
	GetObject(context.Context, *models.Principal, strfmt.UUID, traverser.UnderscoreProperties) (*models.Object, error)
	GetObjects(context.Context, *models.Principal, *int64, traverser.UnderscoreProperties) ([]*models.Object, error)
	UpdateObject(context.Context, *models.Principal, strfmt.UUID, *models.Object) (*models.Object, error)
	MergeObject(context.Context, *models.Principal, strfmt.UUID, *models.Object) error
	DeleteObject(context.Context, *models.Principal, strfmt.UUID) error
	AddObjectReference(context.Context, *models.Principal, strfmt.UUID, string, *models.SingleRef) error
	UpdateObjectReferences(context.Context, *models.Principal, strfmt.UUID, string, models.MultipleRef) error
	DeleteObjectReference(context.Context, *models.Principal, strfmt.UUID, string, *models.SingleRef) error
}

func (h *kindHandlers) addObject(params objects.ObjectsCreateParams,
	principal *models.Principal) middleware.Responder {
	object, err := h.manager.AddObject(params.HTTPRequest.Context(), principal, params.Body)
	if err != nil {
		switch err.(type) {
		case errors.Forbidden:
			return objects.NewObjectsCreateForbidden().
				WithPayload(errPayloadFromSingleErr(err))
		case kinds.ErrInvalidUserInput:
			return objects.NewObjectsCreateUnprocessableEntity().
				WithPayload(errPayloadFromSingleErr(err))
		default:
			return objects.NewObjectsCreateInternalServerError().
				WithPayload(errPayloadFromSingleErr(err))
		}
	}

	schemaMap, ok := object.Schema.(map[string]interface{})
	if ok {
		object.Schema = h.extendSchemaWithAPILinks(schemaMap)
	}

	return objects.NewObjectsCreateOK().WithPayload(object)
}

func (h *kindHandlers) validateObject(params objects.ObjectsValidateParams,
	principal *models.Principal) middleware.Responder {
	err := h.manager.ValidateObject(params.HTTPRequest.Context(), principal, params.Body)
	if err != nil {
		switch err.(type) {
		case errors.Forbidden:
			return objects.NewObjectsValidateForbidden().
				WithPayload(errPayloadFromSingleErr(err))
		case kinds.ErrInvalidUserInput:
			return objects.NewObjectsValidateUnprocessableEntity().
				WithPayload(errPayloadFromSingleErr(err))
		default:
			return objects.NewObjectsValidateInternalServerError().
				WithPayload(errPayloadFromSingleErr(err))
		}
	}

	return objects.NewObjectsValidateOK()
}

func (h *kindHandlers) getObject(params objects.ObjectsGetParams,
	principal *models.Principal) middleware.Responder {
	underscores, err := parseIncludeParam(params.Include)
	if err != nil {
		return objects.NewObjectsGetBadRequest().
			WithPayload(errPayloadFromSingleErr(err))
	}

	object, err := h.manager.GetObject(params.HTTPRequest.Context(), principal, params.ID, underscores)
	if err != nil {
		switch err.(type) {
		case errors.Forbidden:
			return objects.NewObjectsGetForbidden().
				WithPayload(errPayloadFromSingleErr(err))
		case kinds.ErrNotFound:
			return objects.NewObjectsGetNotFound()
		default:
			return objects.NewObjectsGetInternalServerError().
				WithPayload(errPayloadFromSingleErr(err))
		}
	}

	schemaMap, ok := object.Schema.(map[string]interface{})
	if ok {
		object.Schema = h.extendSchemaWithAPILinks(schemaMap)
	}

	return objects.NewObjectsGetOK().WithPayload(object)
}

func (h *kindHandlers) getObjects(params objects.ObjectsListParams,
	principal *models.Principal) middleware.Responder {
	underscores, err := parseIncludeParam(params.Include)
	if err != nil {
		return objects.NewObjectsListBadRequest().
			WithPayload(errPayloadFromSingleErr(err))
	}

	var deprecationsRes []*models.Deprecation

	list, err := h.manager.GetObjects(params.HTTPRequest.Context(), principal, params.Limit, underscores)
	if err != nil {
		switch err.(type) {
		case errors.Forbidden:
			return objects.NewObjectsListForbidden().
				WithPayload(errPayloadFromSingleErr(err))
		default:
			return objects.NewObjectsListInternalServerError().
				WithPayload(errPayloadFromSingleErr(err))
		}
	}

	for i, object := range list {
		schemaMap, ok := object.Schema.(map[string]interface{})
		if ok {
			list[i].Schema = h.extendSchemaWithAPILinks(schemaMap)
		}
	}

	return objects.NewObjectsListOK().
		WithPayload(&models.ObjectsListResponse{
			Objects:      list,
			TotalResults: int64(len(list)),
			Deprecations: deprecationsRes,
		})
}

func (h *kindHandlers) updateObject(params objects.ObjectsUpdateParams,
	principal *models.Principal) middleware.Responder {
	object, err := h.manager.UpdateObject(params.HTTPRequest.Context(), principal, params.ID, params.Body)
	if err != nil {
		switch err.(type) {
		case errors.Forbidden:
			return objects.NewObjectsUpdateForbidden().
				WithPayload(errPayloadFromSingleErr(err))
		case kinds.ErrInvalidUserInput:
			return objects.NewObjectsUpdateUnprocessableEntity().
				WithPayload(errPayloadFromSingleErr(err))
		default:
			return objects.NewObjectsUpdateInternalServerError().
				WithPayload(errPayloadFromSingleErr(err))
		}
	}

	schemaMap, ok := object.Schema.(map[string]interface{})
	if ok {
		object.Schema = h.extendSchemaWithAPILinks(schemaMap)
	}

	return objects.NewObjectsUpdateOK().WithPayload(object)
}

func (h *kindHandlers) deleteObject(params objects.ObjectsDeleteParams,
	principal *models.Principal) middleware.Responder {
	err := h.manager.DeleteObject(params.HTTPRequest.Context(), principal, params.ID)
	if err != nil {
		switch err.(type) {
		case errors.Forbidden:
			return objects.NewObjectsDeleteForbidden().
				WithPayload(errPayloadFromSingleErr(err))
		case kinds.ErrNotFound:
			return objects.NewObjectsDeleteNotFound()
		default:
			return objects.NewObjectsDeleteInternalServerError().
				WithPayload(errPayloadFromSingleErr(err))
		}
	}

	return objects.NewObjectsDeleteNoContent()
}

func (h *kindHandlers) patchObject(params objects.ObjectsPatchParams, principal *models.Principal) middleware.Responder {
	err := h.manager.MergeObject(params.HTTPRequest.Context(), principal, params.ID, params.Body)
	if err != nil {
		switch err.(type) {
		case errors.Forbidden:
			return objects.NewObjectsPatchForbidden().
				WithPayload(errPayloadFromSingleErr(err))
		case kinds.ErrInvalidUserInput:
			return objects.NewObjectsUpdateUnprocessableEntity().
				WithPayload(errPayloadFromSingleErr(err))
		default:
			return objects.NewObjectsUpdateInternalServerError().
				WithPayload(errPayloadFromSingleErr(err))
		}
	}

	return objects.NewObjectsPatchNoContent()
}

func (h *kindHandlers) addObjectReference(params objects.ObjectsReferencesCreateParams,
	principal *models.Principal) middleware.Responder {
	err := h.manager.AddObjectReference(params.HTTPRequest.Context(), principal, params.ID, params.PropertyName, params.Body)
	if err != nil {
		switch err.(type) {
		case errors.Forbidden:
			return objects.NewObjectsReferencesCreateForbidden().
				WithPayload(errPayloadFromSingleErr(err))
		case kinds.ErrNotFound, kinds.ErrInvalidUserInput:
			return objects.NewObjectsReferencesCreateUnprocessableEntity().
				WithPayload(errPayloadFromSingleErr(err))
		default:
			return objects.NewObjectsReferencesCreateInternalServerError().
				WithPayload(errPayloadFromSingleErr(err))
		}
	}

	return objects.NewObjectsReferencesCreateOK()
}

func (h *kindHandlers) updateObjectReferences(params objects.ObjectsReferencesUpdateParams,
	principal *models.Principal) middleware.Responder {
	err := h.manager.UpdateObjectReferences(params.HTTPRequest.Context(), principal, params.ID, params.PropertyName, params.Body)
	if err != nil {
		switch err.(type) {
		case errors.Forbidden:
			return objects.NewObjectsReferencesUpdateForbidden().
				WithPayload(errPayloadFromSingleErr(err))
		case kinds.ErrNotFound, kinds.ErrInvalidUserInput:
			return objects.NewObjectsReferencesUpdateUnprocessableEntity().
				WithPayload(errPayloadFromSingleErr(err))
		default:
			return objects.NewObjectsReferencesUpdateInternalServerError().
				WithPayload(errPayloadFromSingleErr(err))
		}
	}

	return objects.NewObjectsReferencesUpdateOK()
}

func (h *kindHandlers) deleteObjectReference(params objects.ObjectsReferencesDeleteParams,
	principal *models.Principal) middleware.Responder {
	err := h.manager.DeleteObjectReference(params.HTTPRequest.Context(), principal, params.ID, params.PropertyName, params.Body)
	if err != nil {
		switch err.(type) {
		case errors.Forbidden:
			return objects.NewObjectsReferencesDeleteForbidden().
				WithPayload(errPayloadFromSingleErr(err))
		case kinds.ErrNotFound, kinds.ErrInvalidUserInput:
			return objects.NewObjectsReferencesDeleteNotFound().
				WithPayload(errPayloadFromSingleErr(err))
		default:
			return objects.NewObjectsReferencesDeleteInternalServerError().
				WithPayload(errPayloadFromSingleErr(err))
		}
	}

	return objects.NewObjectsReferencesDeleteNoContent()
}

func setupKindHandlers(api *operations.WeaviateAPI,
	manager *kinds.Manager, config config.Config, logger logrus.FieldLogger) {
	h := &kindHandlers{manager, logger, config}

	api.ObjectsObjectsCreateHandler = objects.
		ObjectsCreateHandlerFunc(h.addObject)
	api.ObjectsObjectsValidateHandler = objects.
		ObjectsValidateHandlerFunc(h.validateObject)
	api.ObjectsObjectsGetHandler = objects.
		ObjectsGetHandlerFunc(h.getObject)
	api.ObjectsObjectsDeleteHandler = objects.
		ObjectsDeleteHandlerFunc(h.deleteObject)
	api.ObjectsObjectsListHandler = objects.
		ObjectsListHandlerFunc(h.getObjects)
	api.ObjectsObjectsUpdateHandler = objects.
		ObjectsUpdateHandlerFunc(h.updateObject)
	api.ObjectsObjectsPatchHandler = objects.
		ObjectsPatchHandlerFunc(h.patchObject)
	api.ObjectsObjectsReferencesCreateHandler = objects.
		ObjectsReferencesCreateHandlerFunc(h.addObjectReference)
	api.ObjectsObjectsReferencesDeleteHandler = objects.
		ObjectsReferencesDeleteHandlerFunc(h.deleteObjectReference)
	api.ObjectsObjectsReferencesUpdateHandler = objects.
		ObjectsReferencesUpdateHandlerFunc(h.updateObjectReferences)
}

func (h *kindHandlers) extendSchemaWithAPILinks(schema map[string]interface{}) map[string]interface{} {
	if schema == nil {
		return schema
	}

	for key, value := range schema {
		asMultiRef, ok := value.(models.MultipleRef)
		if !ok {
			continue
		}

		schema[key] = h.extendReferencesWithAPILinks(asMultiRef)
	}
	return schema
}

func (h *kindHandlers) extendReferencesWithAPILinks(refs models.MultipleRef) models.MultipleRef {
	for i, ref := range refs {
		refs[i] = h.extendReferenceWithAPILink(ref)
	}

	return refs
}

func (h *kindHandlers) extendReferenceWithAPILink(ref *models.SingleRef) *models.SingleRef {
	parsed, err := crossref.Parse(ref.Beacon.String())
	if err != nil {
		// ignore return unchanged
		return ref
	}

	ref.Href = strfmt.URI(fmt.Sprintf("%s/v1/%ss/%s", h.config.Origin, parsed.Kind.Name(), parsed.TargetID))
	return ref
}

func parseIncludeParam(in *string) (traverser.UnderscoreProperties, error) {
	out := traverser.UnderscoreProperties{}
	if in == nil {
		return out, nil
	}

	parts := strings.Split(*in, ",")

	for _, prop := range parts {
		switch prop {
		case "_classification", "classification":
			out.Classification = true
			out.RefMeta = true
		case "_interpretation", "interpretation":
			out.Interpretation = true
		case "_nearestNeighbors", "nearestNeighbors", "nearestneighbors", "_nearestneighbors", "nearest-neighbors", "nearest_neighbors", "_nearest_neighbors":
			out.NearestNeighbors = true
		case "_featureProjection", "featureProjection", "featureprojection", "_featureprojection", "feature-projection", "feature_projection", "_feature_projection":
			out.FeatureProjection = &projector.Params{}
		case "_vector", "vector":
			out.Vector = true

		default:
			return out, fmt.Errorf("unrecognized property '%s' in ?include list", prop)
		}
	}

	return out, nil
}
