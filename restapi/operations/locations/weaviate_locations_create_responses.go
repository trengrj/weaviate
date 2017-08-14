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

// WeaviateLocationsCreateAcceptedCode is the HTTP code returned for type WeaviateLocationsCreateAccepted
const WeaviateLocationsCreateAcceptedCode int = 202

/*WeaviateLocationsCreateAccepted Successfully received.

swagger:response weaviateLocationsCreateAccepted
*/
type WeaviateLocationsCreateAccepted struct {

	/*
	  In: Body
	*/
	Payload *models.LocationGetResponse `json:"body,omitempty"`
}

// NewWeaviateLocationsCreateAccepted creates WeaviateLocationsCreateAccepted with default headers values
func NewWeaviateLocationsCreateAccepted() *WeaviateLocationsCreateAccepted {
	return &WeaviateLocationsCreateAccepted{}
}

// WithPayload adds the payload to the weaviate locations create accepted response
func (o *WeaviateLocationsCreateAccepted) WithPayload(payload *models.LocationGetResponse) *WeaviateLocationsCreateAccepted {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the weaviate locations create accepted response
func (o *WeaviateLocationsCreateAccepted) SetPayload(payload *models.LocationGetResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *WeaviateLocationsCreateAccepted) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(202)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// WeaviateLocationsCreateUnauthorizedCode is the HTTP code returned for type WeaviateLocationsCreateUnauthorized
const WeaviateLocationsCreateUnauthorizedCode int = 401

/*WeaviateLocationsCreateUnauthorized Unauthorized or invalid credentials.

swagger:response weaviateLocationsCreateUnauthorized
*/
type WeaviateLocationsCreateUnauthorized struct {
}

// NewWeaviateLocationsCreateUnauthorized creates WeaviateLocationsCreateUnauthorized with default headers values
func NewWeaviateLocationsCreateUnauthorized() *WeaviateLocationsCreateUnauthorized {
	return &WeaviateLocationsCreateUnauthorized{}
}

// WriteResponse to the client
func (o *WeaviateLocationsCreateUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
}

// WeaviateLocationsCreateForbiddenCode is the HTTP code returned for type WeaviateLocationsCreateForbidden
const WeaviateLocationsCreateForbiddenCode int = 403

/*WeaviateLocationsCreateForbidden The used API-key has insufficient permissions.

swagger:response weaviateLocationsCreateForbidden
*/
type WeaviateLocationsCreateForbidden struct {
}

// NewWeaviateLocationsCreateForbidden creates WeaviateLocationsCreateForbidden with default headers values
func NewWeaviateLocationsCreateForbidden() *WeaviateLocationsCreateForbidden {
	return &WeaviateLocationsCreateForbidden{}
}

// WriteResponse to the client
func (o *WeaviateLocationsCreateForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
}

// WeaviateLocationsCreateUnprocessableEntityCode is the HTTP code returned for type WeaviateLocationsCreateUnprocessableEntity
const WeaviateLocationsCreateUnprocessableEntityCode int = 422

/*WeaviateLocationsCreateUnprocessableEntity Can not validate, check the body.

swagger:response weaviateLocationsCreateUnprocessableEntity
*/
type WeaviateLocationsCreateUnprocessableEntity struct {
}

// NewWeaviateLocationsCreateUnprocessableEntity creates WeaviateLocationsCreateUnprocessableEntity with default headers values
func NewWeaviateLocationsCreateUnprocessableEntity() *WeaviateLocationsCreateUnprocessableEntity {
	return &WeaviateLocationsCreateUnprocessableEntity{}
}

// WriteResponse to the client
func (o *WeaviateLocationsCreateUnprocessableEntity) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(422)
}

// WeaviateLocationsCreateNotImplementedCode is the HTTP code returned for type WeaviateLocationsCreateNotImplemented
const WeaviateLocationsCreateNotImplementedCode int = 501

/*WeaviateLocationsCreateNotImplemented Not (yet) implemented.

swagger:response weaviateLocationsCreateNotImplemented
*/
type WeaviateLocationsCreateNotImplemented struct {
}

// NewWeaviateLocationsCreateNotImplemented creates WeaviateLocationsCreateNotImplemented with default headers values
func NewWeaviateLocationsCreateNotImplemented() *WeaviateLocationsCreateNotImplemented {
	return &WeaviateLocationsCreateNotImplemented{}
}

// WriteResponse to the client
func (o *WeaviateLocationsCreateNotImplemented) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(501)
}
