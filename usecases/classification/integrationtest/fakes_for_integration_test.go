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

// +build integrationTest

package classification_integration_test

import (
	"context"
	"fmt"
	"sync"

	"github.com/go-openapi/strfmt"
	uuid "github.com/satori/go.uuid"
	"github.com/semi-technologies/weaviate/adapters/repos/db/vector/hnsw"
	"github.com/semi-technologies/weaviate/entities/models"
	"github.com/semi-technologies/weaviate/entities/schema"
	"github.com/semi-technologies/weaviate/entities/search"
)

type fakeSchemaGetter struct {
	schema schema.Schema
}

func (f *fakeSchemaGetter) GetSchemaSkipAuth() schema.Schema {
	return f.schema
}

type fakeClassificationRepo struct {
	sync.Mutex
	db map[strfmt.UUID]models.Classification
}

func newFakeClassificationRepo() *fakeClassificationRepo {
	return &fakeClassificationRepo{
		db: map[strfmt.UUID]models.Classification{},
	}
}

func (f *fakeClassificationRepo) Put(ctx context.Context, class models.Classification) error {
	f.Lock()
	defer f.Unlock()

	f.db[class.ID] = class
	return nil
}

func (f *fakeClassificationRepo) Get(ctx context.Context, id strfmt.UUID) (*models.Classification, error) {
	f.Lock()
	defer f.Unlock()

	class, ok := f.db[id]
	if !ok {
		return nil, nil
	}

	return &class, nil
}

func testSchema() schema.Schema {
	return schema.Schema{
		Objects: &models.Schema{
			Classes: []*models.Class{
				{
					Class:               "ExactCategory",
					VectorIndexConfig:   hnsw.NewDefaultUserConfig(),
					InvertedIndexConfig: invertedConfig(),
				},
				{
					Class:               "MainCategory",
					VectorIndexConfig:   hnsw.NewDefaultUserConfig(),
					InvertedIndexConfig: invertedConfig(),
				},
				{
					Class:               "Article",
					VectorIndexConfig:   hnsw.NewDefaultUserConfig(),
					InvertedIndexConfig: invertedConfig(),
					Properties: []*models.Property{
						{
							Name:     "description",
							DataType: []string{string(schema.DataTypeText)},
						},
						{
							Name:     "name",
							DataType: []string{string(schema.DataTypeString)},
						},
						{
							Name:     "exactCategory",
							DataType: []string{"ExactCategory"},
						},
						{
							Name:     "mainCategory",
							DataType: []string{"MainCategory"},
						},
						{
							Name:     "categories",
							DataType: []string{"ExactCategory"},
						},
						{
							Name:     "anyCategory",
							DataType: []string{"MainCategory", "ExactCategory"},
						},
					},
				},
			},
		},
	}
}

// only used for knn-type
func testDataAlreadyClassified() search.Results {
	return search.Results{
		search.Result{
			ID:        "8aeecd06-55a0-462c-9853-81b31a284d80",
			ClassName: "Article",
			Vector:    []float32{1, 0, 0},
			Schema: map[string]interface{}{
				"description":   "This article talks about politics",
				"exactCategory": models.MultipleRef{beaconRef(idCategoryPolitics)},
				"mainCategory":  models.MultipleRef{beaconRef(idMainCategoryPoliticsAndSociety)},
			},
		},
		search.Result{
			ID:        "9f4c1847-2567-4de7-8861-34cf47a071ae",
			ClassName: "Article",
			Vector:    []float32{0, 1, 0},
			Schema: map[string]interface{}{
				"description":   "This articles talks about society",
				"exactCategory": models.MultipleRef{beaconRef(idCategorySociety)},
				"mainCategory":  models.MultipleRef{beaconRef(idMainCategoryPoliticsAndSociety)},
			},
		},
		search.Result{
			ID:        "926416ec-8fb1-4e40-ab8c-37b226b3d68e",
			ClassName: "Article",
			Vector:    []float32{0, 0, 1},
			Schema: map[string]interface{}{
				"description":   "This article talks about food",
				"exactCategory": models.MultipleRef{beaconRef(idCategoryFoodAndDrink)},
				"mainCategory":  models.MultipleRef{beaconRef(idMainCategoryFoodAndDrink)},
			},
		},
	}
}

func mustUUID() strfmt.UUID {
	id, err := uuid.NewV4()
	if err != nil {
		panic(err)
	}

	return strfmt.UUID(id.String())
}

func largeTestDataSize(size int) search.Results {
	out := make(search.Results, size)

	for i := range out {
		out[i] = search.Result{
			ID:        mustUUID(),
			ClassName: "Article",
			Vector:    []float32{0.02, 0, 0},
			Schema: map[string]interface{}{
				"description": "does not matter much",
			},
		}
	}
	return out
}

type fakeAuthorizer struct{}

func (f *fakeAuthorizer) Authorize(principal *models.Principal, verb, resource string) error {
	return nil
}

func beaconRef(target string) *models.SingleRef {
	beacon := fmt.Sprintf("weaviate://localhost/%s", target)
	return &models.SingleRef{Beacon: strfmt.URI(beacon)}
}

const (
	idMainCategoryPoliticsAndSociety = "39c6abe3-4bbe-4c4e-9e60-ca5e99ec6b4e"
	idMainCategoryFoodAndDrink       = "5a3d909a-4f0d-4168-8f5c-cd3074d1e79a"
	idCategoryPolitics               = "1b204f16-7da6-44fd-bbd2-8cc4a7414bc3"
	idCategorySociety                = "ec500f39-1dc9-4580-9bd1-55a8ea8e37a2"
	idCategoryFoodAndDrink           = "027b708a-31ca-43ea-9001-88bec864c79c"
)

func invertedConfig() *models.InvertedIndexConfig {
	return &models.InvertedIndexConfig{
		CleanupIntervalSeconds: 60,
	}
}
