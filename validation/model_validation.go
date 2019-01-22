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

package validation

import (
	"context"
	"fmt"

	"github.com/creativesoftwarefdn/weaviate/config"
	dbconnector "github.com/creativesoftwarefdn/weaviate/database/connectors"
	connutils "github.com/creativesoftwarefdn/weaviate/database/connectors/utils"
	"github.com/creativesoftwarefdn/weaviate/database/schema"
	"github.com/creativesoftwarefdn/weaviate/database/schema/crossref"
	"github.com/creativesoftwarefdn/weaviate/database/schema/kind"
	"github.com/creativesoftwarefdn/weaviate/models"
	"github.com/creativesoftwarefdn/weaviate/network"
	"github.com/creativesoftwarefdn/weaviate/network/crossrefs"
)

const (
	// ErrorMissingActionThings message
	ErrorMissingActionThings string = "no things, object and subject, are added. Add 'things' by using the 'things' key in the root of the JSON"
	// ErrorMissingActionThingsObject message
	ErrorMissingActionThingsObject string = "no object-thing is added. Add the 'object' inside the 'things' part of the JSON"
	// ErrorMissingActionThingsSubject message
	ErrorMissingActionThingsSubject string = "no subject-thing is added. Add the 'subject' inside the 'things' part of the JSON"
	// ErrorMissingActionThingsObjectLocation message
	ErrorMissingActionThingsObjectLocation string = "no 'locationURL' is found in the object-thing. Add the 'locationURL' inside the 'object-thing' part of the JSON"
	// ErrorMissingActionThingsObjectType message
	ErrorMissingActionThingsObjectType string = "no 'type' is found in the object-thing. Add the 'type' inside the 'object-thing' part of the JSON"
	// ErrorInvalidActionThingsObjectType message
	ErrorInvalidActionThingsObjectType string = "object-thing requires one of the following values in 'type': '%s', '%s' or '%s'"
	// ErrorMissingActionThingsSubjectLocation message
	ErrorMissingActionThingsSubjectLocation string = "no 'locationURL' is found in the subject-thing. Add the 'locationURL' inside the 'subject-thing' part of the JSON"
	// ErrorMissingActionThingsSubjectType message
	ErrorMissingActionThingsSubjectType string = "no 'type' is found in the subject-thing. Add the 'type' inside the 'subject-thing' part of the JSON"
	// ErrorInvalidActionThingsSubjectType message
	ErrorInvalidActionThingsSubjectType string = "subject-thing requires one of the following values in 'type': '%s', '%s' or '%s'"
	// ErrorMissingClass message
	ErrorMissingClass string = "the given class is empty"
	// ErrorMissingContext message
	ErrorMissingContext string = "the given context is empty"
	// ErrorNoExternalCredentials message
	ErrorNoExternalCredentials string = "no credentials available for the Weaviate instance for %s given in the %s"
	// ErrorExternalNotFound message
	ErrorExternalNotFound string = "given statuscode of '%s' is '%d', but 200 was expected for LocationURL given in the %s"
	// ErrorInvalidCRefType message
	ErrorInvalidCRefType string = "'cref' type '%s' does not exists"
	// ErrorNotFoundInDatabase message
	ErrorNotFoundInDatabase string = "error finding the '%s' in the database: '%s' at %s"
)

// ValidateThingBody Validates a thing body using the 'ThingCreate' object.
func ValidateThingBody(ctx context.Context, thing *models.ThingCreate, databaseSchema schema.WeaviateSchema,
	dbConnector dbconnector.DatabaseConnector, network network.Network, serverConfig *config.WeaviateConfig) error {
	// Validate the body
	bve := validateBody(thing.AtClass, thing.AtContext)

	// Return error if possible
	if bve != nil {
		return bve
	}

	// Return the schema validation error
	sve := ValidateSchemaInBody(ctx, databaseSchema.ThingSchema.Schema, thing, connutils.RefTypeThing,
		dbConnector, network, serverConfig)

	return sve
}

// ValidateActionBody Validates a action body using the 'ActionCreate' object.
func ValidateActionBody(ctx context.Context, action *models.ActionCreate, databaseSchema schema.WeaviateSchema,
	dbConnector dbconnector.DatabaseConnector, network network.Network, serverConfig *config.WeaviateConfig,
) error {
	// Validate the body
	bve := validateBody(action.AtClass, action.AtContext)

	// Return error if possible
	if bve != nil {
		return bve
	}

	// Return the schema validation error
	sve := ValidateSchemaInBody(ctx, databaseSchema.ActionSchema.Schema, action, connutils.RefTypeAction,
		dbConnector, network, serverConfig)

	return sve
}

// validateBody Validates the overlapping body values
func validateBody(class string, context string) error {
	// If the given class is empty, return an error
	if class == "" {
		return fmt.Errorf(ErrorMissingClass)
	}

	// If the given context is empty, return an error
	if context == "" {
		return fmt.Errorf(ErrorMissingContext)
	}

	// No error
	return nil
}

// validateRefType validates the reference type with one of the existing reference types
func validateRefType(s string) bool {
	return (s == "things" || s == "actions")
}

// ValidateSingleRef validates a single ref based on location URL and existence of the object in the database
func ValidateSingleRef(ctx context.Context, serverConfig *config.WeaviateConfig, cref *models.SingleRef, dbConnector dbconnector.DatabaseConnector, network network.Network, errorVal string) error {

	ref, err := crossref.ParseSingleRef(cref)
	if err != nil {
		return fmt.Errorf("invalid reference: %s", err)
	}

	if !ref.Local {
		return validateNetworkRef(network, ref)
	}

	return validateLocalRef(ctx, dbConnector, ref, errorVal)
}

func validateLocalRef(ctx context.Context, dbConnector dbconnector.DatabaseConnector, ref *crossref.Ref, errorVal string) error {
	// Check whether the given Object exists in the DB
	var err error
	switch ref.Kind {
	case kind.THING_KIND:
		obj := &models.ThingGetResponse{}
		err = dbConnector.GetThing(ctx, ref.TargetID, obj)
	case kind.ACTION_KIND:
		obj := &models.ActionGetResponse{}
		err = dbConnector.GetAction(ctx, ref.TargetID, obj)
	}

	if err != nil {
		return fmt.Errorf(ErrorNotFoundInDatabase, ref.Kind.Name(), err, errorVal)
	}

	return nil
}

func validateNetworkRef(network network.Network, ref *crossref.Ref) error {
	// Network ref
	peers, err := network.ListPeers()
	if err != nil {
		return fmt.Errorf("could not validate network reference: could not list network peers: %s", err)
	}

	_, err = peers.RemoteKind(crossrefs.NetworkKind{Kind: ref.Kind, ID: ref.TargetID, PeerName: ref.PeerName})
	if err != nil {
		return fmt.Errorf("invalid network reference: %s", err)
	}

	return nil
}

func ValidateMultipleRef(ctx context.Context, serverConfig *config.WeaviateConfig,
	refs *models.MultipleRef, dbConnector dbconnector.DatabaseConnector, network network.Network,
	errorVal string) error {
	if refs == nil {
		return nil
	}

	for _, ref := range *refs {
		err := ValidateSingleRef(ctx, serverConfig, ref, dbConnector, network, errorVal)
		if err != nil {
			return err
		}
	}
	return nil
}
