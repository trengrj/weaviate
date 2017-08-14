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
   

package commands

 
 

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/weaviate/weaviate/models"
)

// WeaviateCommandsUpdateOKCode is the HTTP code returned for type WeaviateCommandsUpdateOK
const WeaviateCommandsUpdateOKCode int = 200

/*WeaviateCommandsUpdateOK Successful updated.

swagger:response weaviateCommandsUpdateOK
*/
type WeaviateCommandsUpdateOK struct {

	/*
	  In: Body
	*/
	Payload *models.CommandGetResponse `json:"body,omitempty"`
}

// NewWeaviateCommandsUpdateOK creates WeaviateCommandsUpdateOK with default headers values
func NewWeaviateCommandsUpdateOK() *WeaviateCommandsUpdateOK {
	return &WeaviateCommandsUpdateOK{}
}

// WithPayload adds the payload to the weaviate commands update o k response
func (o *WeaviateCommandsUpdateOK) WithPayload(payload *models.CommandGetResponse) *WeaviateCommandsUpdateOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the weaviate commands update o k response
func (o *WeaviateCommandsUpdateOK) SetPayload(payload *models.CommandGetResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *WeaviateCommandsUpdateOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// WeaviateCommandsUpdateUnauthorizedCode is the HTTP code returned for type WeaviateCommandsUpdateUnauthorized
const WeaviateCommandsUpdateUnauthorizedCode int = 401

/*WeaviateCommandsUpdateUnauthorized Unauthorized or invalid credentials.

swagger:response weaviateCommandsUpdateUnauthorized
*/
type WeaviateCommandsUpdateUnauthorized struct {
}

// NewWeaviateCommandsUpdateUnauthorized creates WeaviateCommandsUpdateUnauthorized with default headers values
func NewWeaviateCommandsUpdateUnauthorized() *WeaviateCommandsUpdateUnauthorized {
	return &WeaviateCommandsUpdateUnauthorized{}
}

// WriteResponse to the client
func (o *WeaviateCommandsUpdateUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
}

// WeaviateCommandsUpdateForbiddenCode is the HTTP code returned for type WeaviateCommandsUpdateForbidden
const WeaviateCommandsUpdateForbiddenCode int = 403

/*WeaviateCommandsUpdateForbidden The used API-key has insufficient permissions.

swagger:response weaviateCommandsUpdateForbidden
*/
type WeaviateCommandsUpdateForbidden struct {
}

// NewWeaviateCommandsUpdateForbidden creates WeaviateCommandsUpdateForbidden with default headers values
func NewWeaviateCommandsUpdateForbidden() *WeaviateCommandsUpdateForbidden {
	return &WeaviateCommandsUpdateForbidden{}
}

// WriteResponse to the client
func (o *WeaviateCommandsUpdateForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
}

// WeaviateCommandsUpdateNotFoundCode is the HTTP code returned for type WeaviateCommandsUpdateNotFound
const WeaviateCommandsUpdateNotFoundCode int = 404

/*WeaviateCommandsUpdateNotFound Successful query result but no resource was found.

swagger:response weaviateCommandsUpdateNotFound
*/
type WeaviateCommandsUpdateNotFound struct {
}

// NewWeaviateCommandsUpdateNotFound creates WeaviateCommandsUpdateNotFound with default headers values
func NewWeaviateCommandsUpdateNotFound() *WeaviateCommandsUpdateNotFound {
	return &WeaviateCommandsUpdateNotFound{}
}

// WriteResponse to the client
func (o *WeaviateCommandsUpdateNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
}

// WeaviateCommandsUpdateUnprocessableEntityCode is the HTTP code returned for type WeaviateCommandsUpdateUnprocessableEntity
const WeaviateCommandsUpdateUnprocessableEntityCode int = 422

/*WeaviateCommandsUpdateUnprocessableEntity Can not validate, check the body.

swagger:response weaviateCommandsUpdateUnprocessableEntity
*/
type WeaviateCommandsUpdateUnprocessableEntity struct {
}

// NewWeaviateCommandsUpdateUnprocessableEntity creates WeaviateCommandsUpdateUnprocessableEntity with default headers values
func NewWeaviateCommandsUpdateUnprocessableEntity() *WeaviateCommandsUpdateUnprocessableEntity {
	return &WeaviateCommandsUpdateUnprocessableEntity{}
}

// WriteResponse to the client
func (o *WeaviateCommandsUpdateUnprocessableEntity) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(422)
}

// WeaviateCommandsUpdateNotImplementedCode is the HTTP code returned for type WeaviateCommandsUpdateNotImplemented
const WeaviateCommandsUpdateNotImplementedCode int = 501

/*WeaviateCommandsUpdateNotImplemented Not (yet) implemented.

swagger:response weaviateCommandsUpdateNotImplemented
*/
type WeaviateCommandsUpdateNotImplemented struct {
}

// NewWeaviateCommandsUpdateNotImplemented creates WeaviateCommandsUpdateNotImplemented with default headers values
func NewWeaviateCommandsUpdateNotImplemented() *WeaviateCommandsUpdateNotImplemented {
	return &WeaviateCommandsUpdateNotImplemented{}
}

// WriteResponse to the client
func (o *WeaviateCommandsUpdateNotImplemented) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(501)
}
