/*                          _       _
 *__      _____  __ ___   ___  __ _| |_ ___
 *\ \ /\ / / _ \/ _` \ \ / / |/ _` | __/ _ \
 * \ V  V /  __/ (_| |\ V /| | (_| | ||  __/
 *  \_/\_/ \___|\__,_| \_/ |_|\__,_|\__\___|
 *
 * Copyright © 2016 Weaviate. All rights reserved.
 * LICENSE: https://github.com/creativesoftwarefdn/weaviate/blob/develop/LICENSE.md
 * AUTHOR: Bob van Luijt (bob@weaviate.com)
 * See www.weaviate.com for details
 * Contact: @ CreativeSofwFdn / yourfriends@weaviate.com
 */

// Package restapi with all rest API functions.
package restapi

import (
	"crypto/tls"
	"encoding/json"
	errors_ "errors"
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"net/http"
	"os"
	"strings"

	"github.com/creativesoftwarefdn/weaviate/restapi/operations/graphql"
	"github.com/creativesoftwarefdn/weaviate/restapi/operations/meta"

	jsonpatch "github.com/evanphx/json-patch"
	errors "github.com/go-openapi/errors"
	runtime "github.com/go-openapi/runtime"
	middleware "github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/runtime/yamlpc"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	gographql "github.com/graphql-go/graphql"
	graceful "github.com/tylerb/graceful"
	"google.golang.org/grpc/grpclog"

	"github.com/creativesoftwarefdn/weaviate/config"
	"github.com/creativesoftwarefdn/weaviate/connectors"
	"github.com/creativesoftwarefdn/weaviate/connectors/dgraph"
	"github.com/creativesoftwarefdn/weaviate/connectors/foobar"
	"github.com/creativesoftwarefdn/weaviate/connectors/gremlin"
	"github.com/creativesoftwarefdn/weaviate/connectors/kvcache"
	"github.com/creativesoftwarefdn/weaviate/connectors/utils"
	"github.com/creativesoftwarefdn/weaviate/graphqlapi"
	"github.com/creativesoftwarefdn/weaviate/messages"
	"github.com/creativesoftwarefdn/weaviate/models"
	"github.com/creativesoftwarefdn/weaviate/restapi/operations"
	"github.com/creativesoftwarefdn/weaviate/restapi/operations/actions"
	"github.com/creativesoftwarefdn/weaviate/restapi/operations/keys"
	"github.com/creativesoftwarefdn/weaviate/restapi/operations/things"
	"github.com/creativesoftwarefdn/weaviate/schema"
	"github.com/creativesoftwarefdn/weaviate/validation"
)

const pageOverride int = 1

var connectorOptionGroup *swag.CommandLineOptionsGroup
var databaseSchema schema.WeaviateSchema
var serverConfig config.WeaviateConfig
var dbConnector dbconnector.DatabaseConnector
var graphQLSchema *graphqlapi.GraphQLSchema
var messaging *messages.Messaging

func init() {
	discard := ioutil.Discard
	myGRPCLogger := log.New(discard, "", log.LstdFlags)
	grpclog.SetLogger(myGRPCLogger)

	// Create temp folder if it does not exist
	tempFolder := "temp"
	if _, err := os.Stat(tempFolder); os.IsNotExist(err) {
		messaging.InfoMessage("Temp folder created...")
		os.Mkdir(tempFolder, 0766)
	}
}

// getLimit returns the maximized limit
func getLimit(paramMaxResults *int64) int {
	maxResults := int64(connutils.DefaultFirst)
	// Get the max results from params, if exists
	if paramMaxResults != nil {
		maxResults = *paramMaxResults
	}

	// Max results form URL, otherwise max = connutils.DefaultFirst.
	return int(math.Min(float64(maxResults), float64(connutils.DefaultFirst)))
}

// getPage returns the page if set
func getPage(paramPage *int64) int {
	page := int64(pageOverride)
	// Get the page from params, if exists
	if paramPage != nil {
		page = *paramPage
	}

	// Page form URL, otherwise max = connutils.DefaultFirst.
	return int(page)
}

// isOwnKeyOrLowerInTree returns whether a key is his own or in his children
func isOwnKeyOrLowerInTree(currentKey models.KeyTokenGetResponse, userKeyID strfmt.UUID, databaseConnector dbconnector.DatabaseConnector) bool {
	// If is own key, return true
	if strings.EqualFold(string(userKeyID), string(currentKey.KeyID)) {
		return true
	}

	// Get all child id's
	childIDs := []strfmt.UUID{}
	childIDs, _ = GetKeyChildrenUUIDs(databaseConnector, currentKey.KeyID, true, childIDs, 0, 0)

	// Check ID is in childIds
	isChildID := false
	for _, childID := range childIDs {
		if childID == userKeyID {
			isChildID = true
		}
	}

	// If ID is in the child ID's, you are allowed to do the action
	if isChildID {
		return true
	}

	return false
}

// GetKeyChildrenUUIDs returns children recursivly based on its parameters.
func GetKeyChildrenUUIDs(databaseConnector dbconnector.DatabaseConnector, parentUUID strfmt.UUID, filterOutDeleted bool, allIDs []strfmt.UUID, maxDepth int, depth int) ([]strfmt.UUID, error) {
	// Append on every depth
	if depth > 0 {
		allIDs = append(allIDs, parentUUID)
	}

	// Init children var
	children := []*models.KeyTokenGetResponse{}

	// Get children from the db-connector
	err := databaseConnector.GetKeyChildren(parentUUID, &children)

	// Return error
	if err != nil {
		return allIDs, err
	}

	// For every depth, get the ID's
	if maxDepth == 0 || depth < maxDepth {
		for _, child := range children {
			allIDs, err = GetKeyChildrenUUIDs(databaseConnector, child.KeyID, filterOutDeleted, allIDs, maxDepth, depth+1)
		}
	}

	return allIDs, err
}

