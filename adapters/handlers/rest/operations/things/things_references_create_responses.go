//                           _       _
// __      _____  __ ___   ___  __ _| |_ ___
// \ \ /\ / / _ \/ _` \ \ / / |/ _` | __/ _ \
//  \ V  V /  __/ (_| |\ V /| | (_| | ||  __/
//   \_/\_/ \___|\__,_| \_/ |_|\__,_|\__\___|
//
//  Copyright © 2016 - 2020 SeMI Technologies B.V. All rights reserved.
//
//  CONTACT: hello@semi.technology
//

// Code generated by go-swagger; DO NOT EDIT.

package things

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/semi-technologies/weaviate/entities/models"
)

// ThingsReferencesCreateOKCode is the HTTP code returned for type ThingsReferencesCreateOK
const ThingsReferencesCreateOKCode int = 200

/*ThingsReferencesCreateOK Successfully added the reference.

swagger:response thingsReferencesCreateOK
*/
type ThingsReferencesCreateOK struct {
}

// NewThingsReferencesCreateOK creates ThingsReferencesCreateOK with default headers values
func NewThingsReferencesCreateOK() *ThingsReferencesCreateOK {

	return &ThingsReferencesCreateOK{}
}

// WriteResponse to the client
func (o *ThingsReferencesCreateOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(200)
}

// ThingsReferencesCreateUnauthorizedCode is the HTTP code returned for type ThingsReferencesCreateUnauthorized
const ThingsReferencesCreateUnauthorizedCode int = 401

/*ThingsReferencesCreateUnauthorized Unauthorized or invalid credentials.

swagger:response thingsReferencesCreateUnauthorized
*/
type ThingsReferencesCreateUnauthorized struct {
}

// NewThingsReferencesCreateUnauthorized creates ThingsReferencesCreateUnauthorized with default headers values
func NewThingsReferencesCreateUnauthorized() *ThingsReferencesCreateUnauthorized {

	return &ThingsReferencesCreateUnauthorized{}
}

// WriteResponse to the client
func (o *ThingsReferencesCreateUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}

// ThingsReferencesCreateForbiddenCode is the HTTP code returned for type ThingsReferencesCreateForbidden
const ThingsReferencesCreateForbiddenCode int = 403

/*ThingsReferencesCreateForbidden Forbidden

swagger:response thingsReferencesCreateForbidden
*/
type ThingsReferencesCreateForbidden struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewThingsReferencesCreateForbidden creates ThingsReferencesCreateForbidden with default headers values
func NewThingsReferencesCreateForbidden() *ThingsReferencesCreateForbidden {

	return &ThingsReferencesCreateForbidden{}
}

// WithPayload adds the payload to the things references create forbidden response
func (o *ThingsReferencesCreateForbidden) WithPayload(payload *models.ErrorResponse) *ThingsReferencesCreateForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the things references create forbidden response
func (o *ThingsReferencesCreateForbidden) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ThingsReferencesCreateForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ThingsReferencesCreateUnprocessableEntityCode is the HTTP code returned for type ThingsReferencesCreateUnprocessableEntity
const ThingsReferencesCreateUnprocessableEntityCode int = 422

/*ThingsReferencesCreateUnprocessableEntity Request body is well-formed (i.e., syntactically correct), but semantically erroneous. Are you sure the property exists or that it is a class?

swagger:response thingsReferencesCreateUnprocessableEntity
*/
type ThingsReferencesCreateUnprocessableEntity struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewThingsReferencesCreateUnprocessableEntity creates ThingsReferencesCreateUnprocessableEntity with default headers values
func NewThingsReferencesCreateUnprocessableEntity() *ThingsReferencesCreateUnprocessableEntity {

	return &ThingsReferencesCreateUnprocessableEntity{}
}

// WithPayload adds the payload to the things references create unprocessable entity response
func (o *ThingsReferencesCreateUnprocessableEntity) WithPayload(payload *models.ErrorResponse) *ThingsReferencesCreateUnprocessableEntity {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the things references create unprocessable entity response
func (o *ThingsReferencesCreateUnprocessableEntity) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ThingsReferencesCreateUnprocessableEntity) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(422)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ThingsReferencesCreateInternalServerErrorCode is the HTTP code returned for type ThingsReferencesCreateInternalServerError
const ThingsReferencesCreateInternalServerErrorCode int = 500

/*ThingsReferencesCreateInternalServerError An error has occurred while trying to fulfill the request. Most likely the ErrorResponse will contain more information about the error.

swagger:response thingsReferencesCreateInternalServerError
*/
type ThingsReferencesCreateInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewThingsReferencesCreateInternalServerError creates ThingsReferencesCreateInternalServerError with default headers values
func NewThingsReferencesCreateInternalServerError() *ThingsReferencesCreateInternalServerError {

	return &ThingsReferencesCreateInternalServerError{}
}

// WithPayload adds the payload to the things references create internal server error response
func (o *ThingsReferencesCreateInternalServerError) WithPayload(payload *models.ErrorResponse) *ThingsReferencesCreateInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the things references create internal server error response
func (o *ThingsReferencesCreateInternalServerError) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ThingsReferencesCreateInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
