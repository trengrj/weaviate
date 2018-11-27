// Code generated by go-swagger; DO NOT EDIT.

package things

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/creativesoftwarefdn/weaviate/models"
)

// WeaviateThingsPatchReader is a Reader for the WeaviateThingsPatch structure.
type WeaviateThingsPatchReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *WeaviateThingsPatchReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewWeaviateThingsPatchOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 202:
		result := NewWeaviateThingsPatchAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewWeaviateThingsPatchBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewWeaviateThingsPatchUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewWeaviateThingsPatchForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewWeaviateThingsPatchNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 422:
		result := NewWeaviateThingsPatchUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewWeaviateThingsPatchOK creates a WeaviateThingsPatchOK with default headers values
func NewWeaviateThingsPatchOK() *WeaviateThingsPatchOK {
	return &WeaviateThingsPatchOK{}
}

/*WeaviateThingsPatchOK handles this case with default header values.

Successfully applied.
*/
type WeaviateThingsPatchOK struct {
	Payload *models.ThingGetResponse
}

func (o *WeaviateThingsPatchOK) Error() string {
	return fmt.Sprintf("[PATCH /things/{thingId}][%d] weaviateThingsPatchOK  %+v", 200, o.Payload)
}

func (o *WeaviateThingsPatchOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ThingGetResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewWeaviateThingsPatchAccepted creates a WeaviateThingsPatchAccepted with default headers values
func NewWeaviateThingsPatchAccepted() *WeaviateThingsPatchAccepted {
	return &WeaviateThingsPatchAccepted{}
}

/*WeaviateThingsPatchAccepted handles this case with default header values.

Successfully received.
*/
type WeaviateThingsPatchAccepted struct {
	Payload *models.ThingGetResponse
}

func (o *WeaviateThingsPatchAccepted) Error() string {
	return fmt.Sprintf("[PATCH /things/{thingId}][%d] weaviateThingsPatchAccepted  %+v", 202, o.Payload)
}

func (o *WeaviateThingsPatchAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ThingGetResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewWeaviateThingsPatchBadRequest creates a WeaviateThingsPatchBadRequest with default headers values
func NewWeaviateThingsPatchBadRequest() *WeaviateThingsPatchBadRequest {
	return &WeaviateThingsPatchBadRequest{}
}

/*WeaviateThingsPatchBadRequest handles this case with default header values.

The patch-JSON is malformed.
*/
type WeaviateThingsPatchBadRequest struct {
}

func (o *WeaviateThingsPatchBadRequest) Error() string {
	return fmt.Sprintf("[PATCH /things/{thingId}][%d] weaviateThingsPatchBadRequest ", 400)
}

func (o *WeaviateThingsPatchBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewWeaviateThingsPatchUnauthorized creates a WeaviateThingsPatchUnauthorized with default headers values
func NewWeaviateThingsPatchUnauthorized() *WeaviateThingsPatchUnauthorized {
	return &WeaviateThingsPatchUnauthorized{}
}

/*WeaviateThingsPatchUnauthorized handles this case with default header values.

Unauthorized or invalid credentials.
*/
type WeaviateThingsPatchUnauthorized struct {
}

func (o *WeaviateThingsPatchUnauthorized) Error() string {
	return fmt.Sprintf("[PATCH /things/{thingId}][%d] weaviateThingsPatchUnauthorized ", 401)
}

func (o *WeaviateThingsPatchUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewWeaviateThingsPatchForbidden creates a WeaviateThingsPatchForbidden with default headers values
func NewWeaviateThingsPatchForbidden() *WeaviateThingsPatchForbidden {
	return &WeaviateThingsPatchForbidden{}
}

/*WeaviateThingsPatchForbidden handles this case with default header values.

The used API-key has insufficient permissions.
*/
type WeaviateThingsPatchForbidden struct {
}

func (o *WeaviateThingsPatchForbidden) Error() string {
	return fmt.Sprintf("[PATCH /things/{thingId}][%d] weaviateThingsPatchForbidden ", 403)
}

func (o *WeaviateThingsPatchForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewWeaviateThingsPatchNotFound creates a WeaviateThingsPatchNotFound with default headers values
func NewWeaviateThingsPatchNotFound() *WeaviateThingsPatchNotFound {
	return &WeaviateThingsPatchNotFound{}
}

/*WeaviateThingsPatchNotFound handles this case with default header values.

Successful query result but no resource was found.
*/
type WeaviateThingsPatchNotFound struct {
}

func (o *WeaviateThingsPatchNotFound) Error() string {
	return fmt.Sprintf("[PATCH /things/{thingId}][%d] weaviateThingsPatchNotFound ", 404)
}

func (o *WeaviateThingsPatchNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewWeaviateThingsPatchUnprocessableEntity creates a WeaviateThingsPatchUnprocessableEntity with default headers values
func NewWeaviateThingsPatchUnprocessableEntity() *WeaviateThingsPatchUnprocessableEntity {
	return &WeaviateThingsPatchUnprocessableEntity{}
}

/*WeaviateThingsPatchUnprocessableEntity handles this case with default header values.

The patch-JSON is valid but unprocessable.
*/
type WeaviateThingsPatchUnprocessableEntity struct {
	Payload *models.ErrorResponse
}

func (o *WeaviateThingsPatchUnprocessableEntity) Error() string {
	return fmt.Sprintf("[PATCH /things/{thingId}][%d] weaviateThingsPatchUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *WeaviateThingsPatchUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
