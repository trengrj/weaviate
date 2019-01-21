package getmeta

import (
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
							"sum": 10000.0,
						},
					},
				},
			},
			expectedResults: []result{{
				pathToField:   []string{"GetMeta", "PeerA", "Things", "Car", "horsepower", "sum"},
				expectedValue: 10000.0,
			}},
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
