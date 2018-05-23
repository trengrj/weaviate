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

// WeaviatePeersQuestionsCreateAcceptedCode is the HTTP code returned for type WeaviatePeersQuestionsCreateAccepted
const WeaviatePeersQuestionsCreateAcceptedCode int = 202

/*WeaviatePeersQuestionsCreateAccepted Successfully received the question and answer might be send back.

swagger:response weaviatePeersQuestionsCreateAccepted
*/
type WeaviatePeersQuestionsCreateAccepted struct {

	/*
	  In: Body
	*/
	Payload *models.QuestionResponse `json:"body,omitempty"`
}

// NewWeaviatePeersQuestionsCreateAccepted creates WeaviatePeersQuestionsCreateAccepted with default headers values
func NewWeaviatePeersQuestionsCreateAccepted() *WeaviatePeersQuestionsCreateAccepted {
	return &WeaviatePeersQuestionsCreateAccepted{}
}

// WithPayload adds the payload to the weaviate peers questions create accepted response
func (o *WeaviatePeersQuestionsCreateAccepted) WithPayload(payload *models.QuestionResponse) *WeaviatePeersQuestionsCreateAccepted {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the weaviate peers questions create accepted response
func (o *WeaviatePeersQuestionsCreateAccepted) SetPayload(payload *models.QuestionResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *WeaviatePeersQuestionsCreateAccepted) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(202)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// WeaviatePeersQuestionsCreateForbiddenCode is the HTTP code returned for type WeaviatePeersQuestionsCreateForbidden
const WeaviatePeersQuestionsCreateForbiddenCode int = 403

/*WeaviatePeersQuestionsCreateForbidden You are not allowed on the network.

swagger:response weaviatePeersQuestionsCreateForbidden
*/
type WeaviatePeersQuestionsCreateForbidden struct {
}

// NewWeaviatePeersQuestionsCreateForbidden creates WeaviatePeersQuestionsCreateForbidden with default headers values
func NewWeaviatePeersQuestionsCreateForbidden() *WeaviatePeersQuestionsCreateForbidden {
	return &WeaviatePeersQuestionsCreateForbidden{}
}

// WriteResponse to the client
func (o *WeaviatePeersQuestionsCreateForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
}

// WeaviatePeersQuestionsCreateNotImplementedCode is the HTTP code returned for type WeaviatePeersQuestionsCreateNotImplemented
const WeaviatePeersQuestionsCreateNotImplementedCode int = 501

/*WeaviatePeersQuestionsCreateNotImplemented Not (yet) implemented.

swagger:response weaviatePeersQuestionsCreateNotImplemented
*/
type WeaviatePeersQuestionsCreateNotImplemented struct {
}

// NewWeaviatePeersQuestionsCreateNotImplemented creates WeaviatePeersQuestionsCreateNotImplemented with default headers values
func NewWeaviatePeersQuestionsCreateNotImplemented() *WeaviatePeersQuestionsCreateNotImplemented {
	return &WeaviatePeersQuestionsCreateNotImplemented{}
}

// WriteResponse to the client
func (o *WeaviatePeersQuestionsCreateNotImplemented) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(501)
}
