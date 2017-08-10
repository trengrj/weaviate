/*                          _       _
 *__      _____  __ ___   ___  __ _| |_ ___
 *\ \ /\ / / _ \/ _` \ \ / / |/ _` | __/ _ \
 * \ V  V /  __/ (_| |\ V /| | (_| | ||  __/
 *  \_/\_/ \___|\__,_| \_/ |_|\__,_|\__\___|
 *
 * Copyright © 2016 Weaviate. All rights reserved.
 * LICENSE: https://github.com/weaviate/weaviate/blob/master/LICENSE
 * AUTHOR: Bob van Luijt (bob@weaviate.com)
 * See www.weaviate.com for details
 * Contact: @weaviate_iot / yourfriends@weaviate.com
 */
   

package locations

 
 

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/weaviate/weaviate/models"
)

// WeaviateLocationsListOKCode is the HTTP code returned for type WeaviateLocationsListOK
const WeaviateLocationsListOKCode int = 200

/*WeaviateLocationsListOK Successful response.

swagger:response weaviateLocationsListOK
*/
type WeaviateLocationsListOK struct {

	/*
	  In: Body
	*/
	Payload *models.LocationsListResponse `json:"body,omitempty"`
}

// NewWeaviateLocationsListOK creates WeaviateLocationsListOK with default headers values
func NewWeaviateLocationsListOK() *WeaviateLocationsListOK {
	return &WeaviateLocationsListOK{}
}

// WithPayload adds the payload to the weaviate locations list o k response
func (o *WeaviateLocationsListOK) WithPayload(payload *models.LocationsListResponse) *WeaviateLocationsListOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the weaviate locations list o k response
func (o *WeaviateLocationsListOK) SetPayload(payload *models.LocationsListResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *WeaviateLocationsListOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// WeaviateLocationsListUnauthorizedCode is the HTTP code returned for type WeaviateLocationsListUnauthorized
const WeaviateLocationsListUnauthorizedCode int = 401

/*WeaviateLocationsListUnauthorized Unauthorized or invalid credentials.

swagger:response weaviateLocationsListUnauthorized
*/
type WeaviateLocationsListUnauthorized struct {
}

// NewWeaviateLocationsListUnauthorized creates WeaviateLocationsListUnauthorized with default headers values
func NewWeaviateLocationsListUnauthorized() *WeaviateLocationsListUnauthorized {
	return &WeaviateLocationsListUnauthorized{}
}

// WriteResponse to the client
func (o *WeaviateLocationsListUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
}

// WeaviateLocationsListForbiddenCode is the HTTP code returned for type WeaviateLocationsListForbidden
const WeaviateLocationsListForbiddenCode int = 403

/*WeaviateLocationsListForbidden The used API-key has insufficient permissions.

swagger:response weaviateLocationsListForbidden
*/
type WeaviateLocationsListForbidden struct {
}

// NewWeaviateLocationsListForbidden creates WeaviateLocationsListForbidden with default headers values
func NewWeaviateLocationsListForbidden() *WeaviateLocationsListForbidden {
	return &WeaviateLocationsListForbidden{}
}

// WriteResponse to the client
func (o *WeaviateLocationsListForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
}

// WeaviateLocationsListNotFoundCode is the HTTP code returned for type WeaviateLocationsListNotFound
const WeaviateLocationsListNotFoundCode int = 404

/*WeaviateLocationsListNotFound Successful query result but no resource was found.

swagger:response weaviateLocationsListNotFound
*/
type WeaviateLocationsListNotFound struct {
}

// NewWeaviateLocationsListNotFound creates WeaviateLocationsListNotFound with default headers values
func NewWeaviateLocationsListNotFound() *WeaviateLocationsListNotFound {
	return &WeaviateLocationsListNotFound{}
}

// WriteResponse to the client
func (o *WeaviateLocationsListNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
}

// WeaviateLocationsListNotImplementedCode is the HTTP code returned for type WeaviateLocationsListNotImplemented
const WeaviateLocationsListNotImplementedCode int = 501

/*WeaviateLocationsListNotImplemented Not (yet) implemented.

swagger:response weaviateLocationsListNotImplemented
*/
type WeaviateLocationsListNotImplemented struct {
}

// NewWeaviateLocationsListNotImplemented creates WeaviateLocationsListNotImplemented with default headers values
func NewWeaviateLocationsListNotImplemented() *WeaviateLocationsListNotImplemented {
	return &WeaviateLocationsListNotImplemented{}
}

// WriteResponse to the client
func (o *WeaviateLocationsListNotImplemented) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(501)
}
