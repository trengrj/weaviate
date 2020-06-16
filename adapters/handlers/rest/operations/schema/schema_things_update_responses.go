//                           _       _
// __      _____  __ ___   ___  __ _| |_ ___
// \ \ /\ / / _ \/ _` \ \ / / |/ _` | __/ _ \
//  \ V  V /  __/ (_| |\ V /| | (_| | ||  __/
//   \_/\_/ \___|\__,_| \_/ |_|\__,_|\__\___|
//
//  Copyright © 2016 - 2020 SeMI Holding B.V. (registered @ Dutch Chamber of Commerce no 75221632). All rights reserved.
//  LICENSE WEAVIATE OPEN SOURCE: https://www.semi.technology/playbook/playbook/contract-weaviate-OSS.html
//  LICENSE WEAVIATE ENTERPRISE: https://www.semi.technology/playbook/contract-weaviate-enterprise.html
//  CONCEPT: Bob van Luijt (@bobvanluijt)
//  CONTACT: hello@semi.technology
//

// Code generated by go-swagger; DO NOT EDIT.

package schema

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/semi-technologies/weaviate/entities/models"
)

// SchemaThingsUpdateOKCode is the HTTP code returned for type SchemaThingsUpdateOK
const SchemaThingsUpdateOKCode int = 200

/*SchemaThingsUpdateOK Changes applied.

swagger:response schemaThingsUpdateOK
*/
type SchemaThingsUpdateOK struct {
}

// NewSchemaThingsUpdateOK creates SchemaThingsUpdateOK with default headers values
func NewSchemaThingsUpdateOK() *SchemaThingsUpdateOK {

	return &SchemaThingsUpdateOK{}
}

// WriteResponse to the client
func (o *SchemaThingsUpdateOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(200)
}

// SchemaThingsUpdateUnauthorizedCode is the HTTP code returned for type SchemaThingsUpdateUnauthorized
const SchemaThingsUpdateUnauthorizedCode int = 401

/*SchemaThingsUpdateUnauthorized Unauthorized or invalid credentials.

swagger:response schemaThingsUpdateUnauthorized
*/
type SchemaThingsUpdateUnauthorized struct {
}

// NewSchemaThingsUpdateUnauthorized creates SchemaThingsUpdateUnauthorized with default headers values
func NewSchemaThingsUpdateUnauthorized() *SchemaThingsUpdateUnauthorized {

	return &SchemaThingsUpdateUnauthorized{}
}

// WriteResponse to the client
func (o *SchemaThingsUpdateUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}

// SchemaThingsUpdateForbiddenCode is the HTTP code returned for type SchemaThingsUpdateForbidden
const SchemaThingsUpdateForbiddenCode int = 403

/*SchemaThingsUpdateForbidden Forbidden

swagger:response schemaThingsUpdateForbidden
*/
type SchemaThingsUpdateForbidden struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewSchemaThingsUpdateForbidden creates SchemaThingsUpdateForbidden with default headers values
func NewSchemaThingsUpdateForbidden() *SchemaThingsUpdateForbidden {

	return &SchemaThingsUpdateForbidden{}
}

// WithPayload adds the payload to the schema things update forbidden response
func (o *SchemaThingsUpdateForbidden) WithPayload(payload *models.ErrorResponse) *SchemaThingsUpdateForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the schema things update forbidden response
func (o *SchemaThingsUpdateForbidden) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *SchemaThingsUpdateForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// SchemaThingsUpdateUnprocessableEntityCode is the HTTP code returned for type SchemaThingsUpdateUnprocessableEntity
const SchemaThingsUpdateUnprocessableEntityCode int = 422

/*SchemaThingsUpdateUnprocessableEntity Invalid update.

swagger:response schemaThingsUpdateUnprocessableEntity
*/
type SchemaThingsUpdateUnprocessableEntity struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewSchemaThingsUpdateUnprocessableEntity creates SchemaThingsUpdateUnprocessableEntity with default headers values
func NewSchemaThingsUpdateUnprocessableEntity() *SchemaThingsUpdateUnprocessableEntity {

	return &SchemaThingsUpdateUnprocessableEntity{}
}

// WithPayload adds the payload to the schema things update unprocessable entity response
func (o *SchemaThingsUpdateUnprocessableEntity) WithPayload(payload *models.ErrorResponse) *SchemaThingsUpdateUnprocessableEntity {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the schema things update unprocessable entity response
func (o *SchemaThingsUpdateUnprocessableEntity) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *SchemaThingsUpdateUnprocessableEntity) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(422)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// SchemaThingsUpdateInternalServerErrorCode is the HTTP code returned for type SchemaThingsUpdateInternalServerError
const SchemaThingsUpdateInternalServerErrorCode int = 500

/*SchemaThingsUpdateInternalServerError An error has occurred while trying to fulfill the request. Most likely the ErrorResponse will contain more information about the error.

swagger:response schemaThingsUpdateInternalServerError
*/
type SchemaThingsUpdateInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewSchemaThingsUpdateInternalServerError creates SchemaThingsUpdateInternalServerError with default headers values
func NewSchemaThingsUpdateInternalServerError() *SchemaThingsUpdateInternalServerError {

	return &SchemaThingsUpdateInternalServerError{}
}

// WithPayload adds the payload to the schema things update internal server error response
func (o *SchemaThingsUpdateInternalServerError) WithPayload(payload *models.ErrorResponse) *SchemaThingsUpdateInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the schema things update internal server error response
func (o *SchemaThingsUpdateInternalServerError) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *SchemaThingsUpdateInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
