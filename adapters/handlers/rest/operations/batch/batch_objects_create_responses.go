//                           _       _
// __      _____  __ ___   ___  __ _| |_ ___
// \ \ /\ / / _ \/ _` \ \ / / |/ _` | __/ _ \
//  \ V  V /  __/ (_| |\ V /| | (_| | ||  __/
//   \_/\_/ \___|\__,_| \_/ |_|\__,_|\__\___|
//
//  Copyright © 2016 - 2022 SeMI Technologies B.V. All rights reserved.
//
//  CONTACT: hello@semi.technology
//

// Code generated by go-swagger; DO NOT EDIT.

package batch

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/semi-technologies/weaviate/entities/models"
)

// BatchObjectsCreateOKCode is the HTTP code returned for type BatchObjectsCreateOK
const BatchObjectsCreateOKCode int = 200

/*BatchObjectsCreateOK Request succeeded, see response body to get detailed information about each batched item.

swagger:response batchObjectsCreateOK
*/
type BatchObjectsCreateOK struct {

	/*
	  In: Body
	*/
	Payload []*models.ObjectsGetResponse `json:"body,omitempty"`
}

// NewBatchObjectsCreateOK creates BatchObjectsCreateOK with default headers values
func NewBatchObjectsCreateOK() *BatchObjectsCreateOK {

	return &BatchObjectsCreateOK{}
}

// WithPayload adds the payload to the batch objects create o k response
func (o *BatchObjectsCreateOK) WithPayload(payload []*models.ObjectsGetResponse) *BatchObjectsCreateOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the batch objects create o k response
func (o *BatchObjectsCreateOK) SetPayload(payload []*models.ObjectsGetResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *BatchObjectsCreateOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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

// BatchObjectsCreateUnauthorizedCode is the HTTP code returned for type BatchObjectsCreateUnauthorized
const BatchObjectsCreateUnauthorizedCode int = 401

/*BatchObjectsCreateUnauthorized Unauthorized or invalid credentials.

swagger:response batchObjectsCreateUnauthorized
*/
type BatchObjectsCreateUnauthorized struct {
}

// NewBatchObjectsCreateUnauthorized creates BatchObjectsCreateUnauthorized with default headers values
func NewBatchObjectsCreateUnauthorized() *BatchObjectsCreateUnauthorized {

	return &BatchObjectsCreateUnauthorized{}
}

// WriteResponse to the client
func (o *BatchObjectsCreateUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}

// BatchObjectsCreateForbiddenCode is the HTTP code returned for type BatchObjectsCreateForbidden
const BatchObjectsCreateForbiddenCode int = 403

/*BatchObjectsCreateForbidden Forbidden

swagger:response batchObjectsCreateForbidden
*/
type BatchObjectsCreateForbidden struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewBatchObjectsCreateForbidden creates BatchObjectsCreateForbidden with default headers values
func NewBatchObjectsCreateForbidden() *BatchObjectsCreateForbidden {

	return &BatchObjectsCreateForbidden{}
}

// WithPayload adds the payload to the batch objects create forbidden response
func (o *BatchObjectsCreateForbidden) WithPayload(payload *models.ErrorResponse) *BatchObjectsCreateForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the batch objects create forbidden response
func (o *BatchObjectsCreateForbidden) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *BatchObjectsCreateForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// BatchObjectsCreateUnprocessableEntityCode is the HTTP code returned for type BatchObjectsCreateUnprocessableEntity
const BatchObjectsCreateUnprocessableEntityCode int = 422

/*BatchObjectsCreateUnprocessableEntity Request body is well-formed (i.e., syntactically correct), but semantically erroneous. Are you sure the class is defined in the configuration file?

swagger:response batchObjectsCreateUnprocessableEntity
*/
type BatchObjectsCreateUnprocessableEntity struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewBatchObjectsCreateUnprocessableEntity creates BatchObjectsCreateUnprocessableEntity with default headers values
func NewBatchObjectsCreateUnprocessableEntity() *BatchObjectsCreateUnprocessableEntity {

	return &BatchObjectsCreateUnprocessableEntity{}
}

// WithPayload adds the payload to the batch objects create unprocessable entity response
func (o *BatchObjectsCreateUnprocessableEntity) WithPayload(payload *models.ErrorResponse) *BatchObjectsCreateUnprocessableEntity {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the batch objects create unprocessable entity response
func (o *BatchObjectsCreateUnprocessableEntity) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *BatchObjectsCreateUnprocessableEntity) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(422)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// BatchObjectsCreateInternalServerErrorCode is the HTTP code returned for type BatchObjectsCreateInternalServerError
const BatchObjectsCreateInternalServerErrorCode int = 500

/*BatchObjectsCreateInternalServerError An error has occurred while trying to fulfill the request. Most likely the ErrorResponse will contain more information about the error.

swagger:response batchObjectsCreateInternalServerError
*/
type BatchObjectsCreateInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewBatchObjectsCreateInternalServerError creates BatchObjectsCreateInternalServerError with default headers values
func NewBatchObjectsCreateInternalServerError() *BatchObjectsCreateInternalServerError {

	return &BatchObjectsCreateInternalServerError{}
}

// WithPayload adds the payload to the batch objects create internal server error response
func (o *BatchObjectsCreateInternalServerError) WithPayload(payload *models.ErrorResponse) *BatchObjectsCreateInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the batch objects create internal server error response
func (o *BatchObjectsCreateInternalServerError) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *BatchObjectsCreateInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
