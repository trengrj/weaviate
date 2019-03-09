/*                          _       _
 *__      _____  __ ___   ___  __ _| |_ ___
 *\ \ /\ / / _ \/ _` \ \ / / |/ _` | __/ _ \
 * \ V  V /  __/ (_| |\ V /| | (_| | ||  __/
 *  \_/\_/ \___|\__,_| \_/ |_|\__,_|\__\___|
 *
 * Copyright © 2016 - 2019 Weaviate. All rights reserved.
 * LICENSE: https://github.com/creativesoftwarefdn/weaviate/blob/develop/LICENSE.md
 * DESIGN & CONCEPT: Bob van Luijt (@bobvanluijt)
 * CONTACT: hello@creativesoftwarefdn.org
 */ // Code generated by go-swagger; DO NOT EDIT.

package actions

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/creativesoftwarefdn/weaviate/models"
)

// WeaviateActionsPropertiesDeleteReader is a Reader for the WeaviateActionsPropertiesDelete structure.
type WeaviateActionsPropertiesDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *WeaviateActionsPropertiesDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 204:
		result := NewWeaviateActionsPropertiesDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewWeaviateActionsPropertiesDeleteUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewWeaviateActionsPropertiesDeleteForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewWeaviateActionsPropertiesDeleteNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewWeaviateActionsPropertiesDeleteInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewWeaviateActionsPropertiesDeleteNoContent creates a WeaviateActionsPropertiesDeleteNoContent with default headers values
func NewWeaviateActionsPropertiesDeleteNoContent() *WeaviateActionsPropertiesDeleteNoContent {
	return &WeaviateActionsPropertiesDeleteNoContent{}
}

/*WeaviateActionsPropertiesDeleteNoContent handles this case with default header values.

Successfully deleted.
*/
type WeaviateActionsPropertiesDeleteNoContent struct {
}

func (o *WeaviateActionsPropertiesDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /actions/{actionId}/properties/{propertyName}][%d] weaviateActionsPropertiesDeleteNoContent ", 204)
}

func (o *WeaviateActionsPropertiesDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewWeaviateActionsPropertiesDeleteUnauthorized creates a WeaviateActionsPropertiesDeleteUnauthorized with default headers values
func NewWeaviateActionsPropertiesDeleteUnauthorized() *WeaviateActionsPropertiesDeleteUnauthorized {
	return &WeaviateActionsPropertiesDeleteUnauthorized{}
}

/*WeaviateActionsPropertiesDeleteUnauthorized handles this case with default header values.

Unauthorized or invalid credentials.
*/
type WeaviateActionsPropertiesDeleteUnauthorized struct {
}

func (o *WeaviateActionsPropertiesDeleteUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /actions/{actionId}/properties/{propertyName}][%d] weaviateActionsPropertiesDeleteUnauthorized ", 401)
}

func (o *WeaviateActionsPropertiesDeleteUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewWeaviateActionsPropertiesDeleteForbidden creates a WeaviateActionsPropertiesDeleteForbidden with default headers values
func NewWeaviateActionsPropertiesDeleteForbidden() *WeaviateActionsPropertiesDeleteForbidden {
	return &WeaviateActionsPropertiesDeleteForbidden{}
}

/*WeaviateActionsPropertiesDeleteForbidden handles this case with default header values.

Insufficient permissions.
*/
type WeaviateActionsPropertiesDeleteForbidden struct {
}

func (o *WeaviateActionsPropertiesDeleteForbidden) Error() string {
	return fmt.Sprintf("[DELETE /actions/{actionId}/properties/{propertyName}][%d] weaviateActionsPropertiesDeleteForbidden ", 403)
}

func (o *WeaviateActionsPropertiesDeleteForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewWeaviateActionsPropertiesDeleteNotFound creates a WeaviateActionsPropertiesDeleteNotFound with default headers values
func NewWeaviateActionsPropertiesDeleteNotFound() *WeaviateActionsPropertiesDeleteNotFound {
	return &WeaviateActionsPropertiesDeleteNotFound{}
}

/*WeaviateActionsPropertiesDeleteNotFound handles this case with default header values.

Successful query result but no resource was found.
*/
type WeaviateActionsPropertiesDeleteNotFound struct {
	Payload *models.ErrorResponse
}

func (o *WeaviateActionsPropertiesDeleteNotFound) Error() string {
	return fmt.Sprintf("[DELETE /actions/{actionId}/properties/{propertyName}][%d] weaviateActionsPropertiesDeleteNotFound  %+v", 404, o.Payload)
}

func (o *WeaviateActionsPropertiesDeleteNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewWeaviateActionsPropertiesDeleteInternalServerError creates a WeaviateActionsPropertiesDeleteInternalServerError with default headers values
func NewWeaviateActionsPropertiesDeleteInternalServerError() *WeaviateActionsPropertiesDeleteInternalServerError {
	return &WeaviateActionsPropertiesDeleteInternalServerError{}
}

/*WeaviateActionsPropertiesDeleteInternalServerError handles this case with default header values.

An error has occurred while trying to fulfill the request. Most likely the ErrorResponse will contain more information about the error.
*/
type WeaviateActionsPropertiesDeleteInternalServerError struct {
	Payload *models.ErrorResponse
}

func (o *WeaviateActionsPropertiesDeleteInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /actions/{actionId}/properties/{propertyName}][%d] weaviateActionsPropertiesDeleteInternalServerError  %+v", 500, o.Payload)
}

func (o *WeaviateActionsPropertiesDeleteInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