func generateMultipleRefObject(keyIDs []strfmt.UUID) models.MultipleRef {
	// Init the response
	refs := models.MultipleRef{}

	// Init localhost
	url := serverConfig.GetHostAddress()

	// Generate SingleRefs
	for _, keyID := range keyIDs {
		refs = append(refs, &models.SingleRef{
			LocationURL:  &url,
			NrDollarCref: keyID,
			Type:         connutils.RefTypeKey,
		})
	}

	return refs
}

func deleteKey(databaseConnector dbconnector.DatabaseConnector, parentUUID strfmt.UUID) {
	// Find its children
	var allIDs []strfmt.UUID

	// Get all the children to remove
	allIDs, _ = GetKeyChildrenUUIDs(databaseConnector, parentUUID, false, allIDs, 0, 0)

	// Append the children to the parent UUIDs to remove all
	allIDs = append(allIDs, parentUUID)

	// Delete for every child
	for _, keyID := range allIDs {
		go databaseConnector.DeleteKey(keyID)
	}
}

// GetAllConnectors contains all available connectors
func GetAllConnectors() []dbconnector.DatabaseConnector {
	// Set all existing connectors
	connectors := []dbconnector.DatabaseConnector{
		&dgraph.Dgraph{},
		&gremlin.Gremlin{},
		&foobar.Foobar{},
	}

	return connectors
}

// GetAllCacheConnectors contains all available cache-connectors
func GetAllCacheConnectors() []dbconnector.CacheConnector {
	// Set all existing connectors
	connectors := []dbconnector.CacheConnector{
		&kvcache.KVCache{},
	}

	return connectors
}

// CreateDatabaseConnector gets the database connector by name from config
func CreateDatabaseConnector(env *config.Environment) dbconnector.DatabaseConnector {
	// Get all connectors
	connectors := GetAllConnectors()
	cacheConnectors := GetAllCacheConnectors()

	// Init the db-connector variable
	var connector dbconnector.DatabaseConnector

	// Loop through all connectors and determine its name
	for _, c := range connectors {
		if c.GetName() == env.Database.Name {
			messaging.InfoMessage(fmt.Sprintf("Using database '%s'", env.Database.Name))
			connector = c
			break
		}
	}

	// Loop through all cache-connectors and determine its name
	for _, cc := range cacheConnectors {
		if cc.GetName() == env.Cache.Name {
			messaging.InfoMessage(fmt.Sprintf("Using cache layer '%s'", env.Cache.Name))
			cc.SetDatabaseConnector(connector)
			connector = cc
			break
		}
	}

	return connector
}

// ActionsAllowed returns information whether an action is allowed based on given several input vars.
func ActionsAllowed(actions []string, validateObject interface{}, databaseConnector dbconnector.DatabaseConnector, objectOwnerUUID interface{}) (bool, error) {
	// Get the user by the given principal
	keyObject := validateObject.(models.KeyTokenGetResponse)

	// Check whether the given owner of the object is in the children, if the ownerID is given
	correctChild := false
	if objectOwnerUUID != nil {
		correctChild = isOwnKeyOrLowerInTree(keyObject, objectOwnerUUID.(strfmt.UUID), databaseConnector)
	} else {
		correctChild = true
	}

	// Return false if the object's owner is not the logged in user or one of its childs.
	if !correctChild {
		return false, errors_.New("the object does not belong to the given token or to one of the token's children")
	}

	// All possible actions in a map to check it more easily
	actionsToCheck := map[string]bool{
		"read":    false,
		"write":   false,
		"execute": false,
		"delete":  false,
	}

	// Add 'true' if an action has to be checked on its rights.
	for _, action := range actions {
		actionsToCheck[action] = true
	}

	// Check every action on its rights, if rights are needed and the key has not that kind of rights, return false.
	if actionsToCheck["read"] && !keyObject.Read {
		return false, errors_.New("read rights are needed to perform this action")
	}

	// Idem
	if actionsToCheck["write"] && !keyObject.Write {
		return false, errors_.New("write rights are needed to perform this action")
	}

	// Idem
	if actionsToCheck["delete"] && !keyObject.Delete {
		return false, errors_.New("delete rights are needed to perform this action")
	}

	// Idem
	if actionsToCheck["execute"] && !keyObject.Execute {
		return false, errors_.New("execute rights are needed to perform this action")
	}

	return true, nil
}

func configureFlags(api *operations.WeaviateAPI) {
	connectorOptionGroup = config.GetConfigOptionGroup()

	api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{
		*connectorOptionGroup,
	}
}

// createErrorResponseObject is a common function to create an error response
func createErrorResponseObject(message string) *models.ErrorResponse {
	// Initialize return value
	er := &models.ErrorResponse{}

	// Fill the error with the message
	er.Error = &models.ErrorResponseError{
		Message: message,
	}

	return er
}

