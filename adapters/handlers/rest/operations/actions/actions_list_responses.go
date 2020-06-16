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

package actions

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/semi-technologies/weaviate/entities/models"
)

// ActionsListOKCode is the HTTP code returned for type ActionsListOK
const ActionsListOKCode int = 200

/*ActionsListOK Successful response.

swagger:response actionsListOK
*/
type ActionsListOK struct {

	/*
	  In: Body
	*/
	Payload *models.ActionsListResponse `json:"body,omitempty"`
}

// NewActionsListOK creates ActionsListOK with default headers values
func NewActionsListOK() *ActionsListOK {

	return &ActionsListOK{}
}

// WithPayload adds the payload to the actions list o k response
func (o *ActionsListOK) WithPayload(payload *models.ActionsListResponse) *ActionsListOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the actions list o k response
func (o *ActionsListOK) SetPayload(payload *models.ActionsListResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ActionsListOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ActionsListBadRequestCode is the HTTP code returned for type ActionsListBadRequest
const ActionsListBadRequestCode int = 400

/*ActionsListBadRequest Malformed request.

swagger:response actionsListBadRequest
*/
type ActionsListBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewActionsListBadRequest creates ActionsListBadRequest with default headers values
func NewActionsListBadRequest() *ActionsListBadRequest {

	return &ActionsListBadRequest{}
}

// WithPayload adds the payload to the actions list bad request response
func (o *ActionsListBadRequest) WithPayload(payload *models.ErrorResponse) *ActionsListBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the actions list bad request response
func (o *ActionsListBadRequest) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ActionsListBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ActionsListUnauthorizedCode is the HTTP code returned for type ActionsListUnauthorized
const ActionsListUnauthorizedCode int = 401

/*ActionsListUnauthorized Unauthorized or invalid credentials.

swagger:response actionsListUnauthorized
*/
type ActionsListUnauthorized struct {
}

// NewActionsListUnauthorized creates ActionsListUnauthorized with default headers values
func NewActionsListUnauthorized() *ActionsListUnauthorized {

	return &ActionsListUnauthorized{}
}

// WriteResponse to the client
func (o *ActionsListUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}

// ActionsListForbiddenCode is the HTTP code returned for type ActionsListForbidden
const ActionsListForbiddenCode int = 403

/*ActionsListForbidden Forbidden

swagger:response actionsListForbidden
*/
type ActionsListForbidden struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewActionsListForbidden creates ActionsListForbidden with default headers values
func NewActionsListForbidden() *ActionsListForbidden {

	return &ActionsListForbidden{}
}

// WithPayload adds the payload to the actions list forbidden response
func (o *ActionsListForbidden) WithPayload(payload *models.ErrorResponse) *ActionsListForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the actions list forbidden response
func (o *ActionsListForbidden) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ActionsListForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ActionsListNotFoundCode is the HTTP code returned for type ActionsListNotFound
const ActionsListNotFoundCode int = 404

/*ActionsListNotFound Successful query result but no resource was found.

swagger:response actionsListNotFound
*/
type ActionsListNotFound struct {
}

// NewActionsListNotFound creates ActionsListNotFound with default headers values
func NewActionsListNotFound() *ActionsListNotFound {

	return &ActionsListNotFound{}
}

// WriteResponse to the client
func (o *ActionsListNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(404)
}

// ActionsListInternalServerErrorCode is the HTTP code returned for type ActionsListInternalServerError
const ActionsListInternalServerErrorCode int = 500

/*ActionsListInternalServerError An error has occurred while trying to fulfill the request. Most likely the ErrorResponse will contain more information about the error.

swagger:response actionsListInternalServerError
*/
type ActionsListInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewActionsListInternalServerError creates ActionsListInternalServerError with default headers values
func NewActionsListInternalServerError() *ActionsListInternalServerError {

	return &ActionsListInternalServerError{}
}

// WithPayload adds the payload to the actions list internal server error response
func (o *ActionsListInternalServerError) WithPayload(payload *models.ErrorResponse) *ActionsListInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the actions list internal server error response
func (o *ActionsListInternalServerError) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ActionsListInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
