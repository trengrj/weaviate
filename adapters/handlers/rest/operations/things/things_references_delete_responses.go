//                           _       _
// __      _____  __ ___   ___  __ _| |_ ___
// \ \ /\ / / _ \/ _` \ \ / / |/ _` | __/ _ \
//  \ V  V /  __/ (_| |\ V /| | (_| | ||  __/
//   \_/\_/ \___|\__,_| \_/ |_|\__,_|\__\___|
//
//  Copyright © 2016 - 2019 SeMI Holding B.V. (registered @ Dutch Chamber of Commerce no 75221632). All rights reserved.
//  LICENSE WEAVIATE OPEN SOURCE: https://www.semi.technology/playbook/playbook/contract-weaviate-OSS.html
//  LICENSE WEAVIATE ENTERPRISE: https://www.semi.technology/playbook/contract-weaviate-enterprise.html
//  CONCEPT: Bob van Luijt (@bobvanluijt)
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

// ThingsReferencesDeleteNoContentCode is the HTTP code returned for type ThingsReferencesDeleteNoContent
const ThingsReferencesDeleteNoContentCode int = 204

/*ThingsReferencesDeleteNoContent Successfully deleted.

swagger:response thingsReferencesDeleteNoContent
*/
type ThingsReferencesDeleteNoContent struct {
}

// NewThingsReferencesDeleteNoContent creates ThingsReferencesDeleteNoContent with default headers values
func NewThingsReferencesDeleteNoContent() *ThingsReferencesDeleteNoContent {

	return &ThingsReferencesDeleteNoContent{}
}

// WriteResponse to the client
func (o *ThingsReferencesDeleteNoContent) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(204)
}

// ThingsReferencesDeleteUnauthorizedCode is the HTTP code returned for type ThingsReferencesDeleteUnauthorized
const ThingsReferencesDeleteUnauthorizedCode int = 401

/*ThingsReferencesDeleteUnauthorized Unauthorized or invalid credentials.

swagger:response thingsReferencesDeleteUnauthorized
*/
type ThingsReferencesDeleteUnauthorized struct {
}

// NewThingsReferencesDeleteUnauthorized creates ThingsReferencesDeleteUnauthorized with default headers values
func NewThingsReferencesDeleteUnauthorized() *ThingsReferencesDeleteUnauthorized {

	return &ThingsReferencesDeleteUnauthorized{}
}

// WriteResponse to the client
func (o *ThingsReferencesDeleteUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}

// ThingsReferencesDeleteForbiddenCode is the HTTP code returned for type ThingsReferencesDeleteForbidden
const ThingsReferencesDeleteForbiddenCode int = 403

/*ThingsReferencesDeleteForbidden Forbidden

swagger:response thingsReferencesDeleteForbidden
*/
type ThingsReferencesDeleteForbidden struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewThingsReferencesDeleteForbidden creates ThingsReferencesDeleteForbidden with default headers values
func NewThingsReferencesDeleteForbidden() *ThingsReferencesDeleteForbidden {

	return &ThingsReferencesDeleteForbidden{}
}

// WithPayload adds the payload to the things references delete forbidden response
func (o *ThingsReferencesDeleteForbidden) WithPayload(payload *models.ErrorResponse) *ThingsReferencesDeleteForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the things references delete forbidden response
func (o *ThingsReferencesDeleteForbidden) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ThingsReferencesDeleteForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ThingsReferencesDeleteNotFoundCode is the HTTP code returned for type ThingsReferencesDeleteNotFound
const ThingsReferencesDeleteNotFoundCode int = 404

/*ThingsReferencesDeleteNotFound Successful query result but no resource was found.

swagger:response thingsReferencesDeleteNotFound
*/
type ThingsReferencesDeleteNotFound struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewThingsReferencesDeleteNotFound creates ThingsReferencesDeleteNotFound with default headers values
func NewThingsReferencesDeleteNotFound() *ThingsReferencesDeleteNotFound {

	return &ThingsReferencesDeleteNotFound{}
}

// WithPayload adds the payload to the things references delete not found response
func (o *ThingsReferencesDeleteNotFound) WithPayload(payload *models.ErrorResponse) *ThingsReferencesDeleteNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the things references delete not found response
func (o *ThingsReferencesDeleteNotFound) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ThingsReferencesDeleteNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ThingsReferencesDeleteInternalServerErrorCode is the HTTP code returned for type ThingsReferencesDeleteInternalServerError
const ThingsReferencesDeleteInternalServerErrorCode int = 500

/*ThingsReferencesDeleteInternalServerError An error has occurred while trying to fulfill the request. Most likely the ErrorResponse will contain more information about the error.

swagger:response thingsReferencesDeleteInternalServerError
*/
type ThingsReferencesDeleteInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewThingsReferencesDeleteInternalServerError creates ThingsReferencesDeleteInternalServerError with default headers values
func NewThingsReferencesDeleteInternalServerError() *ThingsReferencesDeleteInternalServerError {

	return &ThingsReferencesDeleteInternalServerError{}
}

// WithPayload adds the payload to the things references delete internal server error response
func (o *ThingsReferencesDeleteInternalServerError) WithPayload(payload *models.ErrorResponse) *ThingsReferencesDeleteInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the things references delete internal server error response
func (o *ThingsReferencesDeleteInternalServerError) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ThingsReferencesDeleteInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
