/*                          _       _
 *__      _____  __ ___   ___  __ _| |_ ___
 *\ \ /\ / / _ \/ _` \ \ / / |/ _` | __/ _ \
 * \ V  V /  __/ (_| |\ V /| | (_| | ||  __/
 *  \_/\_/ \___|\__,_| \_/ |_|\__,_|\__\___|
 *
 * Copyright © 2016 Weaviate. All rights reserved.
 * LICENSE: https://github.com/weaviate/weaviate/blob/master/LICENSE
 * AUTHOR: Bob van Luijt (bob@weaviate.com)
 * See www.weaviate.com for details
 * Contact: @weaviate_iot / yourfriends@weaviate.com
 */

package memory

import (
	"errors"
	"github.com/go-openapi/strfmt"
	"github.com/weaviate/weaviate/models"
	"log"

	"math"
	"sort"

	"github.com/hashicorp/go-memdb"
	"github.com/weaviate/weaviate/connectors/config"
	"github.com/weaviate/weaviate/connectors/utils"
)

// Datastore has some basic variables.
type Memory struct {
	client *memdb.MemDB
	kind   string
}

// GetName returns a unique connector name
func (f *Memory) GetName() string {
	return "memory"
}

// SetConfig is used to fill in a struct with config variables
func (f *Memory) SetConfig(configInput connectorConfig.Environment) {
	// NOTHING
}

// Creates connection and tables if not already available (which is never because it is in memory)
func (f *Memory) Connect() error {

	// Create the weaviate DB schema
	schema := &memdb.DBSchema{
		Tables: map[string]*memdb.TableSchema{
			// create `weaviate` DB
			"weaviate": &memdb.TableSchema{
				Name: "weaviate",
				Indexes: map[string]*memdb.IndexSchema{
					"id": &memdb.IndexSchema{
						Name:    "id",
						Unique:  true,
						Indexer: &memdb.StringFieldIndex{Field: "Uuid"},
					},
					"Deleted": &memdb.IndexSchema{
						Name:    "Deleted",
						Unique:  true,
						Indexer: &memdb.StringFieldIndex{Field: "Deleted"},
					},
					"CreateTimeMs": &memdb.IndexSchema{
						Name:    "CreateTimeMs",
						Unique:  true,
						Indexer: &memdb.StringFieldIndex{Field: "CreateTimeMs"},
					},
					"Object": &memdb.IndexSchema{
						Name:    "Object",
						Unique:  true,
						Indexer: &memdb.StringFieldIndex{Field: "Object"},
					},
					"Owner": &memdb.IndexSchema{
						Name:    "Owner",
						Unique:  true,
						Indexer: &memdb.StringFieldIndex{Field: "Owner"},
					},
					"RefType": &memdb.IndexSchema{
						Name:    "RefType",
						Unique:  true,
						Indexer: &memdb.StringFieldIndex{Field: "RefType"},
					},
					"Uuid": &memdb.IndexSchema{
						Name:    "Uuid",
						Unique:  true,
						Indexer: &memdb.StringFieldIndex{Field: "Uuid"},
					},
				},
			},
			// create `weaviate` DB
			"weaviate_history": &memdb.TableSchema{
				Name: "weaviate_history",
				Indexes: map[string]*memdb.IndexSchema{
					"id": &memdb.IndexSchema{
						Name:    "id",
						Unique:  true,
						Indexer: &memdb.StringFieldIndex{Field: "Uuid"},
					},

					"Deleted": &memdb.IndexSchema{
						Name:    "Deleted",
						Unique:  true,
						Indexer: &memdb.StringFieldIndex{Field: "Deleted"},
					},
					"CreateTimeMs": &memdb.IndexSchema{
						Name:    "CreateTimeMs",
						Unique:  true,
						Indexer: &memdb.StringFieldIndex{Field: "CreateTimeMs"},
					},
					"Object": &memdb.IndexSchema{
						Name:    "Object",
						Unique:  true,
						Indexer: &memdb.StringFieldIndex{Field: "Object"},
					},
					"Owner": &memdb.IndexSchema{
						Name:    "Owner",
						Unique:  false,
						Indexer: &memdb.StringFieldIndex{Field: "Owner"},
					},
					"RefType": &memdb.IndexSchema{
						Name:    "RefType",
						Unique:  false,
						Indexer: &memdb.StringFieldIndex{Field: "RefType"},
					},
					"Uuid": &memdb.IndexSchema{
						Name:    "Uuid",
						Unique:  false,
						Indexer: &memdb.StringFieldIndex{Field: "Uuid"},
					},
				},
			},
			// create `weaviate_users` DB
			"weaviate_users": &memdb.TableSchema{
				Name: "weaviate_users",
				Indexes: map[string]*memdb.IndexSchema{
					"id": &memdb.IndexSchema{
						Name:    "id",
						Unique:  true,
						Indexer: &memdb.StringFieldIndex{Field: "Uuid"},
					},
					"KeyExpiresUnix": &memdb.IndexSchema{
						Name:    "KeyExpiresUnix",
						Unique:  true,
						Indexer: &memdb.StringFieldIndex{Field: "KeyExpiresUnix"},
					},
					"KeyToken": &memdb.IndexSchema{
						Name:    "KeyToken",
						Unique:  true,
						Indexer: &memdb.StringFieldIndex{Field: "KeyToken"},
					},
					"Object": &memdb.IndexSchema{
						Name:    "Object",
						Unique:  true,
						Indexer: &memdb.StringFieldIndex{Field: "Object"},
					},
					"Parent": &memdb.IndexSchema{
						Name:    "Parent",
						Unique:  false,
						Indexer: &memdb.StringFieldIndex{Field: "Parent"},
					},
					"Uuid": &memdb.IndexSchema{
						Name:    "Uuid",
						Unique:  true,
						Indexer: &memdb.StringFieldIndex{Field: "Uuid"},
					},
				},
			},
		},
	}

	// Create a new data base
	client, err := memdb.NewMemDB(schema)

	// If error, return it. Otherwise set client.
	if err != nil {
		return err
	}

	f.client = client

	log.Println("INFO: In memory database is used for testing / development purposes only")

	return nil

}

