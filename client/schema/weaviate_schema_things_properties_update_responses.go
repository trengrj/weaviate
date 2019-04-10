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

package schema

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/creativesoftwarefdn/weaviate/models"
)

// WeaviateSchemaThingsPropertiesUpdateReader is a Reader for the WeaviateSchemaThingsPropertiesUpdate structure.
type WeaviateSchemaThingsPropertiesUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *WeaviateSchemaThingsPropertiesUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewWeaviateSchemaThingsPropertiesUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewWeaviateSchemaThingsPropertiesUpdateUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewWeaviateSchemaThingsPropertiesUpdateForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 422:
		result := NewWeaviateSchemaThingsPropertiesUpdateUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewWeaviateSchemaThingsPropertiesUpdateInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewWeaviateSchemaThingsPropertiesUpdateOK creates a WeaviateSchemaThingsPropertiesUpdateOK with default headers values
func NewWeaviateSchemaThingsPropertiesUpdateOK() *WeaviateSchemaThingsPropertiesUpdateOK {
	return &WeaviateSchemaThingsPropertiesUpdateOK{}
}

/*WeaviateSchemaThingsPropertiesUpdateOK handles this case with default header values.

Changes applied.
*/
type WeaviateSchemaThingsPropertiesUpdateOK struct {
}

func (o *WeaviateSchemaThingsPropertiesUpdateOK) Error() string {
	return fmt.Sprintf("[PUT /schema/things/{className}/properties/{propertyName}][%d] weaviateSchemaThingsPropertiesUpdateOK ", 200)
}

func (o *WeaviateSchemaThingsPropertiesUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewWeaviateSchemaThingsPropertiesUpdateUnauthorized creates a WeaviateSchemaThingsPropertiesUpdateUnauthorized with default headers values
func NewWeaviateSchemaThingsPropertiesUpdateUnauthorized() *WeaviateSchemaThingsPropertiesUpdateUnauthorized {
	return &WeaviateSchemaThingsPropertiesUpdateUnauthorized{}
}

/*WeaviateSchemaThingsPropertiesUpdateUnauthorized handles this case with default header values.

Unauthorized or invalid credentials.
*/
type WeaviateSchemaThingsPropertiesUpdateUnauthorized struct {
}

func (o *WeaviateSchemaThingsPropertiesUpdateUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /schema/things/{className}/properties/{propertyName}][%d] weaviateSchemaThingsPropertiesUpdateUnauthorized ", 401)
}

func (o *WeaviateSchemaThingsPropertiesUpdateUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewWeaviateSchemaThingsPropertiesUpdateForbidden creates a WeaviateSchemaThingsPropertiesUpdateForbidden with default headers values
func NewWeaviateSchemaThingsPropertiesUpdateForbidden() *WeaviateSchemaThingsPropertiesUpdateForbidden {
	return &WeaviateSchemaThingsPropertiesUpdateForbidden{}
}

/*WeaviateSchemaThingsPropertiesUpdateForbidden handles this case with default header values.

Could not find the Thing class or property.
*/
type WeaviateSchemaThingsPropertiesUpdateForbidden struct {
}

func (o *WeaviateSchemaThingsPropertiesUpdateForbidden) Error() string {
	return fmt.Sprintf("[PUT /schema/things/{className}/properties/{propertyName}][%d] weaviateSchemaThingsPropertiesUpdateForbidden ", 403)
}

func (o *WeaviateSchemaThingsPropertiesUpdateForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewWeaviateSchemaThingsPropertiesUpdateUnprocessableEntity creates a WeaviateSchemaThingsPropertiesUpdateUnprocessableEntity with default headers values
func NewWeaviateSchemaThingsPropertiesUpdateUnprocessableEntity() *WeaviateSchemaThingsPropertiesUpdateUnprocessableEntity {
	return &WeaviateSchemaThingsPropertiesUpdateUnprocessableEntity{}
}

/*WeaviateSchemaThingsPropertiesUpdateUnprocessableEntity handles this case with default header values.

Invalid update.
*/
type WeaviateSchemaThingsPropertiesUpdateUnprocessableEntity struct {
	Payload *models.ErrorResponse
}

func (o *WeaviateSchemaThingsPropertiesUpdateUnprocessableEntity) Error() string {
	return fmt.Sprintf("[PUT /schema/things/{className}/properties/{propertyName}][%d] weaviateSchemaThingsPropertiesUpdateUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *WeaviateSchemaThingsPropertiesUpdateUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewWeaviateSchemaThingsPropertiesUpdateInternalServerError creates a WeaviateSchemaThingsPropertiesUpdateInternalServerError with default headers values
func NewWeaviateSchemaThingsPropertiesUpdateInternalServerError() *WeaviateSchemaThingsPropertiesUpdateInternalServerError {
	return &WeaviateSchemaThingsPropertiesUpdateInternalServerError{}
}

/*WeaviateSchemaThingsPropertiesUpdateInternalServerError handles this case with default header values.

An error has occurred while trying to fulfill the request. Most likely the ErrorResponse will contain more information about the error.
*/
type WeaviateSchemaThingsPropertiesUpdateInternalServerError struct {
	Payload *models.ErrorResponse
}

func (o *WeaviateSchemaThingsPropertiesUpdateInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /schema/things/{className}/properties/{propertyName}][%d] weaviateSchemaThingsPropertiesUpdateInternalServerError  %+v", 500, o.Payload)
}

func (o *WeaviateSchemaThingsPropertiesUpdateInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
