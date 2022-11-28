//                           _       _
// __      _____  __ ___   ___  __ _| |_ ___
// \ \ /\ / / _ \/ _` \ \ / / |/ _` | __/ _ \
//  \ V  V /  __/ (_| |\ V /| | (_| | ||  __/
//   \_/\_/ \___|\__,_| \_/ |_|\__,_|\__\___|
//
//  Copyright © 2016 - 2022 SeMI Technologies B.V. All rights reserved.
//
//  CONTACT: hello@semi.technology
//

// Code generated by go-swagger; DO NOT EDIT.

package schema

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/semi-technologies/weaviate/models"
)

// SchemaObjectsShardsUpdateOKCode is the HTTP code returned for type SchemaObjectsShardsUpdateOK
const SchemaObjectsShardsUpdateOKCode int = 200

/*
SchemaObjectsShardsUpdateOK Shard status was updated successfully

swagger:response schemaObjectsShardsUpdateOK
*/
type SchemaObjectsShardsUpdateOK struct {

	/*
	  In: Body
	*/
	Payload *models.ShardStatus `json:"body,omitempty"`
}

// NewSchemaObjectsShardsUpdateOK creates SchemaObjectsShardsUpdateOK with default headers values
func NewSchemaObjectsShardsUpdateOK() *SchemaObjectsShardsUpdateOK {

	return &SchemaObjectsShardsUpdateOK{}
}

// WithPayload adds the payload to the schema objects shards update o k response
func (o *SchemaObjectsShardsUpdateOK) WithPayload(payload *models.ShardStatus) *SchemaObjectsShardsUpdateOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the schema objects shards update o k response
func (o *SchemaObjectsShardsUpdateOK) SetPayload(payload *models.ShardStatus) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *SchemaObjectsShardsUpdateOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// SchemaObjectsShardsUpdateUnauthorizedCode is the HTTP code returned for type SchemaObjectsShardsUpdateUnauthorized
const SchemaObjectsShardsUpdateUnauthorizedCode int = 401

/*
SchemaObjectsShardsUpdateUnauthorized Unauthorized or invalid credentials.

swagger:response schemaObjectsShardsUpdateUnauthorized
*/
type SchemaObjectsShardsUpdateUnauthorized struct {
}

// NewSchemaObjectsShardsUpdateUnauthorized creates SchemaObjectsShardsUpdateUnauthorized with default headers values
func NewSchemaObjectsShardsUpdateUnauthorized() *SchemaObjectsShardsUpdateUnauthorized {

	return &SchemaObjectsShardsUpdateUnauthorized{}
}

// WriteResponse to the client
func (o *SchemaObjectsShardsUpdateUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}

// SchemaObjectsShardsUpdateForbiddenCode is the HTTP code returned for type SchemaObjectsShardsUpdateForbidden
const SchemaObjectsShardsUpdateForbiddenCode int = 403

/*
SchemaObjectsShardsUpdateForbidden Forbidden

swagger:response schemaObjectsShardsUpdateForbidden
*/
type SchemaObjectsShardsUpdateForbidden struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewSchemaObjectsShardsUpdateForbidden creates SchemaObjectsShardsUpdateForbidden with default headers values
func NewSchemaObjectsShardsUpdateForbidden() *SchemaObjectsShardsUpdateForbidden {

	return &SchemaObjectsShardsUpdateForbidden{}
}

// WithPayload adds the payload to the schema objects shards update forbidden response
func (o *SchemaObjectsShardsUpdateForbidden) WithPayload(payload *models.ErrorResponse) *SchemaObjectsShardsUpdateForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the schema objects shards update forbidden response
func (o *SchemaObjectsShardsUpdateForbidden) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *SchemaObjectsShardsUpdateForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// SchemaObjectsShardsUpdateNotFoundCode is the HTTP code returned for type SchemaObjectsShardsUpdateNotFound
const SchemaObjectsShardsUpdateNotFoundCode int = 404

/*
SchemaObjectsShardsUpdateNotFound Shard to be updated does not exist

swagger:response schemaObjectsShardsUpdateNotFound
*/
type SchemaObjectsShardsUpdateNotFound struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewSchemaObjectsShardsUpdateNotFound creates SchemaObjectsShardsUpdateNotFound with default headers values
func NewSchemaObjectsShardsUpdateNotFound() *SchemaObjectsShardsUpdateNotFound {

	return &SchemaObjectsShardsUpdateNotFound{}
}

// WithPayload adds the payload to the schema objects shards update not found response
func (o *SchemaObjectsShardsUpdateNotFound) WithPayload(payload *models.ErrorResponse) *SchemaObjectsShardsUpdateNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the schema objects shards update not found response
func (o *SchemaObjectsShardsUpdateNotFound) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *SchemaObjectsShardsUpdateNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// SchemaObjectsShardsUpdateUnprocessableEntityCode is the HTTP code returned for type SchemaObjectsShardsUpdateUnprocessableEntity
const SchemaObjectsShardsUpdateUnprocessableEntityCode int = 422

/*
SchemaObjectsShardsUpdateUnprocessableEntity Invalid update attempt

swagger:response schemaObjectsShardsUpdateUnprocessableEntity
*/
type SchemaObjectsShardsUpdateUnprocessableEntity struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewSchemaObjectsShardsUpdateUnprocessableEntity creates SchemaObjectsShardsUpdateUnprocessableEntity with default headers values
func NewSchemaObjectsShardsUpdateUnprocessableEntity() *SchemaObjectsShardsUpdateUnprocessableEntity {

	return &SchemaObjectsShardsUpdateUnprocessableEntity{}
}

// WithPayload adds the payload to the schema objects shards update unprocessable entity response
func (o *SchemaObjectsShardsUpdateUnprocessableEntity) WithPayload(payload *models.ErrorResponse) *SchemaObjectsShardsUpdateUnprocessableEntity {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the schema objects shards update unprocessable entity response
func (o *SchemaObjectsShardsUpdateUnprocessableEntity) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *SchemaObjectsShardsUpdateUnprocessableEntity) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(422)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// SchemaObjectsShardsUpdateInternalServerErrorCode is the HTTP code returned for type SchemaObjectsShardsUpdateInternalServerError
const SchemaObjectsShardsUpdateInternalServerErrorCode int = 500

/*
SchemaObjectsShardsUpdateInternalServerError An error has occurred while trying to fulfill the request. Most likely the ErrorResponse will contain more information about the error.

swagger:response schemaObjectsShardsUpdateInternalServerError
*/
type SchemaObjectsShardsUpdateInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewSchemaObjectsShardsUpdateInternalServerError creates SchemaObjectsShardsUpdateInternalServerError with default headers values
func NewSchemaObjectsShardsUpdateInternalServerError() *SchemaObjectsShardsUpdateInternalServerError {

	return &SchemaObjectsShardsUpdateInternalServerError{}
}

// WithPayload adds the payload to the schema objects shards update internal server error response
func (o *SchemaObjectsShardsUpdateInternalServerError) WithPayload(payload *models.ErrorResponse) *SchemaObjectsShardsUpdateInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the schema objects shards update internal server error response
func (o *SchemaObjectsShardsUpdateInternalServerError) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *SchemaObjectsShardsUpdateInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
