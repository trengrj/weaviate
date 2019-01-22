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

package actions

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// WeaviateActionsDeleteReader is a Reader for the WeaviateActionsDelete structure.
type WeaviateActionsDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *WeaviateActionsDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 204:
		result := NewWeaviateActionsDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewWeaviateActionsDeleteUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewWeaviateActionsDeleteForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewWeaviateActionsDeleteNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewWeaviateActionsDeleteNoContent creates a WeaviateActionsDeleteNoContent with default headers values
func NewWeaviateActionsDeleteNoContent() *WeaviateActionsDeleteNoContent {
	return &WeaviateActionsDeleteNoContent{}
}

/*WeaviateActionsDeleteNoContent handles this case with default header values.

Successfully deleted.
*/
type WeaviateActionsDeleteNoContent struct {
}

func (o *WeaviateActionsDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /actions/{actionId}][%d] weaviateActionsDeleteNoContent ", 204)
}

func (o *WeaviateActionsDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewWeaviateActionsDeleteUnauthorized creates a WeaviateActionsDeleteUnauthorized with default headers values
func NewWeaviateActionsDeleteUnauthorized() *WeaviateActionsDeleteUnauthorized {
	return &WeaviateActionsDeleteUnauthorized{}
}

/*WeaviateActionsDeleteUnauthorized handles this case with default header values.

Unauthorized or invalid credentials.
*/
type WeaviateActionsDeleteUnauthorized struct {
}

func (o *WeaviateActionsDeleteUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /actions/{actionId}][%d] weaviateActionsDeleteUnauthorized ", 401)
}

func (o *WeaviateActionsDeleteUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewWeaviateActionsDeleteForbidden creates a WeaviateActionsDeleteForbidden with default headers values
func NewWeaviateActionsDeleteForbidden() *WeaviateActionsDeleteForbidden {
	return &WeaviateActionsDeleteForbidden{}
}

/*WeaviateActionsDeleteForbidden handles this case with default header values.

The used API-key has insufficient permissions.
*/
type WeaviateActionsDeleteForbidden struct {
}

func (o *WeaviateActionsDeleteForbidden) Error() string {
	return fmt.Sprintf("[DELETE /actions/{actionId}][%d] weaviateActionsDeleteForbidden ", 403)
}

func (o *WeaviateActionsDeleteForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewWeaviateActionsDeleteNotFound creates a WeaviateActionsDeleteNotFound with default headers values
func NewWeaviateActionsDeleteNotFound() *WeaviateActionsDeleteNotFound {
	return &WeaviateActionsDeleteNotFound{}
}

/*WeaviateActionsDeleteNotFound handles this case with default header values.

Successful query result but no resource was found.
*/
type WeaviateActionsDeleteNotFound struct {
}

func (o *WeaviateActionsDeleteNotFound) Error() string {
	return fmt.Sprintf("[DELETE /actions/{actionId}][%d] weaviateActionsDeleteNotFound ", 404)
}

func (o *WeaviateActionsDeleteNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
