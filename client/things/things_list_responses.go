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

// ThingsListReader is a Reader for the ThingsList structure.
type ThingsListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ThingsListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewThingsListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewThingsListUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewThingsListForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewThingsListNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewThingsListInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewThingsListOK creates a ThingsListOK with default headers values
func NewThingsListOK() *ThingsListOK {
	return &ThingsListOK{}
}

/*ThingsListOK handles this case with default header values.

Successful response.
*/
type ThingsListOK struct {
	Payload *models.ThingsListResponse
}

func (o *ThingsListOK) Error() string {
	return fmt.Sprintf("[GET /things][%d] thingsListOK  %+v", 200, o.Payload)
}

func (o *ThingsListOK) GetPayload() *models.ThingsListResponse {
	return o.Payload
}

func (o *ThingsListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ThingsListResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewThingsListUnauthorized creates a ThingsListUnauthorized with default headers values
func NewThingsListUnauthorized() *ThingsListUnauthorized {
	return &ThingsListUnauthorized{}
}

/*ThingsListUnauthorized handles this case with default header values.

Unauthorized or invalid credentials.
*/
type ThingsListUnauthorized struct {
}

func (o *ThingsListUnauthorized) Error() string {
	return fmt.Sprintf("[GET /things][%d] thingsListUnauthorized ", 401)
}

func (o *ThingsListUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewThingsListForbidden creates a ThingsListForbidden with default headers values
func NewThingsListForbidden() *ThingsListForbidden {
	return &ThingsListForbidden{}
}

/*ThingsListForbidden handles this case with default header values.

Forbidden
*/
type ThingsListForbidden struct {
	Payload *models.ErrorResponse
}

func (o *ThingsListForbidden) Error() string {
	return fmt.Sprintf("[GET /things][%d] thingsListForbidden  %+v", 403, o.Payload)
}

func (o *ThingsListForbidden) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *ThingsListForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewThingsListNotFound creates a ThingsListNotFound with default headers values
func NewThingsListNotFound() *ThingsListNotFound {
	return &ThingsListNotFound{}
}

/*ThingsListNotFound handles this case with default header values.

Successful query result but no resource was found.
*/
type ThingsListNotFound struct {
}

func (o *ThingsListNotFound) Error() string {
	return fmt.Sprintf("[GET /things][%d] thingsListNotFound ", 404)
}

func (o *ThingsListNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewThingsListInternalServerError creates a ThingsListInternalServerError with default headers values
func NewThingsListInternalServerError() *ThingsListInternalServerError {
	return &ThingsListInternalServerError{}
}

/*ThingsListInternalServerError handles this case with default header values.

An error has occurred while trying to fulfill the request. Most likely the ErrorResponse will contain more information about the error.
*/
type ThingsListInternalServerError struct {
	Payload *models.ErrorResponse
}

func (o *ThingsListInternalServerError) Error() string {
	return fmt.Sprintf("[GET /things][%d] thingsListInternalServerError  %+v", 500, o.Payload)
}

func (o *ThingsListInternalServerError) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *ThingsListInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
