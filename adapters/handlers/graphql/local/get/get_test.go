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

// These tests verify that the parameters to the resolver are properly extracted from a GraphQL query.

package get

import (
	"testing"

	"github.com/go-openapi/strfmt"
	test_helper "github.com/semi-technologies/weaviate/adapters/handlers/graphql/test/helper"
	"github.com/semi-technologies/weaviate/entities/filters"
	"github.com/semi-technologies/weaviate/entities/models"
	"github.com/semi-technologies/weaviate/entities/schema/kind"
	"github.com/semi-technologies/weaviate/usecases/projector"
	"github.com/semi-technologies/weaviate/usecases/sempath"
	"github.com/semi-technologies/weaviate/usecases/traverser"
	"github.com/stretchr/testify/assert"
)

func TestSimpleFieldParamsOK(t *testing.T) {
	t.Parallel()
	resolver := newMockResolver()
	expectedParams := traverser.GetParams{
		Kind:       kind.Object,
		ClassName:  "SomeAction",
		Properties: []traverser.SelectProperty{{Name: "intField", IsPrimitive: true}},
	}

	resolver.On("GetClass", expectedParams).
		Return(test_helper.EmptyList(), nil).Once()

	resolver.AssertResolve(t, "{ Get { SomeAction { intField } } }")
}

func TestExtractIntField(t *testing.T) {
	t.Parallel()

	resolver := newMockResolver()

	expectedParams := traverser.GetParams{
		Kind:       kind.Object,
		ClassName:  "SomeAction",
		Properties: []traverser.SelectProperty{{Name: "intField", IsPrimitive: true}},
	}

	resolver.On("GetClass", expectedParams).
		Return(test_helper.EmptyList(), nil).Once()

	query := "{ Get { SomeAction { intField } } }"
	resolver.AssertResolve(t, query)
}

func TestExtractGeoCoordinatesField(t *testing.T) {
	t.Parallel()

	resolver := newMockResolver()

	expectedParams := traverser.GetParams{
		Kind:       kind.Object,
		ClassName:  "SomeAction",
		Properties: []traverser.SelectProperty{{Name: "location", IsPrimitive: true}},
	}

	resolverReturn := []interface{}{
		map[string]interface{}{
			"location": &models.GeoCoordinates{Latitude: ptFloat32(0.5), Longitude: ptFloat32(0.6)},
		},
	}

	resolver.On("GetClass", expectedParams).
		Return(resolverReturn, nil).Once()

	query := "{ Get { SomeAction { location { latitude longitude } } } }"
	result := resolver.AssertResolve(t, query)

	expectedLocation := map[string]interface{}{
		"location": map[string]interface{}{
			"latitude":  float32(0.5),
			"longitude": float32(0.6),
		},
	}

	assert.Equal(t, expectedLocation, result.Get("Get", "SomeAction").Result.([]interface{})[0])
}

