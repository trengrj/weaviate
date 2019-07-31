//                           _       _
// __      _____  __ ___   ___  __ _| |_ ___
// \ \ /\ / / _ \/ _` \ \ / / |/ _` | __/ _ \
//  \ V  V /  __/ (_| |\ V /| | (_| | ||  __/
//   \_/\_/ \___|\__,_| \_/ |_|\__,_|\__\___|
//
//  Copyright © 2016 - 2019 Weaviate. All rights reserved.
//  LICENSE: https://github.com/semi-technologies/weaviate/blob/develop/LICENSE.md
//  DESIGN & CONCEPT: Bob van Luijt (@bobvanluijt)
//  CONTACT: hello@semi.technology
//

// Code generated by go-swagger; DO NOT EDIT.

package things

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "github.com/semi-technologies/weaviate/entities/models"
)

// ThingsGetOKCode is the HTTP code returned for type ThingsGetOK
const ThingsGetOKCode int = 200

/*ThingsGetOK Successful response.

swagger:response thingsGetOK
*/
type ThingsGetOK struct {

	/*
	  In: Body
	*/
	Payload *models.Thing `json:"body,omitempty"`
}

// NewThingsGetOK creates ThingsGetOK with default headers values
func NewThingsGetOK() *ThingsGetOK {

	return &ThingsGetOK{}
}

// WithPayload adds the payload to the things get o k response
func (o *ThingsGetOK) WithPayload(payload *models.Thing) *ThingsGetOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the things get o k response
func (o *ThingsGetOK) SetPayload(payload *models.Thing) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ThingsGetOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ThingsGetUnauthorizedCode is the HTTP code returned for type ThingsGetUnauthorized
const ThingsGetUnauthorizedCode int = 401

/*ThingsGetUnauthorized Unauthorized or invalid credentials.

swagger:response thingsGetUnauthorized
*/
type ThingsGetUnauthorized struct {
}

// NewThingsGetUnauthorized creates ThingsGetUnauthorized with default headers values
func NewThingsGetUnauthorized() *ThingsGetUnauthorized {

	return &ThingsGetUnauthorized{}
}

// WriteResponse to the client
func (o *ThingsGetUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}

// ThingsGetForbiddenCode is the HTTP code returned for type ThingsGetForbidden
const ThingsGetForbiddenCode int = 403

/*ThingsGetForbidden Forbidden

swagger:response thingsGetForbidden
*/
type ThingsGetForbidden struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewThingsGetForbidden creates ThingsGetForbidden with default headers values
func NewThingsGetForbidden() *ThingsGetForbidden {

	return &ThingsGetForbidden{}
}

// WithPayload adds the payload to the things get forbidden response
func (o *ThingsGetForbidden) WithPayload(payload *models.ErrorResponse) *ThingsGetForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the things get forbidden response
func (o *ThingsGetForbidden) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ThingsGetForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ThingsGetNotFoundCode is the HTTP code returned for type ThingsGetNotFound
const ThingsGetNotFoundCode int = 404

/*ThingsGetNotFound Successful query result but no resource was found.

swagger:response thingsGetNotFound
*/
type ThingsGetNotFound struct {
}

// NewThingsGetNotFound creates ThingsGetNotFound with default headers values
func NewThingsGetNotFound() *ThingsGetNotFound {

	return &ThingsGetNotFound{}
}

// WriteResponse to the client
func (o *ThingsGetNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(404)
}

// ThingsGetInternalServerErrorCode is the HTTP code returned for type ThingsGetInternalServerError
const ThingsGetInternalServerErrorCode int = 500

/*ThingsGetInternalServerError An error has occurred while trying to fulfill the request. Most likely the ErrorResponse will contain more information about the error.

swagger:response thingsGetInternalServerError
*/
type ThingsGetInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewThingsGetInternalServerError creates ThingsGetInternalServerError with default headers values
func NewThingsGetInternalServerError() *ThingsGetInternalServerError {

	return &ThingsGetInternalServerError{}
}

// WithPayload adds the payload to the things get internal server error response
func (o *ThingsGetInternalServerError) WithPayload(payload *models.ErrorResponse) *ThingsGetInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the things get internal server error response
func (o *ThingsGetInternalServerError) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ThingsGetInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
