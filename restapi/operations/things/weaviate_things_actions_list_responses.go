/*                          _       _
 *__      _____  __ ___   ___  __ _| |_ ___
 *\ \ /\ / / _ \/ _` \ \ / / |/ _` | __/ _ \
 * \ V  V /  __/ (_| |\ V /| | (_| | ||  __/
 *  \_/\_/ \___|\__,_| \_/ |_|\__,_|\__\___|
 *
 * Copyright © 2016 Weaviate. All rights reserved.
 * LICENSE: https://github.com/creativesoftwarefdn/weaviate/blob/develop/LICENSE.md
 * AUTHOR: Bob van Luijt (bob@weaviate.com)
 * See www.weaviate.com for details
 * Contact: @CreativeSofwFdn / yourfriends@weaviate.com
 */

package things

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/creativesoftwarefdn/weaviate/models"
)

// WeaviateThingsActionsListOKCode is the HTTP code returned for type WeaviateThingsActionsListOK
const WeaviateThingsActionsListOKCode int = 200

/*WeaviateThingsActionsListOK Successful response.

swagger:response weaviateThingsActionsListOK
*/
type WeaviateThingsActionsListOK struct {

	/*
	  In: Body
	*/
	Payload *models.ActionsListResponse `json:"body,omitempty"`
}

// NewWeaviateThingsActionsListOK creates WeaviateThingsActionsListOK with default headers values
func NewWeaviateThingsActionsListOK() *WeaviateThingsActionsListOK {
	return &WeaviateThingsActionsListOK{}
}

// WithPayload adds the payload to the weaviate things actions list o k response
func (o *WeaviateThingsActionsListOK) WithPayload(payload *models.ActionsListResponse) *WeaviateThingsActionsListOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the weaviate things actions list o k response
func (o *WeaviateThingsActionsListOK) SetPayload(payload *models.ActionsListResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *WeaviateThingsActionsListOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// WeaviateThingsActionsListUnauthorizedCode is the HTTP code returned for type WeaviateThingsActionsListUnauthorized
const WeaviateThingsActionsListUnauthorizedCode int = 401

/*WeaviateThingsActionsListUnauthorized Unauthorized or invalid credentials.

swagger:response weaviateThingsActionsListUnauthorized
*/
type WeaviateThingsActionsListUnauthorized struct {
}

// NewWeaviateThingsActionsListUnauthorized creates WeaviateThingsActionsListUnauthorized with default headers values
func NewWeaviateThingsActionsListUnauthorized() *WeaviateThingsActionsListUnauthorized {
	return &WeaviateThingsActionsListUnauthorized{}
}

// WriteResponse to the client
func (o *WeaviateThingsActionsListUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
}

// WeaviateThingsActionsListForbiddenCode is the HTTP code returned for type WeaviateThingsActionsListForbidden
const WeaviateThingsActionsListForbiddenCode int = 403

/*WeaviateThingsActionsListForbidden The used API-key has insufficient permissions.

swagger:response weaviateThingsActionsListForbidden
*/
type WeaviateThingsActionsListForbidden struct {
}

// NewWeaviateThingsActionsListForbidden creates WeaviateThingsActionsListForbidden with default headers values
func NewWeaviateThingsActionsListForbidden() *WeaviateThingsActionsListForbidden {
	return &WeaviateThingsActionsListForbidden{}
}

// WriteResponse to the client
func (o *WeaviateThingsActionsListForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
}

// WeaviateThingsActionsListNotFoundCode is the HTTP code returned for type WeaviateThingsActionsListNotFound
const WeaviateThingsActionsListNotFoundCode int = 404

/*WeaviateThingsActionsListNotFound Successful query result but no resource was found.

swagger:response weaviateThingsActionsListNotFound
*/
type WeaviateThingsActionsListNotFound struct {
}

// NewWeaviateThingsActionsListNotFound creates WeaviateThingsActionsListNotFound with default headers values
func NewWeaviateThingsActionsListNotFound() *WeaviateThingsActionsListNotFound {
	return &WeaviateThingsActionsListNotFound{}
}

// WriteResponse to the client
func (o *WeaviateThingsActionsListNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
}

// WeaviateThingsActionsListNotImplementedCode is the HTTP code returned for type WeaviateThingsActionsListNotImplemented
const WeaviateThingsActionsListNotImplementedCode int = 501

/*WeaviateThingsActionsListNotImplemented Not (yet) implemented.

swagger:response weaviateThingsActionsListNotImplemented
*/
type WeaviateThingsActionsListNotImplemented struct {
}

// NewWeaviateThingsActionsListNotImplemented creates WeaviateThingsActionsListNotImplemented with default headers values
func NewWeaviateThingsActionsListNotImplemented() *WeaviateThingsActionsListNotImplemented {
	return &WeaviateThingsActionsListNotImplemented{}
}

// WriteResponse to the client
func (o *WeaviateThingsActionsListNotImplemented) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(501)
}
