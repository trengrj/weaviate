/*                          _       _
 *__      _____  __ ___   ___  __ _| |_ ___
 *\ \ /\ / / _ \/ _` \ \ / / |/ _` | __/ _ \
 * \ V  V /  __/ (_| |\ V /| | (_| | ||  __/
 *  \_/\_/ \___|\__,_| \_/ |_|\__,_|\__\___|
 *
 * Copyright © 2016 - 2019 Weaviate. All rights reserved.
 * LICENSE: https://github.com/semi-technologies/weaviate/blob/develop/LICENSE.md
 * DESIGN & CONCEPT: Bob van Luijt (@bobvanluijt)
 * CONTACT: hello@semi.technology
 */ // Code generated by go-swagger; DO NOT EDIT.

package schema

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/semi-technologies/weaviate/entities/models"
)

// WeaviateSchemaDumpReader is a Reader for the WeaviateSchemaDump structure.
type WeaviateSchemaDumpReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *WeaviateSchemaDumpReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewWeaviateSchemaDumpOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewWeaviateSchemaDumpUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewWeaviateSchemaDumpForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewWeaviateSchemaDumpInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewWeaviateSchemaDumpOK creates a WeaviateSchemaDumpOK with default headers values
func NewWeaviateSchemaDumpOK() *WeaviateSchemaDumpOK {
	return &WeaviateSchemaDumpOK{}
}

/*WeaviateSchemaDumpOK handles this case with default header values.

Successfully dumped the database schema.
*/
type WeaviateSchemaDumpOK struct {
	Payload *WeaviateSchemaDumpOKBody
}

func (o *WeaviateSchemaDumpOK) Error() string {
	return fmt.Sprintf("[GET /schema][%d] weaviateSchemaDumpOK  %+v", 200, o.Payload)
}

func (o *WeaviateSchemaDumpOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(WeaviateSchemaDumpOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewWeaviateSchemaDumpUnauthorized creates a WeaviateSchemaDumpUnauthorized with default headers values
func NewWeaviateSchemaDumpUnauthorized() *WeaviateSchemaDumpUnauthorized {
	return &WeaviateSchemaDumpUnauthorized{}
}

/*WeaviateSchemaDumpUnauthorized handles this case with default header values.

Unauthorized or invalid credentials.
*/
type WeaviateSchemaDumpUnauthorized struct {
}

func (o *WeaviateSchemaDumpUnauthorized) Error() string {
	return fmt.Sprintf("[GET /schema][%d] weaviateSchemaDumpUnauthorized ", 401)
}

func (o *WeaviateSchemaDumpUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewWeaviateSchemaDumpForbidden creates a WeaviateSchemaDumpForbidden with default headers values
func NewWeaviateSchemaDumpForbidden() *WeaviateSchemaDumpForbidden {
	return &WeaviateSchemaDumpForbidden{}
}

/*WeaviateSchemaDumpForbidden handles this case with default header values.

Forbidden
*/
type WeaviateSchemaDumpForbidden struct {
	Payload *models.ErrorResponse
}

func (o *WeaviateSchemaDumpForbidden) Error() string {
	return fmt.Sprintf("[GET /schema][%d] weaviateSchemaDumpForbidden  %+v", 403, o.Payload)
}

func (o *WeaviateSchemaDumpForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewWeaviateSchemaDumpInternalServerError creates a WeaviateSchemaDumpInternalServerError with default headers values
func NewWeaviateSchemaDumpInternalServerError() *WeaviateSchemaDumpInternalServerError {
	return &WeaviateSchemaDumpInternalServerError{}
}

/*WeaviateSchemaDumpInternalServerError handles this case with default header values.

An error has occurred while trying to fulfill the request. Most likely the ErrorResponse will contain more information about the error.
*/
type WeaviateSchemaDumpInternalServerError struct {
	Payload *models.ErrorResponse
}

func (o *WeaviateSchemaDumpInternalServerError) Error() string {
	return fmt.Sprintf("[GET /schema][%d] weaviateSchemaDumpInternalServerError  %+v", 500, o.Payload)
}

func (o *WeaviateSchemaDumpInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*WeaviateSchemaDumpOKBody weaviate schema dump o k body
swagger:model WeaviateSchemaDumpOKBody
*/
type WeaviateSchemaDumpOKBody struct {

	// actions
	Actions *models.Schema `json:"actions,omitempty"`

	// things
	Things *models.Schema `json:"things,omitempty"`
}

// Validate validates this weaviate schema dump o k body
func (o *WeaviateSchemaDumpOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateActions(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateThings(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *WeaviateSchemaDumpOKBody) validateActions(formats strfmt.Registry) error {

	if swag.IsZero(o.Actions) { // not required
		return nil
	}

	if o.Actions != nil {
		if err := o.Actions.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("weaviateSchemaDumpOK" + "." + "actions")
			}
			return err
		}
	}

	return nil
}

func (o *WeaviateSchemaDumpOKBody) validateThings(formats strfmt.Registry) error {

	if swag.IsZero(o.Things) { // not required
		return nil
	}

	if o.Things != nil {
		if err := o.Things.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("weaviateSchemaDumpOK" + "." + "things")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *WeaviateSchemaDumpOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *WeaviateSchemaDumpOKBody) UnmarshalBinary(b []byte) error {
	var res WeaviateSchemaDumpOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
