/*                          _       _
 *__      _____  __ ___   ___  __ _| |_ ___
 *\ \ /\ / / _ \/ _` \ \ / / |/ _` | __/ _ \
 * \ V  V /  __/ (_| |\ V /| | (_| | ||  __/
 *  \_/\_/ \___|\__,_| \_/ |_|\__,_|\__\___|
 *
 * Copyright © 2016 - 2019 Weaviate. All rights reserved.
 * LICENSE: https://github.com/creativesoftwarefdn/weaviate/blob/develop/LICENSE.md
 * DESIGN & CONCEPT: Bob van Luijt (@bobvanluijt)
 * CONTACT: hello@creativesoftwarefdn.org
 */ // Code generated by go-swagger; DO NOT EDIT.

package things

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "github.com/creativesoftwarefdn/weaviate/models"
)

// WeaviateThingsValidateOKCode is the HTTP code returned for type WeaviateThingsValidateOK
const WeaviateThingsValidateOKCode int = 200

/*WeaviateThingsValidateOK Successfully validated.

swagger:response weaviateThingsValidateOK
*/
type WeaviateThingsValidateOK struct {
}

// NewWeaviateThingsValidateOK creates WeaviateThingsValidateOK with default headers values
func NewWeaviateThingsValidateOK() *WeaviateThingsValidateOK {

	return &WeaviateThingsValidateOK{}
}

// WriteResponse to the client
func (o *WeaviateThingsValidateOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(200)
}

// WeaviateThingsValidateUnauthorizedCode is the HTTP code returned for type WeaviateThingsValidateUnauthorized
const WeaviateThingsValidateUnauthorizedCode int = 401

/*WeaviateThingsValidateUnauthorized Unauthorized or invalid credentials.

swagger:response weaviateThingsValidateUnauthorized
*/
type WeaviateThingsValidateUnauthorized struct {
}

// NewWeaviateThingsValidateUnauthorized creates WeaviateThingsValidateUnauthorized with default headers values
func NewWeaviateThingsValidateUnauthorized() *WeaviateThingsValidateUnauthorized {

	return &WeaviateThingsValidateUnauthorized{}
}

// WriteResponse to the client
func (o *WeaviateThingsValidateUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}

// WeaviateThingsValidateForbiddenCode is the HTTP code returned for type WeaviateThingsValidateForbidden
const WeaviateThingsValidateForbiddenCode int = 403

/*WeaviateThingsValidateForbidden The used API-key has insufficient permissions.

swagger:response weaviateThingsValidateForbidden
*/
type WeaviateThingsValidateForbidden struct {
}

// NewWeaviateThingsValidateForbidden creates WeaviateThingsValidateForbidden with default headers values
func NewWeaviateThingsValidateForbidden() *WeaviateThingsValidateForbidden {

	return &WeaviateThingsValidateForbidden{}
}

// WriteResponse to the client
func (o *WeaviateThingsValidateForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(403)
}

// WeaviateThingsValidateUnprocessableEntityCode is the HTTP code returned for type WeaviateThingsValidateUnprocessableEntity
const WeaviateThingsValidateUnprocessableEntityCode int = 422

/*WeaviateThingsValidateUnprocessableEntity Request body is well-formed (i.e., syntactically correct), but semantically erroneous. Are you sure the class is defined in the configuration file?

swagger:response weaviateThingsValidateUnprocessableEntity
*/
type WeaviateThingsValidateUnprocessableEntity struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewWeaviateThingsValidateUnprocessableEntity creates WeaviateThingsValidateUnprocessableEntity with default headers values
func NewWeaviateThingsValidateUnprocessableEntity() *WeaviateThingsValidateUnprocessableEntity {

	return &WeaviateThingsValidateUnprocessableEntity{}
}

// WithPayload adds the payload to the weaviate things validate unprocessable entity response
func (o *WeaviateThingsValidateUnprocessableEntity) WithPayload(payload *models.ErrorResponse) *WeaviateThingsValidateUnprocessableEntity {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the weaviate things validate unprocessable entity response
func (o *WeaviateThingsValidateUnprocessableEntity) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *WeaviateThingsValidateUnprocessableEntity) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(422)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