func TestExtractPhoneNumberField(t *testing.T) {
	// We need to explicitly test all cases of asking for just one sub-property
	// at a time, because the AST-parsing uses known fields of known props to
	// distinguish a complex primitive prop from a reference prop
	//
	// See "isPrimitive()" and "fieldNameIsOfObjectButNonReferenceType" in
	// class_builder_fields.go for more details

	type test struct {
		name           string
		query          string
		expectedParams traverser.GetParams
		resolverReturn interface{}
		expectedResult interface{}
	}

	tests := []test{
		test{
			name:  "with only input requested",
			query: "{ Get { SomeAction { phone { input } } } }",
			expectedParams: traverser.GetParams{
				Kind:       kind.Object,
				ClassName:  "SomeAction",
				Properties: []traverser.SelectProperty{{Name: "phone", IsPrimitive: true}},
			},
			resolverReturn: []interface{}{
				map[string]interface{}{
					"phone": &models.PhoneNumber{Input: "+49 171 1234567"},
				},
			},
			expectedResult: map[string]interface{}{
				"phone": map[string]interface{}{
					"input": "+49 171 1234567",
				},
			},
		},
		test{
			name:  "with only internationalFormatted requested",
			query: "{ Get { SomeAction { phone { internationalFormatted } } } }",
			expectedParams: traverser.GetParams{
				Kind:       kind.Object,
				ClassName:  "SomeAction",
				Properties: []traverser.SelectProperty{{Name: "phone", IsPrimitive: true}},
			},
			resolverReturn: []interface{}{
				map[string]interface{}{
					"phone": &models.PhoneNumber{InternationalFormatted: "+49 171 1234567"},
				},
			},
			expectedResult: map[string]interface{}{
				"phone": map[string]interface{}{
					"internationalFormatted": "+49 171 1234567",
				},
			},
		},
		test{
			name:  "with only nationalFormatted requested",
			query: "{ Get { SomeAction { phone { nationalFormatted } } } }",
			expectedParams: traverser.GetParams{
				Kind:       kind.Object,
				ClassName:  "SomeAction",
				Properties: []traverser.SelectProperty{{Name: "phone", IsPrimitive: true}},
			},
			resolverReturn: []interface{}{
				map[string]interface{}{
					"phone": &models.PhoneNumber{NationalFormatted: "0171 1234567"},
				},
			},
			expectedResult: map[string]interface{}{
				"phone": map[string]interface{}{
					"nationalFormatted": "0171 1234567",
				},
			},
		},
		test{
			name:  "with only national requested",
			query: "{ Get { SomeAction { phone { national } } } }",
			expectedParams: traverser.GetParams{
				Kind:       kind.Object,
				ClassName:  "SomeAction",
				Properties: []traverser.SelectProperty{{Name: "phone", IsPrimitive: true}},
			},
			resolverReturn: []interface{}{
				map[string]interface{}{
					"phone": &models.PhoneNumber{National: 0o1711234567},
				},
			},
			expectedResult: map[string]interface{}{
				"phone": map[string]interface{}{
					"national": 0o1711234567,
				},
			},
		},
		test{
			name:  "with only valid requested",
			query: "{ Get { SomeAction { phone { valid } } } }",
			expectedParams: traverser.GetParams{
				Kind:       kind.Object,
				ClassName:  "SomeAction",
				Properties: []traverser.SelectProperty{{Name: "phone", IsPrimitive: true}},
			},
			resolverReturn: []interface{}{
				map[string]interface{}{
					"phone": &models.PhoneNumber{Valid: true},
				},
			},
			expectedResult: map[string]interface{}{
				"phone": map[string]interface{}{
					"valid": true,
				},
			},
		},
		test{
			name:  "with only countryCode requested",
			query: "{ Get { SomeAction { phone { countryCode } } } }",
			expectedParams: traverser.GetParams{
				Kind:       kind.Object,
				ClassName:  "SomeAction",
				Properties: []traverser.SelectProperty{{Name: "phone", IsPrimitive: true}},
			},
			resolverReturn: []interface{}{
				map[string]interface{}{
					"phone": &models.PhoneNumber{CountryCode: 49},
				},
			},
			expectedResult: map[string]interface{}{
				"phone": map[string]interface{}{
					"countryCode": 49,
				},
			},
		},
		test{
			name:  "with only defaultCountry requested",
			query: "{ Get { SomeAction { phone { defaultCountry } } } }",
			expectedParams: traverser.GetParams{
				Kind:       kind.Object,
				ClassName:  "SomeAction",
				Properties: []traverser.SelectProperty{{Name: "phone", IsPrimitive: true}},
			},
			resolverReturn: []interface{}{
				map[string]interface{}{
					"phone": &models.PhoneNumber{DefaultCountry: "DE"},
				},
			},
			expectedResult: map[string]interface{}{
				"phone": map[string]interface{}{
					"defaultCountry": "DE",
				},
			},
		},
		test{
			name: "with multiple fields set",
			query: "{ Get { SomeAction { phone { input internationalFormatted " +
				"nationalFormatted defaultCountry national countryCode valid } } } }",
			expectedParams: traverser.GetParams{
				Kind:       kind.Object,
				ClassName:  "SomeAction",
				Properties: []traverser.SelectProperty{{Name: "phone", IsPrimitive: true}},
			},
			resolverReturn: []interface{}{
				map[string]interface{}{
					"phone": &models.PhoneNumber{
						DefaultCountry:         "DE",
						CountryCode:            49,
						NationalFormatted:      "0171 123456",
						InternationalFormatted: "+49 171 123456",
						National:               171123456,
						Input:                  "0171123456",
						Valid:                  true,
					},
				},
			},
			expectedResult: map[string]interface{}{
				"phone": map[string]interface{}{
					"defaultCountry":         "DE",
					"countryCode":            49,
					"nationalFormatted":      "0171 123456",
					"internationalFormatted": "+49 171 123456",
					"national":               171123456,
					"input":                  "0171123456",
					"valid":                  true,
				},
			},
		},
	}

	for _, test := range tests {
		resolver := newMockResolver()

		resolver.On("GetClass", test.expectedParams).
			Return(test.resolverReturn, nil).Once()
		result := resolver.AssertResolve(t, test.query)
		assert.Equal(t, test.expectedResult, result.Get("Get", "SomeAction").Result.([]interface{})[0])
	}
}

