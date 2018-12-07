package test

// Acceptance tests for actions

import (
	"github.com/creativesoftwarefdn/weaviate/client/actions"
	"github.com/creativesoftwarefdn/weaviate/models"
	"github.com/creativesoftwarefdn/weaviate/test/acceptance/helper"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCanAddAPropertyIndividually(t *testing.T) {
	t.Parallel()

	toPointToUuid := assertCreateAction(t, "TestAction", map[string]interface{}{})

	uuid := assertCreateAction(t, "TestActionTwo", map[string]interface{}{})

	// Verify that testCrefs is empty
	updatedAction := assertGetAction(t, uuid)
	updatedSchema := updatedAction.Schema.(map[string]interface{})
	assert.Nil(t, updatedSchema["testCrefs"])

	// Append a property reference
	wurl := helper.GetWeaviateURL()
	params := actions.NewWeaviateActionsPropertiesCreateParams().
		WithActionID(uuid).
		WithPropertyName("testCrefs").
		WithBody(&models.SingleRef{
			NrDollarCref: toPointToUuid,
			LocationURL:  &wurl,
			Type:         "Action",
		})

	updateResp, err := helper.Client(t).Actions.WeaviateActionsPropertiesCreate(params, helper.RootAuth)
	helper.AssertRequestOk(t, updateResp, err, nil)

	// Get the property again.
	updatedAction = assertGetAction(t, uuid)
	updatedSchema = updatedAction.Schema.(map[string]interface{})
	assert.NotNil(t, updatedSchema["testCrefs"])
}

func TestCanReplaceAllProperties(t *testing.T) {
	t.Parallel()

	toPointToUuidFirst := assertCreateAction(t, "TestAction", map[string]interface{}{})
	toPointToUuidLater := assertCreateAction(t, "TestAction", map[string]interface{}{})

	wurl := helper.GetWeaviateURL()
	uuid := assertCreateAction(t, "TestActionTwo", map[string]interface{}{
		"testCrefs": &models.MultipleRef{
			&models.SingleRef{
				NrDollarCref: toPointToUuidFirst,
				LocationURL:  &wurl,
				Type:         "Action",
			},
		},
	})

	// Verify that testCrefs is empty
	updatedAction := assertGetAction(t, uuid)
	updatedSchema := updatedAction.Schema.(map[string]interface{})
	assert.NotNil(t, updatedSchema["testCrefs"])

	// Replace
	params := actions.NewWeaviateActionsPropertiesUpdateParams().
		WithActionID(uuid).
		WithPropertyName("testCrefs").
		WithBody(models.MultipleRef{
			&models.SingleRef{
				NrDollarCref: toPointToUuidLater,
				LocationURL:  &wurl,
				Type:         "Action",
			},
		})

	updateResp, err := helper.Client(t).Actions.WeaviateActionsPropertiesUpdate(params, helper.RootAuth)
	helper.AssertRequestOk(t, updateResp, err, nil)

	// Get the property again.
	updatedAction = assertGetAction(t, uuid)
	updatedSchema = updatedAction.Schema.(map[string]interface{})
	assert.NotNil(t, updatedSchema["testCrefs"])
}

func TestRemovePropertyIndividually(t *testing.T) {
	t.Parallel()

	toPointToUuid := assertCreateAction(t, "TestAction", map[string]interface{}{})

	wurl := helper.GetWeaviateURL()
	uuid := assertCreateAction(t, "TestActionTwo", map[string]interface{}{
		"testCrefs": &models.MultipleRef{
			&models.SingleRef{
				NrDollarCref: toPointToUuid,
				LocationURL:  &wurl,
				Type:         "Action",
			},
		},
	})

	// Verify that testCrefs is not empty
	updatedAction := assertGetAction(t, uuid)
	updatedSchema := updatedAction.Schema.(map[string]interface{})
	assert.NotNil(t, updatedSchema["testCrefs"])

	// Append a property reference
	params := actions.NewWeaviateActionsPropertiesDeleteParams().
		WithActionID(uuid).
		WithPropertyName("testCrefs").
		WithBody(&models.SingleRef{
			NrDollarCref: toPointToUuid,
			LocationURL:  &wurl,
			Type:         "Action",
		})

	updateResp, err := helper.Client(t).Actions.WeaviateActionsPropertiesDelete(params, helper.RootAuth)
	helper.AssertRequestOk(t, updateResp, err, nil)

	// Get the property again.
	updatedAction = assertGetAction(t, uuid)
	updatedSchema = updatedAction.Schema.(map[string]interface{})
	assert.Nil(t, updatedSchema["testCrefs"])
}
