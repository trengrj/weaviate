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
 */

// Package rest with all rest API functions.
package rest

import (
	"crypto/tls"
	"io/ioutil"
	"log"

	"github.com/creativesoftwarefdn/weaviate/adapters/handlers/rest/state"

	"github.com/go-openapi/swag"
	"google.golang.org/grpc/grpclog"

	"github.com/creativesoftwarefdn/weaviate/adapters/handlers/rest/operations"
	"github.com/creativesoftwarefdn/weaviate/database"
	"github.com/creativesoftwarefdn/weaviate/usecases/config"

	libcontextionary "github.com/creativesoftwarefdn/weaviate/contextionary"
)

var connectorOptionGroup *swag.CommandLineOptionsGroup

// rawContextionary is the contextionary as we read it from the files. It is
// not extended by schema builds. It is important to keep this untouched copy,
// so that we can rebuild a clean contextionary on every schema change based on
// this contextionary and the current schema
var rawContextionary libcontextionary.Contextionary

// contextionary is the contextionary we keep amending on every schema change
var contextionary libcontextionary.Contextionary

var appState *state.State

var db database.Database

func init() {
	appState = &state.State{}

	discard := ioutil.Discard
	myGRPCLogger := log.New(discard, "", log.LstdFlags)
	grpclog.SetLogger(myGRPCLogger)
}

// configureAPI -> see configure_api.go

// configureServer -> see configure_server.go

func configureFlags(api *operations.WeaviateAPI) {
	connectorOptionGroup = config.GetConfigOptionGroup()

	api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{
		*connectorOptionGroup,
	}
}

// The TLS configuration before HTTPS server starts.
func configureTLS(tlsConfig *tls.Config) {
	// Make all necessary changes to the TLS configuration here.
}
