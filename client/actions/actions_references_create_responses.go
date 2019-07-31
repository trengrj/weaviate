//                           _       _
// __      _____  __ ___   ___  __ _| |_ ___
// \ \ /\ / / _ \/ _` \ \ / / |/ _` | __/ _ \
//  \ V  V /  __/ (_| |\ V /| | (_| | ||  __/
//   \_/\_/ \___|\__,_| \_/ |_|\__,_|\__\___|
//
//  Copyright © 2016 - 2019 Weaviate. All rights reserved.
//  LICENSE: https://github.com/semi-technologies/weaviate/blob/develop/LICENSE.md
//  DESIGN & CONCEPT: Bob van Luijt (@bobvanluijt)
//  CONTACT: hello@semi.technology
//

// Code generated by go-swagger; DO NOT EDIT.

package actions

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/semi-technologies/weaviate/entities/models"
)

// ActionsReferencesCreateReader is a Reader for the ActionsReferencesCreate structure.
type ActionsReferencesCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ActionsReferencesCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewActionsReferencesCreateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewActionsReferencesCreateUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewActionsReferencesCreateForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 422:
		result := NewActionsReferencesCreateUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewActionsReferencesCreateInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewActionsReferencesCreateOK creates a ActionsReferencesCreateOK with default headers values
func NewActionsReferencesCreateOK() *ActionsReferencesCreateOK {
	return &ActionsReferencesCreateOK{}
}

/*ActionsReferencesCreateOK handles this case with default header values.

Successfully added the reference.
*/
type ActionsReferencesCreateOK struct {
}

func (o *ActionsReferencesCreateOK) Error() string {
	return fmt.Sprintf("[POST /actions/{id}/references/{propertyName}][%d] actionsReferencesCreateOK ", 200)
}

func (o *ActionsReferencesCreateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewActionsReferencesCreateUnauthorized creates a ActionsReferencesCreateUnauthorized with default headers values
func NewActionsReferencesCreateUnauthorized() *ActionsReferencesCreateUnauthorized {
	return &ActionsReferencesCreateUnauthorized{}
}

/*ActionsReferencesCreateUnauthorized handles this case with default header values.

Unauthorized or invalid credentials.
*/
type ActionsReferencesCreateUnauthorized struct {
}

func (o *ActionsReferencesCreateUnauthorized) Error() string {
	return fmt.Sprintf("[POST /actions/{id}/references/{propertyName}][%d] actionsReferencesCreateUnauthorized ", 401)
}

func (o *ActionsReferencesCreateUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewActionsReferencesCreateForbidden creates a ActionsReferencesCreateForbidden with default headers values
func NewActionsReferencesCreateForbidden() *ActionsReferencesCreateForbidden {
	return &ActionsReferencesCreateForbidden{}
}

/*ActionsReferencesCreateForbidden handles this case with default header values.

Forbidden
*/
type ActionsReferencesCreateForbidden struct {
	Payload *models.ErrorResponse
}

func (o *ActionsReferencesCreateForbidden) Error() string {
	return fmt.Sprintf("[POST /actions/{id}/references/{propertyName}][%d] actionsReferencesCreateForbidden  %+v", 403, o.Payload)
}

func (o *ActionsReferencesCreateForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewActionsReferencesCreateUnprocessableEntity creates a ActionsReferencesCreateUnprocessableEntity with default headers values
func NewActionsReferencesCreateUnprocessableEntity() *ActionsReferencesCreateUnprocessableEntity {
	return &ActionsReferencesCreateUnprocessableEntity{}
}

/*ActionsReferencesCreateUnprocessableEntity handles this case with default header values.

Request body is well-formed (i.e., syntactically correct), but semantically erroneous. Are you sure the property exists or that it is a class?
*/
type ActionsReferencesCreateUnprocessableEntity struct {
	Payload *models.ErrorResponse
}

func (o *ActionsReferencesCreateUnprocessableEntity) Error() string {
	return fmt.Sprintf("[POST /actions/{id}/references/{propertyName}][%d] actionsReferencesCreateUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *ActionsReferencesCreateUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewActionsReferencesCreateInternalServerError creates a ActionsReferencesCreateInternalServerError with default headers values
func NewActionsReferencesCreateInternalServerError() *ActionsReferencesCreateInternalServerError {
	return &ActionsReferencesCreateInternalServerError{}
}

/*ActionsReferencesCreateInternalServerError handles this case with default header values.

An error has occurred while trying to fulfill the request. Most likely the ErrorResponse will contain more information about the error.
*/
type ActionsReferencesCreateInternalServerError struct {
	Payload *models.ErrorResponse
}

func (o *ActionsReferencesCreateInternalServerError) Error() string {
	return fmt.Sprintf("[POST /actions/{id}/references/{propertyName}][%d] actionsReferencesCreateInternalServerError  %+v", 500, o.Payload)
}

func (o *ActionsReferencesCreateInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
