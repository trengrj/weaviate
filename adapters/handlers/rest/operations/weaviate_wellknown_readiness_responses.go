//                           _       _
// __      _____  __ ___   ___  __ _| |_ ___
// \ \ /\ / / _ \/ _` \ \ / / |/ _` | __/ _ \
//  \ V  V /  __/ (_| |\ V /| | (_| | ||  __/
//   \_/\_/ \___|\__,_| \_/ |_|\__,_|\__\___|
//
//  Copyright © 2016 - 2021 SeMI Technologies B.V. All rights reserved.
//
//  CONTACT: hello@semi.technology
//

// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
)

// WeaviateWellknownReadinessOKCode is the HTTP code returned for type WeaviateWellknownReadinessOK
const WeaviateWellknownReadinessOKCode int = 200

/*WeaviateWellknownReadinessOK The application has completed its start-up routine and is ready to accept traffic.

swagger:response weaviateWellknownReadinessOK
*/
type WeaviateWellknownReadinessOK struct {
}

// NewWeaviateWellknownReadinessOK creates WeaviateWellknownReadinessOK with default headers values
func NewWeaviateWellknownReadinessOK() *WeaviateWellknownReadinessOK {

	return &WeaviateWellknownReadinessOK{}
}

// WriteResponse to the client
func (o *WeaviateWellknownReadinessOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(200)
}

// WeaviateWellknownReadinessServiceUnavailableCode is the HTTP code returned for type WeaviateWellknownReadinessServiceUnavailable
const WeaviateWellknownReadinessServiceUnavailableCode int = 503

/*WeaviateWellknownReadinessServiceUnavailable The application is currently not able to serve traffic. If other horizontal replicas of weaviate are available and they are capable of receving traffic, all traffic should be redirected there instead.

swagger:response weaviateWellknownReadinessServiceUnavailable
*/
type WeaviateWellknownReadinessServiceUnavailable struct {
}

// NewWeaviateWellknownReadinessServiceUnavailable creates WeaviateWellknownReadinessServiceUnavailable with default headers values
func NewWeaviateWellknownReadinessServiceUnavailable() *WeaviateWellknownReadinessServiceUnavailable {

	return &WeaviateWellknownReadinessServiceUnavailable{}
}

// WriteResponse to the client
func (o *WeaviateWellknownReadinessServiceUnavailable) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(503)
}
