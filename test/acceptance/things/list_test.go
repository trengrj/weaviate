package test

// Acceptance tests for things.

import (
	//	"fmt"
	"testing"

	//	"sort"
	//	"time"

	//	"github.com/go-openapi/runtime"
	//	"github.com/go-openapi/strfmt"

	"github.com/stretchr/testify/assert"

	"github.com/creativesoftwarefdn/weaviate/client/things"
	"github.com/creativesoftwarefdn/weaviate/models"
	//	"github.com/creativesoftwarefdn/weaviate/schema"
	"github.com/creativesoftwarefdn/weaviate/test/acceptance/helper"
	//	"github.com/creativesoftwarefdn/weaviate/validation"
	//	connutils "github.com/creativesoftwarefdn/weaviate/connectors/utils"
)

// Test that we can properly list things.
// Create two things, and check that the list all contains them all.
func TestListAll(t *testing.T) {
	t.Parallel()

	params1 := things.NewWeaviateThingsCreateParams().WithBody(things.WeaviateThingsCreateBody{
		Thing: &models.ThingCreate{
			AtContext: "http://example.org",
			AtClass:   "TestThing",
			Schema:    map[string]interface{}{},
		},
	})
	resp1, _, err := helper.Client(t).Things.WeaviateThingsCreate(params1, helper.RootAuth)
	assert.Nil(t, err, "creation should succeed")

	params2 := things.NewWeaviateThingsCreateParams().WithBody(things.WeaviateThingsCreateBody{
		Thing: &models.ThingCreate{
			AtContext: "http://example.org",
			AtClass:   "TestThing",
			Schema:    map[string]interface{}{},
		},
	})
	resp2, _, err := helper.Client(t).Things.WeaviateThingsCreate(params2, helper.RootAuth)
	assert.Nil(t, err, "creation should succeed")

	listParams := things.NewWeaviateThingsListParams()
	resp, err := helper.Client(t).Things.WeaviateThingsList(listParams, helper.RootAuth)

	found1 := false
	found2 := false

	for _, thing := range resp.Payload.Things {
		if thing.ThingID == resp1.Payload.ThingID {
			assert.False(t, found1, "found double ID for thing 1!")
			found1 = true
		}

		if thing.ThingID == resp2.Payload.ThingID {
			assert.False(t, found2, "found double ID for thing 2!")
			found2 = true
		}
	}

	assert.True(t, found1, "Did not find thing 1")
	assert.True(t, found2, "Did not find thing 2")
}
