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

package graphql

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/creativesoftwarefdn/weaviate/models"
)

// WeaviateGraphqlPostOKCode is the HTTP code returned for type WeaviateGraphqlPostOK
const WeaviateGraphqlPostOKCode int = 200

/*WeaviateGraphqlPostOK Succesful query (with select).

swagger:response weaviateGraphqlPostOK
*/
type WeaviateGraphqlPostOK struct {

	/*
	  In: Body
	*/
	Payload *models.GraphQLResponse `json:"body,omitempty"`
}

// NewWeaviateGraphqlPostOK creates WeaviateGraphqlPostOK with default headers values
func NewWeaviateGraphqlPostOK() *WeaviateGraphqlPostOK {
	return &WeaviateGraphqlPostOK{}
}

// WithPayload adds the payload to the weaviate graphql post o k response
func (o *WeaviateGraphqlPostOK) WithPayload(payload *models.GraphQLResponse) *WeaviateGraphqlPostOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the weaviate graphql post o k response
func (o *WeaviateGraphqlPostOK) SetPayload(payload *models.GraphQLResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *WeaviateGraphqlPostOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// WeaviateGraphqlPostUnauthorizedCode is the HTTP code returned for type WeaviateGraphqlPostUnauthorized
const WeaviateGraphqlPostUnauthorizedCode int = 401

/*WeaviateGraphqlPostUnauthorized Unauthorized or invalid credentials.

swagger:response weaviateGraphqlPostUnauthorized
*/
type WeaviateGraphqlPostUnauthorized struct {
}

// NewWeaviateGraphqlPostUnauthorized creates WeaviateGraphqlPostUnauthorized with default headers values
func NewWeaviateGraphqlPostUnauthorized() *WeaviateGraphqlPostUnauthorized {
	return &WeaviateGraphqlPostUnauthorized{}
}

// WriteResponse to the client
func (o *WeaviateGraphqlPostUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
}

// WeaviateGraphqlPostForbiddenCode is the HTTP code returned for type WeaviateGraphqlPostForbidden
const WeaviateGraphqlPostForbiddenCode int = 403

/*WeaviateGraphqlPostForbidden The used API-key has insufficient permissions.

swagger:response weaviateGraphqlPostForbidden
*/
type WeaviateGraphqlPostForbidden struct {
}

// NewWeaviateGraphqlPostForbidden creates WeaviateGraphqlPostForbidden with default headers values
func NewWeaviateGraphqlPostForbidden() *WeaviateGraphqlPostForbidden {
	return &WeaviateGraphqlPostForbidden{}
}

// WriteResponse to the client
func (o *WeaviateGraphqlPostForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
}

// WeaviateGraphqlPostUnprocessableEntityCode is the HTTP code returned for type WeaviateGraphqlPostUnprocessableEntity
const WeaviateGraphqlPostUnprocessableEntityCode int = 422

/*WeaviateGraphqlPostUnprocessableEntity Request body contains well-formed (i.e., syntactically correct), but semantically erroneous. Are you sure the class is defined in the configuration file?

swagger:response weaviateGraphqlPostUnprocessableEntity
*/
type WeaviateGraphqlPostUnprocessableEntity struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewWeaviateGraphqlPostUnprocessableEntity creates WeaviateGraphqlPostUnprocessableEntity with default headers values
func NewWeaviateGraphqlPostUnprocessableEntity() *WeaviateGraphqlPostUnprocessableEntity {
	return &WeaviateGraphqlPostUnprocessableEntity{}
}

// WithPayload adds the payload to the weaviate graphql post unprocessable entity response
func (o *WeaviateGraphqlPostUnprocessableEntity) WithPayload(payload *models.ErrorResponse) *WeaviateGraphqlPostUnprocessableEntity {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the weaviate graphql post unprocessable entity response
func (o *WeaviateGraphqlPostUnprocessableEntity) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *WeaviateGraphqlPostUnprocessableEntity) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(422)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// WeaviateGraphqlPostNotImplementedCode is the HTTP code returned for type WeaviateGraphqlPostNotImplemented
const WeaviateGraphqlPostNotImplementedCode int = 501

/*WeaviateGraphqlPostNotImplemented Not (yet) implemented.

swagger:response weaviateGraphqlPostNotImplemented
*/
type WeaviateGraphqlPostNotImplemented struct {
}

// NewWeaviateGraphqlPostNotImplemented creates WeaviateGraphqlPostNotImplemented with default headers values
func NewWeaviateGraphqlPostNotImplemented() *WeaviateGraphqlPostNotImplemented {
	return &WeaviateGraphqlPostNotImplemented{}
}

// WriteResponse to the client
func (o *WeaviateGraphqlPostNotImplemented) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(501)
}
