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

// ThingsPatchReader is a Reader for the ThingsPatch structure.
type ThingsPatchReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ThingsPatchReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewThingsPatchOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewThingsPatchBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewThingsPatchUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewThingsPatchForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewThingsPatchNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 422:
		result := NewThingsPatchUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewThingsPatchInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewThingsPatchOK creates a ThingsPatchOK with default headers values
func NewThingsPatchOK() *ThingsPatchOK {
	return &ThingsPatchOK{}
}

/*ThingsPatchOK handles this case with default header values.

Successfully applied.
*/
type ThingsPatchOK struct {
	Payload *models.Thing
}

func (o *ThingsPatchOK) Error() string {
	return fmt.Sprintf("[PATCH /things/{id}][%d] thingsPatchOK  %+v", 200, o.Payload)
}

func (o *ThingsPatchOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Thing)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewThingsPatchBadRequest creates a ThingsPatchBadRequest with default headers values
func NewThingsPatchBadRequest() *ThingsPatchBadRequest {
	return &ThingsPatchBadRequest{}
}

/*ThingsPatchBadRequest handles this case with default header values.

The patch-JSON is malformed.
*/
type ThingsPatchBadRequest struct {
}

func (o *ThingsPatchBadRequest) Error() string {
	return fmt.Sprintf("[PATCH /things/{id}][%d] thingsPatchBadRequest ", 400)
}

func (o *ThingsPatchBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewThingsPatchUnauthorized creates a ThingsPatchUnauthorized with default headers values
func NewThingsPatchUnauthorized() *ThingsPatchUnauthorized {
	return &ThingsPatchUnauthorized{}
}

/*ThingsPatchUnauthorized handles this case with default header values.

Unauthorized or invalid credentials.
*/
type ThingsPatchUnauthorized struct {
}

func (o *ThingsPatchUnauthorized) Error() string {
	return fmt.Sprintf("[PATCH /things/{id}][%d] thingsPatchUnauthorized ", 401)
}

func (o *ThingsPatchUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewThingsPatchForbidden creates a ThingsPatchForbidden with default headers values
func NewThingsPatchForbidden() *ThingsPatchForbidden {
	return &ThingsPatchForbidden{}
}

/*ThingsPatchForbidden handles this case with default header values.

Forbidden
*/
type ThingsPatchForbidden struct {
	Payload *models.ErrorResponse
}

func (o *ThingsPatchForbidden) Error() string {
	return fmt.Sprintf("[PATCH /things/{id}][%d] thingsPatchForbidden  %+v", 403, o.Payload)
}

func (o *ThingsPatchForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewThingsPatchNotFound creates a ThingsPatchNotFound with default headers values
func NewThingsPatchNotFound() *ThingsPatchNotFound {
	return &ThingsPatchNotFound{}
}

/*ThingsPatchNotFound handles this case with default header values.

Successful query result but no resource was found.
*/
type ThingsPatchNotFound struct {
}

func (o *ThingsPatchNotFound) Error() string {
	return fmt.Sprintf("[PATCH /things/{id}][%d] thingsPatchNotFound ", 404)
}

func (o *ThingsPatchNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewThingsPatchUnprocessableEntity creates a ThingsPatchUnprocessableEntity with default headers values
func NewThingsPatchUnprocessableEntity() *ThingsPatchUnprocessableEntity {
	return &ThingsPatchUnprocessableEntity{}
}

/*ThingsPatchUnprocessableEntity handles this case with default header values.

The patch-JSON is valid but unprocessable.
*/
type ThingsPatchUnprocessableEntity struct {
	Payload *models.ErrorResponse
}

func (o *ThingsPatchUnprocessableEntity) Error() string {
	return fmt.Sprintf("[PATCH /things/{id}][%d] thingsPatchUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *ThingsPatchUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewThingsPatchInternalServerError creates a ThingsPatchInternalServerError with default headers values
func NewThingsPatchInternalServerError() *ThingsPatchInternalServerError {
	return &ThingsPatchInternalServerError{}
}

/*ThingsPatchInternalServerError handles this case with default header values.

An error has occurred while trying to fulfill the request. Most likely the ErrorResponse will contain more information about the error.
*/
type ThingsPatchInternalServerError struct {
	Payload *models.ErrorResponse
}

func (o *ThingsPatchInternalServerError) Error() string {
	return fmt.Sprintf("[PATCH /things/{id}][%d] thingsPatchInternalServerError  %+v", 500, o.Payload)
}

func (o *ThingsPatchInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
