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

package schema

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "github.com/semi-technologies/weaviate/entities/models"
)

// WeaviateSchemaActionsPropertiesAddOKCode is the HTTP code returned for type WeaviateSchemaActionsPropertiesAddOK
const WeaviateSchemaActionsPropertiesAddOKCode int = 200

/*WeaviateSchemaActionsPropertiesAddOK Added the property.

swagger:response weaviateSchemaActionsPropertiesAddOK
*/
type WeaviateSchemaActionsPropertiesAddOK struct {

	/*
	  In: Body
	*/
	Payload *models.Property `json:"body,omitempty"`
}

// NewWeaviateSchemaActionsPropertiesAddOK creates WeaviateSchemaActionsPropertiesAddOK with default headers values
func NewWeaviateSchemaActionsPropertiesAddOK() *WeaviateSchemaActionsPropertiesAddOK {

	return &WeaviateSchemaActionsPropertiesAddOK{}
}

// WithPayload adds the payload to the weaviate schema actions properties add o k response
func (o *WeaviateSchemaActionsPropertiesAddOK) WithPayload(payload *models.Property) *WeaviateSchemaActionsPropertiesAddOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the weaviate schema actions properties add o k response
func (o *WeaviateSchemaActionsPropertiesAddOK) SetPayload(payload *models.Property) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *WeaviateSchemaActionsPropertiesAddOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// WeaviateSchemaActionsPropertiesAddUnauthorizedCode is the HTTP code returned for type WeaviateSchemaActionsPropertiesAddUnauthorized
const WeaviateSchemaActionsPropertiesAddUnauthorizedCode int = 401

/*WeaviateSchemaActionsPropertiesAddUnauthorized Unauthorized or invalid credentials.

swagger:response weaviateSchemaActionsPropertiesAddUnauthorized
*/
type WeaviateSchemaActionsPropertiesAddUnauthorized struct {
}

// NewWeaviateSchemaActionsPropertiesAddUnauthorized creates WeaviateSchemaActionsPropertiesAddUnauthorized with default headers values
func NewWeaviateSchemaActionsPropertiesAddUnauthorized() *WeaviateSchemaActionsPropertiesAddUnauthorized {

	return &WeaviateSchemaActionsPropertiesAddUnauthorized{}
}

// WriteResponse to the client
func (o *WeaviateSchemaActionsPropertiesAddUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}

// WeaviateSchemaActionsPropertiesAddForbiddenCode is the HTTP code returned for type WeaviateSchemaActionsPropertiesAddForbidden
const WeaviateSchemaActionsPropertiesAddForbiddenCode int = 403

/*WeaviateSchemaActionsPropertiesAddForbidden Forbidden

swagger:response weaviateSchemaActionsPropertiesAddForbidden
*/
type WeaviateSchemaActionsPropertiesAddForbidden struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewWeaviateSchemaActionsPropertiesAddForbidden creates WeaviateSchemaActionsPropertiesAddForbidden with default headers values
func NewWeaviateSchemaActionsPropertiesAddForbidden() *WeaviateSchemaActionsPropertiesAddForbidden {

	return &WeaviateSchemaActionsPropertiesAddForbidden{}
}

// WithPayload adds the payload to the weaviate schema actions properties add forbidden response
func (o *WeaviateSchemaActionsPropertiesAddForbidden) WithPayload(payload *models.ErrorResponse) *WeaviateSchemaActionsPropertiesAddForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the weaviate schema actions properties add forbidden response
func (o *WeaviateSchemaActionsPropertiesAddForbidden) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *WeaviateSchemaActionsPropertiesAddForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// WeaviateSchemaActionsPropertiesAddUnprocessableEntityCode is the HTTP code returned for type WeaviateSchemaActionsPropertiesAddUnprocessableEntity
const WeaviateSchemaActionsPropertiesAddUnprocessableEntityCode int = 422

/*WeaviateSchemaActionsPropertiesAddUnprocessableEntity Invalid property.

swagger:response weaviateSchemaActionsPropertiesAddUnprocessableEntity
*/
type WeaviateSchemaActionsPropertiesAddUnprocessableEntity struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewWeaviateSchemaActionsPropertiesAddUnprocessableEntity creates WeaviateSchemaActionsPropertiesAddUnprocessableEntity with default headers values
func NewWeaviateSchemaActionsPropertiesAddUnprocessableEntity() *WeaviateSchemaActionsPropertiesAddUnprocessableEntity {

	return &WeaviateSchemaActionsPropertiesAddUnprocessableEntity{}
}

// WithPayload adds the payload to the weaviate schema actions properties add unprocessable entity response
func (o *WeaviateSchemaActionsPropertiesAddUnprocessableEntity) WithPayload(payload *models.ErrorResponse) *WeaviateSchemaActionsPropertiesAddUnprocessableEntity {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the weaviate schema actions properties add unprocessable entity response
func (o *WeaviateSchemaActionsPropertiesAddUnprocessableEntity) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *WeaviateSchemaActionsPropertiesAddUnprocessableEntity) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(422)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// WeaviateSchemaActionsPropertiesAddInternalServerErrorCode is the HTTP code returned for type WeaviateSchemaActionsPropertiesAddInternalServerError
const WeaviateSchemaActionsPropertiesAddInternalServerErrorCode int = 500

/*WeaviateSchemaActionsPropertiesAddInternalServerError An error has occurred while trying to fulfill the request. Most likely the ErrorResponse will contain more information about the error.

swagger:response weaviateSchemaActionsPropertiesAddInternalServerError
*/
type WeaviateSchemaActionsPropertiesAddInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewWeaviateSchemaActionsPropertiesAddInternalServerError creates WeaviateSchemaActionsPropertiesAddInternalServerError with default headers values
func NewWeaviateSchemaActionsPropertiesAddInternalServerError() *WeaviateSchemaActionsPropertiesAddInternalServerError {

	return &WeaviateSchemaActionsPropertiesAddInternalServerError{}
}

// WithPayload adds the payload to the weaviate schema actions properties add internal server error response
func (o *WeaviateSchemaActionsPropertiesAddInternalServerError) WithPayload(payload *models.ErrorResponse) *WeaviateSchemaActionsPropertiesAddInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the weaviate schema actions properties add internal server error response
func (o *WeaviateSchemaActionsPropertiesAddInternalServerError) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *WeaviateSchemaActionsPropertiesAddInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
