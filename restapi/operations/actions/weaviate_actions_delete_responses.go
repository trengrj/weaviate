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

package actions

import (
	"net/http"

	"github.com/go-openapi/runtime"
)

// WeaviateActionsDeleteNoContentCode is the HTTP code returned for type WeaviateActionsDeleteNoContent
const WeaviateActionsDeleteNoContentCode int = 204

/*WeaviateActionsDeleteNoContent Successful deleted.

swagger:response weaviateActionsDeleteNoContent
*/
type WeaviateActionsDeleteNoContent struct {
}

// NewWeaviateActionsDeleteNoContent creates WeaviateActionsDeleteNoContent with default headers values
func NewWeaviateActionsDeleteNoContent() *WeaviateActionsDeleteNoContent {
	return &WeaviateActionsDeleteNoContent{}
}

// WriteResponse to the client
func (o *WeaviateActionsDeleteNoContent) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(204)
}

// WeaviateActionsDeleteUnauthorizedCode is the HTTP code returned for type WeaviateActionsDeleteUnauthorized
const WeaviateActionsDeleteUnauthorizedCode int = 401

/*WeaviateActionsDeleteUnauthorized Unauthorized or invalid credentials.

swagger:response weaviateActionsDeleteUnauthorized
*/
type WeaviateActionsDeleteUnauthorized struct {
}

// NewWeaviateActionsDeleteUnauthorized creates WeaviateActionsDeleteUnauthorized with default headers values
func NewWeaviateActionsDeleteUnauthorized() *WeaviateActionsDeleteUnauthorized {
	return &WeaviateActionsDeleteUnauthorized{}
}

// WriteResponse to the client
func (o *WeaviateActionsDeleteUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
}

// WeaviateActionsDeleteForbiddenCode is the HTTP code returned for type WeaviateActionsDeleteForbidden
const WeaviateActionsDeleteForbiddenCode int = 403

/*WeaviateActionsDeleteForbidden The used API-key has insufficient permissions.

swagger:response weaviateActionsDeleteForbidden
*/
type WeaviateActionsDeleteForbidden struct {
}

// NewWeaviateActionsDeleteForbidden creates WeaviateActionsDeleteForbidden with default headers values
func NewWeaviateActionsDeleteForbidden() *WeaviateActionsDeleteForbidden {
	return &WeaviateActionsDeleteForbidden{}
}

// WriteResponse to the client
func (o *WeaviateActionsDeleteForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
}

// WeaviateActionsDeleteNotFoundCode is the HTTP code returned for type WeaviateActionsDeleteNotFound
const WeaviateActionsDeleteNotFoundCode int = 404

/*WeaviateActionsDeleteNotFound Successful query result but no resource was found.

swagger:response weaviateActionsDeleteNotFound
*/
type WeaviateActionsDeleteNotFound struct {
}

// NewWeaviateActionsDeleteNotFound creates WeaviateActionsDeleteNotFound with default headers values
func NewWeaviateActionsDeleteNotFound() *WeaviateActionsDeleteNotFound {
	return &WeaviateActionsDeleteNotFound{}
}

// WriteResponse to the client
func (o *WeaviateActionsDeleteNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
}

// WeaviateActionsDeleteNotImplementedCode is the HTTP code returned for type WeaviateActionsDeleteNotImplemented
const WeaviateActionsDeleteNotImplementedCode int = 501

/*WeaviateActionsDeleteNotImplemented Not (yet) implemented.

swagger:response weaviateActionsDeleteNotImplemented
*/
type WeaviateActionsDeleteNotImplemented struct {
}

// NewWeaviateActionsDeleteNotImplemented creates WeaviateActionsDeleteNotImplemented with default headers values
func NewWeaviateActionsDeleteNotImplemented() *WeaviateActionsDeleteNotImplemented {
	return &WeaviateActionsDeleteNotImplemented{}
}

// WriteResponse to the client
func (o *WeaviateActionsDeleteNotImplemented) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(501)
}
