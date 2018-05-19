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
)

// WeaviatePeersEchoOKCode is the HTTP code returned for type WeaviatePeersEchoOK
const WeaviatePeersEchoOKCode int = 200

/*WeaviatePeersEchoOK Alive and kicking!

swagger:response weaviatePeersEchoOK
*/
type WeaviatePeersEchoOK struct {
}

// NewWeaviatePeersEchoOK creates WeaviatePeersEchoOK with default headers values
func NewWeaviatePeersEchoOK() *WeaviatePeersEchoOK {
	return &WeaviatePeersEchoOK{}
}

// WriteResponse to the client
func (o *WeaviatePeersEchoOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
}

// WeaviatePeersEchoNotImplementedCode is the HTTP code returned for type WeaviatePeersEchoNotImplemented
const WeaviatePeersEchoNotImplementedCode int = 501

/*WeaviatePeersEchoNotImplemented Not (yet) implemented.

swagger:response weaviatePeersEchoNotImplemented
*/
type WeaviatePeersEchoNotImplemented struct {
}

// NewWeaviatePeersEchoNotImplemented creates WeaviatePeersEchoNotImplemented with default headers values
func NewWeaviatePeersEchoNotImplemented() *WeaviatePeersEchoNotImplemented {
	return &WeaviatePeersEchoNotImplemented{}
}

// WriteResponse to the client
func (o *WeaviatePeersEchoNotImplemented) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(501)
}
