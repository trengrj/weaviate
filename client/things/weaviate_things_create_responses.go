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

package things

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/creativesoftwarefdn/weaviate/models"
)

// WeaviateThingsCreateReader is a Reader for the WeaviateThingsCreate structure.
type WeaviateThingsCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *WeaviateThingsCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewWeaviateThingsCreateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 202:
		result := NewWeaviateThingsCreateAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewWeaviateThingsCreateUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewWeaviateThingsCreateForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 422:
		result := NewWeaviateThingsCreateUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewWeaviateThingsCreateInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewWeaviateThingsCreateOK creates a WeaviateThingsCreateOK with default headers values
func NewWeaviateThingsCreateOK() *WeaviateThingsCreateOK {
	return &WeaviateThingsCreateOK{}
}

/*WeaviateThingsCreateOK handles this case with default header values.

Thing created.
*/
type WeaviateThingsCreateOK struct {
	Payload *models.ThingGetResponse
}

func (o *WeaviateThingsCreateOK) Error() string {
	return fmt.Sprintf("[POST /things][%d] weaviateThingsCreateOK  %+v", 200, o.Payload)
}

func (o *WeaviateThingsCreateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ThingGetResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewWeaviateThingsCreateAccepted creates a WeaviateThingsCreateAccepted with default headers values
func NewWeaviateThingsCreateAccepted() *WeaviateThingsCreateAccepted {
	return &WeaviateThingsCreateAccepted{}
}

/*WeaviateThingsCreateAccepted handles this case with default header values.

Successfully received.
*/
type WeaviateThingsCreateAccepted struct {
	Payload *models.ThingGetResponse
}

func (o *WeaviateThingsCreateAccepted) Error() string {
	return fmt.Sprintf("[POST /things][%d] weaviateThingsCreateAccepted  %+v", 202, o.Payload)
}

func (o *WeaviateThingsCreateAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ThingGetResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewWeaviateThingsCreateUnauthorized creates a WeaviateThingsCreateUnauthorized with default headers values
func NewWeaviateThingsCreateUnauthorized() *WeaviateThingsCreateUnauthorized {
	return &WeaviateThingsCreateUnauthorized{}
}

/*WeaviateThingsCreateUnauthorized handles this case with default header values.

Unauthorized or invalid credentials.
*/
type WeaviateThingsCreateUnauthorized struct {
}

func (o *WeaviateThingsCreateUnauthorized) Error() string {
	return fmt.Sprintf("[POST /things][%d] weaviateThingsCreateUnauthorized ", 401)
}

func (o *WeaviateThingsCreateUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewWeaviateThingsCreateForbidden creates a WeaviateThingsCreateForbidden with default headers values
func NewWeaviateThingsCreateForbidden() *WeaviateThingsCreateForbidden {
	return &WeaviateThingsCreateForbidden{}
}

/*WeaviateThingsCreateForbidden handles this case with default header values.

The used API-key has insufficient permissions.
*/
type WeaviateThingsCreateForbidden struct {
}

func (o *WeaviateThingsCreateForbidden) Error() string {
	return fmt.Sprintf("[POST /things][%d] weaviateThingsCreateForbidden ", 403)
}

func (o *WeaviateThingsCreateForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewWeaviateThingsCreateUnprocessableEntity creates a WeaviateThingsCreateUnprocessableEntity with default headers values
func NewWeaviateThingsCreateUnprocessableEntity() *WeaviateThingsCreateUnprocessableEntity {
	return &WeaviateThingsCreateUnprocessableEntity{}
}

/*WeaviateThingsCreateUnprocessableEntity handles this case with default header values.

Request body is well-formed (i.e., syntactically correct), but semantically erroneous. Are you sure the class is defined in the configuration file?
*/
type WeaviateThingsCreateUnprocessableEntity struct {
	Payload *models.ErrorResponse
}

func (o *WeaviateThingsCreateUnprocessableEntity) Error() string {
	return fmt.Sprintf("[POST /things][%d] weaviateThingsCreateUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *WeaviateThingsCreateUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewWeaviateThingsCreateInternalServerError creates a WeaviateThingsCreateInternalServerError with default headers values
func NewWeaviateThingsCreateInternalServerError() *WeaviateThingsCreateInternalServerError {
	return &WeaviateThingsCreateInternalServerError{}
}

/*WeaviateThingsCreateInternalServerError handles this case with default header values.

An error has occurred while trying to fulfill the request. Most likely the ErrorResponse will contain more information about the error.
*/
type WeaviateThingsCreateInternalServerError struct {
	Payload *models.ErrorResponse
}

func (o *WeaviateThingsCreateInternalServerError) Error() string {
	return fmt.Sprintf("[POST /things][%d] weaviateThingsCreateInternalServerError  %+v", 500, o.Payload)
}

func (o *WeaviateThingsCreateInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*WeaviateThingsCreateBody weaviate things create body
swagger:model WeaviateThingsCreateBody
*/
type WeaviateThingsCreateBody struct {

	// If `async` is true, return a 202 with the new ID of the Thing. You will receive this response before the data is made persistent. If `async` is false, you will receive confirmation after the value is made persistent. The value of `async` defaults to false.
	Async bool `json:"async,omitempty"`

	// thing
	Thing *models.ThingCreate `json:"thing,omitempty"`
}

// Validate validates this weaviate things create body
func (o *WeaviateThingsCreateBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateThing(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *WeaviateThingsCreateBody) validateThing(formats strfmt.Registry) error {

	if swag.IsZero(o.Thing) { // not required
		return nil
	}

	if o.Thing != nil {
		if err := o.Thing.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("body" + "." + "thing")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *WeaviateThingsCreateBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *WeaviateThingsCreateBody) UnmarshalBinary(b []byte) error {
	var res WeaviateThingsCreateBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
