package test

import (
	"github.com/creativesoftwarefdn/weaviate/test/acceptance/helper"
	"testing"
)

// Helper function to get all the names of Thing classes.
func GetThingClassNames(t *testing.T) []string {
	resp, err := helper.Client(t).Schema.WeaviateSchemaDump(nil, helper.RootAuth)
	var names []string

	// Extract all names
	helper.AssertRequestOk(t, resp, err, func() {
		for _, class := range resp.Payload.Things.Classes {
			names = append(names, class.Class)
		}
	})

	return names
}

// Helper function to get all the names of Action classes.
func GetActionClassNames(t *testing.T) []string {
	resp, err := helper.Client(t).Schema.WeaviateSchemaDump(nil, helper.RootAuth)
	var names []string

	// Extract all names
	helper.AssertRequestOk(t, resp, err, func() {
		for _, class := range resp.Payload.Actions.Classes {
			names = append(names, class.Class)
		}
	})

	return names
}