func configureAPI(api *operations.WeaviateAPI) http.Handler {
	api.ServeError = errors.ServeError

	api.JSONConsumer = runtime.JSONConsumer()

	api.BinConsumer = runtime.ByteStreamConsumer()

	api.UrlformConsumer = runtime.DiscardConsumer

	api.YamlConsumer = yamlpc.YAMLConsumer()

	api.XMLConsumer = runtime.XMLConsumer()

	api.MultipartformConsumer = runtime.DiscardConsumer

	api.TxtConsumer = runtime.TextConsumer()

	api.JSONProducer = runtime.JSONProducer()

	api.BinProducer = runtime.ByteStreamProducer()

	api.UrlformProducer = runtime.DiscardProducer

	api.YamlProducer = yamlpc.YAMLProducer()

	api.XMLProducer = runtime.XMLProducer()

	api.MultipartformProducer = runtime.DiscardProducer

	api.TxtProducer = runtime.TextProducer()

	/*
	 * HANDLE X-API-KEY
	 */
	// Applies when the "X-API-KEY" header is set
	api.APIKeyAuth = func(token string) (interface{}, error) {
		// Create key
		validatedKey := models.KeyTokenGetResponse{}

		// Check if the user has access, true if yes
		err := dbConnector.ValidateToken(strfmt.UUID(token), &validatedKey)

		// Error printing
		if err != nil {
			return nil, errors.New(401, err.Error())
		}

		// Validate the key on expiry time
		currentUnix := connutils.NowUnix()
		if validatedKey.KeyExpiresUnix != -1 && validatedKey.KeyExpiresUnix < currentUnix {
			return nil, errors.New(401, "Provided key has expired.")
		}

		// key is valid, next step is allowing per Handler handling
		return validatedKey, nil
	}

	/*
	 * HANDLE EVENTS
	 */
	api.ActionsWeaviateActionsGetHandler = actions.WeaviateActionsGetHandlerFunc(func(params actions.WeaviateActionsGetParams, principal interface{}) middleware.Responder {
		// Initialize response
		actionGetResponse := models.ActionGetResponse{}
		actionGetResponse.Schema = map[string]models.JSONObject{}
		actionGetResponse.Things = &models.ObjectSubject{}

		// Get item from database
		err := dbConnector.GetAction(params.ActionID, &actionGetResponse)

		// Object is deleted
		if err != nil {
			return actions.NewWeaviateActionsGetNotFound()
		}

		// This is a read function, validate if allowed to read?
		if allowed, _ := ActionsAllowed([]string{"read"}, principal, dbConnector, actionGetResponse.Key.NrDollarCref); !allowed {
			return actions.NewWeaviateActionsGetForbidden()
		}

		// Get is successful
		return actions.NewWeaviateActionsGetOK().WithPayload(&actionGetResponse)
	})
	api.ActionsWeaviateActionsPatchHandler = actions.WeaviateActionsPatchHandlerFunc(func(params actions.WeaviateActionsPatchParams, principal interface{}) middleware.Responder {
		// Initialize response
		actionGetResponse := models.ActionGetResponse{}
		actionGetResponse.Schema = map[string]models.JSONObject{}
		actionGetResponse.Things = &models.ObjectSubject{}

		// Get and transform object
		UUID := strfmt.UUID(params.ActionID)
		errGet := dbConnector.GetAction(UUID, &actionGetResponse)
		actionGetResponse.LastUpdateTimeUnix = connutils.NowUnix()

		// Return error if UUID is not found.
		if errGet != nil {
			return actions.NewWeaviateActionsPatchNotFound()
		}

		// This is a write function, validate if allowed to write?
		if allowed, _ := ActionsAllowed([]string{"write"}, principal, dbConnector, actionGetResponse.Key.NrDollarCref); !allowed {
			return actions.NewWeaviateActionsPatchForbidden()
		}

		// Get PATCH params in format RFC 6902
		jsonBody, marshalErr := json.Marshal(params.Body)
		patchObject, decodeErr := jsonpatch.DecodePatch([]byte(jsonBody))

		if marshalErr != nil || decodeErr != nil {
			return actions.NewWeaviateActionsPatchBadRequest()
		}

		// Convert ActionGetResponse object to JSON
		actionUpdateJSON, marshalErr := json.Marshal(actionGetResponse)
		if marshalErr != nil {
			return actions.NewWeaviateActionsPatchBadRequest()
		}

		// Apply the patch
		updatedJSON, applyErr := patchObject.Apply(actionUpdateJSON)

		if applyErr != nil {
			return actions.NewWeaviateActionsPatchUnprocessableEntity().WithPayload(createErrorResponseObject(applyErr.Error()))
		}

		// Turn it into a Action object
		action := &models.Action{}
		json.Unmarshal([]byte(updatedJSON), &action)

		// Validate schema made after patching with the weaviate schema
		validatedErr := validation.ValidateActionBody(&action.ActionCreate, databaseSchema, dbConnector)
		if validatedErr != nil {
			return actions.NewWeaviateActionsPatchUnprocessableEntity().WithPayload(createErrorResponseObject(validatedErr.Error()))
		}

		// Update the database
		insertErr := dbConnector.UpdateAction(action, UUID)
		if insertErr != nil {
			messaging.ErrorMessage(insertErr)
		}

		// Create return Object
		actionGetResponse.Action = *action

		return actions.NewWeaviateActionsPatchOK().WithPayload(&actionGetResponse)
	})
	api.ActionsWeaviateActionsValidateHandler = actions.WeaviateActionsValidateHandlerFunc(func(params actions.WeaviateActionsValidateParams, principal interface{}) middleware.Responder {
		// Validate schema given in body with the weaviate schema
		validatedErr := validation.ValidateActionBody(&params.Body.ActionCreate, databaseSchema, dbConnector)
		if validatedErr != nil {
			return actions.NewWeaviateActionsValidateUnprocessableEntity().WithPayload(createErrorResponseObject(validatedErr.Error()))
		}

		return actions.NewWeaviateActionsValidateOK()
	})
	api.ActionsWeaviateActionsCreateHandler = actions.WeaviateActionsCreateHandlerFunc(func(params actions.WeaviateActionsCreateParams, principal interface{}) middleware.Responder {
		// This is a read function, validate if allowed to read?
		if allowed, _ := ActionsAllowed([]string{"write"}, principal, dbConnector, nil); !allowed {
			return actions.NewWeaviateActionsCreateForbidden()
		}

		// Generate UUID for the new object
		UUID := connutils.GenerateUUID()

		// Validate schema given in body with the weaviate schema
		validatedErr := validation.ValidateActionBody(params.Body, databaseSchema, dbConnector)
		if validatedErr != nil {
			return actions.NewWeaviateActionsCreateUnprocessableEntity().WithPayload(createErrorResponseObject(validatedErr.Error()))
		}

		// Create Key-ref-Object
		url := serverConfig.GetHostAddress()
		keyRef := &models.SingleRef{
			LocationURL:  &url,
			NrDollarCref: principal.(models.KeyTokenGetResponse).KeyID,
			Type:         connutils.RefTypeKey,
		}

		// Make Action-Object
		actionCreateJSON, _ := json.Marshal(params.Body)
		action := &models.Action{}
		json.Unmarshal([]byte(actionCreateJSON), action)

		action.CreationTimeUnix = connutils.NowUnix()
		action.LastUpdateTimeUnix = 0
		action.Key = keyRef

		// Save to DB, this needs to be a Go routine because we will return an accepted
		insertErr := dbConnector.AddAction(action, UUID)
		if insertErr != nil {
			messaging.ErrorMessage(insertErr)
		}

		// Initialize a response object
		responseObject := &models.ActionGetResponse{}
		responseObject.Action = *action
		responseObject.ActionID = UUID

		// Return SUCCESS (NOTE: this is ACCEPTED, so the databaseConnector.Add should have a go routine)
		return actions.NewWeaviateActionsCreateAccepted().WithPayload(responseObject)
	})
	api.ActionsWeaviateActionsDeleteHandler = actions.WeaviateActionsDeleteHandlerFunc(func(params actions.WeaviateActionsDeleteParams, principal interface{}) middleware.Responder {
		// Initialize response
		actionGetResponse := models.ActionGetResponse{}
		actionGetResponse.Schema = map[string]models.JSONObject{}
		actionGetResponse.Things = &models.ObjectSubject{}

		// Get item from database
		errGet := dbConnector.GetAction(params.ActionID, &actionGetResponse)

		// Not found
		if errGet != nil {
			return actions.NewWeaviateActionsDeleteNotFound()
		}

		// This is a delete function, validate if allowed to delete?
		if allowed, _ := ActionsAllowed([]string{"delete"}, principal, dbConnector, actionGetResponse.Key.NrDollarCref); !allowed {
			return things.NewWeaviateThingsDeleteForbidden()
		}

		// Add new row as GO-routine
		go dbConnector.DeleteAction(params.ActionID)

		// Return 'No Content'
		return actions.NewWeaviateActionsDeleteNoContent()
	})

	/*
	 * HANDLE KEYS
	 */
	api.KeysWeaviateKeyCreateHandler = keys.WeaviateKeyCreateHandlerFunc(func(params keys.WeaviateKeyCreateParams, principal interface{}) middleware.Responder {
		// Create current User object from principal
		key := principal.(models.KeyTokenGetResponse)

		// Fill the new User object
		url := serverConfig.GetHostAddress()
		newKey := &models.KeyTokenGetResponse{}
		newKey.KeyID = connutils.GenerateUUID()
		newKey.Token = connutils.GenerateUUID()
		newKey.Parent = &models.SingleRef{
			LocationURL:  &url,
			NrDollarCref: principal.(models.KeyTokenGetResponse).KeyID,
			Type:         connutils.RefTypeKey,
		}
		newKey.KeyCreate = *params.Body

		// Key expiry time is in the past
		currentUnix := connutils.NowUnix()
		if newKey.KeyExpiresUnix != -1 && newKey.KeyExpiresUnix < currentUnix {
			return keys.NewWeaviateKeyCreateUnprocessableEntity().WithPayload(createErrorResponseObject("Key expiry time is in the past."))
		}

		// Key expiry time is later than the expiry time of parent
		if key.KeyExpiresUnix != -1 && key.KeyExpiresUnix < newKey.KeyExpiresUnix {
			return keys.NewWeaviateKeyCreateUnprocessableEntity().WithPayload(createErrorResponseObject("Key expiry time is later than the expiry time of parent."))
		}

		// Save to DB, this needs to be a Go routine because we will return an accepted
		insertErr := dbConnector.AddKey(&newKey.Key, newKey.KeyID, newKey.Token)
		if insertErr != nil {
			messaging.ErrorMessage(insertErr)
		}

		// Return SUCCESS (NOTE: this is ACCEPTED, so the databaseConnector.Add should have a go routine)
		return keys.NewWeaviateKeyCreateAccepted().WithPayload(newKey)
	})
	api.KeysWeaviateKeysChildrenGetHandler = keys.WeaviateKeysChildrenGetHandlerFunc(func(params keys.WeaviateKeysChildrenGetParams, principal interface{}) middleware.Responder {
		// Initialize response
		keyResponse := models.KeyTokenGetResponse{}

		// First check on 'not found', otherwise it will say 'forbidden' in stead of 'not found'
		errGet := dbConnector.GetKey(params.KeyID, &keyResponse)

		// Not found
		if errGet != nil {
			return keys.NewWeaviateKeysChildrenGetNotFound()
		}

		// Check on permissions
		keyObject, _ := principal.(models.KeyTokenGetResponse)
		if !isOwnKeyOrLowerInTree(keyObject, params.KeyID, dbConnector) {
			return keys.NewWeaviateKeysChildrenGetForbidden()
		}

		// Get the children
		childIDs := []strfmt.UUID{}
		childIDs, _ = GetKeyChildrenUUIDs(dbConnector, params.KeyID, true, childIDs, 1, 0)

		// Initiate response object
		responseObject := &models.KeyChildrenGetResponse{}
		responseObject.Children = generateMultipleRefObject(childIDs)

		// Return children with 'OK'
		return keys.NewWeaviateKeysChildrenGetOK().WithPayload(responseObject)
	})
	api.KeysWeaviateKeysDeleteHandler = keys.WeaviateKeysDeleteHandlerFunc(func(params keys.WeaviateKeysDeleteParams, principal interface{}) middleware.Responder {
		// Initialize response
		keyResponse := models.KeyTokenGetResponse{}

		// First check on 'not found', otherwise it will say 'forbidden' in stead of 'not found'
		errGet := dbConnector.GetKey(params.KeyID, &keyResponse)

		// Not found
		if errGet != nil {
			return keys.NewWeaviateKeysDeleteNotFound()
		}

		// Check on permissions, only delete allowed if lower in tree (not own key)
		keyObject, _ := principal.(models.KeyTokenGetResponse)
		if !isOwnKeyOrLowerInTree(keyObject, params.KeyID, dbConnector) || keyObject.KeyID == params.KeyID {
			return keys.NewWeaviateKeysDeleteForbidden()
		}

		// Remove key from database if found
		deleteKey(dbConnector, params.KeyID)

		// Return 'No Content'
		return keys.NewWeaviateKeysDeleteNoContent()
	})
	api.KeysWeaviateKeysGetHandler = keys.WeaviateKeysGetHandlerFunc(func(params keys.WeaviateKeysGetParams, principal interface{}) middleware.Responder {
		// Initialize response
		keyResponse := models.KeyTokenGetResponse{}

		// Get item from database
		err := dbConnector.GetKey(params.KeyID, &keyResponse)

		// Object is deleted or not-existing
		if err != nil {
			return keys.NewWeaviateKeysGetNotFound()
		}

		// Check on permissions
		keyObject, _ := principal.(models.KeyTokenGetResponse)
		if !isOwnKeyOrLowerInTree(keyObject, params.KeyID, dbConnector) {
			return keys.NewWeaviateKeysGetForbidden()
		}

		// Get is successful
		return keys.NewWeaviateKeysGetOK().WithPayload(&keyResponse.KeyGetResponse)
	})
	api.KeysWeaviateKeysMeChildrenGetHandler = keys.WeaviateKeysMeChildrenGetHandlerFunc(func(params keys.WeaviateKeysMeChildrenGetParams, principal interface{}) middleware.Responder {
		// First check on 'not found', otherwise it will say 'forbidden' in stead of 'not found'
		currentKey := principal.(models.KeyTokenGetResponse)

		// Get the children
		childIDs := []strfmt.UUID{}
		childIDs, _ = GetKeyChildrenUUIDs(dbConnector, currentKey.KeyID, true, childIDs, 1, 0)

		// Initiate response object
		responseObject := &models.KeyChildrenGetResponse{}
		responseObject.Children = generateMultipleRefObject(childIDs)

		// Return children with 'OK'
		return keys.NewWeaviateKeysMeChildrenGetOK().WithPayload(responseObject)
	})
	api.KeysWeaviateKeysMeGetHandler = keys.WeaviateKeysMeGetHandlerFunc(func(params keys.WeaviateKeysMeGetParams, principal interface{}) middleware.Responder {
		// Initialize response object
		tokenResponseObject := models.KeyTokenGetResponse{}

		// Get item from database
		err := dbConnector.GetKey(principal.(models.KeyTokenGetResponse).KeyID, &tokenResponseObject)

		// Object is deleted or not-existing
		if err != nil {
			return keys.NewWeaviateKeysGetNotFound()
		}

		// Get is successful
		return keys.NewWeaviateKeysMeGetOK().WithPayload(&tokenResponseObject)
	})

	/*
	 * HANDLE THINGS
	 */
	api.ThingsWeaviateThingsCreateHandler = things.WeaviateThingsCreateHandlerFunc(func(params things.WeaviateThingsCreateParams, principal interface{}) middleware.Responder {
		// This is a write function, validate if allowed to write?
		if allowed, _ := ActionsAllowed([]string{"write"}, principal, dbConnector, nil); !allowed {
			return things.NewWeaviateThingsCreateForbidden()
		}

		// Generate UUID for the new object
		UUID := connutils.GenerateUUID()

		// Validate schema given in body with the weaviate schema
		validatedErr := validation.ValidateThingBody(params.Body, databaseSchema, dbConnector)
		if validatedErr != nil {
			return things.NewWeaviateThingsCreateUnprocessableEntity().WithPayload(createErrorResponseObject(validatedErr.Error()))
		}

		// Create Key-ref-Object
		url := serverConfig.GetHostAddress()
		keyRef := &models.SingleRef{
			LocationURL:  &url,
			NrDollarCref: principal.(models.KeyTokenGetResponse).KeyID,
			Type:         connutils.RefTypeKey,
		}

		// Make Thing-Object
		thingCreateJSON, _ := json.Marshal(params.Body)
		thing := &models.Thing{}
		json.Unmarshal([]byte(thingCreateJSON), thing)
		thing.CreationTimeUnix = connutils.NowUnix()
		thing.LastUpdateTimeUnix = 0
		thing.Key = keyRef

		// Save to DB, this needs to be a Go routine because we will return an accepted
		insertErr := dbConnector.AddThing(thing, UUID)
		if insertErr != nil {
			messaging.ErrorMessage(insertErr)
		}

		// Create response Object from create object.
		responseObject := &models.ThingGetResponse{}
		responseObject.Thing = *thing
		responseObject.ThingID = UUID

		// Return SUCCESS (NOTE: this is ACCEPTED, so the dbConnector.Add should have a go routine)
		return things.NewWeaviateThingsCreateAccepted().WithPayload(responseObject)
	})
	api.ThingsWeaviateThingsDeleteHandler = things.WeaviateThingsDeleteHandlerFunc(func(params things.WeaviateThingsDeleteParams, principal interface{}) middleware.Responder {
		// Initialize response
		thingGetResponse := models.ThingGetResponse{}
		thingGetResponse.Schema = map[string]models.JSONObject{}

		// Get item from database
		errGet := dbConnector.GetThing(params.ThingID, &thingGetResponse)

		// Not found
		if errGet != nil {
			return things.NewWeaviateThingsDeleteNotFound()
		}

		// This is a delete function, validate if allowed to delete?
		if allowed, _ := ActionsAllowed([]string{"delete"}, principal, dbConnector, thingGetResponse.Key.NrDollarCref); !allowed {
			return things.NewWeaviateThingsDeleteForbidden()
		}

		// Delete the Actions
		actionsExist := true
		for actionsExist {
			actions := models.ActionsListResponse{}
			dbConnector.ListActions(params.ThingID, 50, 0, []*connutils.WhereQuery{}, &actions)
			for _, v := range actions.Actions {
				go dbConnector.DeleteAction(v.ActionID)
			}
			actionsExist = actions.TotalResults > 0
		}

		// Add new row as GO-routine
		go dbConnector.DeleteThing(params.ThingID)

		// Return 'No Content'
		return things.NewWeaviateThingsDeleteNoContent()
	})
	api.ThingsWeaviateThingsGetHandler = things.WeaviateThingsGetHandlerFunc(func(params things.WeaviateThingsGetParams, principal interface{}) middleware.Responder {
		// Initialize response
		responseObject := models.ThingGetResponse{}
		responseObject.Schema = map[string]models.JSONObject{}

		// Get item from database
		err := dbConnector.GetThing(strfmt.UUID(params.ThingID), &responseObject)

		// Object is not found
		if err != nil {
			return things.NewWeaviateThingsGetNotFound()
		}

		// This is a read function, validate if allowed to read?
		if allowed, _ := ActionsAllowed([]string{"read"}, principal, dbConnector, responseObject.Key.NrDollarCref); !allowed {
			return things.NewWeaviateThingsGetForbidden()
		}

		// Get is successful
		return things.NewWeaviateThingsGetOK().WithPayload(&responseObject)
	})
	api.ThingsWeaviateThingsListHandler = things.WeaviateThingsListHandlerFunc(func(params things.WeaviateThingsListParams, principal interface{}) middleware.Responder {
		// This is a read function, validate if allowed to read?
		if allowed, _ := ActionsAllowed([]string{"read"}, principal, dbConnector, nil); !allowed {
			return things.NewWeaviateThingsListForbidden()
		}

		// Get limit and page
		limit := getLimit(params.MaxResults)
		page := getPage(params.Page)

		// Get user out of principal
		keyID := principal.(models.KeyTokenGetResponse).KeyID

		// Initialize response
		thingsResponse := models.ThingsListResponse{}
		thingsResponse.Things = []*models.ThingGetResponse{}

		// List all results
		err := dbConnector.ListThings(limit, (page-1)*limit, keyID, []*connutils.WhereQuery{}, &thingsResponse)

		if err != nil {
			messaging.ErrorMessage(err)
		}

		return things.NewWeaviateThingsListOK().WithPayload(&thingsResponse)
	})
	api.ThingsWeaviateThingsPatchHandler = things.WeaviateThingsPatchHandlerFunc(func(params things.WeaviateThingsPatchParams, principal interface{}) middleware.Responder {
		// Initialize response
		thingGetResponse := models.ThingGetResponse{}
		thingGetResponse.Schema = map[string]models.JSONObject{}

		// Get and transform object
		UUID := strfmt.UUID(params.ThingID)
		errGet := dbConnector.GetThing(UUID, &thingGetResponse)
		thingGetResponse.LastUpdateTimeUnix = connutils.NowUnix()

		// Return error if UUID is not found.
		if errGet != nil {
			return things.NewWeaviateThingsPatchNotFound()
		}

		// This is a write function, validate if allowed to write?
		if allowed, _ := ActionsAllowed([]string{"write"}, principal, dbConnector, thingGetResponse.Key.NrDollarCref); !allowed {
			return things.NewWeaviateThingsPatchForbidden()
		}

		// Get PATCH params in format RFC 6902
		jsonBody, marshalErr := json.Marshal(params.Body)
		patchObject, decodeErr := jsonpatch.DecodePatch([]byte(jsonBody))

		if marshalErr != nil || decodeErr != nil {
			return things.NewWeaviateThingsPatchBadRequest()
		}

		// Convert ThingGetResponse object to JSON
		thingUpdateJSON, marshalErr := json.Marshal(thingGetResponse)
		if marshalErr != nil {
			return things.NewWeaviateThingsPatchBadRequest()
		}

		// Apply the patch
		updatedJSON, applyErr := patchObject.Apply(thingUpdateJSON)

		if applyErr != nil {
			return things.NewWeaviateThingsPatchUnprocessableEntity().WithPayload(createErrorResponseObject(applyErr.Error()))
		}

		// Turn it into a Thing object
		thing := &models.Thing{}
		json.Unmarshal([]byte(updatedJSON), &thing)

		// Validate schema made after patching with the weaviate schema
		validatedErr := validation.ValidateThingBody(&thing.ThingCreate, databaseSchema, dbConnector)
		if validatedErr != nil {
			return things.NewWeaviateThingsPatchUnprocessableEntity().WithPayload(createErrorResponseObject(validatedErr.Error()))
		}

		// Update the database
		insertErr := dbConnector.UpdateThing(thing, UUID)
		if insertErr != nil {
			messaging.ErrorMessage(insertErr)
		}

		// Create return Object
		thingGetResponse.Thing = *thing

		return things.NewWeaviateThingsPatchOK().WithPayload(&thingGetResponse)
	})
	api.ThingsWeaviateThingsUpdateHandler = things.WeaviateThingsUpdateHandlerFunc(func(params things.WeaviateThingsUpdateParams, principal interface{}) middleware.Responder {
		// Initialize response
		thingGetResponse := models.ThingGetResponse{}
		thingGetResponse.Schema = map[string]models.JSONObject{}

		// Get item from database
		UUID := params.ThingID
		errGet := dbConnector.GetThing(UUID, &thingGetResponse)

		// If there are no results, there is an error
		if errGet != nil {
			// Object not found response.
			return things.NewWeaviateThingsUpdateNotFound()
		}

		// This is a write function, validate if allowed to write?
		if allowed, _ := ActionsAllowed([]string{"write"}, principal, dbConnector, thingGetResponse.Key.NrDollarCref); !allowed {
			return things.NewWeaviateThingsUpdateForbidden()
		}

		// Validate schema given in body with the weaviate schema
		validatedErr := validation.ValidateThingBody(&params.Body.ThingCreate, databaseSchema, dbConnector)
		if validatedErr != nil {
			return things.NewWeaviateThingsUpdateUnprocessableEntity().WithPayload(createErrorResponseObject(validatedErr.Error()))
		}

		// Update the database
		params.Body.LastUpdateTimeUnix = connutils.NowUnix()
		params.Body.CreationTimeUnix = thingGetResponse.CreationTimeUnix
		insertErr := dbConnector.UpdateThing(&params.Body.Thing, UUID)
		if insertErr != nil {
			messaging.ErrorMessage(insertErr)
		}

		// Create object to return
		responseObject := &models.ThingGetResponse{}
		responseObject.Thing = params.Body.Thing
		responseObject.ThingID = UUID

		// Return SUCCESS (NOTE: this is ACCEPTED, so the dbConnector.Add should have a go routine)
		return things.NewWeaviateThingsUpdateOK().WithPayload(responseObject)
	})
	api.ThingsWeaviateThingsValidateHandler = things.WeaviateThingsValidateHandlerFunc(func(params things.WeaviateThingsValidateParams, principal interface{}) middleware.Responder {
		// Validate schema given in body with the weaviate schema
		validatedErr := validation.ValidateThingBody(params.Body, databaseSchema, dbConnector)
		if validatedErr != nil {
			return things.NewWeaviateThingsValidateUnprocessableEntity().WithPayload(createErrorResponseObject(validatedErr.Error()))
		}

		return things.NewWeaviateThingsValidateOK()
	})
	api.MetaWeaviateMetaGetHandler = meta.WeaviateMetaGetHandlerFunc(func(params meta.WeaviateMetaGetParams, principal interface{}) middleware.Responder {
		// Create response object
		metaResponse := &models.Meta{}

		// Set the response object's values
		metaResponse.Hostname = serverConfig.GetHostAddress()
		metaResponse.ActionsSchema = databaseSchema.ActionSchema.Schema
		metaResponse.ThingsSchema = databaseSchema.ThingSchema.Schema

		return meta.NewWeaviateMetaGetOK().WithPayload(metaResponse)
	})
	api.ThingsWeaviateThingsActionsListHandler = things.WeaviateThingsActionsListHandlerFunc(func(params things.WeaviateThingsActionsListParams, principal interface{}) middleware.Responder {
		// This is a read function, validate if allowed to read?
		if allowed, _ := ActionsAllowed([]string{"read"}, principal, dbConnector, nil); !allowed {
			return things.NewWeaviateThingsActionsListForbidden()
		}

		// Get limit and page
		limit := getLimit(params.MaxResults)
		page := getPage(params.Page)

		// Get key-object
		keyToken := principal.(models.KeyTokenGetResponse)

		// This is a read function, validate if allowed to read?
		if allowed, _ := ActionsAllowed([]string{"read"}, principal, dbConnector, keyToken.KeyID); !allowed {
			return things.NewWeaviateThingsActionsListForbidden()
		}

		// Initialize response
		thingGetResponse := models.ThingGetResponse{}
		thingGetResponse.Schema = map[string]models.JSONObject{}
		errGet := dbConnector.GetThing(params.ThingID, &thingGetResponse)

		// If there are no results, there is an error
		if errGet != nil {
			// Object not found response.
			return things.NewWeaviateThingsActionsListNotFound()
		}

		// Initialize response
		actionsResponse := models.ActionsListResponse{}
		actionsResponse.Actions = []*models.ActionGetResponse{}

		// List all results
		err := dbConnector.ListActions(params.ThingID, limit, (page-1)*limit, []*connutils.WhereQuery{}, &actionsResponse)

		if err != nil {
			messaging.ErrorMessage(err)
		}

		return things.NewWeaviateThingsActionsListOK().WithPayload(&actionsResponse)
	})
	api.GraphqlWeaviateGraphqlPostHandler = graphql.WeaviateGraphqlPostHandlerFunc(func(params graphql.WeaviateGraphqlPostParams, principal interface{}) middleware.Responder {
		messaging.DebugMessage("Starting GraphQL resolving")

		// Get all input from the body of the request, as it is a POST.
		query := params.Body.Query
		operationName := params.Body.OperationName

		// If query is empty, the request is unprocessable
		if query == "" {
			return graphql.NewWeaviateGraphqlPostUnprocessableEntity()
		}

		// Only set variables if exists in request
		var variables map[string]interface{}
		if params.Body.Variables != nil {
			variables = params.Body.Variables.(map[string]interface{})
		}

		// Get the results by doing a request with the given parameters and the initialized schema.
		graphQLSchema.SetKey(principal.(models.KeyTokenGetResponse))
		gqlSchema, _ := graphQLSchema.GetGraphQLSchema()

		// Do the request
		result := gographql.Do(gographql.Params{
			Schema:         gqlSchema,
			RequestString:  query,
			OperationName:  operationName,
			VariableValues: variables,
		})

		// Marshal the JSON
		resultJSON, jsonErr := json.Marshal(result)

		// If json gave error, return nothing.
		if jsonErr != nil {
			return graphql.NewWeaviateGraphqlPostUnprocessableEntity()
		}

		// Put the data in a response ready object
		graphQLResponse := &models.GraphQLResponse{}
		err := json.Unmarshal(resultJSON, graphQLResponse)

		// If json gave error, return nothing.
		if err != nil {
			return graphql.NewWeaviateGraphqlPostUnprocessableEntity()
		}

		messaging.DebugMessage("Ending GraphQL resolving")

		// Return the response
		return graphql.NewWeaviateGraphqlPostOK().WithPayload(graphQLResponse)
	})

	api.ServerShutdown = func() {}

	return setupGlobalMiddleware(api.Serve(setupMiddlewares))
}

