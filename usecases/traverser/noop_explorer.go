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

package traverser

import (
	"context"

	"github.com/semi-technologies/weaviate/entities/search"
)

// NoOpExplorer errors if an explore operation is attempted
type NoOpExplorer struct {
	err error
}

// GetClass errors
func (n *NoOpExplorer) GetClass(ctx context.Context,
	params GetParams) ([]interface{}, error) {
	return nil, n.err
}

// Concepts errors
func (n *NoOpExplorer) Concepts(ctx context.Context,
	params NearTextParams) ([]search.Result, error) {
	return nil, n.err
}

// NewNoOpExplorer with variable error
func NewNoOpExplorer(err error) *NoOpExplorer {
	return &NoOpExplorer{err: err}
}
