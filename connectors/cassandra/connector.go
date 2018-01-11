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
 * Contact: @CreativeSofwFdn / yourfriends@weaviate.com
 */

/*
 * THIS IS A DEMO CONNECTOR!
 * USE IT TO LEARN HOW TO CREATE YOUR OWN CONNECTOR.
 */

/*
When starting Weaviate, functions are called in the following order;
(find the function in this document to understand what it is that they do)
 - GetName
 - SetConfig
 - SetSchema
 - SetMessaging
 - SetServerAddress
 - Connect
 - Init

All other function are called on the API request

After creating the connector, make sure to add the name of the connector to: func GetAllConnectors() in configure_weaviate.go

*/

package cassandra

import (
	errors_ "errors"
	"fmt"
	"strings"
	"time"

	"github.com/go-openapi/strfmt"
	"github.com/gocql/gocql"
	"github.com/mitchellh/mapstructure"

	"github.com/creativesoftwarefdn/weaviate/config"
	"github.com/creativesoftwarefdn/weaviate/connectors/utils"
	"github.com/creativesoftwarefdn/weaviate/messages"
	"github.com/creativesoftwarefdn/weaviate/models"
	"github.com/creativesoftwarefdn/weaviate/schema"
)

const objectTableName = "object_data"

// IDColumn constant column name
const IDColumn string = "id"

// UUIDColumn constant column name
const UUIDColumn string = "uuid"

// TypeColumn constant column name
const TypeColumn string = "type"

// ClassColumn constant column name
const ClassColumn string = "class"

// PropertyKeyColumn constant column name
const PropertyKeyColumn string = "property_key"

// PropertyValueStringColumn constant column name
const PropertyValueStringColumn string = "property_val_string"

// PropertyValueBoolColumn constant column name
const PropertyValueBoolColumn string = "property_val_bool"

// PropertyValueTimeStampColumn constant column name
const PropertyValueTimeStampColumn string = "property_val_timestamp"

// PropertyValueIntColumn constant column name
const PropertyValueIntColumn string = "property_val_int"

// PropertyValueFloatColumn constant column name
const PropertyValueFloatColumn string = "property_val_float"

// PropertyRefColumn constant column name
const PropertyRefColumn string = "property_ref"

// TimeStampColumn constant column name
const TimeStampColumn string = "timestamp"

// DeletedColumn constant column name
const DeletedColumn string = "deleted"

// Global insert statement
const insertStatement = `
	INSERT INTO %v (` + IDColumn + `, ` + UUIDColumn + `, ` + TypeColumn + `, ` + ClassColumn + `, ` + PropertyKeyColumn + `, %s, ` + PropertyRefColumn + `, ` + TimeStampColumn + `, ` + DeletedColumn + `) 
	VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)
`

const selectStatement = `
	SELECT id, uuid, type, class, property_key, property_val_string, property_val_bool, property_val_timestamp, property_val_int, property_val_float, property_ref, timestamp, deleted 
	FROM ` + objectTableName + ` WHERE uuid = ?
`

// Cassandra has some basic variables.
// This is mandatory, only change it if you need aditional, global variables
type Cassandra struct {
	client *gocql.Session
	kind   string

	config        Config
	serverAddress string
	schema        *schema.WeaviateSchema
	messaging     *messages.Messaging
}

// Config represents the config outline for Cassandra. The Database config shoud be of the following form:
// "database_config" : {
//     "host": "127.0.0.1",
//     "port": 9080
// }
// Notice that the port is the GRPC-port.
type Config struct {
	Host string
	Port int
}

// GetName returns a unique connector name, this name is used to define the connector in the weaviate config
func (f *Cassandra) GetName() string {
	return "cassandra"
}

// SetConfig sets variables, which can be placed in the config file section "database_config: {}"
// can be custom for any connector, in the example below there is only host and port available.
//
// Important to bear in mind;
// 1. You need to add these to the struct Config in this document.
// 2. They will become available via f.config.[variable-name]
//
// 	"database": {
// 		"name": "cassandra",
// 		"database_config" : {
// 			"host": "127.0.0.1",
// 			"port": 9080
// 		}
// 	},
func (f *Cassandra) SetConfig(configInput *config.Environment) error {

	// Mandatory: needed to add the JSON config represented as a map in f.config
	err := mapstructure.Decode(configInput.Database.DatabaseConfig, &f.config)

	// Example to: Validate if the essential  config is available, like host and port.
	if err != nil || len(f.config.Host) == 0 || f.config.Port == 0 {
		return errors_.New("could not get Cassandra host/port from config")
	}

	// If success return nil, otherwise return the error (see above)
	return nil
}

