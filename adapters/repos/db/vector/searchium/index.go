//                           _       _
// __      _____  __ ___   ___  __ _| |_ ___
// \ \ /\ / / _ \/ _` \ \ / / |/ _` | __/ _ \
//  \ V  V /  __/ (_| |\ V /| | (_| | ||  __/
//   \_/\_/ \___|\__,_| \_/ |_|\__,_|\__\___|
//
//  Copyright © 2016 - 2022 SeMI Technologies B.V. All rights reserved.
//
//  CONTACT: hello@semi.technology
//

package searchium

import (
	"context"

	"github.com/pkg/errors"
	"github.com/semi-technologies/weaviate/adapters/repos/db/helpers"
	"github.com/semi-technologies/weaviate/entities/schema"
)

type Index struct{}

func NewIndex() *Index {
	return &Index{}
}

func (i *Index) Add(id uint64, vector []float32) error {
	// silently ignore
	return nil
}

func (i *Index) Delete(id uint64) error {
	// silently ignore
	return nil
}

func (i *Index) SearchByVector(vector []float32, k int, allow helpers.AllowList) ([]uint64, []float32, error) {

	ids := make([]uint64, 2)
	ids[0] = 0
	ids[1] = 1

	distances := make([]float32, 2)
	distances[0] = 0.2
	distances[1] = 0.3
	return ids, distances, nil
}

func (i *Index) SearchByVectorDistance(vector []float32, dist float32, maxLimit int64, allow helpers.AllowList) ([]uint64, []float32, error) {
	ids := make([]uint64, 2)
	ids[0] = 0
	ids[1] = 1

	distances := make([]float32, 2)
	distances[0] = 0.2
	distances[1] = 0.3
	return ids, distances, nil
}

func (i *Index) UpdateUserConfig(updated schema.VectorIndexConfig) error {
	return errors.Errorf("cannot update vector index config on a non-indexed class. Delete and re-create without skip property")
}

func (i *Index) Drop(context.Context) error {
	// silently ignore
	return nil
}

func (i *Index) Flush() error {
	return nil
}

func (i *Index) Shutdown(context.Context) error {
	return nil
}

func (i *Index) PauseMaintenance(context.Context) error {
	return nil
}

func (i *Index) SwitchCommitLogs(context.Context) error {
	return nil
}

func (i *Index) ListFiles(context.Context) ([]string, error) {
	return nil, nil
}

func (i *Index) ResumeMaintenance(context.Context) error {
	return nil
}

func (i *Index) Dump(labels ...string) {
}
