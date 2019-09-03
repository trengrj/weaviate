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

package schema

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "github.com/semi-technologies/weaviate/entities/models"
)

// SchemaThingsPropertiesAddOKCode is the HTTP code returned for type SchemaThingsPropertiesAddOK
const SchemaThingsPropertiesAddOKCode int = 200

/*SchemaThingsPropertiesAddOK Added the property.

swagger:response schemaThingsPropertiesAddOK
*/
type SchemaThingsPropertiesAddOK struct {

	/*
	  In: Body
	*/
	Payload *models.Property `json:"body,omitempty"`
}

// NewSchemaThingsPropertiesAddOK creates SchemaThingsPropertiesAddOK with default headers values
func NewSchemaThingsPropertiesAddOK() *SchemaThingsPropertiesAddOK {

	return &SchemaThingsPropertiesAddOK{}
}

// WithPayload adds the payload to the schema things properties add o k response
func (o *SchemaThingsPropertiesAddOK) WithPayload(payload *models.Property) *SchemaThingsPropertiesAddOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the schema things properties add o k response
func (o *SchemaThingsPropertiesAddOK) SetPayload(payload *models.Property) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *SchemaThingsPropertiesAddOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// SchemaThingsPropertiesAddUnauthorizedCode is the HTTP code returned for type SchemaThingsPropertiesAddUnauthorized
const SchemaThingsPropertiesAddUnauthorizedCode int = 401

/*SchemaThingsPropertiesAddUnauthorized Unauthorized or invalid credentials.

swagger:response schemaThingsPropertiesAddUnauthorized
*/
type SchemaThingsPropertiesAddUnauthorized struct {
}

// NewSchemaThingsPropertiesAddUnauthorized creates SchemaThingsPropertiesAddUnauthorized with default headers values
func NewSchemaThingsPropertiesAddUnauthorized() *SchemaThingsPropertiesAddUnauthorized {

	return &SchemaThingsPropertiesAddUnauthorized{}
}

// WriteResponse to the client
func (o *SchemaThingsPropertiesAddUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}

// SchemaThingsPropertiesAddForbiddenCode is the HTTP code returned for type SchemaThingsPropertiesAddForbidden
const SchemaThingsPropertiesAddForbiddenCode int = 403

/*SchemaThingsPropertiesAddForbidden Forbidden

swagger:response schemaThingsPropertiesAddForbidden
*/
type SchemaThingsPropertiesAddForbidden struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewSchemaThingsPropertiesAddForbidden creates SchemaThingsPropertiesAddForbidden with default headers values
func NewSchemaThingsPropertiesAddForbidden() *SchemaThingsPropertiesAddForbidden {

	return &SchemaThingsPropertiesAddForbidden{}
}

// WithPayload adds the payload to the schema things properties add forbidden response
func (o *SchemaThingsPropertiesAddForbidden) WithPayload(payload *models.ErrorResponse) *SchemaThingsPropertiesAddForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the schema things properties add forbidden response
func (o *SchemaThingsPropertiesAddForbidden) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *SchemaThingsPropertiesAddForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// SchemaThingsPropertiesAddUnprocessableEntityCode is the HTTP code returned for type SchemaThingsPropertiesAddUnprocessableEntity
const SchemaThingsPropertiesAddUnprocessableEntityCode int = 422

/*SchemaThingsPropertiesAddUnprocessableEntity Invalid property.

swagger:response schemaThingsPropertiesAddUnprocessableEntity
*/
type SchemaThingsPropertiesAddUnprocessableEntity struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewSchemaThingsPropertiesAddUnprocessableEntity creates SchemaThingsPropertiesAddUnprocessableEntity with default headers values
func NewSchemaThingsPropertiesAddUnprocessableEntity() *SchemaThingsPropertiesAddUnprocessableEntity {

	return &SchemaThingsPropertiesAddUnprocessableEntity{}
}

// WithPayload adds the payload to the schema things properties add unprocessable entity response
func (o *SchemaThingsPropertiesAddUnprocessableEntity) WithPayload(payload *models.ErrorResponse) *SchemaThingsPropertiesAddUnprocessableEntity {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the schema things properties add unprocessable entity response
func (o *SchemaThingsPropertiesAddUnprocessableEntity) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *SchemaThingsPropertiesAddUnprocessableEntity) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(422)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// SchemaThingsPropertiesAddInternalServerErrorCode is the HTTP code returned for type SchemaThingsPropertiesAddInternalServerError
const SchemaThingsPropertiesAddInternalServerErrorCode int = 500

/*SchemaThingsPropertiesAddInternalServerError An error has occurred while trying to fulfill the request. Most likely the ErrorResponse will contain more information about the error.

swagger:response schemaThingsPropertiesAddInternalServerError
*/
type SchemaThingsPropertiesAddInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewSchemaThingsPropertiesAddInternalServerError creates SchemaThingsPropertiesAddInternalServerError with default headers values
func NewSchemaThingsPropertiesAddInternalServerError() *SchemaThingsPropertiesAddInternalServerError {

	return &SchemaThingsPropertiesAddInternalServerError{}
}

// WithPayload adds the payload to the schema things properties add internal server error response
func (o *SchemaThingsPropertiesAddInternalServerError) WithPayload(payload *models.ErrorResponse) *SchemaThingsPropertiesAddInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the schema things properties add internal server error response
func (o *SchemaThingsPropertiesAddInternalServerError) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *SchemaThingsPropertiesAddInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
