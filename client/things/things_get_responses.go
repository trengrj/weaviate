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

// ThingsGetReader is a Reader for the ThingsGet structure.
type ThingsGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ThingsGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewThingsGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewThingsGetUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewThingsGetForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewThingsGetNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewThingsGetInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewThingsGetOK creates a ThingsGetOK with default headers values
func NewThingsGetOK() *ThingsGetOK {
	return &ThingsGetOK{}
}

/*ThingsGetOK handles this case with default header values.

Successful response.
*/
type ThingsGetOK struct {
	Payload *models.Thing
}

func (o *ThingsGetOK) Error() string {
	return fmt.Sprintf("[GET /things/{id}][%d] thingsGetOK  %+v", 200, o.Payload)
}

func (o *ThingsGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Thing)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewThingsGetUnauthorized creates a ThingsGetUnauthorized with default headers values
func NewThingsGetUnauthorized() *ThingsGetUnauthorized {
	return &ThingsGetUnauthorized{}
}

/*ThingsGetUnauthorized handles this case with default header values.

Unauthorized or invalid credentials.
*/
type ThingsGetUnauthorized struct {
}

func (o *ThingsGetUnauthorized) Error() string {
	return fmt.Sprintf("[GET /things/{id}][%d] thingsGetUnauthorized ", 401)
}

func (o *ThingsGetUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewThingsGetForbidden creates a ThingsGetForbidden with default headers values
func NewThingsGetForbidden() *ThingsGetForbidden {
	return &ThingsGetForbidden{}
}

/*ThingsGetForbidden handles this case with default header values.

Forbidden
*/
type ThingsGetForbidden struct {
	Payload *models.ErrorResponse
}

func (o *ThingsGetForbidden) Error() string {
	return fmt.Sprintf("[GET /things/{id}][%d] thingsGetForbidden  %+v", 403, o.Payload)
}

func (o *ThingsGetForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewThingsGetNotFound creates a ThingsGetNotFound with default headers values
func NewThingsGetNotFound() *ThingsGetNotFound {
	return &ThingsGetNotFound{}
}

/*ThingsGetNotFound handles this case with default header values.

Successful query result but no resource was found.
*/
type ThingsGetNotFound struct {
}

func (o *ThingsGetNotFound) Error() string {
	return fmt.Sprintf("[GET /things/{id}][%d] thingsGetNotFound ", 404)
}

func (o *ThingsGetNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewThingsGetInternalServerError creates a ThingsGetInternalServerError with default headers values
func NewThingsGetInternalServerError() *ThingsGetInternalServerError {
	return &ThingsGetInternalServerError{}
}

/*ThingsGetInternalServerError handles this case with default header values.

An error has occurred while trying to fulfill the request. Most likely the ErrorResponse will contain more information about the error.
*/
type ThingsGetInternalServerError struct {
	Payload *models.ErrorResponse
}

func (o *ThingsGetInternalServerError) Error() string {
	return fmt.Sprintf("[GET /things/{id}][%d] thingsGetInternalServerError  %+v", 500, o.Payload)
}

func (o *ThingsGetInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
