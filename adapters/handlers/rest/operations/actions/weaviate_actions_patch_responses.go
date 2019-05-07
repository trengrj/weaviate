/*                          _       _
 *__      _____  __ ___   ___  __ _| |_ ___
 *\ \ /\ / / _ \/ _` \ \ / / |/ _` | __/ _ \
 * \ V  V /  __/ (_| |\ V /| | (_| | ||  __/
 *  \_/\_/ \___|\__,_| \_/ |_|\__,_|\__\___|
 *
 * Copyright © 2016 - 2019 Weaviate. All rights reserved.
 * LICENSE: https://github.com/semi-technologies/weaviate/blob/develop/LICENSE.md
 * DESIGN & CONCEPT: Bob van Luijt (@bobvanluijt)
 * CONTACT: hello@semi.technology
 */ // Code generated by go-swagger; DO NOT EDIT.

package actions

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "github.com/semi-technologies/weaviate/entities/models"
)

// WeaviateActionsPatchOKCode is the HTTP code returned for type WeaviateActionsPatchOK
const WeaviateActionsPatchOKCode int = 200

/*WeaviateActionsPatchOK Successfully applied.

swagger:response weaviateActionsPatchOK
*/
type WeaviateActionsPatchOK struct {

	/*
	  In: Body
	*/
	Payload *models.Action `json:"body,omitempty"`
}

// NewWeaviateActionsPatchOK creates WeaviateActionsPatchOK with default headers values
func NewWeaviateActionsPatchOK() *WeaviateActionsPatchOK {

	return &WeaviateActionsPatchOK{}
}

// WithPayload adds the payload to the weaviate actions patch o k response
func (o *WeaviateActionsPatchOK) WithPayload(payload *models.Action) *WeaviateActionsPatchOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the weaviate actions patch o k response
func (o *WeaviateActionsPatchOK) SetPayload(payload *models.Action) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *WeaviateActionsPatchOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// WeaviateActionsPatchBadRequestCode is the HTTP code returned for type WeaviateActionsPatchBadRequest
const WeaviateActionsPatchBadRequestCode int = 400

/*WeaviateActionsPatchBadRequest The patch-JSON is malformed.

swagger:response weaviateActionsPatchBadRequest
*/
type WeaviateActionsPatchBadRequest struct {
}

// NewWeaviateActionsPatchBadRequest creates WeaviateActionsPatchBadRequest with default headers values
func NewWeaviateActionsPatchBadRequest() *WeaviateActionsPatchBadRequest {

	return &WeaviateActionsPatchBadRequest{}
}

// WriteResponse to the client
func (o *WeaviateActionsPatchBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(400)
}

// WeaviateActionsPatchUnauthorizedCode is the HTTP code returned for type WeaviateActionsPatchUnauthorized
const WeaviateActionsPatchUnauthorizedCode int = 401

/*WeaviateActionsPatchUnauthorized Unauthorized or invalid credentials.

swagger:response weaviateActionsPatchUnauthorized
*/
type WeaviateActionsPatchUnauthorized struct {
}

// NewWeaviateActionsPatchUnauthorized creates WeaviateActionsPatchUnauthorized with default headers values
func NewWeaviateActionsPatchUnauthorized() *WeaviateActionsPatchUnauthorized {

	return &WeaviateActionsPatchUnauthorized{}
}

// WriteResponse to the client
func (o *WeaviateActionsPatchUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}

// WeaviateActionsPatchForbiddenCode is the HTTP code returned for type WeaviateActionsPatchForbidden
const WeaviateActionsPatchForbiddenCode int = 403

/*WeaviateActionsPatchForbidden Insufficient permissions.

swagger:response weaviateActionsPatchForbidden
*/
type WeaviateActionsPatchForbidden struct {
}

// NewWeaviateActionsPatchForbidden creates WeaviateActionsPatchForbidden with default headers values
func NewWeaviateActionsPatchForbidden() *WeaviateActionsPatchForbidden {

	return &WeaviateActionsPatchForbidden{}
}

// WriteResponse to the client
func (o *WeaviateActionsPatchForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(403)
}

// WeaviateActionsPatchNotFoundCode is the HTTP code returned for type WeaviateActionsPatchNotFound
const WeaviateActionsPatchNotFoundCode int = 404

/*WeaviateActionsPatchNotFound Successful query result but no resource was found.

swagger:response weaviateActionsPatchNotFound
*/
type WeaviateActionsPatchNotFound struct {
}

// NewWeaviateActionsPatchNotFound creates WeaviateActionsPatchNotFound with default headers values
func NewWeaviateActionsPatchNotFound() *WeaviateActionsPatchNotFound {

	return &WeaviateActionsPatchNotFound{}
}

// WriteResponse to the client
func (o *WeaviateActionsPatchNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(404)
}

// WeaviateActionsPatchUnprocessableEntityCode is the HTTP code returned for type WeaviateActionsPatchUnprocessableEntity
const WeaviateActionsPatchUnprocessableEntityCode int = 422

/*WeaviateActionsPatchUnprocessableEntity The patch-JSON is valid but unprocessable.

swagger:response weaviateActionsPatchUnprocessableEntity
*/
type WeaviateActionsPatchUnprocessableEntity struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewWeaviateActionsPatchUnprocessableEntity creates WeaviateActionsPatchUnprocessableEntity with default headers values
func NewWeaviateActionsPatchUnprocessableEntity() *WeaviateActionsPatchUnprocessableEntity {

	return &WeaviateActionsPatchUnprocessableEntity{}
}

// WithPayload adds the payload to the weaviate actions patch unprocessable entity response
func (o *WeaviateActionsPatchUnprocessableEntity) WithPayload(payload *models.ErrorResponse) *WeaviateActionsPatchUnprocessableEntity {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the weaviate actions patch unprocessable entity response
func (o *WeaviateActionsPatchUnprocessableEntity) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *WeaviateActionsPatchUnprocessableEntity) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(422)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// WeaviateActionsPatchInternalServerErrorCode is the HTTP code returned for type WeaviateActionsPatchInternalServerError
const WeaviateActionsPatchInternalServerErrorCode int = 500

/*WeaviateActionsPatchInternalServerError An error has occurred while trying to fulfill the request. Most likely the ErrorResponse will contain more information about the error.

swagger:response weaviateActionsPatchInternalServerError
*/
type WeaviateActionsPatchInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewWeaviateActionsPatchInternalServerError creates WeaviateActionsPatchInternalServerError with default headers values
func NewWeaviateActionsPatchInternalServerError() *WeaviateActionsPatchInternalServerError {

	return &WeaviateActionsPatchInternalServerError{}
}

// WithPayload adds the payload to the weaviate actions patch internal server error response
func (o *WeaviateActionsPatchInternalServerError) WithPayload(payload *models.ErrorResponse) *WeaviateActionsPatchInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the weaviate actions patch internal server error response
func (o *WeaviateActionsPatchInternalServerError) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *WeaviateActionsPatchInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