// Creates a root key, normally this should be validaded, but because it is an inmemory DB it is created always
func (f *Memory) Init() error {
	// Generate a basic DB object and print it's key.
	dbObject := connector_utils.CreateFirstUserObject()

	// Create a write transaction
	txn := f.client.Txn(true)

	// Saves the new entity.
	if err := txn.Insert("weaviate_users", dbObject); err != nil {
		return err
	}

	// commit transaction
	txn.Commit()

	return nil
}

func (f *Memory) Add(dbObject connector_utils.DatabaseObject) (string, error) {

	// Create a write transaction
	txn := f.client.Txn(true)

	// Saves the new entity.
	if err := txn.Insert("weaviate", dbObject); err != nil {
		return "Error", err
	}

	// commit transaction
	txn.Commit()

	// Return the ID that is used to create.
	return dbObject.Uuid, nil

}

func (f *Memory) Get(Uuid string) (connector_utils.DatabaseObject, error) {

	// Create read-only transaction
	txn := f.client.Txn(false)
	defer txn.Abort()

	// Lookup by Uuid
	result, err := txn.First("weaviate", "Uuid", Uuid)
	if err != nil {
		return connector_utils.DatabaseObject{}, err
	}

	// Return 'not found'
	if result == nil {
		notFoundErr := errors.New("no object with such UUID found")
		return connector_utils.DatabaseObject{}, notFoundErr
	}

	// Return found object
	return result.(connector_utils.DatabaseObject), nil

}

// return a list
func (f *Memory) List(refType string, ownerUUID string, limit int, page int, referenceFilter *connector_utils.ObjectReferences) (connector_utils.DatabaseObjects, int64, error) {
	dataObjs := connector_utils.DatabaseObjects{}

	// Create read-only transaction
	txn := f.client.Txn(false)
	defer txn.Abort()

	// Lookup by Uuid
	result, err := txn.Get("weaviate", "id")

	// return the error
	if err != nil {
		return dataObjs, 0, err
	}

	if result != nil {

		// loop through the results
		singleResult := result.Next()
		for singleResult != nil {
			// only store if refType and owner is correct and object is not deleted
			if singleResult.(connector_utils.DatabaseObject).RefType == refType &&
				singleResult.(connector_utils.DatabaseObject).Owner == ownerUUID &&
				!singleResult.(connector_utils.DatabaseObject).Deleted {

				if referenceFilter != nil {
					// check for extra filters
					if referenceFilter.ThingID != "" &&
						singleResult.(connector_utils.DatabaseObject).RelatedObjects.ThingID == referenceFilter.ThingID {
						dataObjs = append(dataObjs, singleResult.(connector_utils.DatabaseObject))
					}
				} else {
					dataObjs = append(dataObjs, singleResult.(connector_utils.DatabaseObject))
				}
			}
			singleResult = result.Next()
		}

		// Sorting on CreateTimeMs
		sort.Sort(dataObjs)

		// count total
		totalResults := len(dataObjs)

		// calculate the amount to chop off totalResults-limit
		offset := (limit * (page - 1))
		end := int(math.Min(float64(limit*(page)), float64(totalResults)))
		dataObjs := dataObjs[offset:end]

		// return found set
		return dataObjs, int64(totalResults), err
	}

	// nothing found
	return dataObjs, 0, nil
}

