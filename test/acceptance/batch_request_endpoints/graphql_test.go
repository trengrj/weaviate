package batch_request_endpoints

// Acceptance tests for the batch GraphQL endpoint

// There is a helper struct called GraphQLResult that helps to navigate through the output,
// a query generator and a few helper functions to access the GraphQL endpoint.
// See the end of this file for more details on how those work.

import (
	"fmt"
	"testing"

	graphql_client "github.com/creativesoftwarefdn/weaviate/client/graphql"
	"github.com/creativesoftwarefdn/weaviate/models"
	"github.com/creativesoftwarefdn/weaviate/test/acceptance/helper"
	"github.com/go-openapi/runtime"
	"github.com/stretchr/testify/assert"
)

// TODO: change this test to simulate a successful query response when the test dataset is implemented.

// Check if batch results are returned in the correct order by comparing result equality to predefined outcomes.
// This includes testing whether individual requests and the batch request are handled correctly
func TestBatchGraphQLResultsOrder(t *testing.T) {
	t.Parallel()

	queryOneName := "testQuery"
	queryTwoName := "testQuery2"
	expectedResult := "Syntax Error GraphQL request (1:1) Unexpected Name \"%s\"\n\n1: %s\n   ^\n"

	// perform the query
	gqlResponse, err := queryBatchEndpoint(t, helper.RootAuth)

	if err != nil {
		t.Fatal("The returned schema is not an JSON object")
	}
	// check if the batch response contains two batched responses
	assert.Equal(t, 2, len(gqlResponse))

	// check if the error message matches the expected outcome (and are therefore returned in the correct order)
	if len(gqlResponse) == 2 {
		responseOne := gqlResponse[0].Errors[0].Message
		responseTwo := gqlResponse[1].Errors[0].Message

		fullExpectedOutcomeOne := fmt.Sprintf(expectedResult, queryOneName, queryOneName)
		assert.Equal(t, fullExpectedOutcomeOne, responseOne)

		fullExpectedOutcomeTwo := fmt.Sprintf(expectedResult, queryTwoName, queryTwoName)
		assert.Equal(t, fullExpectedOutcomeTwo, responseTwo)
	}
}

// Helper functions
// TODO: change this to a successful query when the test dataset is implemented. Make sure to implement a query returning 3 or more elements.
// Perform a batch GraphQL query
func queryBatchEndpoint(t *testing.T, auth runtime.ClientAuthInfoWriterFunc) (models.GraphQLResponses, error) {
	var vars interface{} = nil
	query1 := &models.GraphQLQuery{OperationName: "testQuery", Query: "testQuery", Variables: vars}
	query2 := &models.GraphQLQuery{OperationName: "testQuery2", Query: "testQuery2", Variables: vars}

	queries := models.GraphQLQueries{query1, query2}

	params := graphql_client.NewWeaviateGraphqlBatchParams().WithBody(queries)
	response, err := helper.Client(t).Graphql.WeaviateGraphqlBatch(params, auth)

	if err != nil {
		return nil, err
	}

	return response.Payload, nil
}
