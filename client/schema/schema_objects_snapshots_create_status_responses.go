//                           _       _
// __      _____  __ ___   ___  __ _| |_ ___
// \ \ /\ / / _ \/ _` \ \ / / |/ _` | __/ _ \
//  \ V  V /  __/ (_| |\ V /| | (_| | ||  __/
//   \_/\_/ \___|\__,_| \_/ |_|\__,_|\__\___|
//
//  Copyright © 2016 - 2022 SeMI Technologies B.V. All rights reserved.
//
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
	"github.com/go-openapi/strfmt"

	"github.com/semi-technologies/weaviate/entities/models"
)

// SchemaObjectsSnapshotsCreateStatusReader is a Reader for the SchemaObjectsSnapshotsCreateStatus structure.
type SchemaObjectsSnapshotsCreateStatusReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SchemaObjectsSnapshotsCreateStatusReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSchemaObjectsSnapshotsCreateStatusOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewSchemaObjectsSnapshotsCreateStatusUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewSchemaObjectsSnapshotsCreateStatusForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewSchemaObjectsSnapshotsCreateStatusNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewSchemaObjectsSnapshotsCreateStatusInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewSchemaObjectsSnapshotsCreateStatusOK creates a SchemaObjectsSnapshotsCreateStatusOK with default headers values
func NewSchemaObjectsSnapshotsCreateStatusOK() *SchemaObjectsSnapshotsCreateStatusOK {
	return &SchemaObjectsSnapshotsCreateStatusOK{}
}

/*
SchemaObjectsSnapshotsCreateStatusOK handles this case with default header values.

Snapshot creation status successfully returned
*/
type SchemaObjectsSnapshotsCreateStatusOK struct {
	Payload *models.SnapshotMeta
}

func (o *SchemaObjectsSnapshotsCreateStatusOK) Error() string {
	return fmt.Sprintf("[GET /schema/{className}/snapshots/{storageName}/{id}][%d] schemaObjectsSnapshotsCreateStatusOK  %+v", 200, o.Payload)
}

func (o *SchemaObjectsSnapshotsCreateStatusOK) GetPayload() *models.SnapshotMeta {
	return o.Payload
}

func (o *SchemaObjectsSnapshotsCreateStatusOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SnapshotMeta)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSchemaObjectsSnapshotsCreateStatusUnauthorized creates a SchemaObjectsSnapshotsCreateStatusUnauthorized with default headers values
func NewSchemaObjectsSnapshotsCreateStatusUnauthorized() *SchemaObjectsSnapshotsCreateStatusUnauthorized {
	return &SchemaObjectsSnapshotsCreateStatusUnauthorized{}
}

/*
SchemaObjectsSnapshotsCreateStatusUnauthorized handles this case with default header values.

Unauthorized or invalid credentials.
*/
type SchemaObjectsSnapshotsCreateStatusUnauthorized struct {
}

func (o *SchemaObjectsSnapshotsCreateStatusUnauthorized) Error() string {
	return fmt.Sprintf("[GET /schema/{className}/snapshots/{storageName}/{id}][%d] schemaObjectsSnapshotsCreateStatusUnauthorized ", 401)
}

func (o *SchemaObjectsSnapshotsCreateStatusUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewSchemaObjectsSnapshotsCreateStatusForbidden creates a SchemaObjectsSnapshotsCreateStatusForbidden with default headers values
func NewSchemaObjectsSnapshotsCreateStatusForbidden() *SchemaObjectsSnapshotsCreateStatusForbidden {
	return &SchemaObjectsSnapshotsCreateStatusForbidden{}
}

/*
SchemaObjectsSnapshotsCreateStatusForbidden handles this case with default header values.

Forbidden
*/
type SchemaObjectsSnapshotsCreateStatusForbidden struct {
	Payload *models.ErrorResponse
}

func (o *SchemaObjectsSnapshotsCreateStatusForbidden) Error() string {
	return fmt.Sprintf("[GET /schema/{className}/snapshots/{storageName}/{id}][%d] schemaObjectsSnapshotsCreateStatusForbidden  %+v", 403, o.Payload)
}

func (o *SchemaObjectsSnapshotsCreateStatusForbidden) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *SchemaObjectsSnapshotsCreateStatusForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSchemaObjectsSnapshotsCreateStatusNotFound creates a SchemaObjectsSnapshotsCreateStatusNotFound with default headers values
func NewSchemaObjectsSnapshotsCreateStatusNotFound() *SchemaObjectsSnapshotsCreateStatusNotFound {
	return &SchemaObjectsSnapshotsCreateStatusNotFound{}
}

/*
SchemaObjectsSnapshotsCreateStatusNotFound handles this case with default header values.

Not Found - Snapshot does not exist
*/
type SchemaObjectsSnapshotsCreateStatusNotFound struct {
}

func (o *SchemaObjectsSnapshotsCreateStatusNotFound) Error() string {
	return fmt.Sprintf("[GET /schema/{className}/snapshots/{storageName}/{id}][%d] schemaObjectsSnapshotsCreateStatusNotFound ", 404)
}

func (o *SchemaObjectsSnapshotsCreateStatusNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewSchemaObjectsSnapshotsCreateStatusInternalServerError creates a SchemaObjectsSnapshotsCreateStatusInternalServerError with default headers values
func NewSchemaObjectsSnapshotsCreateStatusInternalServerError() *SchemaObjectsSnapshotsCreateStatusInternalServerError {
	return &SchemaObjectsSnapshotsCreateStatusInternalServerError{}
}

/*
SchemaObjectsSnapshotsCreateStatusInternalServerError handles this case with default header values.

An error has occurred while trying to fulfill the request. Most likely the ErrorResponse will contain more information about the error.
*/
type SchemaObjectsSnapshotsCreateStatusInternalServerError struct {
	Payload *models.ErrorResponse
}

func (o *SchemaObjectsSnapshotsCreateStatusInternalServerError) Error() string {
	return fmt.Sprintf("[GET /schema/{className}/snapshots/{storageName}/{id}][%d] schemaObjectsSnapshotsCreateStatusInternalServerError  %+v", 500, o.Payload)
}

func (o *SchemaObjectsSnapshotsCreateStatusInternalServerError) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *SchemaObjectsSnapshotsCreateStatusInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
