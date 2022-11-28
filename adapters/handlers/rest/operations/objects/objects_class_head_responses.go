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

package objects

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/semi-technologies/weaviate/models"
)

// ObjectsClassHeadNoContentCode is the HTTP code returned for type ObjectsClassHeadNoContent
const ObjectsClassHeadNoContentCode int = 204

/*
ObjectsClassHeadNoContent Object exists.

swagger:response objectsClassHeadNoContent
*/
type ObjectsClassHeadNoContent struct {
}

// NewObjectsClassHeadNoContent creates ObjectsClassHeadNoContent with default headers values
func NewObjectsClassHeadNoContent() *ObjectsClassHeadNoContent {

	return &ObjectsClassHeadNoContent{}
}

// WriteResponse to the client
func (o *ObjectsClassHeadNoContent) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(204)
}

// ObjectsClassHeadUnauthorizedCode is the HTTP code returned for type ObjectsClassHeadUnauthorized
const ObjectsClassHeadUnauthorizedCode int = 401

/*
ObjectsClassHeadUnauthorized Unauthorized or invalid credentials.

swagger:response objectsClassHeadUnauthorized
*/
type ObjectsClassHeadUnauthorized struct {
}

// NewObjectsClassHeadUnauthorized creates ObjectsClassHeadUnauthorized with default headers values
func NewObjectsClassHeadUnauthorized() *ObjectsClassHeadUnauthorized {

	return &ObjectsClassHeadUnauthorized{}
}

// WriteResponse to the client
func (o *ObjectsClassHeadUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}

// ObjectsClassHeadForbiddenCode is the HTTP code returned for type ObjectsClassHeadForbidden
const ObjectsClassHeadForbiddenCode int = 403

/*
ObjectsClassHeadForbidden Forbidden

swagger:response objectsClassHeadForbidden
*/
type ObjectsClassHeadForbidden struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewObjectsClassHeadForbidden creates ObjectsClassHeadForbidden with default headers values
func NewObjectsClassHeadForbidden() *ObjectsClassHeadForbidden {

	return &ObjectsClassHeadForbidden{}
}

// WithPayload adds the payload to the objects class head forbidden response
func (o *ObjectsClassHeadForbidden) WithPayload(payload *models.ErrorResponse) *ObjectsClassHeadForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the objects class head forbidden response
func (o *ObjectsClassHeadForbidden) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ObjectsClassHeadForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ObjectsClassHeadNotFoundCode is the HTTP code returned for type ObjectsClassHeadNotFound
const ObjectsClassHeadNotFoundCode int = 404

/*
ObjectsClassHeadNotFound Object doesn't exist.

swagger:response objectsClassHeadNotFound
*/
type ObjectsClassHeadNotFound struct {
}

// NewObjectsClassHeadNotFound creates ObjectsClassHeadNotFound with default headers values
func NewObjectsClassHeadNotFound() *ObjectsClassHeadNotFound {

	return &ObjectsClassHeadNotFound{}
}

// WriteResponse to the client
func (o *ObjectsClassHeadNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(404)
}

// ObjectsClassHeadInternalServerErrorCode is the HTTP code returned for type ObjectsClassHeadInternalServerError
const ObjectsClassHeadInternalServerErrorCode int = 500

/*
ObjectsClassHeadInternalServerError An error has occurred while trying to fulfill the request. Most likely the ErrorResponse will contain more information about the error.

swagger:response objectsClassHeadInternalServerError
*/
type ObjectsClassHeadInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewObjectsClassHeadInternalServerError creates ObjectsClassHeadInternalServerError with default headers values
func NewObjectsClassHeadInternalServerError() *ObjectsClassHeadInternalServerError {

	return &ObjectsClassHeadInternalServerError{}
}

// WithPayload adds the payload to the objects class head internal server error response
func (o *ObjectsClassHeadInternalServerError) WithPayload(payload *models.ErrorResponse) *ObjectsClassHeadInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the objects class head internal server error response
func (o *ObjectsClassHeadInternalServerError) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ObjectsClassHeadInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
