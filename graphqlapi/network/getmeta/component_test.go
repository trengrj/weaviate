package getmeta

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/creativesoftwarefdn/weaviate/graphqlapi/test/helper"
	"github.com/creativesoftwarefdn/weaviate/models"
	"github.com/graphql-go/graphql"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type testCase struct {
	name            string
	query           string
	resolverReturn  interface{}
	expectedResults []result
}

type testCases []testCase

type result struct {
	pathToField   []string
	expectedValue interface{}
}

func TestNetworkGetMeta(t *testing.T) {

	tests := testCases{
		testCase{
			name:  "network get meta happy path",
			query: "{ GetMeta { PeerA { Things { Car { horsepower { sum }}}}}}",
			resolverReturn: map[string]interface{}{
				"Things": map[string]interface{}{
					"Car": map[string]interface{}{
						"horsepower": map[string]interface{}{
							"sum": json.Number("10000"),
						},
					},
				},
			},
			expectedResults: []result{{
				pathToField:   []string{"GetMeta", "PeerA", "Things", "Car", "horsepower", "sum"},
				expectedValue: 10000.0,
			}},
		},

		testCase{
			name: "with every possible json number field there is",
			query: `{ 
				GetMeta { 
					PeerA { 
						Things { 
							Car { 
								horsepower { 
									sum highest lowest average count
								}
								weight { 
									sum highest lowest average count
								}
								meta {
								  count
								}
								modelName {
									count
								  topOccurrences {
										occurs
									}
								}
								stillInProduction {
									totalTrue totalFalse percentageTrue percentageFalse count
								}
								MadeBy {
									count
								}
							}
						}
					}
				}
			}
				`,
			resolverReturn: map[string]interface{}{
				"Things": map[string]interface{}{
					"Car": map[string]interface{}{
						"horsepower": map[string]interface{}{
							"sum":     json.Number("10000"),
							"highest": json.Number("10000"),
							"lowest":  json.Number("10000"),
							"average": json.Number("10000"),
							"count":   json.Number("10000"),
						},
						"weight": map[string]interface{}{
							"sum":     json.Number("10000"),
							"highest": json.Number("10000"),
							"lowest":  json.Number("10000"),
							"average": json.Number("10000"),
							"count":   json.Number("10000"),
						},
						"stillInProduction": map[string]interface{}{
							"totalTrue":       json.Number("10000"),
							"totalFalse":      json.Number("10000"),
							"percentageTrue":  json.Number("10000"),
							"percentageFalse": json.Number("10000"),
							"count":           json.Number("10000"),
						},
						"meta": map[string]interface{}{
							"count": json.Number("10000"),
						},
						"MadeBy": map[string]interface{}{
							"count": json.Number("10000"),
						},
						"modelName": map[string]interface{}{
							"count": json.Number("10000"),
							"topOccurrences": []interface{}{
								map[string]interface{}{
									"occurs": json.Number("10000"),
								},
							},
						},
					},
				},
			},
			expectedResults: []result{
				{
					pathToField: []string{"GetMeta", "PeerA", "Things", "Car", "horsepower"},
					expectedValue: map[string]interface{}{
						"sum":     10000.0,
						"highest": 10000.0,
						"lowest":  10000.0,
						"average": 10000.0,
						"count":   10000,
					},
				},
				{
					pathToField: []string{"GetMeta", "PeerA", "Things", "Car", "weight"},
					expectedValue: map[string]interface{}{
						"sum":     10000.0,
						"highest": 10000.0,
						"lowest":  10000.0,
						"average": 10000.0,
						"count":   10000,
					},
				},
				{
					pathToField: []string{"GetMeta", "PeerA", "Things", "Car", "stillInProduction"},
					expectedValue: map[string]interface{}{
						"totalTrue":       10000,
						"totalFalse":      10000,
						"percentageTrue":  10000.0,
						"percentageFalse": 10000.0,
						"count":           10000,
					},
				},
				{
					pathToField:   []string{"GetMeta", "PeerA", "Things", "Car", "meta", "count"},
					expectedValue: 10000,
				},
				{
					pathToField:   []string{"GetMeta", "PeerA", "Things", "Car", "MadeBy", "count"},
					expectedValue: 10000,
				},
				{
					pathToField:   []string{"GetMeta", "PeerA", "Things", "Car", "modelName", "topOccurrences"},
					expectedValue: []interface{}{map[string]interface{}{"occurs": 10000}},
				},
				{
					pathToField:   []string{"GetMeta", "PeerA", "Things", "Car", "modelName", "count"},
					expectedValue: 10000,
				},
			},
		},
	}

	tests.Assert(t)
}

func (tests testCases) Assert(t *testing.T) {
	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			resolver := newMockResolver()

			resolverReturn := &models.GraphQLResponse{
				Data: map[string]models.JSONObject{
					"Local": map[string]interface{}{
						"GetMeta": testCase.resolverReturn,
					},
				},
			}

			resolver.On("ProxyGetMetaInstance", mock.AnythingOfType("Params")).
				Return(resolverReturn, nil).Once()

			result := resolver.AssertResolve(t, testCase.query)

			for _, expectedResult := range testCase.expectedResults {
				value := result.Get(expectedResult.pathToField...).Result

				assert.Equal(t, expectedResult.expectedValue, value)
			}
		})
	}
}

type mockResolver struct {
	helper.MockResolver
}

func newMockResolver() *mockResolver {
	peerA, err := New("PeerA", helper.CarSchema).PeerField()
	if err != nil {
		panic(fmt.Sprintf("could not build graphql test schema: %s", err))
	}

	peerField := &graphql.Field{
		Name: "Peers",
		Type: graphql.NewObject(graphql.ObjectConfig{
			Name:   "PeerAObj",
			Fields: graphql.Fields{"PeerA": peerA},
		}),
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			resolver, ok := p.Source.(map[string]interface{})["NetworkResolver"].(Resolver)
			if !ok {
				return nil, fmt.Errorf("source does not contain a NetworkResolver, but \n%#v", p.Source)
			}

			return resolver, nil
		},
	}

	mocker := &mockResolver{}
	mocker.RootFieldName = "GetMeta"
	mocker.RootField = peerField
	mocker.RootObject = map[string]interface{}{"NetworkResolver": Resolver(mocker)}
	return mocker
}

func (m *mockResolver) ProxyGetMetaInstance(params Params) (*models.GraphQLResponse, error) {
	args := m.Called(params)
	return args.Get(0).(*models.GraphQLResponse), args.Error(1)
}
