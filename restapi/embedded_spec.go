/*                          _       _
 *__      _____  __ ___   ___  __ _| |_ ___
 *\ \ /\ / / _ \/ _` \ \ / / |/ _` | __/ _ \
 * \ V  V /  __/ (_| |\ V /| | (_| | ||  __/
 *  \_/\_/ \___|\__,_| \_/ |_|\__,_|\__\___|
 *
 * Copyright © 2016 - 2018 Weaviate. All rights reserved.
 * LICENSE: https://github.com/creativesoftwarefdn/weaviate/blob/develop/LICENSE.md
 * AUTHOR: Bob van Luijt (bob@kub.design)
 * See www.creativesoftwarefdn.org for details
 * Contact: @CreativeSofwFdn / bob@kub.design
 */

package restapi

import (
	"encoding/json"
)

// SwaggerJSON embedded version of the swagger document used at generation time
var SwaggerJSON json.RawMessage

func init() {
	SwaggerJSON = json.RawMessage([]byte(`{
  "consumes": [
    "application/json"
  ],
  "produces": [
    "application/json"
  ],
  "schemes": [
    "https"
  ],
  "swagger": "2.0",
  "info": {
    "description": "Weaviate - Semantic Graphql, RESTful Web of Things platform.",
    "title": "Weaviate - Semantic Graphql, RESTful Web of Things platform.",
    "contact": {
      "name": "Weaviate",
      "url": "https://github.com/creativesoftwarefdn",
      "email": "hello@creativesoftwarefdn.org"
    },
    "version": "0.9.2"
  },
  "basePath": "/weaviate/v1",
  "paths": {
    "/actions": {
      "post": {
        "description": "Registers a new action. Given meta-data and schema values are validated.",
        "tags": [
          "actions"
        ],
        "summary": "Create actions between two things (object and subject).",
        "operationId": "weaviate.actions.create",
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/ActionCreate"
            }
          }
        ],
        "responses": {
          "202": {
            "description": "Successfully received.",
            "schema": {
              "$ref": "#/definitions/ActionGetResponse"
            }
          },
          "401": {
            "description": "Unauthorized or invalid credentials."
          },
          "403": {
            "description": "The used API-key has insufficient permissions."
          },
          "422": {
            "description": "Request body contains well-formed (i.e., syntactically correct), but semantically erroneous. Are you sure the class is defined in the configuration file?",
            "schema": {
              "$ref": "#/definitions/ErrorResponse"
            }
          },
          "501": {
            "description": "Not (yet) implemented."
          }
        },
        "x-available-in-mqtt": false,
        "x-available-in-websocket": false
      }
    },
    "/actions/validate": {
      "post": {
        "description": "Validate an action's schema and meta-data. It has to be based on a schema, which is related to the given action to be accepted by this validation.",
        "tags": [
          "actions"
        ],
        "summary": "Validate an action based on a schema.",
        "operationId": "weaviate.actions.validate",
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/ActionValidate"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "Successful validated."
          },
          "401": {
            "description": "Unauthorized or invalid credentials."
          },
          "403": {
            "description": "The used API-key has insufficient permissions."
          },
          "422": {
            "description": "Request body contains well-formed (i.e., syntactically correct), but semantically erroneous. Are you sure the class is defined in the configuration file?",
            "schema": {
              "$ref": "#/definitions/ErrorResponse"
            }
          },
          "501": {
            "description": "Not (yet) implemented."
          }
        },
        "x-available-in-mqtt": false,
        "x-available-in-websocket": false
      }
    },
    "/actions/{actionId}": {
      "get": {
        "description": "Lists actions.",
        "tags": [
          "actions"
        ],
        "summary": "Get a specific action based on its uuid and a thing uuid related to this key. Also available as Websocket bus.",
        "operationId": "weaviate.actions.get",
        "parameters": [
          {
            "type": "string",
            "format": "uuid",
            "description": "Unique ID of the action.",
            "name": "actionId",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "Successful response.",
            "schema": {
              "$ref": "#/definitions/ActionGetResponse"
            }
          },
          "401": {
            "description": "Unauthorized or invalid credentials."
          },
          "403": {
            "description": "The used API-key has insufficient permissions."
          },
          "404": {
            "description": "Successful query result but no resource was found."
          },
          "501": {
            "description": "Not (yet) implemented."
          }
        },
        "x-available-in-mqtt": false,
        "x-available-in-websocket": false
      },
      "put": {
        "description": "Updates an action's data. Given meta-data and schema values are validated. LastUpdateTime is set to the time this function is called.",
        "tags": [
          "actions"
        ],
        "summary": "Update an action based on its uuid related to this key.",
        "operationId": "weaviate.action.update",
        "parameters": [
          {
            "type": "string",
            "format": "uuid",
            "description": "Unique ID of the action.",
            "name": "actionId",
            "in": "path",
            "required": true
          },
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/ActionUpdate"
            }
          }
        ],
        "responses": {
          "202": {
            "description": "Successfully received.",
            "schema": {
              "$ref": "#/definitions/ActionGetResponse"
            }
          },
          "401": {
            "description": "Unauthorized or invalid credentials."
          },
          "403": {
            "description": "The used API-key has insufficient permissions."
          },
          "404": {
            "description": "Successful query result but no resource was found."
          },
          "422": {
            "description": "Request body contains well-formed (i.e., syntactically correct), but semantically erroneous. Are you sure the class is defined in the configuration file?",
            "schema": {
              "$ref": "#/definitions/ErrorResponse"
            }
          },
          "501": {
            "description": "Not (yet) implemented."
          }
        },
        "x-available-in-mqtt": false,
        "x-available-in-websocket": false
      },
      "delete": {
        "description": "Deletes an action from the system.",
        "tags": [
          "actions"
        ],
        "summary": "Delete an action based on its uuid related to this key.",
        "operationId": "weaviate.actions.delete",
        "parameters": [
          {
            "type": "string",
            "format": "uuid",
            "description": "Unique ID of the thing.",
            "name": "actionId",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "204": {
            "description": "Successful deleted."
          },
          "401": {
            "description": "Unauthorized or invalid credentials."
          },
          "403": {
            "description": "The used API-key has insufficient permissions."
          },
          "404": {
            "description": "Successful query result but no resource was found."
          },
          "501": {
            "description": "Not (yet) implemented."
          }
        },
        "x-available-in-mqtt": true,
        "x-available-in-websocket": true
      },
      "patch": {
        "description": "Updates an action. This method supports patch semantics. Given meta-data and schema values are validated. LastUpdateTime is set to the time this function is called.",
        "tags": [
          "actions"
        ],
        "summary": "Update an action based on its uuid (using patch semantics) related to this key.",
        "operationId": "weaviate.actions.patch",
        "parameters": [
          {
            "type": "string",
            "format": "uuid",
            "description": "Unique ID of the action.",
            "name": "actionId",
            "in": "path",
            "required": true
          },
          {
            "description": "JSONPatch document as defined by RFC 6902.",
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "type": "array",
              "items": {
                "$ref": "#/definitions/PatchDocument"
              }
            }
          }
        ],
        "responses": {
          "202": {
            "description": "Successfully received.",
            "schema": {
              "$ref": "#/definitions/ActionGetResponse"
            }
          },
          "400": {
            "description": "The patch-JSON is malformed."
          },
          "401": {
            "description": "Unauthorized or invalid credentials."
          },
          "403": {
            "description": "The used API-key has insufficient permissions."
          },
          "404": {
            "description": "Successful query result but no resource was found."
          },
          "422": {
            "description": "The patch-JSON is valid but unprocessable.",
            "schema": {
              "$ref": "#/definitions/ErrorResponse"
            }
          },
          "501": {
            "description": "Not (yet) implemented."
          }
        },
        "x-available-in-mqtt": false,
        "x-available-in-websocket": false
      }
    },
    "/actions/{actionId}/history": {
      "get": {
        "description": "Returns a particular action history.",
        "tags": [
          "actions"
        ],
        "summary": "Get a action's history based on its uuid related to this key.",
        "operationId": "weaviate.action.history.get",
        "parameters": [
          {
            "type": "string",
            "format": "uuid",
            "description": "Unique ID of the action.",
            "name": "actionId",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "Successful response.",
            "schema": {
              "$ref": "#/definitions/ActionGetHistoryResponse"
            }
          },
          "401": {
            "description": "Unauthorized or invalid credentials."
          },
          "403": {
            "description": "The used API-key has insufficient permissions."
          },
          "404": {
            "description": "Successful query result but no resource was found."
          },
          "501": {
            "description": "Not (yet) implemented."
          }
        },
        "x-available-in-mqtt": false,
        "x-available-in-websocket": false
      }
    },
    "/graphql": {
      "post": {
        "description": "Get an object based on GraphQL",
        "tags": [
          "graphql"
        ],
        "summary": "Get a response based on GraphQL",
        "operationId": "weaviate.graphql.post",
        "parameters": [
          {
            "description": "The GraphQL query request parameters.",
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/GraphQLQuery"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "Succesful query (with select).",
            "schema": {
              "$ref": "#/definitions/GraphQLResponse"
            }
          },
          "401": {
            "description": "Unauthorized or invalid credentials."
          },
          "403": {
            "description": "The used API-key has insufficient permissions."
          },
          "422": {
            "description": "Request body contains well-formed (i.e., syntactically correct), but semantically erroneous. Are you sure the class is defined in the configuration file?",
            "schema": {
              "$ref": "#/definitions/ErrorResponse"
            }
          },
          "501": {
            "description": "Not (yet) implemented."
          }
        },
        "x-available-in-mqtt": false,
        "x-available-in-websocket": false
      }
    },
    "/keys": {
      "post": {
        "description": "Creates a new key. Input expiration date is validated on being in the future and not longer than parent expiration date.",
        "tags": [
          "keys"
        ],
        "summary": "Create a new key related to this key.",
        "operationId": "weaviate.key.create",
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/KeyCreate"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "Successfully created.",
            "schema": {
              "$ref": "#/definitions/KeyTokenGetResponse"
            }
          },
          "401": {
            "description": "Unauthorized or invalid credentials."
          },
          "422": {
            "description": "Request body contains well-formed (i.e., syntactically correct), but semantically erroneous. Are you sure the class is defined in the configuration file?",
            "schema": {
              "$ref": "#/definitions/ErrorResponse"
            }
          },
          "501": {
            "description": "Not (yet) implemented."
          }
        },
        "x-available-in-mqtt": false,
        "x-available-in-websocket": false
      }
    },
    "/keys/me": {
      "get": {
        "description": "Get the key-information of the key used.",
        "tags": [
          "keys"
        ],
        "summary": "Get a key based on the key used to do the request.",
        "operationId": "weaviate.keys.me.get",
        "responses": {
          "200": {
            "description": "Successful response.",
            "schema": {
              "$ref": "#/definitions/KeyGetResponse"
            }
          },
          "401": {
            "description": "Unauthorized or invalid credentials."
          },
          "404": {
            "description": "Successful query result but no resource was found."
          },
          "501": {
            "description": "Not (yet) implemented."
          }
        },
        "x-available-in-mqtt": false,
        "x-available-in-websocket": false
      }
    },
    "/keys/me/children": {
      "get": {
        "description": "Get children of used key, only one step deep. A child can have children of its own.",
        "tags": [
          "keys"
        ],
        "summary": "Get an object of this keys' children related to the key used for request.",
        "operationId": "weaviate.keys.me.children.get",
        "responses": {
          "200": {
            "description": "Successful response.",
            "schema": {
              "$ref": "#/definitions/KeyChildrenGetResponse"
            }
          },
          "401": {
            "description": "Unauthorized or invalid credentials."
          },
          "404": {
            "description": "Successful query result but no resource was found."
          },
          "501": {
            "description": "Not (yet) implemented"
          }
        },
        "x-available-in-mqtt": false,
        "x-available-in-websocket": false
      }
    },
    "/keys/{keyId}": {
      "get": {
        "description": "Get a key.",
        "tags": [
          "keys"
        ],
        "summary": "Get a key based on its uuid related to this key.",
        "operationId": "weaviate.keys.get",
        "parameters": [
          {
            "type": "string",
            "format": "uuid",
            "description": "Unique ID of the key.",
            "name": "keyId",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "Successful response.",
            "schema": {
              "$ref": "#/definitions/KeyGetResponse"
            }
          },
          "401": {
            "description": "Unauthorized or invalid credentials."
          },
          "403": {
            "description": "The used API-key has insufficient permissions."
          },
          "404": {
            "description": "Successful query result but no resource was found."
          },
          "501": {
            "description": "Not (yet) implemented."
          }
        },
        "x-available-in-mqtt": false,
        "x-available-in-websocket": false
      },
      "delete": {
        "description": "Deletes a key. Only parent or self is allowed to delete key. When you delete a key, all its children will be deleted as well.",
        "tags": [
          "keys"
        ],
        "summary": "Delete a key based on its uuid related to this key.",
        "operationId": "weaviate.keys.delete",
        "parameters": [
          {
            "type": "string",
            "format": "uuid",
            "description": "Unique ID of the key.",
            "name": "keyId",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "204": {
            "description": "Successful deleted."
          },
          "401": {
            "description": "Unauthorized or invalid credentials."
          },
          "403": {
            "description": "The used API-key has insufficient permissions."
          },
          "404": {
            "description": "Successful query result but no resource was found."
          },
          "501": {
            "description": "Not (yet) implemented."
          }
        },
        "x-available-in-mqtt": false,
        "x-available-in-websocket": false
      }
    },
    "/keys/{keyId}/children": {
      "get": {
        "description": "Get children of a key, only one step deep. A child can have children of its own.",
        "tags": [
          "keys"
        ],
        "summary": "Get an object of this keys' children related to this key.",
        "operationId": "weaviate.keys.children.get",
        "parameters": [
          {
            "type": "string",
            "format": "uuid",
            "description": "Unique ID of the key.",
            "name": "keyId",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "Successful response.",
            "schema": {
              "$ref": "#/definitions/KeyChildrenGetResponse"
            }
          },
          "401": {
            "description": "Unauthorized or invalid credentials."
          },
          "403": {
            "description": "The used API-key has insufficient permissions."
          },
          "404": {
            "description": "Successful query result but no resource was found."
          },
          "501": {
            "description": "Not (yet) implemented"
          }
        },
        "x-available-in-mqtt": false,
        "x-available-in-websocket": false
      }
    },
    "/keys/{keyId}/renew-token": {
      "put": {
        "description": "Renews the related key. Validates being lower in tree than given key. Can not renew itself, unless being parent.",
        "tags": [
          "keys"
        ],
        "summary": "Renews a key based on the key given in the query string.",
        "operationId": "weaviate.keys.renew.token",
        "parameters": [
          {
            "type": "string",
            "format": "uuid",
            "description": "Unique ID of the key.",
            "name": "keyId",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "Successful response.",
            "schema": {
              "$ref": "#/definitions/KeyTokenGetResponse"
            }
          },
          "401": {
            "description": "Unauthorized or invalid credentials."
          },
          "403": {
            "description": "The used API-key has insufficient permissions."
          },
          "404": {
            "description": "Successful query result but no resource was found."
          },
          "422": {
            "description": "Request body contains well-formed (i.e., syntactically correct), but semantically erroneous. Are you sure the class is defined in the configuration file?",
            "schema": {
              "$ref": "#/definitions/ErrorResponse"
            }
          },
          "501": {
            "description": "Not (yet) implemented."
          }
        },
        "x-available-in-mqtt": false,
        "x-available-in-websocket": false
      }
    },
    "/meta": {
      "get": {
        "description": "Gives meta information about the server and can be used to provide information to another Weaviate instance that wants to interact with the current instance.",
        "produces": [
          "application/json"
        ],
        "tags": [
          "meta"
        ],
        "summary": "Returns meta information of the current Weaviate instance.",
        "operationId": "weaviate.meta.get",
        "responses": {
          "200": {
            "description": "Successful response.",
            "schema": {
              "$ref": "#/definitions/Meta"
            }
          },
          "401": {
            "description": "Unauthorized or invalid credentials."
          },
          "501": {
            "description": "Not (yet) implemented"
          }
        },
        "x-available-in-mqtt": false,
        "x-available-in-websocket": false
      }
    },
    "/peers": {
      "post": {
        "description": "Announce a new peer, authentication not needed (all peers are allowed to try and connect). This endpoint will only be used in M2M communications.",
        "tags": [
          "P2P"
        ],
        "summary": "Announce a new peer.",
        "operationId": "weaviate.peers.announce",
        "parameters": [
          {
            "description": "The announcement message",
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/PeerAnnouncement"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "Successfully registred the peer to the network."
          },
          "403": {
            "description": "You are not allowed on the network."
          },
          "501": {
            "description": "Not (yet) implemented."
          }
        },
        "x-available-in-mqtt": false,
        "x-available-in-websocket": false
      }
    },
    "/peers/answers/{answerId}": {
      "post": {
        "description": "Receive an answer based on a question from a peer in the network.",
        "tags": [
          "P2P"
        ],
        "summary": "Receiving a new answer from a peer.",
        "operationId": "weaviate.peers.answers.create",
        "parameters": [
          {
            "type": "string",
            "format": "uuid",
            "description": "The Uuid of the answer.",
            "name": "answerId",
            "in": "path",
            "required": true
          },
          {
            "description": "The answer.",
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/Schema"
            }
          }
        ],
        "responses": {
          "202": {
            "description": "Successfully received."
          },
          "401": {
            "description": "Unauthorized or invalid credentials."
          },
          "403": {
            "description": "The used API-key has insufficient permissions.",
            "schema": {
              "$ref": "#/definitions/ErrorResponse"
            }
          },
          "501": {
            "description": "Not (yet) implemented."
          }
        },
        "x-available-in-mqtt": false,
        "x-available-in-websocket": false
      }
    },
    "/peers/echo": {
      "get": {
        "description": "Check if a peer is alive.",
        "tags": [
          "P2P"
        ],
        "summary": "Check if a peer is alive.",
        "operationId": "weaviate.peers.echo",
        "responses": {
          "200": {
            "description": "Alive and kicking!"
          },
          "501": {
            "description": "Not (yet) implemented."
          }
        },
        "x-available-in-mqtt": false,
        "x-available-in-websocket": false
      }
    },
    "/peers/questions": {
      "post": {
        "description": "Receive a question from a peer in the network.",
        "tags": [
          "P2P"
        ],
        "summary": "Receive a question from a peer in the network.",
        "operationId": "weaviate.peers.questions.create",
        "parameters": [
          {
            "description": "The question.",
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/QuestionCreate"
            }
          }
        ],
        "responses": {
          "202": {
            "description": "Successfully received the question and answer might be send back.",
            "schema": {
              "$ref": "#/definitions/QuestionResponse"
            }
          },
          "403": {
            "description": "You are not allowed on the network."
          },
          "501": {
            "description": "Not (yet) implemented."
          }
        },
        "x-available-in-mqtt": false,
        "x-available-in-websocket": false
      }
    },
    "/things": {
      "get": {
        "description": "Lists all things in reverse order of creation, owned by the user that belongs to the used token.",
        "tags": [
          "things"
        ],
        "summary": "Get a list of things related to this key.",
        "operationId": "weaviate.things.list",
        "parameters": [
          {
            "$ref": "#/parameters/CommonMaxResultsParameterQuery"
          },
          {
            "$ref": "#/parameters/CommonPageParameterQuery"
          }
        ],
        "responses": {
          "200": {
            "description": "Successful response.",
            "schema": {
              "$ref": "#/definitions/ThingsListResponse"
            }
          },
          "401": {
            "description": "Unauthorized or invalid credentials."
          },
          "403": {
            "description": "The used API-key has insufficient permissions."
          },
          "404": {
            "description": "Successful query result but no resource was found."
          },
          "501": {
            "description": "Not (yet) implemented."
          }
        },
        "x-available-in-mqtt": false,
        "x-available-in-websocket": false
      },
      "post": {
        "description": "Registers a new thing. Given meta-data and schema values are validated.",
        "tags": [
          "things"
        ],
        "summary": "Create a new thing based on a thing template related to this key.",
        "operationId": "weaviate.things.create",
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/ThingCreate"
            }
          }
        ],
        "responses": {
          "202": {
            "description": "Successfully received.",
            "schema": {
              "$ref": "#/definitions/ThingGetResponse"
            }
          },
          "401": {
            "description": "Unauthorized or invalid credentials."
          },
          "403": {
            "description": "The used API-key has insufficient permissions."
          },
          "422": {
            "description": "Request body contains well-formed (i.e., syntactically correct), but semantically erroneous. Are you sure the class is defined in the configuration file?",
            "schema": {
              "$ref": "#/definitions/ErrorResponse"
            }
          },
          "501": {
            "description": "Not (yet) implemented."
          }
        },
        "x-available-in-mqtt": false,
        "x-available-in-websocket": false
      }
    },
    "/things/validate": {
      "post": {
        "description": "Validate a thing's schema and meta-data. It has to be based on a schema, which is related to the given Thing to be accepted by this validation.",
        "tags": [
          "things"
        ],
        "summary": "Validate Things schema.",
        "operationId": "weaviate.things.validate",
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/ThingCreate"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "Successful validated."
          },
          "401": {
            "description": "Unauthorized or invalid credentials."
          },
          "403": {
            "description": "The used API-key has insufficient permissions."
          },
          "422": {
            "description": "Request body contains well-formed (i.e., syntactically correct), but semantically erroneous. Are you sure the class is defined in the configuration file?",
            "schema": {
              "$ref": "#/definitions/ErrorResponse"
            }
          },
          "501": {
            "description": "Not (yet) implemented."
          }
        },
        "x-available-in-mqtt": false,
        "x-available-in-websocket": false
      }
    },
    "/things/{thingId}": {
      "get": {
        "description": "Returns a particular thing data.",
        "tags": [
          "things"
        ],
        "summary": "Get a thing based on its uuid related to this key.",
        "operationId": "weaviate.things.get",
        "parameters": [
          {
            "type": "string",
            "format": "uuid",
            "description": "Unique ID of the thing.",
            "name": "thingId",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "Successful response.",
            "schema": {
              "$ref": "#/definitions/ThingGetResponse"
            }
          },
          "401": {
            "description": "Unauthorized or invalid credentials."
          },
          "403": {
            "description": "The used API-key has insufficient permissions."
          },
          "404": {
            "description": "Successful query result but no resource was found."
          },
          "501": {
            "description": "Not (yet) implemented."
          }
        },
        "x-available-in-mqtt": false,
        "x-available-in-websocket": false
      },
      "put": {
        "description": "Updates a thing data. Given meta-data and schema values are validated. LastUpdateTime is set to the time this function is called.",
        "tags": [
          "things"
        ],
        "summary": "Update a thing based on its uuid related to this key.",
        "operationId": "weaviate.things.update",
        "parameters": [
          {
            "type": "string",
            "format": "uuid",
            "description": "Unique ID of the thing.",
            "name": "thingId",
            "in": "path",
            "required": true
          },
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/ThingUpdate"
            }
          }
        ],
        "responses": {
          "202": {
            "description": "Successfully received.",
            "schema": {
              "$ref": "#/definitions/ThingGetResponse"
            }
          },
          "401": {
            "description": "Unauthorized or invalid credentials."
          },
          "403": {
            "description": "The used API-key has insufficient permissions."
          },
          "404": {
            "description": "Successful query result but no resource was found."
          },
          "422": {
            "description": "Request body contains well-formed (i.e., syntactically correct), but semantically erroneous. Are you sure the class is defined in the configuration file?",
            "schema": {
              "$ref": "#/definitions/ErrorResponse"
            }
          },
          "501": {
            "description": "Not (yet) implemented."
          }
        },
        "x-available-in-mqtt": false,
        "x-available-in-websocket": false
      },
      "delete": {
        "description": "Deletes a thing from the system. All actions pointing to this thing, where the thing is the object of the action, are also being deleted.",
        "tags": [
          "things"
        ],
        "summary": "Delete a thing based on its uuid related to this key.",
        "operationId": "weaviate.things.delete",
        "parameters": [
          {
            "type": "string",
            "format": "uuid",
            "description": "Unique ID of the thing.",
            "name": "thingId",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "204": {
            "description": "Successful deleted."
          },
          "401": {
            "description": "Unauthorized or invalid credentials."
          },
          "403": {
            "description": "The used API-key has insufficient permissions."
          },
          "404": {
            "description": "Successful query result but no resource was found."
          },
          "501": {
            "description": "Not (yet) implemented."
          }
        },
        "x-available-in-mqtt": true,
        "x-available-in-websocket": true
      },
      "patch": {
        "description": "Updates a thing data. This method supports patch semantics. Given meta-data and schema values are validated. LastUpdateTime is set to the time this function is called.",
        "tags": [
          "things"
        ],
        "summary": "Update a thing based on its uuid (using patch semantics) related to this key.",
        "operationId": "weaviate.things.patch",
        "parameters": [
          {
            "type": "string",
            "format": "uuid",
            "description": "Unique ID of the thing.",
            "name": "thingId",
            "in": "path",
            "required": true
          },
          {
            "description": "JSONPatch document as defined by RFC 6902.",
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "type": "array",
              "items": {
                "$ref": "#/definitions/PatchDocument"
              }
            }
          }
        ],
        "responses": {
          "202": {
            "description": "Successfully received.",
            "schema": {
              "$ref": "#/definitions/ThingGetResponse"
            }
          },
          "400": {
            "description": "The patch-JSON is malformed."
          },
          "401": {
            "description": "Unauthorized or invalid credentials."
          },
          "403": {
            "description": "The used API-key has insufficient permissions."
          },
          "404": {
            "description": "Successful query result but no resource was found."
          },
          "422": {
            "description": "The patch-JSON is valid but unprocessable.",
            "schema": {
              "$ref": "#/definitions/ErrorResponse"
            }
          },
          "501": {
            "description": "Not (yet) implemented."
          }
        },
        "x-available-in-mqtt": false,
        "x-available-in-websocket": false
      }
    },
    "/things/{thingId}/actions": {
      "get": {
        "description": "Lists all actions in reverse order of creation, related to the thing that belongs to the used thingId.",
        "tags": [
          "things"
        ],
        "summary": "Get a thing based on its uuid related to this thing. Also available as Websocket.",
        "operationId": "weaviate.things.actions.list",
        "parameters": [
          {
            "type": "string",
            "format": "uuid",
            "description": "Unique ID of the thing.",
            "name": "thingId",
            "in": "path",
            "required": true
          },
          {
            "$ref": "#/parameters/CommonMaxResultsParameterQuery"
          },
          {
            "$ref": "#/parameters/CommonPageParameterQuery"
          }
        ],
        "responses": {
          "200": {
            "description": "Successful response.",
            "schema": {
              "$ref": "#/definitions/ActionsListResponse"
            }
          },
          "401": {
            "description": "Unauthorized or invalid credentials."
          },
          "403": {
            "description": "The used API-key has insufficient permissions."
          },
          "404": {
            "description": "Successful query result but no resource was found."
          },
          "501": {
            "description": "Not (yet) implemented."
          }
        },
        "x-available-in-mqtt": false,
        "x-available-in-websocket": false
      }
    },
    "/things/{thingId}/history": {
      "get": {
        "description": "Returns a particular thing history.",
        "tags": [
          "things"
        ],
        "summary": "Get a thing's history based on its uuid related to this key.",
        "operationId": "weaviate.thing.history.get",
        "parameters": [
          {
            "type": "string",
            "format": "uuid",
            "description": "Unique ID of the thing.",
            "name": "thingId",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "Successful response.",
            "schema": {
              "$ref": "#/definitions/ThingGetHistoryResponse"
            }
          },
          "401": {
            "description": "Unauthorized or invalid credentials."
          },
          "403": {
            "description": "The used API-key has insufficient permissions."
          },
          "404": {
            "description": "Successful query result but no resource was found."
          },
          "501": {
            "description": "Not (yet) implemented."
          }
        },
        "x-available-in-mqtt": false,
        "x-available-in-websocket": false
      }
    }
  },
  "definitions": {
    "Action": {
      "type": "object",
      "allOf": [
        {
          "$ref": "#/definitions/ActionCreate"
        },
        {
          "type": "object",
          "properties": {
            "creationTimeUnix": {
              "description": "Timestamp of creation of this action in milliseconds since epoch UTC.",
              "type": "integer",
              "format": "int64"
            },
            "key": {
              "$ref": "#/definitions/SingleRef"
            },
            "lastUpdateTimeUnix": {
              "description": "Timestamp since epoch of last update made to the action.",
              "type": "integer",
              "format": "int64"
            }
          }
        }
      ]
    },
    "ActionCreate": {
      "type": "object",
      "properties": {
        "@class": {
          "description": "Type of the Action, defined in the schema.",
          "type": "string"
        },
        "@context": {
          "description": "Available context schema.",
          "type": "string"
        },
        "schema": {
          "$ref": "#/definitions/Schema"
        },
        "things": {
          "$ref": "#/definitions/ObjectSubject"
        }
      }
    },
    "ActionGetHistoryResponse": {
      "allOf": [
        {
          "$ref": "#/definitions/ActionHistory"
        },
        {
          "type": "object",
          "properties": {
            "actionId": {
              "type": "string",
              "format": "uuid"
            }
          }
        }
      ]
    },
    "ActionGetResponse": {
      "type": "object",
      "allOf": [
        {
          "$ref": "#/definitions/Action"
        },
        {
          "properties": {
            "actionId": {
              "description": "ID of the action.",
              "type": "string",
              "format": "uuid"
            }
          }
        }
      ]
    },
    "ActionHistory": {
      "type": "object",
      "properties": {
        "deleted": {
          "description": "Indication whether the action is deleted",
          "type": "boolean"
        },
        "key": {
          "$ref": "#/definitions/SingleRef"
        },
        "propertyHistory": {
          "description": "An array with the history of the action.",
          "type": "array",
          "items": {
            "$ref": "#/definitions/ActionHistoryObject"
          }
        }
      }
    },
    "ActionHistoryObject": {
      "allOf": [
        {
          "$ref": "#/definitions/ActionCreate"
        },
        {
          "type": "object",
          "properties": {
            "creationTimeUnix": {
              "description": "Timestamp of creation of this action history in milliseconds since epoch UTC.",
              "type": "integer",
              "format": "int64"
            }
          }
        }
      ]
    },
    "ActionUpdate": {
      "allOf": [
        {
          "$ref": "#/definitions/Action"
        },
        {
          "type": "object"
        }
      ]
    },
    "ActionValidate": {
      "type": "object",
      "allOf": [
        {
          "$ref": "#/definitions/ActionCreate"
        }
      ]
    },
    "ActionsListResponse": {
      "description": "List of actions for specific Thing.",
      "type": "object",
      "properties": {
        "actions": {
          "description": "The actual list of actions.",
          "type": "array",
          "items": {
            "$ref": "#/definitions/ActionGetResponse"
          }
        },
        "totalResults": {
          "description": "The total number of actions for the query. The number of items in a response may be smaller due to paging.",
          "type": "integer",
          "format": "int64"
        }
      }
    },
    "ErrorResponse": {
      "description": "An error response given by Weaviate end-points.",
      "type": "object",
      "properties": {
        "error": {
          "type": "object",
          "properties": {
            "message": {
              "type": "string"
            }
          }
        }
      }
    },
    "GraphQLError": {
      "description": "Error messages responded only if error exists.",
      "properties": {
        "locations": {
          "type": "array",
          "items": {
            "type": "object",
            "properties": {
              "column": {
                "type": "integer",
                "format": "int64"
              },
              "line": {
                "type": "integer",
                "format": "int64"
              }
            }
          }
        },
        "message": {
          "type": "string"
        },
        "path": {
          "type": "array",
          "items": {
            "type": "string"
          }
        }
      }
    },
    "GraphQLQuery": {
      "description": "GraphQL query based on: http://facebook.github.io/graphql/",
      "type": "object",
      "properties": {
        "operationName": {
          "description": "Name of the operation if multiple exist in query.",
          "type": "string"
        },
        "query": {
          "description": "Query based on GraphQL syntax",
          "type": "string"
        },
        "variables": {
          "description": "Additional variables for the query.",
          "type": "object"
        }
      }
    },
    "GraphQLResponse": {
      "description": "GraphQL based repsonse: http://facebook.github.io/graphql/",
      "properties": {
        "data": {
          "description": "GraphQL data object",
          "type": "object",
          "additionalProperties": {
            "$ref": "#/definitions/JsonObject"
          }
        },
        "errors": {
          "description": "Array with errors",
          "type": "array",
          "items": {
            "$ref": "#/definitions/GraphQLError"
          }
        }
      }
    },
    "JsonObject": {
      "description": "JSON object value.",
      "type": "object"
    },
    "Key": {
      "allOf": [
        {
          "$ref": "#/definitions/KeyCreate"
        },
        {
          "properties": {
            "parent": {
              "$ref": "#/definitions/SingleRef"
            }
          }
        }
      ]
    },
    "KeyChildrenGetResponse": {
      "properties": {
        "children": {
          "$ref": "#/definitions/MultipleRef"
        }
      }
    },
    "KeyCreate": {
      "properties": {
        "delete": {
          "description": "Is user allowed to delete.",
          "type": "boolean"
        },
        "email": {
          "description": "Email associated with this account.",
          "type": "string"
        },
        "execute": {
          "description": "Is user allowed to execute.",
          "type": "boolean"
        },
        "ipOrigin": {
          "description": "Origin of the IP using CIDR notation.",
          "type": "array",
          "items": {
            "type": "string"
          }
        },
        "isRoot": {
          "description": "Shows if key is root key",
          "type": "boolean",
          "default": false
        },
        "keyExpiresUnix": {
          "description": "Time as Unix timestamp that the key expires. Set to 0 for never.",
          "type": "integer",
          "format": "int64"
        },
        "read": {
          "description": "Is user allowed to read.",
          "type": "boolean"
        },
        "write": {
          "description": "Is user allowed to write.",
          "type": "boolean"
        }
      }
    },
    "KeyGetResponse": {
      "allOf": [
        {
          "$ref": "#/definitions/Key"
        },
        {
          "properties": {
            "keyId": {
              "description": "Id of the key.",
              "type": "string",
              "format": "uuid"
            }
          }
        }
      ]
    },
    "KeyTokenGetResponse": {
      "allOf": [
        {
          "$ref": "#/definitions/KeyGetResponse"
        },
        {
          "properties": {
            "token": {
              "description": "Key for user to use.",
              "type": "string",
              "format": "uuid"
            }
          }
        }
      ]
    },
    "Meta": {
      "description": "Contains meta information of the current Weaviate instance.",
      "type": "object",
      "properties": {
        "actionsSchema": {
          "$ref": "#/definitions/SemanticSchema"
        },
        "hostname": {
          "description": "The url of the host",
          "type": "string",
          "format": "url"
        },
        "thingsSchema": {
          "$ref": "#/definitions/SemanticSchema"
        }
      }
    },
    "MultipleRef": {
      "description": "Multiple instances of references to other objects.",
      "type": "array",
      "items": {
        "$ref": "#/definitions/SingleRef"
      }
    },
    "ObjectSubject": {
      "description": "returns a ref to the object and the subject",
      "type": "object",
      "properties": {
        "object": {
          "$ref": "#/definitions/SingleRef"
        },
        "subject": {
          "$ref": "#/definitions/SingleRef"
        }
      }
    },
    "PatchDocument": {
      "description": "A JSONPatch document as defined by RFC 6902.",
      "required": [
        "op",
        "path"
      ],
      "properties": {
        "from": {
          "description": "A string containing a JSON Pointer value.",
          "type": "string"
        },
        "op": {
          "description": "The operation to be performed.",
          "type": "string",
          "enum": [
            "add",
            "remove",
            "replace",
            "move",
            "copy",
            "test"
          ]
        },
        "path": {
          "description": "A JSON-Pointer.",
          "type": "string"
        },
        "value": {
          "description": "The value to be used within the operations.",
          "type": "object"
        }
      }
    },
    "PeerAnnouncement": {
      "description": "Announcent of a peer on the network",
      "type": "object",
      "properties": {
        "networkUuid": {
          "description": "Uuid of the network.",
          "type": "string",
          "format": "uuid"
        },
        "networkVoucherUuid": {
          "description": "Voucher that allows access or not to the network.",
          "type": "string",
          "format": "uuid"
        },
        "peerHost": {
          "description": "Host or IP of the peer.",
          "type": "string",
          "format": "hostname"
        },
        "peerName": {
          "description": "Name of the peer in readable format",
          "type": "string"
        },
        "peerUuid": {
          "description": "Uuid of the peer.",
          "type": "string",
          "format": "uuid"
        }
      }
    },
    "QuestionCreate": {
      "type": "object",
      "properties": {
        "answerUuid": {
          "description": "The Uuid of the answer when generated and returned to the /answer endpoint.",
          "type": "string",
          "format": "uuid"
        },
        "question": {
          "$ref": "#/definitions/VectorBasedQuestion"
        },
        "returnTo": {
          "type": "object",
          "properties": {
            "host": {
              "description": "The answer should be returned to which host?",
              "type": "string"
            },
            "port": {
              "description": "The answer should be returned to which port?",
              "type": "string"
            }
          }
        }
      }
    },
    "QuestionResponse": {
      "type": "object",
      "properties": {
        "answerUuid": {
          "description": "The Uuid of the answer when generated and returned to the /answer endpoint.",
          "type": "string",
          "format": "uuid"
        }
      }
    },
    "Schema": {
      "description": "This is an open object, with OpenAPI Specification 3.0 this will be more detailed. See Weaviate docs for more info. In the future this will become a key/value OR a SingleRef definition",
      "type": "object"
    },
    "SchemaHistory": {
      "description": "This is an open object, with OpenAPI Specification 3.0 this will be more detailed. See Weaviate docs for more info. In the future this will become a key/value OR a SingleRef definition",
      "type": "object"
    },
    "SemanticSchema": {
      "description": "Definitions of semantic schemas (also see: https://github.com/creativesoftwarefdn/weaviate-semantic-schemas)",
      "type": "object",
      "properties": {
        "@context": {
          "description": "URL of the context",
          "type": "string",
          "format": "uri"
        },
        "classes": {
          "description": "Semantic classes that are available.",
          "type": "array",
          "items": {
            "$ref": "#/definitions/SemanticSchemaClass"
          }
        },
        "maintainer": {
          "description": "Email of the maintainer.",
          "type": "string",
          "format": "email"
        },
        "name": {
          "description": "Name of the schema",
          "type": "string"
        },
        "type": {
          "description": "Type of schema, should be \"thing\" or \"action\".",
          "type": "string",
          "enum": [
            "thing",
            "action"
          ]
        },
        "version": {
          "description": "Version number of the schema in semver format.",
          "type": "string"
        }
      }
    },
    "SemanticSchemaClass": {
      "type": "object",
      "properties": {
        "class": {
          "description": "Name of the class as URI relative to the schema URL.",
          "type": "string"
        },
        "description": {
          "description": "Description of the class",
          "type": "string"
        },
        "keywords": {
          "description": "Describes the kind of class. For example Geolocation for the class City.",
          "type": "array",
          "items": {
            "type": "object",
            "properties": {
              "kind": {
                "type": "string"
              },
              "weight": {
                "type": "number",
                "format": "float"
              }
            }
          }
        },
        "properties": {
          "description": "The properties of the class.",
          "type": "array",
          "items": {
            "$ref": "#/definitions/SemanticSchemaClassProperty"
          }
        }
      }
    },
    "SemanticSchemaClassProperty": {
      "type": "object",
      "properties": {
        "@dataType": {
          "description": "Can be a reference ($cref) to another type when starts with a capital (for example Person) otherwise \"string\" or \"int\".",
          "type": "array",
          "items": {
            "type": "string"
          }
        },
        "description": {
          "description": "Description of the property",
          "type": "string"
        },
        "keywords": {
          "description": "Describes the kind of class. For example Geolocation for the class City.",
          "type": "array",
          "items": {
            "type": "object",
            "properties": {
              "kind": {
                "type": "string"
              },
              "weight": {
                "type": "number",
                "format": "float"
              }
            }
          }
        },
        "name": {
          "description": "Name of the property as URI relative to the schema URL.",
          "type": "string"
        }
      }
    },
    "SingleRef": {
      "properties": {
        "$cref": {
          "description": "Location of the cross reference.",
          "type": "string",
          "format": "uuid"
        },
        "locationUrl": {
          "description": "url of location. http://localhost means this database. This option can be used to refer to other databases.",
          "type": "string",
          "format": "url",
          "default": "http://localhost/"
        },
        "type": {
          "description": "Type should be Thing, Action or Key",
          "type": "string",
          "enum": [
            "Thing",
            "Action",
            "Key"
          ]
        }
      }
    },
    "Thing": {
      "allOf": [
        {
          "$ref": "#/definitions/ThingCreate"
        },
        {
          "type": "object",
          "properties": {
            "creationTimeUnix": {
              "description": "Timestamp of creation of this thing in milliseconds since epoch UTC.",
              "type": "integer",
              "format": "int64"
            },
            "key": {
              "$ref": "#/definitions/SingleRef"
            },
            "lastUpdateTimeUnix": {
              "description": "Timestamp of the last thing update in milliseconds since epoch UTC.",
              "type": "integer",
              "format": "int64"
            }
          }
        }
      ]
    },
    "ThingCreate": {
      "type": "object",
      "properties": {
        "@class": {
          "description": "Class of the Thing, defined in the schema.",
          "type": "string"
        },
        "@context": {
          "description": "Available context schema.",
          "type": "string"
        },
        "schema": {
          "$ref": "#/definitions/Schema"
        }
      }
    },
    "ThingGetHistoryResponse": {
      "allOf": [
        {
          "$ref": "#/definitions/ThingHistory"
        },
        {
          "type": "object",
          "properties": {
            "thingId": {
              "type": "string",
              "format": "uuid"
            }
          }
        }
      ]
    },
    "ThingGetResponse": {
      "allOf": [
        {
          "$ref": "#/definitions/Thing"
        },
        {
          "type": "object",
          "properties": {
            "thingId": {
              "type": "string",
              "format": "uuid"
            }
          }
        }
      ]
    },
    "ThingHistory": {
      "type": "object",
      "properties": {
        "deleted": {
          "description": "Indication whether the action is deleted",
          "type": "boolean"
        },
        "key": {
          "$ref": "#/definitions/SingleRef"
        },
        "propertyHistory": {
          "description": "An array with the history of the things.",
          "type": "array",
          "items": {
            "$ref": "#/definitions/ThingHistoryObject"
          }
        }
      }
    },
    "ThingHistoryObject": {
      "allOf": [
        {
          "$ref": "#/definitions/ThingCreate"
        },
        {
          "type": "object",
          "properties": {
            "creationTimeUnix": {
              "description": "Timestamp of creation of this thing history in milliseconds since epoch UTC.",
              "type": "integer",
              "format": "int64"
            }
          }
        }
      ]
    },
    "ThingUpdate": {
      "allOf": [
        {
          "$ref": "#/definitions/Thing"
        },
        {
          "type": "object"
        }
      ]
    },
    "ThingsListResponse": {
      "description": "List of things.",
      "type": "object",
      "properties": {
        "things": {
          "description": "The actual list of things.",
          "type": "array",
          "items": {
            "$ref": "#/definitions/ThingGetResponse"
          }
        },
        "totalResults": {
          "description": "The total number of things for the query. The number of items in a response may be smaller due to paging.",
          "type": "integer",
          "format": "int64"
        }
      }
    },
    "VectorBasedQuestion": {
      "description": "receive question based on array of classes, properties and values.",
      "type": "array",
      "items": {
        "type": "object",
        "properties": {
          "classProps": {
            "description": "Vectorized properties.",
            "type": "array",
            "maxItems": 300,
            "minItems": 300,
            "items": {
              "type": "object",
              "properties": {
                "propsVectors": {
                  "type": "array",
                  "items": {
                    "type": "number",
                    "format": "float"
                  }
                },
                "value": {
                  "description": "String with valuename.",
                  "type": "string"
                }
              }
            }
          },
          "classVectors": {
            "description": "Vectorized classname.",
            "type": "array",
            "maxItems": 300,
            "minItems": 300,
            "items": {
              "type": "number",
              "format": "float"
            }
          }
        }
      }
    }
  },
  "parameters": {
    "CommonMaxResultsParameterQuery": {
      "type": "integer",
      "format": "int64",
      "description": "The maximum number of items to be returned per page. Default value is set in Weaviate config.",
      "name": "maxResults",
      "in": "query"
    },
    "CommonPageParameterQuery": {
      "type": "integer",
      "format": "int64",
      "description": "The page number of the items to be returned.",
      "name": "page",
      "in": "query"
    }
  },
  "securityDefinitions": {
    "apiKey": {
      "type": "apiKey",
      "name": "X-API-KEY",
      "in": "header"
    },
    "apiToken": {
      "type": "apiKey",
      "name": "X-API-TOKEN",
      "in": "header"
    }
  },
  "security": [
    {
      "apiKey": [],
      "apiToken": []
    }
  ],
  "tags": [
    {
      "name": "actions"
    },
    {
      "name": "graphql"
    },
    {
      "name": "keys"
    },
    {
      "name": "meta"
    },
    {
      "name": "P2P"
    },
    {
      "name": "things"
    }
  ],
  "externalDocs": {
    "url": "https://github.com/creativesoftwarefdn/weaviate"
  }
}`))
}
