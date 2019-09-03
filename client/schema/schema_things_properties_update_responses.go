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

// SchemaThingsPropertiesUpdateReader is a Reader for the SchemaThingsPropertiesUpdate structure.
type SchemaThingsPropertiesUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SchemaThingsPropertiesUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewSchemaThingsPropertiesUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewSchemaThingsPropertiesUpdateUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewSchemaThingsPropertiesUpdateForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 422:
		result := NewSchemaThingsPropertiesUpdateUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewSchemaThingsPropertiesUpdateInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewSchemaThingsPropertiesUpdateOK creates a SchemaThingsPropertiesUpdateOK with default headers values
func NewSchemaThingsPropertiesUpdateOK() *SchemaThingsPropertiesUpdateOK {
	return &SchemaThingsPropertiesUpdateOK{}
}

/*SchemaThingsPropertiesUpdateOK handles this case with default header values.

Changes applied.
*/
type SchemaThingsPropertiesUpdateOK struct {
}

func (o *SchemaThingsPropertiesUpdateOK) Error() string {
	return fmt.Sprintf("[PUT /schema/things/{className}/properties/{propertyName}][%d] schemaThingsPropertiesUpdateOK ", 200)
}

func (o *SchemaThingsPropertiesUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewSchemaThingsPropertiesUpdateUnauthorized creates a SchemaThingsPropertiesUpdateUnauthorized with default headers values
func NewSchemaThingsPropertiesUpdateUnauthorized() *SchemaThingsPropertiesUpdateUnauthorized {
	return &SchemaThingsPropertiesUpdateUnauthorized{}
}

/*SchemaThingsPropertiesUpdateUnauthorized handles this case with default header values.

Unauthorized or invalid credentials.
*/
type SchemaThingsPropertiesUpdateUnauthorized struct {
}

func (o *SchemaThingsPropertiesUpdateUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /schema/things/{className}/properties/{propertyName}][%d] schemaThingsPropertiesUpdateUnauthorized ", 401)
}

func (o *SchemaThingsPropertiesUpdateUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewSchemaThingsPropertiesUpdateForbidden creates a SchemaThingsPropertiesUpdateForbidden with default headers values
func NewSchemaThingsPropertiesUpdateForbidden() *SchemaThingsPropertiesUpdateForbidden {
	return &SchemaThingsPropertiesUpdateForbidden{}
}

/*SchemaThingsPropertiesUpdateForbidden handles this case with default header values.

Forbidden
*/
type SchemaThingsPropertiesUpdateForbidden struct {
	Payload *models.ErrorResponse
}

func (o *SchemaThingsPropertiesUpdateForbidden) Error() string {
	return fmt.Sprintf("[PUT /schema/things/{className}/properties/{propertyName}][%d] schemaThingsPropertiesUpdateForbidden  %+v", 403, o.Payload)
}

func (o *SchemaThingsPropertiesUpdateForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSchemaThingsPropertiesUpdateUnprocessableEntity creates a SchemaThingsPropertiesUpdateUnprocessableEntity with default headers values
func NewSchemaThingsPropertiesUpdateUnprocessableEntity() *SchemaThingsPropertiesUpdateUnprocessableEntity {
	return &SchemaThingsPropertiesUpdateUnprocessableEntity{}
}

/*SchemaThingsPropertiesUpdateUnprocessableEntity handles this case with default header values.

Invalid update.
*/
type SchemaThingsPropertiesUpdateUnprocessableEntity struct {
	Payload *models.ErrorResponse
}

func (o *SchemaThingsPropertiesUpdateUnprocessableEntity) Error() string {
	return fmt.Sprintf("[PUT /schema/things/{className}/properties/{propertyName}][%d] schemaThingsPropertiesUpdateUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *SchemaThingsPropertiesUpdateUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSchemaThingsPropertiesUpdateInternalServerError creates a SchemaThingsPropertiesUpdateInternalServerError with default headers values
func NewSchemaThingsPropertiesUpdateInternalServerError() *SchemaThingsPropertiesUpdateInternalServerError {
	return &SchemaThingsPropertiesUpdateInternalServerError{}
}

/*SchemaThingsPropertiesUpdateInternalServerError handles this case with default header values.

An error has occurred while trying to fulfill the request. Most likely the ErrorResponse will contain more information about the error.
*/
type SchemaThingsPropertiesUpdateInternalServerError struct {
	Payload *models.ErrorResponse
}

func (o *SchemaThingsPropertiesUpdateInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /schema/things/{className}/properties/{propertyName}][%d] schemaThingsPropertiesUpdateInternalServerError  %+v", 500, o.Payload)
}

func (o *SchemaThingsPropertiesUpdateInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
