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

package traverser

import (
	"context"
	"fmt"
	"testing"

	"github.com/semi-technologies/weaviate/entities/schema/kind"
	"github.com/semi-technologies/weaviate/usecases/config"
	"github.com/sirupsen/logrus/hooks/test"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_ExploreConcepts(t *testing.T) {
	t.Run("with network exploration", func(t *testing.T) {

		authorizer := &fakeAuthorizer{}
		repo := &fakeRepo{}
		locks := &fakeLocks{}
		c11y := &fakeC11y{}
		logger, _ := test.NewNullLogger()
		vectorizer := &fakeVectorizer{}
		vectorSearcher := &fakeVectorSearcher{}
		vectorRepo := &fakeVectorRepo{}
		explorer := NewExplorer(vectorSearcher, vectorizer, vectorRepo)
		traverser := NewTraverser(&config.WeaviateConfig{}, locks, repo, c11y, logger, authorizer,
			vectorizer, vectorSearcher, explorer)
		params := ExploreParams{
			Values:  []string{"a search term", "another"},
			Network: true,
		}

		_, err := traverser.Explore(context.Background(), nil, params)
		assert.Equal(t, fmt.Errorf(
			"explorer: network exploration currently not supported"), err)
	})
	t.Run("with no movements set", func(t *testing.T) {

		authorizer := &fakeAuthorizer{}
		repo := &fakeRepo{}
		locks := &fakeLocks{}
		c11y := &fakeC11y{}
		logger, _ := test.NewNullLogger()
		vectorizer := &fakeVectorizer{}
		vectorSearcher := &fakeVectorSearcher{}
		vectorRepo := &fakeVectorRepo{}
		explorer := NewExplorer(vectorSearcher, vectorizer, vectorRepo)
		traverser := NewTraverser(&config.WeaviateConfig{}, locks, repo, c11y, logger, authorizer,
			vectorizer, vectorSearcher, explorer)
		params := ExploreParams{
			Values: []string{"a search term", "another"},
		}
		vectorSearcher.results = []VectorSearchResult{
			VectorSearchResult{
				ClassName: "BestClass",
				Kind:      kind.Thing,
				ID:        "123-456-789",
			},
			VectorSearchResult{
				ClassName: "AnAction",
				Kind:      kind.Action,
				ID:        "987-654-321",
			},
		}

		res, err := traverser.Explore(context.Background(), nil, params)
		require.Nil(t, err)
		assert.Equal(t, []VectorSearchResult{
			VectorSearchResult{
				ClassName: "BestClass",
				Kind:      kind.Thing,
				ID:        "123-456-789",
				Beacon:    "weaviate://localhost/things/123-456-789",
				Certainty: 0.5,
			},
			VectorSearchResult{
				ClassName: "AnAction",
				Kind:      kind.Action,
				ID:        "987-654-321",
				Beacon:    "weaviate://localhost/actions/987-654-321",
				Certainty: 0.5,
			},
		}, res)

		assert.Equal(t, []float32{1, 2, 3}, vectorSearcher.calledWithVector)
		assert.Equal(t, 20, vectorSearcher.calledWithLimit,
			"uses the default limit if not explicitly set")
	})

	t.Run("with minimum certainty set to 0.6", func(t *testing.T) {

		authorizer := &fakeAuthorizer{}
		repo := &fakeRepo{}
		locks := &fakeLocks{}
		c11y := &fakeC11y{}
		logger, _ := test.NewNullLogger()
		vectorizer := &fakeVectorizer{}
		vectorSearcher := &fakeVectorSearcher{}
		vectorRepo := &fakeVectorRepo{}
		explorer := NewExplorer(vectorSearcher, vectorizer, vectorRepo)
		traverser := NewTraverser(&config.WeaviateConfig{}, locks, repo, c11y, logger, authorizer,
			vectorizer, vectorSearcher, explorer)
		params := ExploreParams{
			Values:    []string{"a search term", "another"},
			Certainty: 0.6,
		}
		vectorSearcher.results = []VectorSearchResult{
			VectorSearchResult{
				ClassName: "BestClass",
				Kind:      kind.Thing,
				ID:        "123-456-789",
			},
			VectorSearchResult{
				ClassName: "AnAction",
				Kind:      kind.Action,
				ID:        "987-654-321",
			},
		}

		res, err := traverser.Explore(context.Background(), nil, params)
		require.Nil(t, err)
		assert.Equal(t, []VectorSearchResult{}, res, "empty result because certainty is not met")
		assert.Equal(t, []float32{1, 2, 3}, vectorSearcher.calledWithVector)
		assert.Equal(t, 20, vectorSearcher.calledWithLimit,
			"uses the default limit if not explicitly set")
	})

	t.Run("with movements set", func(t *testing.T) {

		authorizer := &fakeAuthorizer{}
		repo := &fakeRepo{}
		locks := &fakeLocks{}
		c11y := &fakeC11y{}
		logger, _ := test.NewNullLogger()
		vectorizer := &fakeVectorizer{}
		vectorSearcher := &fakeVectorSearcher{}
		vectorRepo := &fakeVectorRepo{}
		explorer := NewExplorer(vectorSearcher, vectorizer, vectorRepo)
		traverser := NewTraverser(&config.WeaviateConfig{}, locks, repo, c11y, logger, authorizer,
			vectorizer, vectorSearcher, explorer)
		params := ExploreParams{
			Limit:  100,
			Values: []string{"a search term", "another"},
			MoveTo: ExploreMove{
				Values: []string{"foo"},
				Force:  0.7,
			},
			MoveAwayFrom: ExploreMove{
				Values: []string{"bar"},
				Force:  0.7,
			},
		}
		vectorSearcher.results = []VectorSearchResult{
			VectorSearchResult{
				ClassName: "BestClass",
				Kind:      kind.Thing,
				ID:        "123-456-789",
			},
			VectorSearchResult{
				ClassName: "AnAction",
				Kind:      kind.Action,
				ID:        "987-654-321",
			},
		}

		res, err := traverser.Explore(context.Background(), nil, params)
		require.Nil(t, err)
		assert.Equal(t, []VectorSearchResult{
			VectorSearchResult{
				ClassName: "BestClass",
				Kind:      kind.Thing,
				ID:        "123-456-789",
				Beacon:    "weaviate://localhost/things/123-456-789",
				Certainty: 0.5,
			},
			VectorSearchResult{
				ClassName: "AnAction",
				Kind:      kind.Action,
				ID:        "987-654-321",
				Beacon:    "weaviate://localhost/actions/987-654-321",
				Certainty: 0.5,
			},
		}, res)

		// see dummy implemenation of MoveTo and MoveAway for why the vector should
		// be the way it is
		assert.Equal(t, []float32{1.5, 2.5, 3.5}, vectorSearcher.calledWithVector)
		assert.Equal(t, 100, vectorSearcher.calledWithLimit,
			"limit explicitly set")
	})
}
