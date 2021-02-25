//                           _       _
// __      _____  __ ___   ___  __ _| |_ ___
// \ \ /\ / / _ \/ _` \ \ / / |/ _` | __/ _ \
//  \ V  V /  __/ (_| |\ V /| | (_| | ||  __/
//   \_/\_/ \___|\__,_| \_/ |_|\__,_|\__\___|
//
//  Copyright © 2016 - 2021 SeMI Technologies B.V. All rights reserved.
//
//  CONTACT: hello@semi.technology
//

package vectorizer

import (
	"context"

	"github.com/pkg/errors"
)

func (v *Vectorizer) Texts(ctx context.Context, inputs []string,
	settings ClassSettings) ([]float32, error) {
	return nil, errors.Errorf("vectorizers.Texts not implemented yet")
}