func TestExtractAdditionalFields(t *testing.T) {
	// We don't need to explicitly test every subselection as we did on
	// phoneNumber as these fields have fixed keys. So we can simply check for
	// the prop

	type test struct {
		name           string
		query          string
		expectedParams traverser.GetParams
		resolverReturn interface{}
		expectedResult interface{}
	}

	tests := []test{
		test{
			name:  "with _additional certainty",
			query: "{ Get { SomeAction { _additional { certainty } } } }",
			expectedParams: traverser.GetParams{
				Kind:      kind.Object,
				ClassName: "SomeAction",
				AdditionalProperties: traverser.AdditionalProperties{
					Certainty: true,
				},
			},
			resolverReturn: []interface{}{
				map[string]interface{}{
					"_additional": map[string]interface{}{
						"certainty": 0.69,
					},
				},
			},
			expectedResult: map[string]interface{}{
				"_additional": map[string]interface{}{
					"certainty": 0.69,
				},
			},
		},
		test{
			name:  "with _additional classification",
			query: "{ Get { SomeAction { _additional { classification { id completed classifiedFields scope basedOn }  } } } }",
			expectedParams: traverser.GetParams{
				Kind:      kind.Object,
				ClassName: "SomeAction",
				AdditionalProperties: traverser.AdditionalProperties{
					Classification: true,
				},
			},
			resolverReturn: []interface{}{
				map[string]interface{}{
					"_additional": &models.AdditionalProperties{
						Classification: &models.AdditionalPropertiesClassification{
							ID:               "12345",
							BasedOn:          []string{"primitiveProp"},
							Scope:            []string{"refprop1", "refprop2", "refprop3"},
							ClassifiedFields: []string{"refprop3"},
							Completed:        timeMust(strfmt.ParseDateTime("2006-01-02T15:04:05.000Z")),
						},
					},
				},
			},
			expectedResult: map[string]interface{}{
				"_additional": map[string]interface{}{
					"classification": map[string]interface{}{
						"id":               "12345",
						"basedOn":          []interface{}{"primitiveProp"},
						"scope":            []interface{}{"refprop1", "refprop2", "refprop3"},
						"classifiedFields": []interface{}{"refprop3"},
						"completed":        "2006-01-02T15:04:05.000Z",
					},
				},
			},
		},
		test{
			name:  "with _additional interpretation",
			query: "{ Get { SomeAction { _additional { interpretation { source { concept weight occurrence } }  } } } }",
			expectedParams: traverser.GetParams{
				Kind:      kind.Object,
				ClassName: "SomeAction",
				AdditionalProperties: traverser.AdditionalProperties{
					Interpretation: true,
				},
			},
			resolverReturn: []interface{}{
				map[string]interface{}{
					"_additional": map[string]interface{}{
						"interpretation": &models.Interpretation{
							Source: []*models.InterpretationSource{
								&models.InterpretationSource{
									Concept:    "foo",
									Weight:     0.6,
									Occurrence: 1200,
								},
								&models.InterpretationSource{
									Concept:    "bar",
									Weight:     0.9,
									Occurrence: 800,
								},
							},
						},
					},
				},
			},
			expectedResult: map[string]interface{}{
				"_additional": map[string]interface{}{
					"interpretation": map[string]interface{}{
						"source": []interface{}{
							map[string]interface{}{
								"concept":    "foo",
								"weight":     0.6,
								"occurrence": 1200,
							},
							map[string]interface{}{
								"concept":    "bar",
								"weight":     0.9,
								"occurrence": 800,
							},
						},
					},
				},
			},
		},
		test{
			name:  "with _additional nearestNeighbors",
			query: "{ Get { SomeAction { _additional { nearestNeighbors { neighbors { concept distance } }  } } } }",
			expectedParams: traverser.GetParams{
				Kind:      kind.Object,
				ClassName: "SomeAction",
				AdditionalProperties: traverser.AdditionalProperties{
					NearestNeighbors: true,
				},
			},
			resolverReturn: []interface{}{
				map[string]interface{}{
					"_additional": map[string]interface{}{
						"nearestNeighbors": &models.NearestNeighbors{
							Neighbors: []*models.NearestNeighbor{
								&models.NearestNeighbor{
									Concept:  "foo",
									Distance: 0.1,
								},
								&models.NearestNeighbor{
									Concept:  "bar",
									Distance: 0.2,
								},
							},
						},
					},
				},
			},
			expectedResult: map[string]interface{}{
				"_additional": map[string]interface{}{
					"nearestNeighbors": map[string]interface{}{
						"neighbors": []interface{}{
							map[string]interface{}{
								"concept":  "foo",
								"distance": float32(0.1),
							},
							map[string]interface{}{
								"concept":  "bar",
								"distance": float32(0.2),
							},
						},
					},
				},
			},
		},
		test{
			name:  "with _additional featureProjection without any optional parameters",
			query: "{ Get { SomeAction { _additional { featureProjection { vector }  } } } }",
			expectedParams: traverser.GetParams{
				Kind:      kind.Object,
				ClassName: "SomeAction",
				AdditionalProperties: traverser.AdditionalProperties{
					FeatureProjection: &projector.Params{
						Enabled: true,
					},
				},
			},
			resolverReturn: []interface{}{
				map[string]interface{}{
					"_additional": &models.AdditionalProperties{
						FeatureProjection: &models.FeatureProjection{
							Vector: []float32{0.0, 1.1, 2.2},
						},
					},
				},
			},
			expectedResult: map[string]interface{}{
				"_additional": map[string]interface{}{
					"featureProjection": map[string]interface{}{
						"vector": []interface{}{float32(0.0), float32(1.1), float32(2.2)},
					},
				},
			},
		},
		test{
			name:  "with _additional featureProjection with optional parameters",
			query: `{ Get { SomeAction { _additional { featureProjection(algorithm: "tsne", dimensions: 3, learningRate: 15, iterations: 100, perplexity: 10) { vector }  } } } }`,
			expectedParams: traverser.GetParams{
				Kind:      kind.Object,
				ClassName: "SomeAction",
				AdditionalProperties: traverser.AdditionalProperties{
					FeatureProjection: &projector.Params{
						Enabled:      true,
						Algorithm:    ptString("tsne"),
						Dimensions:   ptInt(3),
						Iterations:   ptInt(100),
						LearningRate: ptInt(15),
						Perplexity:   ptInt(10),
					},
				},
			},
			resolverReturn: []interface{}{
				map[string]interface{}{
					"_additional": map[string]interface{}{
						"featureProjection": &models.FeatureProjection{
							Vector: []float32{0.0, 1.1, 2.2},
						},
					},
				},
			},
			expectedResult: map[string]interface{}{
				"_additional": map[string]interface{}{
					"featureProjection": map[string]interface{}{
						"vector": []interface{}{float32(0.0), float32(1.1), float32(2.2)},
					},
				},
			},
		},
		test{
			name:  "with _additional semanticPath set",
			query: `{ Get { SomeAction { _additional { semanticPath { path { concept distanceToQuery distanceToResult distanceToPrevious distanceToNext } } } } } }`,
			expectedParams: traverser.GetParams{
				Kind:      kind.Object,
				ClassName: "SomeAction",
				AdditionalProperties: traverser.AdditionalProperties{
					SemanticPath: &sempath.Params{},
				},
			},
			resolverReturn: []interface{}{
				map[string]interface{}{
					"_additional": &models.AdditionalProperties{
						SemanticPath: &models.SemanticPath{
							Path: []*models.SemanticPathElement{
								&models.SemanticPathElement{
									Concept:            "foo",
									DistanceToNext:     ptFloat32(0.5),
									DistanceToPrevious: nil,
									DistanceToQuery:    0.1,
									DistanceToResult:   0.1,
								},
								&models.SemanticPathElement{
									Concept:            "bar",
									DistanceToPrevious: ptFloat32(0.5),
									DistanceToNext:     nil,
									DistanceToQuery:    0.1,
									DistanceToResult:   0.1,
								},
							},
						},
					},
				},
			},
			expectedResult: map[string]interface{}{
				"_additional": map[string]interface{}{
					"semanticPath": map[string]interface{}{
						"path": []interface{}{
							map[string]interface{}{
								"concept":            "foo",
								"distanceToNext":     float32(0.5),
								"distanceToPrevious": nil,
								"distanceToQuery":    float32(0.1),
								"distanceToResult":   float32(0.1),
							},
							map[string]interface{}{
								"concept":            "bar",
								"distanceToPrevious": float32(0.5),
								"distanceToNext":     nil,
								"distanceToQuery":    float32(0.1),
								"distanceToResult":   float32(0.1),
							},
						},
					},
				},
			},
		},
	}

	for _, test := range tests {
		resolver := newMockResolver()

		resolver.On("GetClass", test.expectedParams).
			Return(test.resolverReturn, nil).Once()
		result := resolver.AssertResolve(t, test.query)
		assert.Equal(t, test.expectedResult, result.Get("Get", "SomeAction").Result.([]interface{})[0])
	}
}

