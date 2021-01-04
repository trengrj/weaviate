//                           _       _
// __      _____  __ ___   ___  __ _| |_ ___
// \ \ /\ / / _ \/ _` \ \ / / |/ _` | __/ _ \
//  \ V  V /  __/ (_| |\ V /| | (_| | ||  __/
//   \_/\_/ \___|\__,_| \_/ |_|\__,_|\__\___|
//
//  Copyright © 2016 - 2020 SeMI Technologies B.V. All rights reserved.
//
//  CONTACT: hello@semi.technology
//

package objects

import (
	"github.com/go-openapi/strfmt"
	"github.com/semi-technologies/weaviate/entities/models"
	"github.com/semi-technologies/weaviate/entities/schema/crossref"
)

// BatchObject is a helper type that groups all the info about one object in a
// batch that belongs together, i.e. uuid, object body and error state.
//
// Consumers of an Object (i.e. database connector) should always check
// whether an error is already present by the time they receive a batch object.
// Errors can be introduced at all levels, e.g. validation.
//
// However, error'd objects are not removed to make sure that the list in
// Objects matches the order and content of the incoming batch request
type BatchObject struct {
	OriginalIndex int
	Err           error
	Object        *models.Object
	UUID          strfmt.UUID
	Vector        []float32
}

// BatchObjects groups many Object items together. The order matches the
// order from the original request. It can be turned into the expected response
// type using the .Response() method
type BatchObjects []BatchObject

// BatchReference is a helper type that groups all the info about one references in a
// batch that belongs together, i.e. from, to, original index and error state
//
// Consumers of an Object (i.e. database connector) should always check
// whether an error is already present by the time they receive a batch object.
// Errors can be introduced at all levels, e.g. validation.
//
// However, error'd objects are not removed to make sure that the list in
// Objects matches the order and content of the incoming batch request
type BatchReference struct {
	OriginalIndex int
	Err           error
	From          *crossref.RefSource
	To            *crossref.Ref
}

// BatchReferences groups many Reference items together. The order matches the
// order from the original request. It can be turned into the expected response
// type using the .Response() method
type BatchReferences []BatchReference