// Validate if a user has access, returns permissions object
func (f *Memory) ValidateKey(token string) ([]connector_utils.DatabaseUsersObject, error) {

	dbUsersObjects := []connector_utils.DatabaseUsersObject{}

	// Create read-only transaction
	txn := f.client.Txn(false)
	defer txn.Abort()

	// Filter on timestamp, deleted and token itself
	result, err := txn.First("weaviate_users", "KeyToken", token)
	if err != nil || result == nil {
		return []connector_utils.DatabaseUsersObject{}, err
	}

	// Add to results
	userObject := result.(connector_utils.DatabaseUsersObject)
	dbUsersObjects = append(dbUsersObjects, userObject)

	// keys are found, return true
	return dbUsersObjects, nil
}

// GetKey returns user object by ID
func (f *Memory) GetKey(Uuid string) (connector_utils.DatabaseUsersObject, error) {
	// Create read-only transaction
	txn := f.client.Txn(false)
	defer txn.Abort()

	// Lookup by Uuid
	result, err := txn.First("weaviate_users", "Uuid", Uuid)
	if err != nil {
		return connector_utils.DatabaseUsersObject{}, err
	}

	// Return 'not found'
	if result == nil {
		notFoundErr := errors.New("No object with such UUID found")
		return connector_utils.DatabaseUsersObject{}, notFoundErr
	}

	// Return found object
	return result.(connector_utils.DatabaseUsersObject), nil

}

// AddUser to DB
func (f *Memory) AddKey(parentUuid string, dbObject connector_utils.DatabaseUsersObject) (connector_utils.DatabaseUsersObject, error) {

	// Create a write transaction
	txn := f.client.Txn(true)

	// Auto set the parent ID
	dbObject.Parent = parentUuid

	// Saves the new entity.
	if err := txn.Insert("weaviate_users", dbObject); err != nil {
		return dbObject, err
	}

	// commit transaction
	txn.Commit()

	// Return the ID that is used to create.
	return dbObject, nil

}

// DeleteKey removes a key from the database
func (f *Memory) DeleteKey(UUID string) error {
	// Create a read transaction
	txn := f.client.Txn(false)
	defer txn.Abort()

	// Lookup all Children
	result, err := txn.First("weaviate_users", "Uuid", UUID)

	// Return the error
	if err != nil {
		return err
	}

	childUserObject := result.(connector_utils.DatabaseUsersObject)
	childUserObject.Deleted = true

	txn2 := f.client.Txn(true)
	// Delete item(s) with given Uuid
	_, errDel := txn2.DeleteAll("weaviate_users", "Uuid", childUserObject.Uuid)
	txn2.Insert("weaviate_users", childUserObject)

	// Commit transaction
	txn2.Commit()

	return errDel
}

// GetChildKeys returns all the child keys
func (f *Memory) GetChildObjects(UUID string, filterOutDeleted bool) ([]connector_utils.DatabaseUsersObject, error) {
	// Create a read transaction
	txn := f.client.Txn(false)
	defer txn.Abort()

	// // Fill children array
	childUserObjects := []connector_utils.DatabaseUsersObject{}

	// Lookup by Uuid
	result, err := txn.Get("weaviate_users", "Parent", UUID)

	// return the error
	if err != nil {
		return childUserObjects, err
	}

	if result != nil {
		// loop through the results
		singleResult := result.Next()
		for singleResult != nil {
			// only store if refType is correct
			if filterOutDeleted {
				if !singleResult.(connector_utils.DatabaseUsersObject).Deleted {
					childUserObjects = append(childUserObjects, singleResult.(connector_utils.DatabaseUsersObject))
				}
			} else {
				childUserObjects = append(childUserObjects, singleResult.(connector_utils.DatabaseUsersObject))
			}

			singleResult = result.Next()
		}
	}

	return childUserObjects, nil
}

func (f *Memory) AddThing(thing *models.ThingCreate, UUID strfmt.UUID) error {
	return nil
}

func (f *Memory) GetThing(UUID strfmt.UUID) (models.ThingGetResponse, error) {
	thingResponse := models.ThingGetResponse{}

	return thingResponse, nil
}
