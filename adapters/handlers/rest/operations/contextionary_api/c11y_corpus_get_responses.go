//                           _       _
// __      _____  __ ___   ___  __ _| |_ ___
// \ \ /\ / / _ \/ _` \ \ / / |/ _` | __/ _ \
//  \ V  V /  __/ (_| |\ V /| | (_| | ||  __/
//   \_/\_/ \___|\__,_| \_/ |_|\__,_|\__\___|
//
//  Copyright © 2016 - 2019 Weaviate. All rights reserved.
//  LICENSE WEAVIATE OPEN SOURCE: https://www.semi.technology/playbook/playbook/contract-weaviate-OSS.html
//  LICENSE WEAVIATE ENTERPRISE: https://www.semi.technology/playbook/contract-weaviate-enterprise.html
//  CONCEPT: Bob van Luijt (@bobvanluijt)
//  CONTACT: hello@semi.technology
//

// Code generated by go-swagger; DO NOT EDIT.

package contextionary_api

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
)

// C11yCorpusGetNotImplementedCode is the HTTP code returned for type C11yCorpusGetNotImplemented
const C11yCorpusGetNotImplementedCode int = 501

/*C11yCorpusGetNotImplemented Not (yet) implemented.

swagger:response c11yCorpusGetNotImplemented
*/
type C11yCorpusGetNotImplemented struct {
}

// NewC11yCorpusGetNotImplemented creates C11yCorpusGetNotImplemented with default headers values
func NewC11yCorpusGetNotImplemented() *C11yCorpusGetNotImplemented {

	return &C11yCorpusGetNotImplemented{}
}

// WriteResponse to the client
func (o *C11yCorpusGetNotImplemented) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(501)
}
