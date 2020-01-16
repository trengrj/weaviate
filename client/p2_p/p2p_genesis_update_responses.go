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

package p2_p

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/semi-technologies/weaviate/entities/models"
)

// P2pGenesisUpdateReader is a Reader for the P2pGenesisUpdate structure.
type P2pGenesisUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *P2pGenesisUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewP2pGenesisUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewP2pGenesisUpdateUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewP2pGenesisUpdateInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewP2pGenesisUpdateOK creates a P2pGenesisUpdateOK with default headers values
func NewP2pGenesisUpdateOK() *P2pGenesisUpdateOK {
	return &P2pGenesisUpdateOK{}
}

/*P2pGenesisUpdateOK handles this case with default header values.

Alive and kicking!
*/
type P2pGenesisUpdateOK struct {
}

func (o *P2pGenesisUpdateOK) Error() string {
	return fmt.Sprintf("[PUT /p2p/genesis][%d] p2pGenesisUpdateOK ", 200)
}

func (o *P2pGenesisUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewP2pGenesisUpdateUnauthorized creates a P2pGenesisUpdateUnauthorized with default headers values
func NewP2pGenesisUpdateUnauthorized() *P2pGenesisUpdateUnauthorized {
	return &P2pGenesisUpdateUnauthorized{}
}

/*P2pGenesisUpdateUnauthorized handles this case with default header values.

Unauthorized update.
*/
type P2pGenesisUpdateUnauthorized struct {
}

func (o *P2pGenesisUpdateUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /p2p/genesis][%d] p2pGenesisUpdateUnauthorized ", 401)
}

func (o *P2pGenesisUpdateUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewP2pGenesisUpdateInternalServerError creates a P2pGenesisUpdateInternalServerError with default headers values
func NewP2pGenesisUpdateInternalServerError() *P2pGenesisUpdateInternalServerError {
	return &P2pGenesisUpdateInternalServerError{}
}

/*P2pGenesisUpdateInternalServerError handles this case with default header values.

An error has occurred while trying to fulfill the request. Most likely the ErrorResponse will contain more information about the error.
*/
type P2pGenesisUpdateInternalServerError struct {
	Payload *models.ErrorResponse
}

func (o *P2pGenesisUpdateInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /p2p/genesis][%d] p2pGenesisUpdateInternalServerError  %+v", 500, o.Payload)
}

func (o *P2pGenesisUpdateInternalServerError) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *P2pGenesisUpdateInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