func TestNearTextRanker(t *testing.T) {
	t.Parallel()

	resolver := newMockResolver()

	t.Run("for actions", func(t *testing.T) {
		query := `{ Get { SomeAction(nearText: {
                concepts: ["c1", "c2", "c3"],
								moveTo: {
									concepts:["positive"],
									force: 0.5
								},
								moveAwayFrom: {
									concepts:["epic"],
									force: 0.25
								}
        			}) { intField } } }`

		expectedParams := traverser.GetParams{
			Kind:       kind.Object,
			ClassName:  "SomeAction",
			Properties: []traverser.SelectProperty{{Name: "intField", IsPrimitive: true}},
			NearText: &traverser.NearTextParams{
				Values: []string{"c1", "c2", "c3"},
				MoveTo: traverser.ExploreMove{
					Values: []string{"positive"},
					Force:  0.5,
				},
				MoveAwayFrom: traverser.ExploreMove{
					Values: []string{"epic"},
					Force:  0.25,
				},
			},
		}

		resolver.On("GetClass", expectedParams).
			Return([]interface{}{}, nil).Once()

		resolver.AssertResolve(t, query)
	})

	t.Run("for things with optional certainty set", func(t *testing.T) {
		query := `{ Get { SomeThing(nearText: {
                concepts: ["c1", "c2", "c3"],
								certainty: 0.4,
								moveTo: {
									concepts:["positive"],
									force: 0.5
								},
								moveAwayFrom: {
									concepts:["epic"],
									force: 0.25
								}
        			}) { intField } } }`

		expectedParams := traverser.GetParams{
			Kind:       kind.Object,
			ClassName:  "SomeThing",
			Properties: []traverser.SelectProperty{{Name: "intField", IsPrimitive: true}},
			NearText: &traverser.NearTextParams{
				Values:    []string{"c1", "c2", "c3"},
				Certainty: 0.4,
				MoveTo: traverser.ExploreMove{
					Values: []string{"positive"},
					Force:  0.5,
				},
				MoveAwayFrom: traverser.ExploreMove{
					Values: []string{"epic"},
					Force:  0.25,
				},
			},
		}
		resolver.On("GetClass", expectedParams).
			Return([]interface{}{}, nil).Once()

		resolver.AssertResolve(t, query)
	})
}

