//                           _       _
// __      _____  __ ___   ___  __ _| |_ ___
// \ \ /\ / / _ \/ _` \ \ / / |/ _` | __/ _ \
//  \ V  V /  __/ (_| |\ V /| | (_| | ||  __/
//   \_/\_/ \___|\__,_| \_/ |_|\__,_|\__\___|
// 
//  Copyright © 2016 - 2019 Weaviate. All rights reserved.
//  LICENSE WEAVIATE OPEN SOURCE: https://www.semi.technology/playbook/playbook/contract-weaviate-OSS.html
//  LICENSE WEAVIATE ENTERPRISE: https://www.semi.technology/playbook/contract-weaviate-enterprise.html
//  CONCEPT: Bob van Luijt (@bobvanluijt)
//  CONTACT: hello@semi.technology
//

// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "github.com/semi-technologies/weaviate/entities/models"
)

// WeaviateBatchingActionsCreateOKCode is the HTTP code returned for type WeaviateBatchingActionsCreateOK
const WeaviateBatchingActionsCreateOKCode int = 200

/*WeaviateBatchingActionsCreateOK Request succeeded, see response body to get detailed information about each batched item.

swagger:response weaviateBatchingActionsCreateOK
*/
type WeaviateBatchingActionsCreateOK struct {

	/*
	  In: Body
	*/
	Payload []*models.ActionsGetResponse `json:"body,omitempty"`
}

// NewWeaviateBatchingActionsCreateOK creates WeaviateBatchingActionsCreateOK with default headers values
func NewWeaviateBatchingActionsCreateOK() *WeaviateBatchingActionsCreateOK {

	return &WeaviateBatchingActionsCreateOK{}
}

// WithPayload adds the payload to the weaviate batching actions create o k response
func (o *WeaviateBatchingActionsCreateOK) WithPayload(payload []*models.ActionsGetResponse) *WeaviateBatchingActionsCreateOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the weaviate batching actions create o k response
func (o *WeaviateBatchingActionsCreateOK) SetPayload(payload []*models.ActionsGetResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *WeaviateBatchingActionsCreateOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if payload == nil {
		payload = make([]*models.ActionsGetResponse, 0, 50)
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}

}

// WeaviateBatchingActionsCreateUnauthorizedCode is the HTTP code returned for type WeaviateBatchingActionsCreateUnauthorized
const WeaviateBatchingActionsCreateUnauthorizedCode int = 401

/*WeaviateBatchingActionsCreateUnauthorized Unauthorized or invalid credentials.

swagger:response weaviateBatchingActionsCreateUnauthorized
*/
type WeaviateBatchingActionsCreateUnauthorized struct {
}

// NewWeaviateBatchingActionsCreateUnauthorized creates WeaviateBatchingActionsCreateUnauthorized with default headers values
func NewWeaviateBatchingActionsCreateUnauthorized() *WeaviateBatchingActionsCreateUnauthorized {

	return &WeaviateBatchingActionsCreateUnauthorized{}
}

// WriteResponse to the client
func (o *WeaviateBatchingActionsCreateUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}

// WeaviateBatchingActionsCreateForbiddenCode is the HTTP code returned for type WeaviateBatchingActionsCreateForbidden
const WeaviateBatchingActionsCreateForbiddenCode int = 403

/*WeaviateBatchingActionsCreateForbidden Forbidden

swagger:response weaviateBatchingActionsCreateForbidden
*/
type WeaviateBatchingActionsCreateForbidden struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewWeaviateBatchingActionsCreateForbidden creates WeaviateBatchingActionsCreateForbidden with default headers values
func NewWeaviateBatchingActionsCreateForbidden() *WeaviateBatchingActionsCreateForbidden {

	return &WeaviateBatchingActionsCreateForbidden{}
}

// WithPayload adds the payload to the weaviate batching actions create forbidden response
func (o *WeaviateBatchingActionsCreateForbidden) WithPayload(payload *models.ErrorResponse) *WeaviateBatchingActionsCreateForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the weaviate batching actions create forbidden response
func (o *WeaviateBatchingActionsCreateForbidden) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *WeaviateBatchingActionsCreateForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// WeaviateBatchingActionsCreateUnprocessableEntityCode is the HTTP code returned for type WeaviateBatchingActionsCreateUnprocessableEntity
const WeaviateBatchingActionsCreateUnprocessableEntityCode int = 422

/*WeaviateBatchingActionsCreateUnprocessableEntity Request body is well-formed (i.e., syntactically correct), but semantically erroneous. Are you sure the class is defined in the configuration file?

swagger:response weaviateBatchingActionsCreateUnprocessableEntity
*/
type WeaviateBatchingActionsCreateUnprocessableEntity struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewWeaviateBatchingActionsCreateUnprocessableEntity creates WeaviateBatchingActionsCreateUnprocessableEntity with default headers values
func NewWeaviateBatchingActionsCreateUnprocessableEntity() *WeaviateBatchingActionsCreateUnprocessableEntity {

	return &WeaviateBatchingActionsCreateUnprocessableEntity{}
}

// WithPayload adds the payload to the weaviate batching actions create unprocessable entity response
func (o *WeaviateBatchingActionsCreateUnprocessableEntity) WithPayload(payload *models.ErrorResponse) *WeaviateBatchingActionsCreateUnprocessableEntity {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the weaviate batching actions create unprocessable entity response
func (o *WeaviateBatchingActionsCreateUnprocessableEntity) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *WeaviateBatchingActionsCreateUnprocessableEntity) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(422)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// WeaviateBatchingActionsCreateInternalServerErrorCode is the HTTP code returned for type WeaviateBatchingActionsCreateInternalServerError
const WeaviateBatchingActionsCreateInternalServerErrorCode int = 500

/*WeaviateBatchingActionsCreateInternalServerError An error has occurred while trying to fulfill the request. Most likely the ErrorResponse will contain more information about the error.

swagger:response weaviateBatchingActionsCreateInternalServerError
*/
type WeaviateBatchingActionsCreateInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewWeaviateBatchingActionsCreateInternalServerError creates WeaviateBatchingActionsCreateInternalServerError with default headers values
func NewWeaviateBatchingActionsCreateInternalServerError() *WeaviateBatchingActionsCreateInternalServerError {

	return &WeaviateBatchingActionsCreateInternalServerError{}
}

// WithPayload adds the payload to the weaviate batching actions create internal server error response
func (o *WeaviateBatchingActionsCreateInternalServerError) WithPayload(payload *models.ErrorResponse) *WeaviateBatchingActionsCreateInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the weaviate batching actions create internal server error response
func (o *WeaviateBatchingActionsCreateInternalServerError) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *WeaviateBatchingActionsCreateInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
