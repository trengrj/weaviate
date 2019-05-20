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

// WeaviateActionsReferencesUpdateOKCode is the HTTP code returned for type WeaviateActionsReferencesUpdateOK
const WeaviateActionsReferencesUpdateOKCode int = 200

/*WeaviateActionsReferencesUpdateOK Successfully replaced all the references.

swagger:response weaviateActionsReferencesUpdateOK
*/
type WeaviateActionsReferencesUpdateOK struct {
}

// NewWeaviateActionsReferencesUpdateOK creates WeaviateActionsReferencesUpdateOK with default headers values
func NewWeaviateActionsReferencesUpdateOK() *WeaviateActionsReferencesUpdateOK {

	return &WeaviateActionsReferencesUpdateOK{}
}

// WriteResponse to the client
func (o *WeaviateActionsReferencesUpdateOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(200)
}

// WeaviateActionsReferencesUpdateUnauthorizedCode is the HTTP code returned for type WeaviateActionsReferencesUpdateUnauthorized
const WeaviateActionsReferencesUpdateUnauthorizedCode int = 401

/*WeaviateActionsReferencesUpdateUnauthorized Unauthorized or invalid credentials.

swagger:response weaviateActionsReferencesUpdateUnauthorized
*/
type WeaviateActionsReferencesUpdateUnauthorized struct {
}

// NewWeaviateActionsReferencesUpdateUnauthorized creates WeaviateActionsReferencesUpdateUnauthorized with default headers values
func NewWeaviateActionsReferencesUpdateUnauthorized() *WeaviateActionsReferencesUpdateUnauthorized {

	return &WeaviateActionsReferencesUpdateUnauthorized{}
}

// WriteResponse to the client
func (o *WeaviateActionsReferencesUpdateUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}

// WeaviateActionsReferencesUpdateForbiddenCode is the HTTP code returned for type WeaviateActionsReferencesUpdateForbidden
const WeaviateActionsReferencesUpdateForbiddenCode int = 403

/*WeaviateActionsReferencesUpdateForbidden Forbidden

swagger:response weaviateActionsReferencesUpdateForbidden
*/
type WeaviateActionsReferencesUpdateForbidden struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewWeaviateActionsReferencesUpdateForbidden creates WeaviateActionsReferencesUpdateForbidden with default headers values
func NewWeaviateActionsReferencesUpdateForbidden() *WeaviateActionsReferencesUpdateForbidden {

	return &WeaviateActionsReferencesUpdateForbidden{}
}

// WithPayload adds the payload to the weaviate actions references update forbidden response
func (o *WeaviateActionsReferencesUpdateForbidden) WithPayload(payload *models.ErrorResponse) *WeaviateActionsReferencesUpdateForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the weaviate actions references update forbidden response
func (o *WeaviateActionsReferencesUpdateForbidden) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *WeaviateActionsReferencesUpdateForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// WeaviateActionsReferencesUpdateUnprocessableEntityCode is the HTTP code returned for type WeaviateActionsReferencesUpdateUnprocessableEntity
const WeaviateActionsReferencesUpdateUnprocessableEntityCode int = 422

/*WeaviateActionsReferencesUpdateUnprocessableEntity Request body is well-formed (i.e., syntactically correct), but semantically erroneous. Are you sure the property exists or that it is a class?

swagger:response weaviateActionsReferencesUpdateUnprocessableEntity
*/
type WeaviateActionsReferencesUpdateUnprocessableEntity struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewWeaviateActionsReferencesUpdateUnprocessableEntity creates WeaviateActionsReferencesUpdateUnprocessableEntity with default headers values
func NewWeaviateActionsReferencesUpdateUnprocessableEntity() *WeaviateActionsReferencesUpdateUnprocessableEntity {

	return &WeaviateActionsReferencesUpdateUnprocessableEntity{}
}

// WithPayload adds the payload to the weaviate actions references update unprocessable entity response
func (o *WeaviateActionsReferencesUpdateUnprocessableEntity) WithPayload(payload *models.ErrorResponse) *WeaviateActionsReferencesUpdateUnprocessableEntity {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the weaviate actions references update unprocessable entity response
func (o *WeaviateActionsReferencesUpdateUnprocessableEntity) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *WeaviateActionsReferencesUpdateUnprocessableEntity) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(422)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// WeaviateActionsReferencesUpdateInternalServerErrorCode is the HTTP code returned for type WeaviateActionsReferencesUpdateInternalServerError
const WeaviateActionsReferencesUpdateInternalServerErrorCode int = 500

/*WeaviateActionsReferencesUpdateInternalServerError An error has occurred while trying to fulfill the request. Most likely the ErrorResponse will contain more information about the error.

swagger:response weaviateActionsReferencesUpdateInternalServerError
*/
type WeaviateActionsReferencesUpdateInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewWeaviateActionsReferencesUpdateInternalServerError creates WeaviateActionsReferencesUpdateInternalServerError with default headers values
func NewWeaviateActionsReferencesUpdateInternalServerError() *WeaviateActionsReferencesUpdateInternalServerError {

	return &WeaviateActionsReferencesUpdateInternalServerError{}
}

// WithPayload adds the payload to the weaviate actions references update internal server error response
func (o *WeaviateActionsReferencesUpdateInternalServerError) WithPayload(payload *models.ErrorResponse) *WeaviateActionsReferencesUpdateInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the weaviate actions references update internal server error response
func (o *WeaviateActionsReferencesUpdateInternalServerError) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *WeaviateActionsReferencesUpdateInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