// SetSchema takes actionSchema and thingsSchema as an input and makes them available globally at f.schema
// In case you want to modify the schema, this is the place to do so.
// Note: When this function is called, the schemas (action + things) are already validated, so you don't have to build the validation.
func (f *Cassandra) SetSchema(schemaInput *schema.WeaviateSchema) error {
	f.schema = schemaInput

	// If success return nil, otherwise return the error
	return nil
}

// SetMessaging is used to send messages to the service.
// Available message types are: f.messaging.Infomessage ...DebugMessage ...ErrorMessage ...ExitError (also exits the service) ...InfoMessage
func (f *Cassandra) SetMessaging(m *messages.Messaging) error {

	// mandatory, adds the message functions to f.messaging to make them globally accessible.
	f.messaging = m

	// If success return nil, otherwise return the error
	return nil
}

// SetServerAddress is used to fill in a global variable with the server address, but can also be used
// to do some custom actions.
// Does not return anything
func (f *Cassandra) SetServerAddress(addr string) {
	f.serverAddress = addr
}

// Connect creates a connection to the database and tables if not already available.
// The connections could not be closed because it is used more often.
func (f *Cassandra) Connect() error {
	// Create a Cassandra cluster
	cluster := gocql.NewCluster("127.0.0.1") // TODO variable

	// Create a session on the cluster for just creating/checking the Keyspace
	session, err := cluster.CreateSession()

	if err != nil {
		return err
	}

	if err := session.Query(`CREATE KEYSPACE IF NOT EXISTS weaviate 
		WITH REPLICATION = { 'class' : 'SimpleStrategy', 'replication_factor' : 1 }`).Exec(); err != nil {
		return err
	} // TODO variable

	// Close session for checking Keyspace
	session.Close()

	// Settings for createing the new Session
	cluster.Keyspace = "weaviate" // TODO variable
	cluster.ConnectTimeout = time.Minute
	cluster.Timeout = time.Hour
	session, err = cluster.CreateSession()

	if err != nil {
		return err
	}

	// Put the session into the client-variable to make is usable everywhere else
	f.client = session

	// If success return nil, otherwise return the error (also see above)
	return nil
}

// Init 1st initializes the schema in the database and 2nd creates a root key.
func (f *Cassandra) Init() error {
	// Add table 'object_data'

	// TODO make const
	// TODO property_ref = UUID type?

	err := f.client.Query(fmt.Sprintf(`
		CREATE TABLE IF NOT EXISTS object_data (
			%s UUID PRIMARY KEY, %s UUID, %s text, %s text,
			%s text, %s text, %s boolean, %s timestamp,
			%s int, %s double, %s text, %s timestamp, %s boolean
		);`, IDColumn, UUIDColumn, TypeColumn, ClassColumn, PropertyKeyColumn, PropertyValueStringColumn, PropertyValueBoolColumn, PropertyValueTimeStampColumn, PropertyValueIntColumn, PropertyValueFloatColumn, PropertyRefColumn, TimeStampColumn, DeletedColumn)).Exec()

	if err != nil {
		return err
	}

	// Create all indexes
	indexes := []string{UUIDColumn, TypeColumn, ClassColumn, PropertyKeyColumn, PropertyValueStringColumn, PropertyValueBoolColumn, PropertyValueTimeStampColumn, PropertyValueIntColumn, PropertyValueFloatColumn, PropertyRefColumn, TimeStampColumn, DeletedColumn}
	for _, prop := range indexes {
		if err := f.client.Query(fmt.Sprintf(`CREATE INDEX IF NOT EXISTS object_%s ON object_data (%s);`, prop, prop)).Exec(); err != nil {
			return err
		}
	}

	// Add ROOT-key if not exists
	// Search for Root key
	var rootCount int

	if err := f.client.Query(fmt.Sprintf(`
		SELECT COUNT(id) AS rootCount FROM object_data WHERE %s = ? AND %s = ? ALLOW FILTERING
	`, PropertyKeyColumn, PropertyValueBoolColumn), "root", true).Scan(&rootCount); err != nil {
		return err
	}

	// If root-key is not found
	if rootCount == 0 {
		f.messaging.InfoMessage("No root-key found.")

		// Create new object and fill it
		keyObject := models.Key{}
		token := connutils.CreateRootKeyObject(&keyObject)

		err = f.AddKey(&keyObject, connutils.GenerateUUID(), token)

		if err != nil {
			return err
		}
	}

	// If success return nil, otherwise return the error
	return nil
}

