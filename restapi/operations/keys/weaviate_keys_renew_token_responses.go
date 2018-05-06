/*                          _       _
 *__      _____  __ ___   ___  __ _| |_ ___
 *\ \ /\ / / _ \/ _` \ \ / / |/ _` | __/ _ \
 * \ V  V /  __/ (_| |\ V /| | (_| | ||  __/
 *  \_/\_/ \___|\__,_| \_/ |_|\__,_|\__\___|
 *
 * Copyright © 2016 - 2018 Weaviate. All rights reserved.
 * LICENSE: https://github.com/creativesoftwarefdn/weaviate/blob/develop/LICENSE.md
 * AUTHOR: Bob van Luijt (bob@kub.design)
 * See www.creativesoftwarefdn.org for details
 * Contact: @CreativeSofwFdn / bob@kub.design
 */

package keys

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/creativesoftwarefdn/weaviate/models"
)

// WeaviateKeysRenewTokenOKCode is the HTTP code returned for type WeaviateKeysRenewTokenOK
const WeaviateKeysRenewTokenOKCode int = 200

/*WeaviateKeysRenewTokenOK Successful response.

swagger:response weaviateKeysRenewTokenOK
*/
type WeaviateKeysRenewTokenOK struct {

	/*
	  In: Body
	*/
	Payload *models.KeyTokenGetResponse `json:"body,omitempty"`
}

// NewWeaviateKeysRenewTokenOK creates WeaviateKeysRenewTokenOK with default headers values
func NewWeaviateKeysRenewTokenOK() *WeaviateKeysRenewTokenOK {
	return &WeaviateKeysRenewTokenOK{}
}

// WithPayload adds the payload to the weaviate keys renew token o k response
func (o *WeaviateKeysRenewTokenOK) WithPayload(payload *models.KeyTokenGetResponse) *WeaviateKeysRenewTokenOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the weaviate keys renew token o k response
func (o *WeaviateKeysRenewTokenOK) SetPayload(payload *models.KeyTokenGetResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *WeaviateKeysRenewTokenOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// WeaviateKeysRenewTokenUnauthorizedCode is the HTTP code returned for type WeaviateKeysRenewTokenUnauthorized
const WeaviateKeysRenewTokenUnauthorizedCode int = 401

/*WeaviateKeysRenewTokenUnauthorized Unauthorized or invalid credentials.

swagger:response weaviateKeysRenewTokenUnauthorized
*/
type WeaviateKeysRenewTokenUnauthorized struct {
}

// NewWeaviateKeysRenewTokenUnauthorized creates WeaviateKeysRenewTokenUnauthorized with default headers values
func NewWeaviateKeysRenewTokenUnauthorized() *WeaviateKeysRenewTokenUnauthorized {
	return &WeaviateKeysRenewTokenUnauthorized{}
}

// WriteResponse to the client
func (o *WeaviateKeysRenewTokenUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
}

// WeaviateKeysRenewTokenForbiddenCode is the HTTP code returned for type WeaviateKeysRenewTokenForbidden
const WeaviateKeysRenewTokenForbiddenCode int = 403

/*WeaviateKeysRenewTokenForbidden The used API-key has insufficient permissions.

swagger:response weaviateKeysRenewTokenForbidden
*/
type WeaviateKeysRenewTokenForbidden struct {
}

// NewWeaviateKeysRenewTokenForbidden creates WeaviateKeysRenewTokenForbidden with default headers values
func NewWeaviateKeysRenewTokenForbidden() *WeaviateKeysRenewTokenForbidden {
	return &WeaviateKeysRenewTokenForbidden{}
}

// WriteResponse to the client
func (o *WeaviateKeysRenewTokenForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
}

// WeaviateKeysRenewTokenNotFoundCode is the HTTP code returned for type WeaviateKeysRenewTokenNotFound
const WeaviateKeysRenewTokenNotFoundCode int = 404

/*WeaviateKeysRenewTokenNotFound Successful query result but no resource was found.

swagger:response weaviateKeysRenewTokenNotFound
*/
type WeaviateKeysRenewTokenNotFound struct {
}

// NewWeaviateKeysRenewTokenNotFound creates WeaviateKeysRenewTokenNotFound with default headers values
func NewWeaviateKeysRenewTokenNotFound() *WeaviateKeysRenewTokenNotFound {
	return &WeaviateKeysRenewTokenNotFound{}
}

// WriteResponse to the client
func (o *WeaviateKeysRenewTokenNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
}

// WeaviateKeysRenewTokenUnprocessableEntityCode is the HTTP code returned for type WeaviateKeysRenewTokenUnprocessableEntity
const WeaviateKeysRenewTokenUnprocessableEntityCode int = 422

/*WeaviateKeysRenewTokenUnprocessableEntity Request body contains well-formed (i.e., syntactically correct), but semantically erroneous. Are you sure the class is defined in the configuration file?

swagger:response weaviateKeysRenewTokenUnprocessableEntity
*/
type WeaviateKeysRenewTokenUnprocessableEntity struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewWeaviateKeysRenewTokenUnprocessableEntity creates WeaviateKeysRenewTokenUnprocessableEntity with default headers values
func NewWeaviateKeysRenewTokenUnprocessableEntity() *WeaviateKeysRenewTokenUnprocessableEntity {
	return &WeaviateKeysRenewTokenUnprocessableEntity{}
}

// WithPayload adds the payload to the weaviate keys renew token unprocessable entity response
func (o *WeaviateKeysRenewTokenUnprocessableEntity) WithPayload(payload *models.ErrorResponse) *WeaviateKeysRenewTokenUnprocessableEntity {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the weaviate keys renew token unprocessable entity response
func (o *WeaviateKeysRenewTokenUnprocessableEntity) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *WeaviateKeysRenewTokenUnprocessableEntity) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(422)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// WeaviateKeysRenewTokenNotImplementedCode is the HTTP code returned for type WeaviateKeysRenewTokenNotImplemented
const WeaviateKeysRenewTokenNotImplementedCode int = 501

/*WeaviateKeysRenewTokenNotImplemented Not (yet) implemented.

swagger:response weaviateKeysRenewTokenNotImplemented
*/
type WeaviateKeysRenewTokenNotImplemented struct {
}

// NewWeaviateKeysRenewTokenNotImplemented creates WeaviateKeysRenewTokenNotImplemented with default headers values
func NewWeaviateKeysRenewTokenNotImplemented() *WeaviateKeysRenewTokenNotImplemented {
	return &WeaviateKeysRenewTokenNotImplemented{}
}

// WriteResponse to the client
func (o *WeaviateKeysRenewTokenNotImplemented) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(501)
}
