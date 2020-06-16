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

package contextionary_api

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/semi-technologies/weaviate/entities/models"
)

// C11yExtensionsReader is a Reader for the C11yExtensions structure.
type C11yExtensionsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *C11yExtensionsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewC11yExtensionsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewC11yExtensionsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewC11yExtensionsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewC11yExtensionsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewC11yExtensionsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 501:
		result := NewC11yExtensionsNotImplemented()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewC11yExtensionsOK creates a C11yExtensionsOK with default headers values
func NewC11yExtensionsOK() *C11yExtensionsOK {
	return &C11yExtensionsOK{}
}

/*C11yExtensionsOK handles this case with default header values.

Successfully extended the contextionary with the custom cocnept
*/
type C11yExtensionsOK struct {
	Payload *models.C11yExtension
}

func (o *C11yExtensionsOK) Error() string {
	return fmt.Sprintf("[POST /c11y/extensions/][%d] c11yExtensionsOK  %+v", 200, o.Payload)
}

func (o *C11yExtensionsOK) GetPayload() *models.C11yExtension {
	return o.Payload
}

func (o *C11yExtensionsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.C11yExtension)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewC11yExtensionsBadRequest creates a C11yExtensionsBadRequest with default headers values
func NewC11yExtensionsBadRequest() *C11yExtensionsBadRequest {
	return &C11yExtensionsBadRequest{}
}

/*C11yExtensionsBadRequest handles this case with default header values.

Incorrect request
*/
type C11yExtensionsBadRequest struct {
	Payload *models.ErrorResponse
}

func (o *C11yExtensionsBadRequest) Error() string {
	return fmt.Sprintf("[POST /c11y/extensions/][%d] c11yExtensionsBadRequest  %+v", 400, o.Payload)
}

func (o *C11yExtensionsBadRequest) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *C11yExtensionsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewC11yExtensionsUnauthorized creates a C11yExtensionsUnauthorized with default headers values
func NewC11yExtensionsUnauthorized() *C11yExtensionsUnauthorized {
	return &C11yExtensionsUnauthorized{}
}

/*C11yExtensionsUnauthorized handles this case with default header values.

Unauthorized or invalid credentials.
*/
type C11yExtensionsUnauthorized struct {
}

func (o *C11yExtensionsUnauthorized) Error() string {
	return fmt.Sprintf("[POST /c11y/extensions/][%d] c11yExtensionsUnauthorized ", 401)
}

func (o *C11yExtensionsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewC11yExtensionsForbidden creates a C11yExtensionsForbidden with default headers values
func NewC11yExtensionsForbidden() *C11yExtensionsForbidden {
	return &C11yExtensionsForbidden{}
}

/*C11yExtensionsForbidden handles this case with default header values.

Forbidden
*/
type C11yExtensionsForbidden struct {
	Payload *models.ErrorResponse
}

func (o *C11yExtensionsForbidden) Error() string {
	return fmt.Sprintf("[POST /c11y/extensions/][%d] c11yExtensionsForbidden  %+v", 403, o.Payload)
}

func (o *C11yExtensionsForbidden) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *C11yExtensionsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewC11yExtensionsInternalServerError creates a C11yExtensionsInternalServerError with default headers values
func NewC11yExtensionsInternalServerError() *C11yExtensionsInternalServerError {
	return &C11yExtensionsInternalServerError{}
}

/*C11yExtensionsInternalServerError handles this case with default header values.

An error has occurred while trying to fulfill the request. Most likely the ErrorResponse will contain more information about the error.
*/
type C11yExtensionsInternalServerError struct {
	Payload *models.ErrorResponse
}

func (o *C11yExtensionsInternalServerError) Error() string {
	return fmt.Sprintf("[POST /c11y/extensions/][%d] c11yExtensionsInternalServerError  %+v", 500, o.Payload)
}

func (o *C11yExtensionsInternalServerError) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *C11yExtensionsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewC11yExtensionsNotImplemented creates a C11yExtensionsNotImplemented with default headers values
func NewC11yExtensionsNotImplemented() *C11yExtensionsNotImplemented {
	return &C11yExtensionsNotImplemented{}
}

/*C11yExtensionsNotImplemented handles this case with default header values.

Not (yet) implemented.
*/
type C11yExtensionsNotImplemented struct {
}

func (o *C11yExtensionsNotImplemented) Error() string {
	return fmt.Sprintf("[POST /c11y/extensions/][%d] c11yExtensionsNotImplemented ", 501)
}

func (o *C11yExtensionsNotImplemented) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
