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

package actions

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/semi-technologies/weaviate/entities/models"
)

// WeaviateActionsReferencesUpdateReader is a Reader for the WeaviateActionsReferencesUpdate structure.
type WeaviateActionsReferencesUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *WeaviateActionsReferencesUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewWeaviateActionsReferencesUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewWeaviateActionsReferencesUpdateUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewWeaviateActionsReferencesUpdateForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 422:
		result := NewWeaviateActionsReferencesUpdateUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewWeaviateActionsReferencesUpdateInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewWeaviateActionsReferencesUpdateOK creates a WeaviateActionsReferencesUpdateOK with default headers values
func NewWeaviateActionsReferencesUpdateOK() *WeaviateActionsReferencesUpdateOK {
	return &WeaviateActionsReferencesUpdateOK{}
}

/*WeaviateActionsReferencesUpdateOK handles this case with default header values.

Successfully replaced all the references.
*/
type WeaviateActionsReferencesUpdateOK struct {
}

func (o *WeaviateActionsReferencesUpdateOK) Error() string {
	return fmt.Sprintf("[PUT /actions/{id}/references/{propertyName}][%d] weaviateActionsReferencesUpdateOK ", 200)
}

func (o *WeaviateActionsReferencesUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewWeaviateActionsReferencesUpdateUnauthorized creates a WeaviateActionsReferencesUpdateUnauthorized with default headers values
func NewWeaviateActionsReferencesUpdateUnauthorized() *WeaviateActionsReferencesUpdateUnauthorized {
	return &WeaviateActionsReferencesUpdateUnauthorized{}
}

/*WeaviateActionsReferencesUpdateUnauthorized handles this case with default header values.

Unauthorized or invalid credentials.
*/
type WeaviateActionsReferencesUpdateUnauthorized struct {
}

func (o *WeaviateActionsReferencesUpdateUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /actions/{id}/references/{propertyName}][%d] weaviateActionsReferencesUpdateUnauthorized ", 401)
}

func (o *WeaviateActionsReferencesUpdateUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewWeaviateActionsReferencesUpdateForbidden creates a WeaviateActionsReferencesUpdateForbidden with default headers values
func NewWeaviateActionsReferencesUpdateForbidden() *WeaviateActionsReferencesUpdateForbidden {
	return &WeaviateActionsReferencesUpdateForbidden{}
}

/*WeaviateActionsReferencesUpdateForbidden handles this case with default header values.

Forbidden
*/
type WeaviateActionsReferencesUpdateForbidden struct {
	Payload *models.ErrorResponse
}

func (o *WeaviateActionsReferencesUpdateForbidden) Error() string {
	return fmt.Sprintf("[PUT /actions/{id}/references/{propertyName}][%d] weaviateActionsReferencesUpdateForbidden  %+v", 403, o.Payload)
}

func (o *WeaviateActionsReferencesUpdateForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewWeaviateActionsReferencesUpdateUnprocessableEntity creates a WeaviateActionsReferencesUpdateUnprocessableEntity with default headers values
func NewWeaviateActionsReferencesUpdateUnprocessableEntity() *WeaviateActionsReferencesUpdateUnprocessableEntity {
	return &WeaviateActionsReferencesUpdateUnprocessableEntity{}
}

/*WeaviateActionsReferencesUpdateUnprocessableEntity handles this case with default header values.

Request body is well-formed (i.e., syntactically correct), but semantically erroneous. Are you sure the property exists or that it is a class?
*/
type WeaviateActionsReferencesUpdateUnprocessableEntity struct {
	Payload *models.ErrorResponse
}

func (o *WeaviateActionsReferencesUpdateUnprocessableEntity) Error() string {
	return fmt.Sprintf("[PUT /actions/{id}/references/{propertyName}][%d] weaviateActionsReferencesUpdateUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *WeaviateActionsReferencesUpdateUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewWeaviateActionsReferencesUpdateInternalServerError creates a WeaviateActionsReferencesUpdateInternalServerError with default headers values
func NewWeaviateActionsReferencesUpdateInternalServerError() *WeaviateActionsReferencesUpdateInternalServerError {
	return &WeaviateActionsReferencesUpdateInternalServerError{}
}

/*WeaviateActionsReferencesUpdateInternalServerError handles this case with default header values.

An error has occurred while trying to fulfill the request. Most likely the ErrorResponse will contain more information about the error.
*/
type WeaviateActionsReferencesUpdateInternalServerError struct {
	Payload *models.ErrorResponse
}

func (o *WeaviateActionsReferencesUpdateInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /actions/{id}/references/{propertyName}][%d] weaviateActionsReferencesUpdateInternalServerError  %+v", 500, o.Payload)
}

func (o *WeaviateActionsReferencesUpdateInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
