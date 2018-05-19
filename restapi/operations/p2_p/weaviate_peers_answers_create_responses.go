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

package p2_p

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/creativesoftwarefdn/weaviate/models"
)

// WeaviatePeersAnswersCreateAcceptedCode is the HTTP code returned for type WeaviatePeersAnswersCreateAccepted
const WeaviatePeersAnswersCreateAcceptedCode int = 202

/*WeaviatePeersAnswersCreateAccepted Successfully received.

swagger:response weaviatePeersAnswersCreateAccepted
*/
type WeaviatePeersAnswersCreateAccepted struct {
}

// NewWeaviatePeersAnswersCreateAccepted creates WeaviatePeersAnswersCreateAccepted with default headers values
func NewWeaviatePeersAnswersCreateAccepted() *WeaviatePeersAnswersCreateAccepted {
	return &WeaviatePeersAnswersCreateAccepted{}
}

// WriteResponse to the client
func (o *WeaviatePeersAnswersCreateAccepted) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(202)
}

// WeaviatePeersAnswersCreateUnauthorizedCode is the HTTP code returned for type WeaviatePeersAnswersCreateUnauthorized
const WeaviatePeersAnswersCreateUnauthorizedCode int = 401

/*WeaviatePeersAnswersCreateUnauthorized Unauthorized or invalid credentials.

swagger:response weaviatePeersAnswersCreateUnauthorized
*/
type WeaviatePeersAnswersCreateUnauthorized struct {
}

// NewWeaviatePeersAnswersCreateUnauthorized creates WeaviatePeersAnswersCreateUnauthorized with default headers values
func NewWeaviatePeersAnswersCreateUnauthorized() *WeaviatePeersAnswersCreateUnauthorized {
	return &WeaviatePeersAnswersCreateUnauthorized{}
}

// WriteResponse to the client
func (o *WeaviatePeersAnswersCreateUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
}

// WeaviatePeersAnswersCreateForbiddenCode is the HTTP code returned for type WeaviatePeersAnswersCreateForbidden
const WeaviatePeersAnswersCreateForbiddenCode int = 403

/*WeaviatePeersAnswersCreateForbidden The used API-key has insufficient permissions.

swagger:response weaviatePeersAnswersCreateForbidden
*/
type WeaviatePeersAnswersCreateForbidden struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewWeaviatePeersAnswersCreateForbidden creates WeaviatePeersAnswersCreateForbidden with default headers values
func NewWeaviatePeersAnswersCreateForbidden() *WeaviatePeersAnswersCreateForbidden {
	return &WeaviatePeersAnswersCreateForbidden{}
}

// WithPayload adds the payload to the weaviate peers answers create forbidden response
func (o *WeaviatePeersAnswersCreateForbidden) WithPayload(payload *models.ErrorResponse) *WeaviatePeersAnswersCreateForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the weaviate peers answers create forbidden response
func (o *WeaviatePeersAnswersCreateForbidden) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *WeaviatePeersAnswersCreateForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// WeaviatePeersAnswersCreateNotImplementedCode is the HTTP code returned for type WeaviatePeersAnswersCreateNotImplemented
const WeaviatePeersAnswersCreateNotImplementedCode int = 501

/*WeaviatePeersAnswersCreateNotImplemented Not (yet) implemented.

swagger:response weaviatePeersAnswersCreateNotImplemented
*/
type WeaviatePeersAnswersCreateNotImplemented struct {
}

// NewWeaviatePeersAnswersCreateNotImplemented creates WeaviatePeersAnswersCreateNotImplemented with default headers values
func NewWeaviatePeersAnswersCreateNotImplemented() *WeaviatePeersAnswersCreateNotImplemented {
	return &WeaviatePeersAnswersCreateNotImplemented{}
}

// WriteResponse to the client
func (o *WeaviatePeersAnswersCreateNotImplemented) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(501)
}
