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

package things

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "github.com/semi-technologies/weaviate/entities/models"
)

// ThingsListOKCode is the HTTP code returned for type ThingsListOK
const ThingsListOKCode int = 200

/*ThingsListOK Successful response.

swagger:response thingsListOK
*/
type ThingsListOK struct {

	/*
	  In: Body
	*/
	Payload *models.ThingsListResponse `json:"body,omitempty"`
}

// NewThingsListOK creates ThingsListOK with default headers values
func NewThingsListOK() *ThingsListOK {

	return &ThingsListOK{}
}

// WithPayload adds the payload to the things list o k response
func (o *ThingsListOK) WithPayload(payload *models.ThingsListResponse) *ThingsListOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the things list o k response
func (o *ThingsListOK) SetPayload(payload *models.ThingsListResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ThingsListOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ThingsListUnauthorizedCode is the HTTP code returned for type ThingsListUnauthorized
const ThingsListUnauthorizedCode int = 401

/*ThingsListUnauthorized Unauthorized or invalid credentials.

swagger:response thingsListUnauthorized
*/
type ThingsListUnauthorized struct {
}

// NewThingsListUnauthorized creates ThingsListUnauthorized with default headers values
func NewThingsListUnauthorized() *ThingsListUnauthorized {

	return &ThingsListUnauthorized{}
}

// WriteResponse to the client
func (o *ThingsListUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}

// ThingsListForbiddenCode is the HTTP code returned for type ThingsListForbidden
const ThingsListForbiddenCode int = 403

/*ThingsListForbidden Forbidden

swagger:response thingsListForbidden
*/
type ThingsListForbidden struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewThingsListForbidden creates ThingsListForbidden with default headers values
func NewThingsListForbidden() *ThingsListForbidden {

	return &ThingsListForbidden{}
}

// WithPayload adds the payload to the things list forbidden response
func (o *ThingsListForbidden) WithPayload(payload *models.ErrorResponse) *ThingsListForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the things list forbidden response
func (o *ThingsListForbidden) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ThingsListForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ThingsListNotFoundCode is the HTTP code returned for type ThingsListNotFound
const ThingsListNotFoundCode int = 404

/*ThingsListNotFound Successful query result but no resource was found.

swagger:response thingsListNotFound
*/
type ThingsListNotFound struct {
}

// NewThingsListNotFound creates ThingsListNotFound with default headers values
func NewThingsListNotFound() *ThingsListNotFound {

	return &ThingsListNotFound{}
}

// WriteResponse to the client
func (o *ThingsListNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(404)
}

// ThingsListInternalServerErrorCode is the HTTP code returned for type ThingsListInternalServerError
const ThingsListInternalServerErrorCode int = 500

/*ThingsListInternalServerError An error has occurred while trying to fulfill the request. Most likely the ErrorResponse will contain more information about the error.

swagger:response thingsListInternalServerError
*/
type ThingsListInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewThingsListInternalServerError creates ThingsListInternalServerError with default headers values
func NewThingsListInternalServerError() *ThingsListInternalServerError {

	return &ThingsListInternalServerError{}
}

// WithPayload adds the payload to the things list internal server error response
func (o *ThingsListInternalServerError) WithPayload(payload *models.ErrorResponse) *ThingsListInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the things list internal server error response
func (o *ThingsListInternalServerError) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ThingsListInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
