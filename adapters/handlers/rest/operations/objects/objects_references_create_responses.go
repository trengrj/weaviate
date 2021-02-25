//                           _       _
// __      _____  __ ___   ___  __ _| |_ ___
// \ \ /\ / / _ \/ _` \ \ / / |/ _` | __/ _ \
//  \ V  V /  __/ (_| |\ V /| | (_| | ||  __/
//   \_/\_/ \___|\__,_| \_/ |_|\__,_|\__\___|
//
//  Copyright © 2016 - 2021 SeMI Technologies B.V. All rights reserved.
//
//  CONTACT: hello@semi.technology
//

// Code generated by go-swagger; DO NOT EDIT.

package objects

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/semi-technologies/weaviate/entities/models"
)

// ObjectsReferencesCreateOKCode is the HTTP code returned for type ObjectsReferencesCreateOK
const ObjectsReferencesCreateOKCode int = 200

/*ObjectsReferencesCreateOK Successfully added the reference.

swagger:response objectsReferencesCreateOK
*/
type ObjectsReferencesCreateOK struct {
}

// NewObjectsReferencesCreateOK creates ObjectsReferencesCreateOK with default headers values
func NewObjectsReferencesCreateOK() *ObjectsReferencesCreateOK {

	return &ObjectsReferencesCreateOK{}
}

// WriteResponse to the client
func (o *ObjectsReferencesCreateOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(200)
}

// ObjectsReferencesCreateUnauthorizedCode is the HTTP code returned for type ObjectsReferencesCreateUnauthorized
const ObjectsReferencesCreateUnauthorizedCode int = 401

/*ObjectsReferencesCreateUnauthorized Unauthorized or invalid credentials.

swagger:response objectsReferencesCreateUnauthorized
*/
type ObjectsReferencesCreateUnauthorized struct {
}

// NewObjectsReferencesCreateUnauthorized creates ObjectsReferencesCreateUnauthorized with default headers values
func NewObjectsReferencesCreateUnauthorized() *ObjectsReferencesCreateUnauthorized {

	return &ObjectsReferencesCreateUnauthorized{}
}

// WriteResponse to the client
func (o *ObjectsReferencesCreateUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}

// ObjectsReferencesCreateForbiddenCode is the HTTP code returned for type ObjectsReferencesCreateForbidden
const ObjectsReferencesCreateForbiddenCode int = 403

/*ObjectsReferencesCreateForbidden Forbidden

swagger:response objectsReferencesCreateForbidden
*/
type ObjectsReferencesCreateForbidden struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewObjectsReferencesCreateForbidden creates ObjectsReferencesCreateForbidden with default headers values
func NewObjectsReferencesCreateForbidden() *ObjectsReferencesCreateForbidden {

	return &ObjectsReferencesCreateForbidden{}
}

// WithPayload adds the payload to the objects references create forbidden response
func (o *ObjectsReferencesCreateForbidden) WithPayload(payload *models.ErrorResponse) *ObjectsReferencesCreateForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the objects references create forbidden response
func (o *ObjectsReferencesCreateForbidden) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ObjectsReferencesCreateForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ObjectsReferencesCreateUnprocessableEntityCode is the HTTP code returned for type ObjectsReferencesCreateUnprocessableEntity
const ObjectsReferencesCreateUnprocessableEntityCode int = 422

/*ObjectsReferencesCreateUnprocessableEntity Request body is well-formed (i.e., syntactically correct), but semantically erroneous. Are you sure the property exists or that it is a class?

swagger:response objectsReferencesCreateUnprocessableEntity
*/
type ObjectsReferencesCreateUnprocessableEntity struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewObjectsReferencesCreateUnprocessableEntity creates ObjectsReferencesCreateUnprocessableEntity with default headers values
func NewObjectsReferencesCreateUnprocessableEntity() *ObjectsReferencesCreateUnprocessableEntity {

	return &ObjectsReferencesCreateUnprocessableEntity{}
}

// WithPayload adds the payload to the objects references create unprocessable entity response
func (o *ObjectsReferencesCreateUnprocessableEntity) WithPayload(payload *models.ErrorResponse) *ObjectsReferencesCreateUnprocessableEntity {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the objects references create unprocessable entity response
func (o *ObjectsReferencesCreateUnprocessableEntity) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ObjectsReferencesCreateUnprocessableEntity) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(422)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ObjectsReferencesCreateInternalServerErrorCode is the HTTP code returned for type ObjectsReferencesCreateInternalServerError
const ObjectsReferencesCreateInternalServerErrorCode int = 500

/*ObjectsReferencesCreateInternalServerError An error has occurred while trying to fulfill the request. Most likely the ErrorResponse will contain more information about the error.

swagger:response objectsReferencesCreateInternalServerError
*/
type ObjectsReferencesCreateInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewObjectsReferencesCreateInternalServerError creates ObjectsReferencesCreateInternalServerError with default headers values
func NewObjectsReferencesCreateInternalServerError() *ObjectsReferencesCreateInternalServerError {

	return &ObjectsReferencesCreateInternalServerError{}
}

// WithPayload adds the payload to the objects references create internal server error response
func (o *ObjectsReferencesCreateInternalServerError) WithPayload(payload *models.ErrorResponse) *ObjectsReferencesCreateInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the objects references create internal server error response
func (o *ObjectsReferencesCreateInternalServerError) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ObjectsReferencesCreateInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
