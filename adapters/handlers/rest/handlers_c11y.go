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

	middleware "github.com/go-openapi/runtime/middleware"
	"github.com/semi-technologies/weaviate/entities/models"

	"github.com/semi-technologies/weaviate/adapters/handlers/rest/operations"
	"github.com/semi-technologies/weaviate/adapters/handlers/rest/operations/contextionary_api"
)

type inspector interface {
	GetWords(ctx context.Context, words string) (*models.C11yWordsResponse, error)
}

type c11yProxy interface {
	AddExtension(ctx context.Context, extension *models.C11yExtension) error
}

func setupC11yHandlers(api *operations.WeaviateAPI, inspector inspector, proxy c11yProxy) {
	api.ContextionaryAPIC11yConceptsHandler = contextionary_api.C11yConceptsHandlerFunc(
		func(params contextionary_api.C11yConceptsParams, principal *models.Principal) middleware.Responder {
			ctx := params.HTTPRequest.Context()
			// Register the request

			res, err := inspector.GetWords(ctx, params.Concept)
			if err != nil {
				return contextionary_api.NewC11yConceptsBadRequest().WithPayload(errPayloadFromSingleErr(err))
			}

			return contextionary_api.NewC11yConceptsOK().WithPayload(res)
		})

	api.ContextionaryAPIC11yExtensionsHandler = contextionary_api.C11yExtensionsHandlerFunc(func(params contextionary_api.C11yExtensionsParams, principal *models.Principal) middleware.Responder {
		ctx := params.HTTPRequest.Context()

		err := proxy.AddExtension(ctx, params.Extension)
		if err != nil {
			// TODO: distinguish between 400 and 500, right now the grpc client always returns the same kind of error
			return contextionary_api.NewC11yExtensionsBadRequest().WithPayload(
				errPayloadFromSingleErr(err))
		}

		return contextionary_api.NewC11yExtensionsOK().WithPayload(params.Extension)
	})
}
