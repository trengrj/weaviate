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

package batching

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/semi-technologies/weaviate/entities/models"
)

// BatchingObjectsCreateOKCode is the HTTP code returned for type BatchingObjectsCreateOK
const BatchingObjectsCreateOKCode int = 200

/*BatchingObjectsCreateOK Request succeeded, see response body to get detailed information about each batched item.

swagger:response batchingObjectsCreateOK
*/
type BatchingObjectsCreateOK struct {

	/*
	  In: Body
	*/
	Payload []*models.ObjectsGetResponse `json:"body,omitempty"`
}

// NewBatchingObjectsCreateOK creates BatchingObjectsCreateOK with default headers values
func NewBatchingObjectsCreateOK() *BatchingObjectsCreateOK {

	return &BatchingObjectsCreateOK{}
}

// WithPayload adds the payload to the batching objects create o k response
func (o *BatchingObjectsCreateOK) WithPayload(payload []*models.ObjectsGetResponse) *BatchingObjectsCreateOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the batching objects create o k response
func (o *BatchingObjectsCreateOK) SetPayload(payload []*models.ObjectsGetResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *BatchingObjectsCreateOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if payload == nil {
		// return empty array
		payload = make([]*models.ObjectsGetResponse, 0, 50)
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// BatchingObjectsCreateUnauthorizedCode is the HTTP code returned for type BatchingObjectsCreateUnauthorized
const BatchingObjectsCreateUnauthorizedCode int = 401

/*BatchingObjectsCreateUnauthorized Unauthorized or invalid credentials.

swagger:response batchingObjectsCreateUnauthorized
*/
type BatchingObjectsCreateUnauthorized struct {
}

// NewBatchingObjectsCreateUnauthorized creates BatchingObjectsCreateUnauthorized with default headers values
func NewBatchingObjectsCreateUnauthorized() *BatchingObjectsCreateUnauthorized {

	return &BatchingObjectsCreateUnauthorized{}
}

// WriteResponse to the client
func (o *BatchingObjectsCreateUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}

// BatchingObjectsCreateForbiddenCode is the HTTP code returned for type BatchingObjectsCreateForbidden
const BatchingObjectsCreateForbiddenCode int = 403

/*BatchingObjectsCreateForbidden Forbidden

swagger:response batchingObjectsCreateForbidden
*/
type BatchingObjectsCreateForbidden struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewBatchingObjectsCreateForbidden creates BatchingObjectsCreateForbidden with default headers values
func NewBatchingObjectsCreateForbidden() *BatchingObjectsCreateForbidden {

	return &BatchingObjectsCreateForbidden{}
}

// WithPayload adds the payload to the batching objects create forbidden response
func (o *BatchingObjectsCreateForbidden) WithPayload(payload *models.ErrorResponse) *BatchingObjectsCreateForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the batching objects create forbidden response
func (o *BatchingObjectsCreateForbidden) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *BatchingObjectsCreateForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// BatchingObjectsCreateUnprocessableEntityCode is the HTTP code returned for type BatchingObjectsCreateUnprocessableEntity
const BatchingObjectsCreateUnprocessableEntityCode int = 422

/*BatchingObjectsCreateUnprocessableEntity Request body is well-formed (i.e., syntactically correct), but semantically erroneous. Are you sure the class is defined in the configuration file?

swagger:response batchingObjectsCreateUnprocessableEntity
*/
type BatchingObjectsCreateUnprocessableEntity struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewBatchingObjectsCreateUnprocessableEntity creates BatchingObjectsCreateUnprocessableEntity with default headers values
func NewBatchingObjectsCreateUnprocessableEntity() *BatchingObjectsCreateUnprocessableEntity {

	return &BatchingObjectsCreateUnprocessableEntity{}
}

// WithPayload adds the payload to the batching objects create unprocessable entity response
func (o *BatchingObjectsCreateUnprocessableEntity) WithPayload(payload *models.ErrorResponse) *BatchingObjectsCreateUnprocessableEntity {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the batching objects create unprocessable entity response
func (o *BatchingObjectsCreateUnprocessableEntity) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *BatchingObjectsCreateUnprocessableEntity) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(422)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// BatchingObjectsCreateInternalServerErrorCode is the HTTP code returned for type BatchingObjectsCreateInternalServerError
const BatchingObjectsCreateInternalServerErrorCode int = 500

/*BatchingObjectsCreateInternalServerError An error has occurred while trying to fulfill the request. Most likely the ErrorResponse will contain more information about the error.

swagger:response batchingObjectsCreateInternalServerError
*/
type BatchingObjectsCreateInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewBatchingObjectsCreateInternalServerError creates BatchingObjectsCreateInternalServerError with default headers values
func NewBatchingObjectsCreateInternalServerError() *BatchingObjectsCreateInternalServerError {

	return &BatchingObjectsCreateInternalServerError{}
}

// WithPayload adds the payload to the batching objects create internal server error response
func (o *BatchingObjectsCreateInternalServerError) WithPayload(payload *models.ErrorResponse) *BatchingObjectsCreateInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the batching objects create internal server error response
func (o *BatchingObjectsCreateInternalServerError) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *BatchingObjectsCreateInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
