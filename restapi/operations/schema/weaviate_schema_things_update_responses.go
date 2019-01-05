/*                          _       _
 *__      _____  __ ___   ___  __ _| |_ ___
 *\ \ /\ / / _ \/ _` \ \ / / |/ _` | __/ _ \
 * \ V  V /  __/ (_| |\ V /| | (_| | ||  __/
 *  \_/\_/ \___|\__,_| \_/ |_|\__,_|\__\___|
 *
 * Copyright © 2016 - 2019 Weaviate. All rights reserved.
 * LICENSE: https://github.com/creativesoftwarefdn/weaviate/blob/develop/LICENSE.md
 * DESIGN: Bob van Luijt (bob@k10y.co)
 */
// Code generated by go-swagger; DO NOT EDIT.

package schema

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "github.com/creativesoftwarefdn/weaviate/models"
)

// WeaviateSchemaThingsUpdateOKCode is the HTTP code returned for type WeaviateSchemaThingsUpdateOK
const WeaviateSchemaThingsUpdateOKCode int = 200

/*WeaviateSchemaThingsUpdateOK Changes applied.

swagger:response weaviateSchemaThingsUpdateOK
*/
type WeaviateSchemaThingsUpdateOK struct {
}

// NewWeaviateSchemaThingsUpdateOK creates WeaviateSchemaThingsUpdateOK with default headers values
func NewWeaviateSchemaThingsUpdateOK() *WeaviateSchemaThingsUpdateOK {

	return &WeaviateSchemaThingsUpdateOK{}
}

// WriteResponse to the client
func (o *WeaviateSchemaThingsUpdateOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(200)
}

// WeaviateSchemaThingsUpdateUnauthorizedCode is the HTTP code returned for type WeaviateSchemaThingsUpdateUnauthorized
const WeaviateSchemaThingsUpdateUnauthorizedCode int = 401

/*WeaviateSchemaThingsUpdateUnauthorized Unauthorized or invalid credentials.

swagger:response weaviateSchemaThingsUpdateUnauthorized
*/
type WeaviateSchemaThingsUpdateUnauthorized struct {
}

// NewWeaviateSchemaThingsUpdateUnauthorized creates WeaviateSchemaThingsUpdateUnauthorized with default headers values
func NewWeaviateSchemaThingsUpdateUnauthorized() *WeaviateSchemaThingsUpdateUnauthorized {

	return &WeaviateSchemaThingsUpdateUnauthorized{}
}

// WriteResponse to the client
func (o *WeaviateSchemaThingsUpdateUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}

// WeaviateSchemaThingsUpdateForbiddenCode is the HTTP code returned for type WeaviateSchemaThingsUpdateForbidden
const WeaviateSchemaThingsUpdateForbiddenCode int = 403

/*WeaviateSchemaThingsUpdateForbidden Could not find the Thing class.

swagger:response weaviateSchemaThingsUpdateForbidden
*/
type WeaviateSchemaThingsUpdateForbidden struct {
}

// NewWeaviateSchemaThingsUpdateForbidden creates WeaviateSchemaThingsUpdateForbidden with default headers values
func NewWeaviateSchemaThingsUpdateForbidden() *WeaviateSchemaThingsUpdateForbidden {

	return &WeaviateSchemaThingsUpdateForbidden{}
}

// WriteResponse to the client
func (o *WeaviateSchemaThingsUpdateForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(403)
}

// WeaviateSchemaThingsUpdateUnprocessableEntityCode is the HTTP code returned for type WeaviateSchemaThingsUpdateUnprocessableEntity
const WeaviateSchemaThingsUpdateUnprocessableEntityCode int = 422

/*WeaviateSchemaThingsUpdateUnprocessableEntity Invalid update.

swagger:response weaviateSchemaThingsUpdateUnprocessableEntity
*/
type WeaviateSchemaThingsUpdateUnprocessableEntity struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewWeaviateSchemaThingsUpdateUnprocessableEntity creates WeaviateSchemaThingsUpdateUnprocessableEntity with default headers values
func NewWeaviateSchemaThingsUpdateUnprocessableEntity() *WeaviateSchemaThingsUpdateUnprocessableEntity {

	return &WeaviateSchemaThingsUpdateUnprocessableEntity{}
}

// WithPayload adds the payload to the weaviate schema things update unprocessable entity response
func (o *WeaviateSchemaThingsUpdateUnprocessableEntity) WithPayload(payload *models.ErrorResponse) *WeaviateSchemaThingsUpdateUnprocessableEntity {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the weaviate schema things update unprocessable entity response
func (o *WeaviateSchemaThingsUpdateUnprocessableEntity) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *WeaviateSchemaThingsUpdateUnprocessableEntity) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(422)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
