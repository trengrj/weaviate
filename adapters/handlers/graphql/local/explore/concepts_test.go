/*                          _       _
 *__      _____  __ ___   ___  __ _| |_ ___
 *\ \ /\ / / _ \/ _` \ \ / / |/ _` | __/ _ \
 * \ V  V /  __/ (_| |\ V /| | (_| | ||  __/
 *  \_/\_/ \___|\__,_| \_/ |_|\__,_|\__\___|
 *
 * Copyright © 2016 - 2019 Weaviate. All rights reserved.
 * LICENSE: https://github.com/semi-technologies/weaviate/blob/develop/LICENSE.md
 * DESIGN & CONCEPT: Bob van Luijt (@bobvanluijt)
 * CONTACT: hello@semi.technology
 */
package explore

import (
	"testing"

	"github.com/semi-technologies/weaviate/usecases/traverser"
	"github.com/stretchr/testify/assert"
)

type testCase struct {
	name                      string
	query                     string
	expectedParamsToTraverser traverser.ExploreConceptsParams
	resolverReturn            []traverser.VectorSearchResult
	expectedResults           []result
}

type testCases []testCase

type result struct {
	pathToField   []string
	expectedValue interface{}
}

func Test_ResolveExploreConcepts(t *testing.T) {
	t.Parallel()

	tests := testCases{
		testCase{
			name: "Resolve Explore Concepts",
			query: `
			{
				Explore {
					Concepts(keywords: ["car", "best brand"]) {
						beacon className
					}
				}
			}`,
			expectedParamsToTraverser: traverser.ExploreConceptsParams{
				Values: []string{"car", "best brand"},
			},
			resolverReturn: []traverser.VectorSearchResult{
				traverser.VectorSearchResult{
					Beacon:    "weaviate://localhost/things/some-uuid",
					ClassName: "bestClass",
				},
			},
			expectedResults: []result{{
				pathToField: []string{"Explore", "Concepts"},
				expectedValue: []interface{}{
					map[string]interface{}{
						"beacon":    "weaviate://localhost/things/some-uuid",
						"className": "bestClass",
					},
				},
			}},
		},

		testCase{
			name: "with optional limit set",
			query: `
			{
				Explore {
					Concepts(keywords: ["car", "best brand"], limit: 17) {
						beacon className
					}
				}
			}`,
			expectedParamsToTraverser: traverser.ExploreConceptsParams{
				Values: []string{"car", "best brand"},
				Limit:  17,
			},
			resolverReturn: []traverser.VectorSearchResult{
				traverser.VectorSearchResult{
					Beacon:    "weaviate://localhost/things/some-uuid",
					ClassName: "bestClass",
				},
			},
			expectedResults: []result{{
				pathToField: []string{"Explore", "Concepts"},
				expectedValue: []interface{}{
					map[string]interface{}{
						"beacon":    "weaviate://localhost/things/some-uuid",
						"className": "bestClass",
					},
				},
			}},
		},

		testCase{
			name: "with moveTo set",
			query: `
			{
				Explore {
					Concepts(
						keywords: ["car", "best brand"]
						limit: 17
						moveTo: {
							keywords: ["mercedes"]
							force: 0.7
						}
						) {
						beacon className
					}
				}
			}`,
			expectedParamsToTraverser: traverser.ExploreConceptsParams{
				Values: []string{"car", "best brand"},
				Limit:  17,
				MoveTo: traverser.ExploreMove{
					Values: []string{"mercedes"},
					Force:  0.7,
				},
			},
			resolverReturn: []traverser.VectorSearchResult{
				traverser.VectorSearchResult{
					Beacon:    "weaviate://localhost/things/some-uuid",
					ClassName: "bestClass",
				},
			},
			expectedResults: []result{{
				pathToField: []string{"Explore", "Concepts"},
				expectedValue: []interface{}{
					map[string]interface{}{
						"beacon":    "weaviate://localhost/things/some-uuid",
						"className": "bestClass",
					},
				},
			}},
		},

		testCase{
			name: "with moveTo and moveAwayFrom set",
			query: `
			{
				Explore {
					Concepts(
						keywords: ["car", "best brand"]
						limit: 17
						moveTo: {
							keywords: ["mercedes"]
							force: 0.7
						}
						moveAwayFrom: {
							keywords: ["van"]
							force: 0.7
						}
						) {
						beacon className
					}
				}
			}`,
			expectedParamsToTraverser: traverser.ExploreConceptsParams{
				Values: []string{"car", "best brand"},
				Limit:  17,
				MoveTo: traverser.ExploreMove{
					Values: []string{"mercedes"},
					Force:  0.7,
				},
				MoveAwayFrom: traverser.ExploreMove{
					Values: []string{"van"},
					Force:  0.7,
				},
			},
			resolverReturn: []traverser.VectorSearchResult{
				traverser.VectorSearchResult{
					Beacon:    "weaviate://localhost/things/some-uuid",
					ClassName: "bestClass",
				},
			},
			expectedResults: []result{{
				pathToField: []string{"Explore", "Concepts"},
				expectedValue: []interface{}{
					map[string]interface{}{
						"beacon":    "weaviate://localhost/things/some-uuid",
						"className": "bestClass",
					},
				},
			}},
		},
	}

	tests.AssertExtraction(t)
}

func (tests testCases) AssertExtraction(t *testing.T) {
	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {

			resolver := newMockResolver()

			resolver.On("ExploreConcepts", testCase.expectedParamsToTraverser).
				Return(testCase.resolverReturn, nil).Once()

			result := resolver.AssertResolve(t, testCase.query)

			for _, expectedResult := range testCase.expectedResults {
				value := result.Get(expectedResult.pathToField...).Result

				assert.Equal(t, expectedResult.expectedValue, value)
			}
		})
	}
}

// func Test__Resolve_MissingOperator(t *testing.T) {
// 	query := `
// 			{
// 				Fetch {
// 					Things(where: {
// 						class: {
// 							name: "bestclass"
// 							certainty: 0.8
// 							keywords: [{value: "foo", weight: 0.9}]
// 						},
// 						properties: {
// 							name: "bestproperty"
// 							certainty: 0.8
// 							keywords: [{value: "bar", weight: 0.9}]
// 							valueString: "some-value"
// 						},
// 					}) {
// 						beacon certainty
// 					}
// 				}
// 			}`
// 	c11y := newEmptyContextionary()
// 	c11y.On("SchemaSearch", mock.Anything).Twice()
// 	resolver := newMockResolver(c11y)
// 	res := resolver.Resolve(query)
// 	require.Len(t, res.Errors, 1)
// 	assert.Equal(t,
// 		`Argument "where" has invalid value {class: {name: "bestclass", certainty: 0.8, keywords: `+
// 			`[{value: "foo", weight: 0.9}]}, properties: {name: "bestproperty", certainty: 0.8, keywords: `+
// 			`[{value: "bar", weight: 0.9}], valueString: "some-value"}}.`+"\n"+
// 			`In field "properties": In field "operator": Expected "WeaviateLocalFetchThingWhereOperatorEnum!", found null.`,
// 		res.Errors[0].Message)
// }
