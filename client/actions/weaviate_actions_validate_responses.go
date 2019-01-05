/*                          _       _
 *__      _____  __ ___   ___  __ _| |_ ___
 *\ \ /\ / / _ \/ _` \ \ / / |/ _` | __/ _ \
 * \ V  V /  __/ (_| |\ V /| | (_| | ||  __/
 *  \_/\_/ \___|\__,_| \_/ |_|\__,_|\__\___|
 *
 * Copyright © 2016 - 2019 Weaviate. All rights reserved.
 * LICENSE: https://github.com/creativesoftwarefdn/weaviate/blob/develop/LICENSE.md
 * DESIGN: Bob van Luijt (bob@k10y.co)
 */
// Code generated by go-swagger; DO NOT EDIT.

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

// WeaviateActionsValidateReader is a Reader for the WeaviateActionsValidate structure.
type WeaviateActionsValidateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *WeaviateActionsValidateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewWeaviateActionsValidateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewWeaviateActionsValidateUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewWeaviateActionsValidateForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 422:
		result := NewWeaviateActionsValidateUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewWeaviateActionsValidateOK creates a WeaviateActionsValidateOK with default headers values
func NewWeaviateActionsValidateOK() *WeaviateActionsValidateOK {
	return &WeaviateActionsValidateOK{}
}

/*WeaviateActionsValidateOK handles this case with default header values.

Successfully validated.
*/
type WeaviateActionsValidateOK struct {
}

func (o *WeaviateActionsValidateOK) Error() string {
	return fmt.Sprintf("[POST /actions/validate][%d] weaviateActionsValidateOK ", 200)
}

func (o *WeaviateActionsValidateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewWeaviateActionsValidateUnauthorized creates a WeaviateActionsValidateUnauthorized with default headers values
func NewWeaviateActionsValidateUnauthorized() *WeaviateActionsValidateUnauthorized {
	return &WeaviateActionsValidateUnauthorized{}
}

/*WeaviateActionsValidateUnauthorized handles this case with default header values.

Unauthorized or invalid credentials.
*/
type WeaviateActionsValidateUnauthorized struct {
}

func (o *WeaviateActionsValidateUnauthorized) Error() string {
	return fmt.Sprintf("[POST /actions/validate][%d] weaviateActionsValidateUnauthorized ", 401)
}

func (o *WeaviateActionsValidateUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewWeaviateActionsValidateForbidden creates a WeaviateActionsValidateForbidden with default headers values
func NewWeaviateActionsValidateForbidden() *WeaviateActionsValidateForbidden {
	return &WeaviateActionsValidateForbidden{}
}

/*WeaviateActionsValidateForbidden handles this case with default header values.

The used API-key has insufficient permissions.
*/
type WeaviateActionsValidateForbidden struct {
}

func (o *WeaviateActionsValidateForbidden) Error() string {
	return fmt.Sprintf("[POST /actions/validate][%d] weaviateActionsValidateForbidden ", 403)
}

func (o *WeaviateActionsValidateForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewWeaviateActionsValidateUnprocessableEntity creates a WeaviateActionsValidateUnprocessableEntity with default headers values
func NewWeaviateActionsValidateUnprocessableEntity() *WeaviateActionsValidateUnprocessableEntity {
	return &WeaviateActionsValidateUnprocessableEntity{}
}

/*WeaviateActionsValidateUnprocessableEntity handles this case with default header values.

Request body is well-formed (i.e., syntactically correct), but semantically erroneous. Are you sure the class is defined in the configuration file?
*/
type WeaviateActionsValidateUnprocessableEntity struct {
	Payload *models.ErrorResponse
}

func (o *WeaviateActionsValidateUnprocessableEntity) Error() string {
	return fmt.Sprintf("[POST /actions/validate][%d] weaviateActionsValidateUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *WeaviateActionsValidateUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