// AddThing adds a thing to the Cassandra database with the given UUID.
// Takes the thing and a UUID as input.
// Thing is already validated against the ontology
func (f *Cassandra) AddThing(thing *models.Thing, UUID strfmt.UUID) error {
	// Create new batch
	batch := f.client.NewBatch(gocql.LoggedBatch)

	// Parse UUID in Cassandra type
	cqlUUID := f.convUUIDtoCQLUUID(UUID)

	// Add first level properties
	f.addFirstLevelQueriesToBatch(connutils.RefTypeThing, thing, cqlUUID, batch)

	// Add Object properties using a callback
	callback := f.createPropertyCallback(batch, cqlUUID, thing.AtClass)
	schema.UpdateObjectSchemaProperties(connutils.RefTypeThing, thing, thing.Schema, f.schema, callback)

	if err := f.client.ExecuteBatch(batch); err != nil {
		return err
	}

	// If success return nil, otherwise return the error
	return nil
}

// GetThing fills the given ThingGetResponse with the values from the database, based on the given UUID.
func (f *Cassandra) GetThing(UUID strfmt.UUID, thingResponse *models.ThingGetResponse) error {
	// Get the iterator
	iter := f.getSelectIteratorByUUID(UUID)

	f.fillResponseWithIter(iter, thingResponse, connutils.RefTypeThing)
	if err := iter.Close(); err != nil {
		return err
	}

	// If success return nil, otherwise return the error
	return nil
}

// ListThings fills the given ThingsListResponse with the values from the database, based on the given parameters.
func (f *Cassandra) ListThings(first int, offset int, keyID strfmt.UUID, wheres []*connutils.WhereQuery, thingsResponse *models.ThingsListResponse) error {

	// thingsResponse should be populated with the response that comes from the DB.
	// thingsResponse = based on the ontology

	// If success return nil, otherwise return the error
	return nil
}

// UpdateThing updates the Thing in the DB at the given UUID.
func (f *Cassandra) UpdateThing(thing *models.Thing, UUID strfmt.UUID) error {

	// Run the query to update the thing based on its UUID.

	// If success return nil, otherwise return the error
	return nil
}

// DeleteThing deletes the Thing in the DB at the given UUID.
func (f *Cassandra) DeleteThing(UUID strfmt.UUID) error {

	// Run the query to delete the thing based on its UUID.

	// If success return nil, otherwise return the error
	return nil
}

// AddAction adds an action to the Cassandra database with the given UUID.
// Takes the action and a UUID as input.
// Action is already validated against the ontology
func (f *Cassandra) AddAction(action *models.Action, UUID strfmt.UUID) error {

	// If success return nil, otherwise return the error
	return nil
}

// GetAction fills the given ActionGetResponse with the values from the database, based on the given UUID.
func (f *Cassandra) GetAction(UUID strfmt.UUID, actionResponse *models.ActionGetResponse) error {
	// actionResponse should be populated with the response that comes from the DB.
	// actionResponse = based on the ontology

	// If success return nil, otherwise return the error
	return nil
}

// ListActions fills the given ActionListResponse with the values from the database, based on the given parameters.
func (f *Cassandra) ListActions(UUID strfmt.UUID, first int, offset int, wheres []*connutils.WhereQuery, actionsResponse *models.ActionsListResponse) error {
	// actionsResponse should be populated with the response that comes from the DB.
	// actionsResponse = based on the ontology

	// If success return nil, otherwise return the error
	return nil
}

// UpdateAction updates the Thing in the DB at the given UUID.
func (f *Cassandra) UpdateAction(action *models.Action, UUID strfmt.UUID) error {

	// If success return nil, otherwise return the error
	return nil
}

