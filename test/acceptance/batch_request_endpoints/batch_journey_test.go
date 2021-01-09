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

package batch_request_endpoints

import (
	"fmt"
	"testing"

	"github.com/go-openapi/strfmt"
	"github.com/google/uuid"
	"github.com/semi-technologies/weaviate/client/batch"
	"github.com/semi-technologies/weaviate/entities/models"
	"github.com/semi-technologies/weaviate/test/acceptance/helper"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func batchJourney(t *testing.T) {
	sourcesSize := 10
	targetsSize := 3
	var sources []*models.Object
	var targets []*models.Object

	t.Run("create some data", func(t *testing.T) {
		sources = make([]*models.Object, sourcesSize)
		for i := range sources {
			sources[i] = &models.Object{
				Class: "BulkTestSource",
				ID:    mustNewUUID(),
				Properties: map[string]interface{}{
					"name": fmt.Sprintf("source%d", i),
				},
			}
		}

		targets = make([]*models.Object, targetsSize)
		for i := range targets {
			targets[i] = &models.Object{
				Class: "BulkTest",
				ID:    mustNewUUID(),
				Properties: map[string]interface{}{
					"name": fmt.Sprintf("target%d", i),
				},
			}
		}
	})

	t.Run("import all data in batch", func(t *testing.T) {
		params := batch.NewBatchObjectsCreateParams().WithBody(
			batch.BatchObjectsCreateBody{
				Objects: append(sources, targets...),
			},
		)
		res, err := helper.Client(t).Batch.BatchObjectsCreate(params, nil)
		require.Nil(t, err)

		for _, elem := range res.Payload {
			assert.Nil(t, elem.Result.Errors)
		}
	})

	t.Run("set one cref each from each source to all targets", func(t *testing.T) {
		body := make([]*models.BatchReference, sourcesSize*targetsSize)
		for i := range sources {
			for j := range targets {
				index := i*targetsSize + j
				body[index] = &models.BatchReference{
					From: strfmt.URI(
						fmt.Sprintf("weaviate://localhost/BulkTestSource/%s/ref", sources[i].ID)),
					To: strfmt.URI(fmt.Sprintf("weaviate://localhost/%s", targets[j].ID)),
				}
			}
		}
		params := batch.NewBatchReferencesCreateParams().WithBody(body)
		res, err := helper.Client(t).Batch.BatchReferencesCreate(params, nil)
		require.Nil(t, err)

		for _, elem := range res.Payload {
			assert.Nil(t, elem.Result.Errors)
		}
	})

	t.Run("verify using GraphQL", func(t *testing.T) {
		result := AssertGraphQL(t, helper.RootAuth, `
		{  Get { BulkTestSource { ref { ... on BulkTest { name }  } } } }
		`)
		items := result.Get("Get", "BulkTestSource").AsSlice()
		assert.Len(t, items, sourcesSize)
		for _, obj := range items {
			refs := obj.(map[string]interface{})["ref"].([]interface{})
			assert.Len(t, refs, targetsSize)
		}
	})
}

func mustNewUUID() strfmt.UUID {
	return strfmt.UUID(uuid.New().String())
}

func Test_BugFlakyResultCountWithVectorSearch(t *testing.T) {
	className := "FlakyBugTestClass"

	// since this bug occurs only in around 1 in 25 cases, we run the test
	// multiple times to increase the chance we're running into it
	amount := 100
	for i := 0; i < amount; i++ {
		t.Run("create schema", func(t *testing.T) {
			createObjectClass(t, &models.Class{
				Class: className,
				Properties: []*models.Property{
					&models.Property{
						Name:     "title",
						DataType: []string{"string"},
					},
					&models.Property{
						Name:     "url",
						DataType: []string{"string"},
					},
					&models.Property{
						Name:     "wordCount",
						DataType: []string{"int"},
					},
				},
			})
		})

		t.Run("create and import some data", func(t *testing.T) {
			objects := []*models.Object{
				&models.Object{
					Class: className,
					Properties: map[string]interface{}{
						"title":     "article 1",
						"url":       "http://articles.local/my-article-1",
						"wordCount": 60,
					},
				},
				&models.Object{
					Class: className,
					Properties: map[string]interface{}{
						"title":     "article 2",
						"url":       "http://articles.local/my-article-2",
						"wordCount": 40,
					},
				},
				&models.Object{
					Class: className,
					Properties: map[string]interface{}{
						"title":     "article 3",
						"url":       "http://articles.local/my-article-3",
						"wordCount": 600,
					},
				},
			}

			params := batch.NewBatchObjectsCreateParams().WithBody(
				batch.BatchObjectsCreateBody{
					Objects: objects,
				},
			)
			res, err := helper.Client(t).Batch.BatchObjectsCreate(params, nil)
			require.Nil(t, err)

			for _, elem := range res.Payload {
				assert.Nil(t, elem.Result.Errors)
			}
		})

		t.Run("verify using GraphQL", func(t *testing.T) {
			result := AssertGraphQL(t, helper.RootAuth, fmt.Sprintf(`
		{  Get { %s(nearText: {concepts: ["news"]}, limit:7) {  
				wordCount title url 
		} } }
		`, className))
			items := result.Get("Get", className).AsSlice()
			assert.Len(t, items, 3)
		})

		t.Run("cleanup", func(t *testing.T) {
			deleteObjectClass(t, className)
		})
	}
}
