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

package p2_p

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/creativesoftwarefdn/weaviate/entities/models"
)

// WeaviateP2pHealthReader is a Reader for the WeaviateP2pHealth structure.
type WeaviateP2pHealthReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *WeaviateP2pHealthReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewWeaviateP2pHealthOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 500:
		result := NewWeaviateP2pHealthInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewWeaviateP2pHealthOK creates a WeaviateP2pHealthOK with default headers values
func NewWeaviateP2pHealthOK() *WeaviateP2pHealthOK {
	return &WeaviateP2pHealthOK{}
}

/*WeaviateP2pHealthOK handles this case with default header values.

Alive and kicking!
*/
type WeaviateP2pHealthOK struct {
}

func (o *WeaviateP2pHealthOK) Error() string {
	return fmt.Sprintf("[GET /p2p/health][%d] weaviateP2pHealthOK ", 200)
}

func (o *WeaviateP2pHealthOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewWeaviateP2pHealthInternalServerError creates a WeaviateP2pHealthInternalServerError with default headers values
func NewWeaviateP2pHealthInternalServerError() *WeaviateP2pHealthInternalServerError {
	return &WeaviateP2pHealthInternalServerError{}
}

/*WeaviateP2pHealthInternalServerError handles this case with default header values.

An error has occurred while trying to fulfill the request. Most likely the ErrorResponse will contain more information about the error.
*/
type WeaviateP2pHealthInternalServerError struct {
	Payload *models.ErrorResponse
}

func (o *WeaviateP2pHealthInternalServerError) Error() string {
	return fmt.Sprintf("[GET /p2p/health][%d] weaviateP2pHealthInternalServerError  %+v", 500, o.Payload)
}

func (o *WeaviateP2pHealthInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