func TestNearVectorRanker(t *testing.T) {
	t.Parallel()

	resolver := newMockResolver()

	t.Run("for actions", func(t *testing.T) {
		query := `{ Get { SomeAction(nearVector: {
							  vector: [0.123, 0.984] 
        			}) { intField } } }`

		expectedParams := traverser.GetParams{
			Kind:       kind.Object,
			ClassName:  "SomeAction",
			Properties: []traverser.SelectProperty{{Name: "intField", IsPrimitive: true}},
			NearVector: &traverser.NearVectorParams{
				Vector: []float32{0.123, 0.984},
			},
		}

		resolver.On("GetClass", expectedParams).
			Return([]interface{}{}, nil).Once()

		resolver.AssertResolve(t, query)
	})

	t.Run("for things with optional certainty set", func(t *testing.T) {
		query := `{ Get { SomeThing(nearVector: {
							  vector: [0.123, 0.984] 
								certainty: 0.4
        			}) { intField } } }`

		expectedParams := traverser.GetParams{
			Kind:       kind.Object,
			ClassName:  "SomeThing",
			Properties: []traverser.SelectProperty{{Name: "intField", IsPrimitive: true}},
			NearVector: &traverser.NearVectorParams{
				Vector:    []float32{0.123, 0.984},
				Certainty: 0.4,
			},
		}
		resolver.On("GetClass", expectedParams).
			Return([]interface{}{}, nil).Once()

		resolver.AssertResolve(t, query)
	})
}

func TestExtractPagination(t *testing.T) {
	t.Parallel()

	resolver := newMockResolver()

	expectedParams := traverser.GetParams{
		Kind:       kind.Object,
		ClassName:  "SomeAction",
		Properties: []traverser.SelectProperty{{Name: "intField", IsPrimitive: true}},
		Pagination: &filters.Pagination{
			Limit: 10,
		},
	}

	resolver.On("GetClass", expectedParams).
		Return(test_helper.EmptyList(), nil).Once()

	query := "{ Get { SomeAction(limit: 10) { intField } } }"
	resolver.AssertResolve(t, query)
}

