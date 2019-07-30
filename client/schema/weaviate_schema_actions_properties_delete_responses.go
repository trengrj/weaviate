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

package schema

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/semi-technologies/weaviate/entities/models"
)

// WeaviateSchemaActionsPropertiesDeleteReader is a Reader for the WeaviateSchemaActionsPropertiesDelete structure.
type WeaviateSchemaActionsPropertiesDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *WeaviateSchemaActionsPropertiesDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewWeaviateSchemaActionsPropertiesDeleteOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewWeaviateSchemaActionsPropertiesDeleteUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewWeaviateSchemaActionsPropertiesDeleteForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewWeaviateSchemaActionsPropertiesDeleteInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewWeaviateSchemaActionsPropertiesDeleteOK creates a WeaviateSchemaActionsPropertiesDeleteOK with default headers values
func NewWeaviateSchemaActionsPropertiesDeleteOK() *WeaviateSchemaActionsPropertiesDeleteOK {
	return &WeaviateSchemaActionsPropertiesDeleteOK{}
}

/*WeaviateSchemaActionsPropertiesDeleteOK handles this case with default header values.

Removed the property from the ontology.
*/
type WeaviateSchemaActionsPropertiesDeleteOK struct {
}

func (o *WeaviateSchemaActionsPropertiesDeleteOK) Error() string {
	return fmt.Sprintf("[DELETE /schema/actions/{className}/properties/{propertyName}][%d] weaviateSchemaActionsPropertiesDeleteOK ", 200)
}

func (o *WeaviateSchemaActionsPropertiesDeleteOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewWeaviateSchemaActionsPropertiesDeleteUnauthorized creates a WeaviateSchemaActionsPropertiesDeleteUnauthorized with default headers values
func NewWeaviateSchemaActionsPropertiesDeleteUnauthorized() *WeaviateSchemaActionsPropertiesDeleteUnauthorized {
	return &WeaviateSchemaActionsPropertiesDeleteUnauthorized{}
}

/*WeaviateSchemaActionsPropertiesDeleteUnauthorized handles this case with default header values.

Unauthorized or invalid credentials.
*/
type WeaviateSchemaActionsPropertiesDeleteUnauthorized struct {
}

func (o *WeaviateSchemaActionsPropertiesDeleteUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /schema/actions/{className}/properties/{propertyName}][%d] weaviateSchemaActionsPropertiesDeleteUnauthorized ", 401)
}

func (o *WeaviateSchemaActionsPropertiesDeleteUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewWeaviateSchemaActionsPropertiesDeleteForbidden creates a WeaviateSchemaActionsPropertiesDeleteForbidden with default headers values
func NewWeaviateSchemaActionsPropertiesDeleteForbidden() *WeaviateSchemaActionsPropertiesDeleteForbidden {
	return &WeaviateSchemaActionsPropertiesDeleteForbidden{}
}

/*WeaviateSchemaActionsPropertiesDeleteForbidden handles this case with default header values.

Forbidden
*/
type WeaviateSchemaActionsPropertiesDeleteForbidden struct {
	Payload *models.ErrorResponse
}

func (o *WeaviateSchemaActionsPropertiesDeleteForbidden) Error() string {
	return fmt.Sprintf("[DELETE /schema/actions/{className}/properties/{propertyName}][%d] weaviateSchemaActionsPropertiesDeleteForbidden  %+v", 403, o.Payload)
}

func (o *WeaviateSchemaActionsPropertiesDeleteForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewWeaviateSchemaActionsPropertiesDeleteInternalServerError creates a WeaviateSchemaActionsPropertiesDeleteInternalServerError with default headers values
func NewWeaviateSchemaActionsPropertiesDeleteInternalServerError() *WeaviateSchemaActionsPropertiesDeleteInternalServerError {
	return &WeaviateSchemaActionsPropertiesDeleteInternalServerError{}
}

/*WeaviateSchemaActionsPropertiesDeleteInternalServerError handles this case with default header values.

An error has occurred while trying to fulfill the request. Most likely the ErrorResponse will contain more information about the error.
*/
type WeaviateSchemaActionsPropertiesDeleteInternalServerError struct {
	Payload *models.ErrorResponse
}

func (o *WeaviateSchemaActionsPropertiesDeleteInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /schema/actions/{className}/properties/{propertyName}][%d] weaviateSchemaActionsPropertiesDeleteInternalServerError  %+v", 500, o.Payload)
}

func (o *WeaviateSchemaActionsPropertiesDeleteInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
