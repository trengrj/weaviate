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

package things

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/semi-technologies/weaviate/entities/models"
)

// WeaviateThingsCreateReader is a Reader for the WeaviateThingsCreate structure.
type WeaviateThingsCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *WeaviateThingsCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewWeaviateThingsCreateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewWeaviateThingsCreateUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewWeaviateThingsCreateForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 422:
		result := NewWeaviateThingsCreateUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewWeaviateThingsCreateInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewWeaviateThingsCreateOK creates a WeaviateThingsCreateOK with default headers values
func NewWeaviateThingsCreateOK() *WeaviateThingsCreateOK {
	return &WeaviateThingsCreateOK{}
}

/*WeaviateThingsCreateOK handles this case with default header values.

Thing created.
*/
type WeaviateThingsCreateOK struct {
	Payload *models.Thing
}

func (o *WeaviateThingsCreateOK) Error() string {
	return fmt.Sprintf("[POST /things][%d] weaviateThingsCreateOK  %+v", 200, o.Payload)
}

func (o *WeaviateThingsCreateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Thing)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewWeaviateThingsCreateUnauthorized creates a WeaviateThingsCreateUnauthorized with default headers values
func NewWeaviateThingsCreateUnauthorized() *WeaviateThingsCreateUnauthorized {
	return &WeaviateThingsCreateUnauthorized{}
}

/*WeaviateThingsCreateUnauthorized handles this case with default header values.

Unauthorized or invalid credentials.
*/
type WeaviateThingsCreateUnauthorized struct {
}

func (o *WeaviateThingsCreateUnauthorized) Error() string {
	return fmt.Sprintf("[POST /things][%d] weaviateThingsCreateUnauthorized ", 401)
}

func (o *WeaviateThingsCreateUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewWeaviateThingsCreateForbidden creates a WeaviateThingsCreateForbidden with default headers values
func NewWeaviateThingsCreateForbidden() *WeaviateThingsCreateForbidden {
	return &WeaviateThingsCreateForbidden{}
}

/*WeaviateThingsCreateForbidden handles this case with default header values.

Insufficient permissions.
*/
type WeaviateThingsCreateForbidden struct {
}

func (o *WeaviateThingsCreateForbidden) Error() string {
	return fmt.Sprintf("[POST /things][%d] weaviateThingsCreateForbidden ", 403)
}

func (o *WeaviateThingsCreateForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewWeaviateThingsCreateUnprocessableEntity creates a WeaviateThingsCreateUnprocessableEntity with default headers values
func NewWeaviateThingsCreateUnprocessableEntity() *WeaviateThingsCreateUnprocessableEntity {
	return &WeaviateThingsCreateUnprocessableEntity{}
}

/*WeaviateThingsCreateUnprocessableEntity handles this case with default header values.

Request body is well-formed (i.e., syntactically correct), but semantically erroneous. Are you sure the class is defined in the configuration file?
*/
type WeaviateThingsCreateUnprocessableEntity struct {
	Payload *models.ErrorResponse
}

func (o *WeaviateThingsCreateUnprocessableEntity) Error() string {
	return fmt.Sprintf("[POST /things][%d] weaviateThingsCreateUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *WeaviateThingsCreateUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewWeaviateThingsCreateInternalServerError creates a WeaviateThingsCreateInternalServerError with default headers values
func NewWeaviateThingsCreateInternalServerError() *WeaviateThingsCreateInternalServerError {
	return &WeaviateThingsCreateInternalServerError{}
}

/*WeaviateThingsCreateInternalServerError handles this case with default header values.

An error has occurred while trying to fulfill the request. Most likely the ErrorResponse will contain more information about the error.
*/
type WeaviateThingsCreateInternalServerError struct {
	Payload *models.ErrorResponse
}

func (o *WeaviateThingsCreateInternalServerError) Error() string {
	return fmt.Sprintf("[POST /things][%d] weaviateThingsCreateInternalServerError  %+v", 500, o.Payload)
}

func (o *WeaviateThingsCreateInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