// DeleteAction deletes the Action in the DB at the given UUID.
func (f *Cassandra) DeleteAction(UUID strfmt.UUID) error {

	// Run the query to delete the action based on its UUID.

	// If success return nil, otherwise return the error
	return nil
}

// AddKey adds a key to the Cassandra database with the given UUID and token.
// UUID  = reference to the key
// token = is the actual access token used in the API's header
func (f *Cassandra) AddKey(key *models.Key, UUID strfmt.UUID, token strfmt.UUID) error {
	batch := f.client.NewBatch(gocql.LoggedBatch)

	keyUUID, _ := gocql.ParseUUID(string(UUID))

	batch.Query(fmt.Sprintf(insertStatement, objectTableName, PropertyValueBoolColumn), gocql.TimeUUID(), keyUUID, connutils.RefTypeKey, "", "delete", key.Delete, "", connutils.NowUnix(), false)
	batch.Query(fmt.Sprintf(insertStatement, objectTableName, PropertyValueStringColumn), gocql.TimeUUID(), keyUUID, connutils.RefTypeKey, "", "email", key.Email, "", connutils.NowUnix(), false)
	batch.Query(fmt.Sprintf(insertStatement, objectTableName, PropertyValueBoolColumn), gocql.TimeUUID(), keyUUID, connutils.RefTypeKey, "", "execute", key.Execute, "", connutils.NowUnix(), false)
	batch.Query(fmt.Sprintf(insertStatement, objectTableName, PropertyValueStringColumn), gocql.TimeUUID(), keyUUID, connutils.RefTypeKey, "", "ipOrigin", strings.Join(key.IPOrigin, "|"), "", connutils.NowUnix(), false)
	batch.Query(fmt.Sprintf(insertStatement, objectTableName, PropertyValueTimeStampColumn), gocql.TimeUUID(), keyUUID, connutils.RefTypeKey, "", "keyExpiresUnix", key.KeyExpiresUnix, "", connutils.NowUnix(), false)
	batch.Query(fmt.Sprintf(insertStatement, objectTableName, PropertyValueBoolColumn), gocql.TimeUUID(), keyUUID, connutils.RefTypeKey, "", "read", key.Read, "", connutils.NowUnix(), false)
	batch.Query(fmt.Sprintf(insertStatement, objectTableName, PropertyValueBoolColumn), gocql.TimeUUID(), keyUUID, connutils.RefTypeKey, "", "write", key.Write, "", connutils.NowUnix(), false)
	batch.Query(fmt.Sprintf(insertStatement, objectTableName, PropertyValueStringColumn), gocql.TimeUUID(), keyUUID, connutils.RefTypeKey, "", "token", token, "", connutils.NowUnix(), false)

	isRoot := key.Parent != nil
	if !isRoot {
		batch.Query(fmt.Sprintf(insertStatement, objectTableName, PropertyValueStringColumn), gocql.TimeUUID(), keyUUID, connutils.RefTypeKey, "", "parent", key.Parent.LocationURL, key.Parent.NrDollarCref, connutils.NowUnix(), false)
	}
	batch.Query(fmt.Sprintf(insertStatement, objectTableName, PropertyValueBoolColumn), gocql.TimeUUID(), keyUUID, connutils.RefTypeKey, "", "root", isRoot, "", connutils.NowUnix(), false)

	if err := f.client.ExecuteBatch(batch); err != nil {
		return err
	}

	// If success return nil, otherwise return the error
	return nil
}

// ValidateToken validates/gets a key to the Cassandra database with the given token (=UUID)
func (f *Cassandra) ValidateToken(token strfmt.UUID, keyResponse *models.KeyTokenGetResponse) error {

	// key (= models.KeyTokenGetResponse) should be populated with the response that comes from the DB.
	// key = based on the ontology

	// in case the key is not found, return an error like:
	// return errors_.New("Key not found in database.")

	var UUID gocql.UUID

	if err := f.client.Query(fmt.Sprintf(`SELECT uuid 
			FROM %s 
			WHERE property_key = 'token' AND property_val_string = ?
			LIMIT 1
			ALLOW FILTERING
		`, objectTableName), token).Consistency(gocql.One).Scan(&UUID); err != nil {
		return err
	}

	// Get all rows for the key
	iter := f.getSelectIteratorByUUID(f.convCQLUUIDtoUUID(UUID))

	f.fillResponseWithIter(iter, keyResponse, connutils.RefTypeKey)
	if err := iter.Close(); err != nil {
		return err
	}

	// If success return nil, otherwise return the error
	return nil
}