func TestExtractGroupParams(t *testing.T) {
	t.Parallel()

	resolver := newMockResolver()

	expectedParams := traverser.GetParams{
		Kind:       kind.Object,
		ClassName:  "SomeAction",
		Properties: []traverser.SelectProperty{{Name: "intField", IsPrimitive: true}},
		Group: &traverser.GroupParams{
			Strategy: "closest",
			Force:    0.3,
		},
	}

	resolver.On("GetClass", expectedParams).
		Return(test_helper.EmptyList(), nil).Once()

	query := "{ Get { SomeAction(group: {type: closest, force: 0.3}) { intField } } }"
	resolver.AssertResolve(t, query)
}

func TestGetRelation(t *testing.T) {
	t.Parallel()

	t.Run("without using custom fragments", func(t *testing.T) {
		resolver := newMockResolver()

		expectedParams := traverser.GetParams{
			Kind:      kind.Object,
			ClassName: "SomeAction",
			Properties: []traverser.SelectProperty{
				{
					Name:        "hasAction",
					IsPrimitive: false,
					Refs: []traverser.SelectClass{
						{
							ClassName: "SomeAction",
							RefProperties: []traverser.SelectProperty{
								{
									Name:        "intField",
									IsPrimitive: true,
								},
								{
									Name:        "hasAction",
									IsPrimitive: false,
									Refs: []traverser.SelectClass{
										{
											ClassName: "SomeAction",
											RefProperties: []traverser.SelectProperty{
												{
													Name:        "intField",
													IsPrimitive: true,
												},
											},
										},
									},
								},
							},
						},
					},
				},
			},
		}

		resolver.On("GetClass", expectedParams).
			Return(test_helper.EmptyList(), nil).Once()

		query := "{ Get { SomeAction { hasAction { ... on SomeAction { intField, hasAction { ... on SomeAction { intField } } } } } } }"
		resolver.AssertResolve(t, query)
	})

	t.Run("with a custom fragment one level deep", func(t *testing.T) {
		resolver := newMockResolver()

		expectedParams := traverser.GetParams{
			Kind:      kind.Object,
			ClassName: "SomeAction",
			Properties: []traverser.SelectProperty{
				{
					Name:        "hasAction",
					IsPrimitive: false,
					Refs: []traverser.SelectClass{
						{
							ClassName: "SomeAction",
							RefProperties: []traverser.SelectProperty{
								{
									Name:        "intField",
									IsPrimitive: true,
								},
							},
						},
					},
				},
			},
		}

		resolver.On("GetClass", expectedParams).
			Return(test_helper.EmptyList(), nil).Once()

		query := "fragment actionFragment on SomeAction { intField } { Get { SomeAction { hasAction { ...actionFragment } } } }"
		resolver.AssertResolve(t, query)
	})

	t.Run("with a custom fragment multiple levels deep", func(t *testing.T) {
		resolver := newMockResolver()

		expectedParams := traverser.GetParams{
			Kind:      kind.Object,
			ClassName: "SomeAction",
			Properties: []traverser.SelectProperty{
				{
					Name:        "hasAction",
					IsPrimitive: false,
					Refs: []traverser.SelectClass{
						{
							ClassName: "SomeAction",
							RefProperties: []traverser.SelectProperty{
								{
									Name:        "intField",
									IsPrimitive: true,
								},
								{
									Name:        "hasAction",
									IsPrimitive: false,
									Refs: []traverser.SelectClass{
										{
											ClassName: "SomeAction",
											RefProperties: []traverser.SelectProperty{
												{
													Name:        "intField",
													IsPrimitive: true,
												},
											},
										},
									},
								},
							},
						},
					},
				},
			},
		}

		resolver.On("GetClass", expectedParams).
			Return(test_helper.EmptyList(), nil).Once()

		query := `
			fragment innerFragment on SomeAction { intField }
			fragment actionFragment on SomeAction { intField hasAction { ...innerFragment } } 
			
			{ Get { SomeAction { hasAction { ...actionFragment } } } }`
		resolver.AssertResolve(t, query)
	})
}

func ptFloat32(in float32) *float32 {
	return &in
}

func timeMust(t strfmt.DateTime, err error) strfmt.DateTime {
	if err != nil {
		panic(err)
	}

	return t
}
