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
 */package restapi

import (
	"log"
	"net/http"

	"github.com/creativesoftwarefdn/weaviate/graphqlapi/graphiql"
	"github.com/creativesoftwarefdn/weaviate/lib/feature_flags"
	"github.com/creativesoftwarefdn/weaviate/restapi/swagger_middleware"
	"github.com/rs/cors"
)

// The middleware configuration is for the handler executors. These do not apply to the swagger.json document.
// The middleware executes after routing but before authentication, binding and validation
func setupMiddlewares(handler http.Handler) http.Handler {
	// Rewrite / workaround because of issue with handling two API keys
	return handler
}

// The middleware configuration happens before anything, this middleware also applies to serving the swagger.json document.
// So this is a good place to plug in a panic handling middleware, logging and metrics
// Contains "x-api-key", "x-api-token" for legacy reasons, older interfaces might need these headers.
func setupGlobalMiddleware(handler http.Handler) http.Handler {
	handleCORS := cors.New(cors.Options{
		OptionsPassthrough: true,
		AllowedHeaders:     []string{"*"},
	}).Handler
	handler = handleCORS(handler)

	if feature_flags.EnableDevUI {
		handler = graphiql.AddMiddleware(handler)
		handler = swagger_middleware.AddMiddleware([]byte(SwaggerJSON), handler)
	}

	handler = addLogging(handler)
	handler = addPreflight(handler)

	return handler
}

func addLogging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if serverConfig.Environment.Debug {
			log.Printf("Received request: %+v %+v\n", r.Method, r.URL)
		}
		next.ServeHTTP(w, r)
	})
}

func addPreflight(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		if r.Method == "OPTIONS" {
			w.Header().Set("Access-Control-Allow-Origin", "*")
			w.Header().Set("Access-Control-Allow-Methods", "*")
			w.Header().Set("Access-Control-Allow-Headers", "*")
			return
		}

		next.ServeHTTP(w, r)
	})
}
