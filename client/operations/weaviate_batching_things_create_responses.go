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

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"
	"io"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/creativesoftwarefdn/weaviate/models"
)

// WeaviateBatchingThingsCreateReader is a Reader for the WeaviateBatchingThingsCreate structure.
type WeaviateBatchingThingsCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *WeaviateBatchingThingsCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewWeaviateBatchingThingsCreateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewWeaviateBatchingThingsCreateUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewWeaviateBatchingThingsCreateForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 422:
		result := NewWeaviateBatchingThingsCreateUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewWeaviateBatchingThingsCreateInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewWeaviateBatchingThingsCreateOK creates a WeaviateBatchingThingsCreateOK with default headers values
func NewWeaviateBatchingThingsCreateOK() *WeaviateBatchingThingsCreateOK {
	return &WeaviateBatchingThingsCreateOK{}
}

/*WeaviateBatchingThingsCreateOK handles this case with default header values.

Request succeeded, see response body to get detailed information about each batched item.
*/
type WeaviateBatchingThingsCreateOK struct {
	Payload []*models.ThingsGetResponse
}

func (o *WeaviateBatchingThingsCreateOK) Error() string {
	return fmt.Sprintf("[POST /batching/things][%d] weaviateBatchingThingsCreateOK  %+v", 200, o.Payload)
}

func (o *WeaviateBatchingThingsCreateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewWeaviateBatchingThingsCreateUnauthorized creates a WeaviateBatchingThingsCreateUnauthorized with default headers values
func NewWeaviateBatchingThingsCreateUnauthorized() *WeaviateBatchingThingsCreateUnauthorized {
	return &WeaviateBatchingThingsCreateUnauthorized{}
}

/*WeaviateBatchingThingsCreateUnauthorized handles this case with default header values.

Unauthorized or invalid credentials.
*/
type WeaviateBatchingThingsCreateUnauthorized struct {
}

func (o *WeaviateBatchingThingsCreateUnauthorized) Error() string {
	return fmt.Sprintf("[POST /batching/things][%d] weaviateBatchingThingsCreateUnauthorized ", 401)
}

func (o *WeaviateBatchingThingsCreateUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewWeaviateBatchingThingsCreateForbidden creates a WeaviateBatchingThingsCreateForbidden with default headers values
func NewWeaviateBatchingThingsCreateForbidden() *WeaviateBatchingThingsCreateForbidden {
	return &WeaviateBatchingThingsCreateForbidden{}
}

/*WeaviateBatchingThingsCreateForbidden handles this case with default header values.

Insufficient permissions.
*/
type WeaviateBatchingThingsCreateForbidden struct {
}

func (o *WeaviateBatchingThingsCreateForbidden) Error() string {
	return fmt.Sprintf("[POST /batching/things][%d] weaviateBatchingThingsCreateForbidden ", 403)
}

func (o *WeaviateBatchingThingsCreateForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewWeaviateBatchingThingsCreateUnprocessableEntity creates a WeaviateBatchingThingsCreateUnprocessableEntity with default headers values
func NewWeaviateBatchingThingsCreateUnprocessableEntity() *WeaviateBatchingThingsCreateUnprocessableEntity {
	return &WeaviateBatchingThingsCreateUnprocessableEntity{}
}

/*WeaviateBatchingThingsCreateUnprocessableEntity handles this case with default header values.

Request body is well-formed (i.e., syntactically correct), but semantically erroneous. Are you sure the class is defined in the configuration file?
*/
type WeaviateBatchingThingsCreateUnprocessableEntity struct {
	Payload *models.ErrorResponse
}

func (o *WeaviateBatchingThingsCreateUnprocessableEntity) Error() string {
	return fmt.Sprintf("[POST /batching/things][%d] weaviateBatchingThingsCreateUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *WeaviateBatchingThingsCreateUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewWeaviateBatchingThingsCreateInternalServerError creates a WeaviateBatchingThingsCreateInternalServerError with default headers values
func NewWeaviateBatchingThingsCreateInternalServerError() *WeaviateBatchingThingsCreateInternalServerError {
	return &WeaviateBatchingThingsCreateInternalServerError{}
}

/*WeaviateBatchingThingsCreateInternalServerError handles this case with default header values.

An error has occurred while trying to fulfill the request. Most likely the ErrorResponse will contain more information about the error.
*/
type WeaviateBatchingThingsCreateInternalServerError struct {
	Payload *models.ErrorResponse
}

func (o *WeaviateBatchingThingsCreateInternalServerError) Error() string {
	return fmt.Sprintf("[POST /batching/things][%d] weaviateBatchingThingsCreateInternalServerError  %+v", 500, o.Payload)
}

func (o *WeaviateBatchingThingsCreateInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*WeaviateBatchingThingsCreateBody weaviate batching things create body
swagger:model WeaviateBatchingThingsCreateBody
*/
type WeaviateBatchingThingsCreateBody struct {

	// Define which fields need to be returned. Default value is ALL
	Fields []*string `json:"fields"`

	// things
	Things []*models.Thing `json:"things"`
}

// Validate validates this weaviate batching things create body
func (o *WeaviateBatchingThingsCreateBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateFields(formats); err != nil {
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

var weaviateBatchingThingsCreateBodyFieldsItemsEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["ALL","class","schema","id","creationTimeUnix"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		weaviateBatchingThingsCreateBodyFieldsItemsEnum = append(weaviateBatchingThingsCreateBodyFieldsItemsEnum, v)
	}
}

func (o *WeaviateBatchingThingsCreateBody) validateFieldsItemsEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, weaviateBatchingThingsCreateBodyFieldsItemsEnum); err != nil {
		return err
	}
	return nil
}

func (o *WeaviateBatchingThingsCreateBody) validateFields(formats strfmt.Registry) error {

	if swag.IsZero(o.Fields) { // not required
		return nil
	}

	for i := 0; i < len(o.Fields); i++ {
		if swag.IsZero(o.Fields[i]) { // not required
			continue
		}

		// value enum
		if err := o.validateFieldsItemsEnum("body"+"."+"fields"+"."+strconv.Itoa(i), "body", *o.Fields[i]); err != nil {
			return err
		}

	}

	return nil
}

func (o *WeaviateBatchingThingsCreateBody) validateThings(formats strfmt.Registry) error {

	if swag.IsZero(o.Things) { // not required
		return nil
	}

	for i := 0; i < len(o.Things); i++ {
		if swag.IsZero(o.Things[i]) { // not required
			continue
		}

		if o.Things[i] != nil {
			if err := o.Things[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("body" + "." + "things" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *WeaviateBatchingThingsCreateBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *WeaviateBatchingThingsCreateBody) UnmarshalBinary(b []byte) error {
	var res WeaviateBatchingThingsCreateBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
