/*                          _       _
 *__      _____  __ ___   ___  __ _| |_ ___
 *\ \ /\ / / _ \/ _` \ \ / / |/ _` | __/ _ \
 * \ V  V /  __/ (_| |\ V /| | (_| | ||  __/
 *  \_/\_/ \___|\__,_| \_/ |_|\__,_|\__\___|
 *
 * Copyright © 2016 - 2019 Weaviate. All rights reserved.
 * LICENSE WEAVIATE OPEN SOURCE: https://www.semi.technology/playbook/playbook/contract-weaviate-OSS.html
 * LICENSE WEAVIATE ENTERPRISE: https://www.semi.technology/playbook/contract-weaviate-enterprise.html
 * CONCEPT: Bob van Luijt (@bobvanluijt)
 * CONTACT: hello@semi.technology
 */ // Code generated by go-swagger; DO NOT EDIT.

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

// WeaviateSchemaThingsUpdateReader is a Reader for the WeaviateSchemaThingsUpdate structure.
type WeaviateSchemaThingsUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *WeaviateSchemaThingsUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewWeaviateSchemaThingsUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewWeaviateSchemaThingsUpdateUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewWeaviateSchemaThingsUpdateForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 422:
		result := NewWeaviateSchemaThingsUpdateUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewWeaviateSchemaThingsUpdateInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewWeaviateSchemaThingsUpdateOK creates a WeaviateSchemaThingsUpdateOK with default headers values
func NewWeaviateSchemaThingsUpdateOK() *WeaviateSchemaThingsUpdateOK {
	return &WeaviateSchemaThingsUpdateOK{}
}

/*WeaviateSchemaThingsUpdateOK handles this case with default header values.

Changes applied.
*/
type WeaviateSchemaThingsUpdateOK struct {
}

func (o *WeaviateSchemaThingsUpdateOK) Error() string {
	return fmt.Sprintf("[PUT /schema/things/{className}][%d] weaviateSchemaThingsUpdateOK ", 200)
}

func (o *WeaviateSchemaThingsUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewWeaviateSchemaThingsUpdateUnauthorized creates a WeaviateSchemaThingsUpdateUnauthorized with default headers values
func NewWeaviateSchemaThingsUpdateUnauthorized() *WeaviateSchemaThingsUpdateUnauthorized {
	return &WeaviateSchemaThingsUpdateUnauthorized{}
}

/*WeaviateSchemaThingsUpdateUnauthorized handles this case with default header values.

Unauthorized or invalid credentials.
*/
type WeaviateSchemaThingsUpdateUnauthorized struct {
}

func (o *WeaviateSchemaThingsUpdateUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /schema/things/{className}][%d] weaviateSchemaThingsUpdateUnauthorized ", 401)
}

func (o *WeaviateSchemaThingsUpdateUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewWeaviateSchemaThingsUpdateForbidden creates a WeaviateSchemaThingsUpdateForbidden with default headers values
func NewWeaviateSchemaThingsUpdateForbidden() *WeaviateSchemaThingsUpdateForbidden {
	return &WeaviateSchemaThingsUpdateForbidden{}
}

/*WeaviateSchemaThingsUpdateForbidden handles this case with default header values.

Forbidden
*/
type WeaviateSchemaThingsUpdateForbidden struct {
	Payload *models.ErrorResponse
}

func (o *WeaviateSchemaThingsUpdateForbidden) Error() string {
	return fmt.Sprintf("[PUT /schema/things/{className}][%d] weaviateSchemaThingsUpdateForbidden  %+v", 403, o.Payload)
}

func (o *WeaviateSchemaThingsUpdateForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewWeaviateSchemaThingsUpdateUnprocessableEntity creates a WeaviateSchemaThingsUpdateUnprocessableEntity with default headers values
func NewWeaviateSchemaThingsUpdateUnprocessableEntity() *WeaviateSchemaThingsUpdateUnprocessableEntity {
	return &WeaviateSchemaThingsUpdateUnprocessableEntity{}
}

/*WeaviateSchemaThingsUpdateUnprocessableEntity handles this case with default header values.

Invalid update.
*/
type WeaviateSchemaThingsUpdateUnprocessableEntity struct {
	Payload *models.ErrorResponse
}

func (o *WeaviateSchemaThingsUpdateUnprocessableEntity) Error() string {
	return fmt.Sprintf("[PUT /schema/things/{className}][%d] weaviateSchemaThingsUpdateUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *WeaviateSchemaThingsUpdateUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewWeaviateSchemaThingsUpdateInternalServerError creates a WeaviateSchemaThingsUpdateInternalServerError with default headers values
func NewWeaviateSchemaThingsUpdateInternalServerError() *WeaviateSchemaThingsUpdateInternalServerError {
	return &WeaviateSchemaThingsUpdateInternalServerError{}
}

/*WeaviateSchemaThingsUpdateInternalServerError handles this case with default header values.

An error has occurred while trying to fulfill the request. Most likely the ErrorResponse will contain more information about the error.
*/
type WeaviateSchemaThingsUpdateInternalServerError struct {
	Payload *models.ErrorResponse
}

func (o *WeaviateSchemaThingsUpdateInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /schema/things/{className}][%d] weaviateSchemaThingsUpdateInternalServerError  %+v", 500, o.Payload)
}

func (o *WeaviateSchemaThingsUpdateInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