// The TLS configuration before HTTPS server starts.
func configureTLS(tlsConfig *tls.Config) {
	// Make all necessary changes to the TLS configuration here.
}

// As soon as server is initialized but not run yet, this function will be called.
// If you need to modify a config, store server instance to stop it individually later, this is the place.
// This function can be called multiple times, depending on the number of serving schemes.
// scheme value will be set accordingly: "http", "https" or "unix"
func configureServer(s *graceful.Server, scheme, addr string) {
	// Create message service
	messaging = &messages.Messaging{}

	// Load the config using the flags
	serverConfig = config.WeaviateConfig{}
	err := serverConfig.LoadConfig(connectorOptionGroup, messaging)

	// Add properties to the config
	serverConfig.Hostname = addr
	serverConfig.Scheme = scheme

	// Fatal error loading config file
	if err != nil {
		messaging.ExitError(78, err.Error())
	}

	// Load the schema using the config
	databaseSchema = schema.WeaviateSchema{}
	err = databaseSchema.LoadSchema(&serverConfig.Environment, messaging)

	// Fatal error loading schema file
	if err != nil {
		messaging.ExitError(78, err.Error())
	}

	// Create the database connector usint the config
	dbConnector = CreateDatabaseConnector(&serverConfig.Environment)

	// Error the system when the database connector returns no connector
	if dbConnector == nil {
		messaging.ExitError(78, "database with the name '"+serverConfig.Environment.Database.Name+"' couldn't be loaded from the config")
	}

	// Set connector vars
	err = dbConnector.SetConfig(&serverConfig.Environment)
	// Fatal error loading config file
	if err != nil {
		messaging.ExitError(78, err.Error())
	}

	err = dbConnector.SetSchema(&databaseSchema)
	// Fatal error loading schema file
	if err != nil {
		messaging.ExitError(78, err.Error())
	}

	err = dbConnector.SetMessaging(messaging)
	// Fatal error setting messaging
	if err != nil {
		messaging.ExitError(78, err.Error())
	}

	dbConnector.SetServerAddress(serverConfig.GetHostAddress())

	// connect the database
	errConnect := dbConnector.Connect()
	if errConnect != nil {
		messaging.ExitError(1, "database with the name '"+serverConfig.Environment.Database.Name+"' gave an error when connecting: "+errConnect.Error())
	}

	// init the database
	errInit := dbConnector.Init()
	if errInit != nil {
		messaging.ExitError(1, "database with the name '"+serverConfig.Environment.Database.Name+"' gave an error when initializing: "+errInit.Error())
	}

	// Init the GraphQL schema
	graphQLSchema = graphqlapi.NewGraphQLSchema(dbConnector, &serverConfig, &databaseSchema, messaging)

	// Error init
	errInitGQL := graphQLSchema.InitSchema()
	if errInitGQL != nil {
		messaging.ExitError(1, "GraphQL schema initialization gave an error when initializing: "+errInitGQL.Error())
	}

}

// The middleware configuration is for the handler executors. These do not apply to the swagger.json document.
// The middleware executes after routing but before authentication, binding and validation
func setupMiddlewares(handler http.Handler) http.Handler {
	return handler
}

// The middleware configuration happens before anything, this middleware also applies to serving the swagger.json document.
// So this is a good place to plug in a panic handling middleware, logging and metrics
func setupGlobalMiddleware(handler http.Handler) http.Handler {
	return handler
}
