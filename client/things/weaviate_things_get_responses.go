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

// WeaviateThingsGetReader is a Reader for the WeaviateThingsGet structure.
type WeaviateThingsGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *WeaviateThingsGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewWeaviateThingsGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewWeaviateThingsGetUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewWeaviateThingsGetForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewWeaviateThingsGetNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewWeaviateThingsGetInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewWeaviateThingsGetOK creates a WeaviateThingsGetOK with default headers values
func NewWeaviateThingsGetOK() *WeaviateThingsGetOK {
	return &WeaviateThingsGetOK{}
}

/*WeaviateThingsGetOK handles this case with default header values.

Successful response.
*/
type WeaviateThingsGetOK struct {
	Payload *models.Thing
}

func (o *WeaviateThingsGetOK) Error() string {
	return fmt.Sprintf("[GET /things/{id}][%d] weaviateThingsGetOK  %+v", 200, o.Payload)
}

func (o *WeaviateThingsGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Thing)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewWeaviateThingsGetUnauthorized creates a WeaviateThingsGetUnauthorized with default headers values
func NewWeaviateThingsGetUnauthorized() *WeaviateThingsGetUnauthorized {
	return &WeaviateThingsGetUnauthorized{}
}

/*WeaviateThingsGetUnauthorized handles this case with default header values.

Unauthorized or invalid credentials.
*/
type WeaviateThingsGetUnauthorized struct {
}

func (o *WeaviateThingsGetUnauthorized) Error() string {
	return fmt.Sprintf("[GET /things/{id}][%d] weaviateThingsGetUnauthorized ", 401)
}

func (o *WeaviateThingsGetUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewWeaviateThingsGetForbidden creates a WeaviateThingsGetForbidden with default headers values
func NewWeaviateThingsGetForbidden() *WeaviateThingsGetForbidden {
	return &WeaviateThingsGetForbidden{}
}

/*WeaviateThingsGetForbidden handles this case with default header values.

Insufficient permissions.
*/
type WeaviateThingsGetForbidden struct {
}

func (o *WeaviateThingsGetForbidden) Error() string {
	return fmt.Sprintf("[GET /things/{id}][%d] weaviateThingsGetForbidden ", 403)
}

func (o *WeaviateThingsGetForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewWeaviateThingsGetNotFound creates a WeaviateThingsGetNotFound with default headers values
func NewWeaviateThingsGetNotFound() *WeaviateThingsGetNotFound {
	return &WeaviateThingsGetNotFound{}
}

/*WeaviateThingsGetNotFound handles this case with default header values.

Successful query result but no resource was found.
*/
type WeaviateThingsGetNotFound struct {
}

func (o *WeaviateThingsGetNotFound) Error() string {
	return fmt.Sprintf("[GET /things/{id}][%d] weaviateThingsGetNotFound ", 404)
}

func (o *WeaviateThingsGetNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewWeaviateThingsGetInternalServerError creates a WeaviateThingsGetInternalServerError with default headers values
func NewWeaviateThingsGetInternalServerError() *WeaviateThingsGetInternalServerError {
	return &WeaviateThingsGetInternalServerError{}
}

/*WeaviateThingsGetInternalServerError handles this case with default header values.

An error has occurred while trying to fulfill the request. Most likely the ErrorResponse will contain more information about the error.
*/
type WeaviateThingsGetInternalServerError struct {
	Payload *models.ErrorResponse
}

func (o *WeaviateThingsGetInternalServerError) Error() string {
	return fmt.Sprintf("[GET /things/{id}][%d] weaviateThingsGetInternalServerError  %+v", 500, o.Payload)
}

func (o *WeaviateThingsGetInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