// GetKey fills the given KeyTokenGetResponse with the values from the database, based on the given UUID.
func (f *Cassandra) GetKey(UUID strfmt.UUID, keyResponse *models.KeyTokenGetResponse) error {
	// Get all rows for the key
	// Get all rows for the key
	iter := f.getSelectIteratorByUUID(UUID)

	f.fillResponseWithIter(iter, keyResponse, connutils.RefTypeKey)

	if err := iter.Close(); err != nil {
		return err
	}

	return nil
}

// DeleteKey deletes the Key in the DB at the given UUID.
func (f *Cassandra) DeleteKey(UUID strfmt.UUID) error {

	return nil
}

// GetKeyChildren fills the given KeyTokenGetResponse array with the values from the database, based on the given UUID.
func (f *Cassandra) GetKeyChildren(UUID strfmt.UUID, children *[]*models.KeyTokenGetResponse) error {

	// for examle: `children = [OBJECT-A, OBJECT-B, OBJECT-C]`
	// Where an OBJECT = models.KeyTokenGetResponse

	return nil
}

func (f *Cassandra) convUUIDtoCQLUUID(UUID strfmt.UUID) gocql.UUID {
	cqlUUID, _ := gocql.ParseUUID(string(UUID))
	return cqlUUID
}

func (f *Cassandra) convCQLUUIDtoUUID(cqlUUID gocql.UUID) strfmt.UUID {
	UUID := strfmt.UUID(cqlUUID.String())
	return UUID
}

func (f *Cassandra) addFirstLevelQueriesToBatch(refType string, object interface{}, UUID gocql.UUID, batch *gocql.Batch) {
	if connutils.RefTypeThing == refType {
		thing := object.(*models.Thing)
		batch.Query(fmt.Sprintf(insertStatement, objectTableName, PropertyValueStringColumn), gocql.TimeUUID(), UUID, refType, thing.AtClass, "@context", thing.AtContext, "", connutils.NowUnix(), false)
		batch.Query(fmt.Sprintf(insertStatement, objectTableName, PropertyValueTimeStampColumn), gocql.TimeUUID(), UUID, refType, thing.AtClass, "creationTimeUnix", thing.CreationTimeUnix, "", connutils.NowUnix(), false)
		batch.Query(fmt.Sprintf(insertStatement, objectTableName, PropertyValueTimeStampColumn), gocql.TimeUUID(), UUID, refType, thing.AtClass, "lastUpdateTimeUnix", thing.LastUpdateTimeUnix, "", connutils.NowUnix(), false)
		batch.Query(fmt.Sprintf(insertStatement, objectTableName, PropertyValueStringColumn), gocql.TimeUUID(), UUID, refType, thing.AtClass, "key", thing.Key.LocationURL, thing.Key.NrDollarCref, connutils.NowUnix(), false)
	}
}

func (f *Cassandra) createPropertyCallback(batch *gocql.Batch, cqlUUID gocql.UUID, className string) func(string, interface{}, *schema.DataType, string) error {
	return func(propKey string, propValue interface{}, dataType *schema.DataType, edgeName string) error {
		dataTypeColumn := ""

		if *dataType == schema.DataTypeBoolean {
			dataTypeColumn = PropertyValueBoolColumn
		} else if *dataType == schema.DataTypeDate {
			dataTypeColumn = PropertyValueTimeStampColumn

			var err error
			propValue, err = time.Parse(time.RFC3339, propValue.(string))

			// Return if there is an error while parsing
			if err != nil {
				return err
			}
		} else if *dataType == schema.DataTypeInt {
			dataTypeColumn = PropertyValueIntColumn
			propValue = int64(propValue.(float64))
		} else if *dataType == schema.DataTypeNumber {
			dataTypeColumn = PropertyValueFloatColumn
		} else if *dataType == schema.DataTypeString {
			dataTypeColumn = PropertyValueStringColumn
		} else if *dataType == schema.DataTypeCRef {
			// TODO
		}

		if dataTypeColumn != "" {
			batch.Query(fmt.Sprintf(insertStatement, objectTableName, dataTypeColumn), gocql.TimeUUID(), cqlUUID, connutils.RefTypeThing, className, edgeName, propValue, "", connutils.NowUnix(), false)
		}

		return nil
	}
}

