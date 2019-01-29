/*                          _       _
 *__      _____  __ ___   ___  __ _| |_ ___
 *\ \ /\ / / _ \/ _` \ \ / / |/ _` | __/ _ \
 * \ V  V /  __/ (_| |\ V /| | (_| | ||  __/
 *  \_/\_/ \___|\__,_| \_/ |_|\__,_|\__\___|
 *
 * Copyright © 2016 - 2018 Weaviate. All rights reserved.
 * LICENSE: https://github.com/creativesoftwarefdn/weaviate/blob/develop/LICENSE.md
 * AUTHOR: Bob van Luijt (bob@kub.design)
 * See www.creativesoftwarefdn.org for details
 * Contact: @CreativeSofwFdn / bob@kub.design
 */
// Code generated by go-swagger; DO NOT EDIT.

package things

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "github.com/creativesoftwarefdn/weaviate/models"
)

// WeaviateThingsCreateOKCode is the HTTP code returned for type WeaviateThingsCreateOK
const WeaviateThingsCreateOKCode int = 200

/*WeaviateThingsCreateOK Thing created.

swagger:response weaviateThingsCreateOK
*/
type WeaviateThingsCreateOK struct {

	/*
	  In: Body
	*/
	Payload *models.ThingGetResponse `json:"body,omitempty"`
}

// NewWeaviateThingsCreateOK creates WeaviateThingsCreateOK with default headers values
func NewWeaviateThingsCreateOK() *WeaviateThingsCreateOK {

	return &WeaviateThingsCreateOK{}
}

// WithPayload adds the payload to the weaviate things create o k response
func (o *WeaviateThingsCreateOK) WithPayload(payload *models.ThingGetResponse) *WeaviateThingsCreateOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the weaviate things create o k response
func (o *WeaviateThingsCreateOK) SetPayload(payload *models.ThingGetResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *WeaviateThingsCreateOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// WeaviateThingsCreateAcceptedCode is the HTTP code returned for type WeaviateThingsCreateAccepted
const WeaviateThingsCreateAcceptedCode int = 202

/*WeaviateThingsCreateAccepted Successfully received.

swagger:response weaviateThingsCreateAccepted
*/
type WeaviateThingsCreateAccepted struct {

	/*
	  In: Body
	*/
	Payload *models.ThingGetResponse `json:"body,omitempty"`
}

// NewWeaviateThingsCreateAccepted creates WeaviateThingsCreateAccepted with default headers values
func NewWeaviateThingsCreateAccepted() *WeaviateThingsCreateAccepted {

	return &WeaviateThingsCreateAccepted{}
}

// WithPayload adds the payload to the weaviate things create accepted response
func (o *WeaviateThingsCreateAccepted) WithPayload(payload *models.ThingGetResponse) *WeaviateThingsCreateAccepted {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the weaviate things create accepted response
func (o *WeaviateThingsCreateAccepted) SetPayload(payload *models.ThingGetResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *WeaviateThingsCreateAccepted) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(202)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// WeaviateThingsCreateUnauthorizedCode is the HTTP code returned for type WeaviateThingsCreateUnauthorized
const WeaviateThingsCreateUnauthorizedCode int = 401

/*WeaviateThingsCreateUnauthorized Unauthorized or invalid credentials.

swagger:response weaviateThingsCreateUnauthorized
*/
type WeaviateThingsCreateUnauthorized struct {
}

// NewWeaviateThingsCreateUnauthorized creates WeaviateThingsCreateUnauthorized with default headers values
func NewWeaviateThingsCreateUnauthorized() *WeaviateThingsCreateUnauthorized {

	return &WeaviateThingsCreateUnauthorized{}
}

// WriteResponse to the client
func (o *WeaviateThingsCreateUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}

// WeaviateThingsCreateForbiddenCode is the HTTP code returned for type WeaviateThingsCreateForbidden
const WeaviateThingsCreateForbiddenCode int = 403

/*WeaviateThingsCreateForbidden The used API-key has insufficient permissions.

swagger:response weaviateThingsCreateForbidden
*/
type WeaviateThingsCreateForbidden struct {
}

// NewWeaviateThingsCreateForbidden creates WeaviateThingsCreateForbidden with default headers values
func NewWeaviateThingsCreateForbidden() *WeaviateThingsCreateForbidden {

	return &WeaviateThingsCreateForbidden{}
}

// WriteResponse to the client
func (o *WeaviateThingsCreateForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(403)
}

// WeaviateThingsCreateUnprocessableEntityCode is the HTTP code returned for type WeaviateThingsCreateUnprocessableEntity
const WeaviateThingsCreateUnprocessableEntityCode int = 422

/*WeaviateThingsCreateUnprocessableEntity Request body is well-formed (i.e., syntactically correct), but semantically erroneous. Are you sure the class is defined in the configuration file?

swagger:response weaviateThingsCreateUnprocessableEntity
*/
type WeaviateThingsCreateUnprocessableEntity struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewWeaviateThingsCreateUnprocessableEntity creates WeaviateThingsCreateUnprocessableEntity with default headers values
func NewWeaviateThingsCreateUnprocessableEntity() *WeaviateThingsCreateUnprocessableEntity {

	return &WeaviateThingsCreateUnprocessableEntity{}
}

// WithPayload adds the payload to the weaviate things create unprocessable entity response
func (o *WeaviateThingsCreateUnprocessableEntity) WithPayload(payload *models.ErrorResponse) *WeaviateThingsCreateUnprocessableEntity {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the weaviate things create unprocessable entity response
func (o *WeaviateThingsCreateUnprocessableEntity) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *WeaviateThingsCreateUnprocessableEntity) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(422)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// WeaviateThingsCreateInternalServerErrorCode is the HTTP code returned for type WeaviateThingsCreateInternalServerError
const WeaviateThingsCreateInternalServerErrorCode int = 500

/*WeaviateThingsCreateInternalServerError An error has occurred while trying to fulfill the request. Most likely the ErrorResponse will contain more information about the error.

swagger:response weaviateThingsCreateInternalServerError
*/
type WeaviateThingsCreateInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewWeaviateThingsCreateInternalServerError creates WeaviateThingsCreateInternalServerError with default headers values
func NewWeaviateThingsCreateInternalServerError() *WeaviateThingsCreateInternalServerError {

	return &WeaviateThingsCreateInternalServerError{}
}

// WithPayload adds the payload to the weaviate things create internal server error response
func (o *WeaviateThingsCreateInternalServerError) WithPayload(payload *models.ErrorResponse) *WeaviateThingsCreateInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the weaviate things create internal server error response
func (o *WeaviateThingsCreateInternalServerError) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *WeaviateThingsCreateInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