func (f *Cassandra) getSelectIteratorByUUID(UUID strfmt.UUID) *gocql.Iter {
	return f.client.Query(selectStatement, string(UUID)).Iter()
}

func (f *Cassandra) fillResponseWithIter(iter *gocql.Iter, response interface{}, refType string) error {
	var sID gocql.UUID
	var sUUID gocql.UUID
	var sType string
	var sClass string
	var sPropKey string
	var sPropValString string
	var sPropValBool bool
	var sPropValTime time.Time
	var sPropValInt int
	var sPropValFloat float64
	var sPropRef string
	var sTimeStamp time.Time
	var sDeleted bool

	responseSchema := make(map[string]interface{})
	var atContext string
	var creationTimeUnix int64
	var lastUpdateTimeUnix int64
	crefObj := models.SingleRef{}
	var UUID strfmt.UUID
	var atClass string

	var delete bool
	var email string
	var execute bool
	var ipOrigin []string
	var keyExpiresUnix int64
	var read bool
	var write bool
	var token strfmt.UUID

	for iter.Scan(&sID, &sUUID, &sType, &sClass, &sPropKey, &sPropValString, &sPropValBool, &sPropValTime, &sPropValInt, &sPropValFloat, &sPropRef, &sTimeStamp, &sDeleted) {
		if isSchema, propKey, dataType, err := schema.TranslateSchemaPropertiesFromDataBase(sPropKey, sClass, f.schema.ThingSchema.Schema); isSchema {
			if err != nil {
				return err
			}

			var propValue interface{}
			if *dataType == schema.DataTypeBoolean {
				propValue = sPropValBool
			} else if *dataType == schema.DataTypeDate {
				propValue = sPropValTime
			} else if *dataType == schema.DataTypeInt {
				propValue = sPropValInt
			} else if *dataType == schema.DataTypeNumber {
				propValue = sPropValFloat
			} else if *dataType == schema.DataTypeString {
				propValue = sPropValString
			}

			responseSchema[propKey] = propValue
		} else {
			switch sPropKey {
			case "@context":
				atContext = sPropValString
			case "creationTimeUnix":
				creationTimeUnix = sPropValTime.Unix()
			case "lastUpdateTimeUnix":
				lastUpdateTimeUnix = sPropValTime.Unix()
			case "key":
				url := sPropValString

				crefObj.NrDollarCref = strfmt.UUID(sPropRef)
				crefObj.LocationURL = &url
				crefObj.Type = connutils.RefTypeKey
			case "delete":
				delete = sPropValBool
			case "email":
				email = sPropValString
			case "execute":
				execute = sPropValBool
			case "ipOrigin":
				ipOrigin = strings.Split(sPropValString, "|") // TODO const seperator
			case "keyExpiresUnix":
				keyExpiresUnix = sPropValTime.Unix()
			case "read":
				read = sPropValBool
			case "write":
				write = sPropValBool
			case "token":
				token = strfmt.UUID(sPropValString)
			}
		}
	}

	UUID = strfmt.UUID(sUUID.String())
	atClass = sClass

	if connutils.RefTypeThing == refType {
		thingResponse := response.(*models.ThingGetResponse)
		thingResponse.Schema = responseSchema
		thingResponse.AtContext = atContext
		thingResponse.CreationTimeUnix = creationTimeUnix
		thingResponse.LastUpdateTimeUnix = lastUpdateTimeUnix
		thingResponse.Key = &crefObj
		thingResponse.ThingID = UUID
		thingResponse.AtClass = atClass
	} else if connutils.RefTypeKey == refType {
		keyResponse := response.(*models.KeyTokenGetResponse)
		keyResponse.KeyID = UUID
		keyResponse.Delete = delete
		keyResponse.Email = email
		keyResponse.Execute = execute
		keyResponse.IPOrigin = ipOrigin
		keyResponse.KeyExpiresUnix = keyExpiresUnix
		keyResponse.Read = read
		keyResponse.Write = write
		keyResponse.Token = token
	}

	return nil
}
