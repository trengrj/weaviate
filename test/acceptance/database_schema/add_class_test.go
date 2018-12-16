/*                          _       _
 *__      _____  __ ___   ___  __ _| |_ ___
 *\ \ /\ / / _ \/ _` \ \ / / |/ _` | __/ _ \
 * \ V  V /  __/ (_| |\ V /| | (_| | ||  __/
 *  \_/\_/ \___|\__,_| \_/ |_|\__,_|\__\___|
 *
 * Copyright © 2016 - 2018 Weaviate. All rights reserved.
 * LICENSE: https://github.com/creativesoftwarefdn/weaviate/blob/develop/LICENSE.md
 * AUTHOR: Bob van Luijt (bob@kub.design)
 * See www.creativesoftwarefdn.org for details
 * Contact: @CreativeSofwFdn / bob@kub.design
 */
package test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/creativesoftwarefdn/weaviate/client/schema"
	"github.com/creativesoftwarefdn/weaviate/models"
	"github.com/creativesoftwarefdn/weaviate/test/acceptance/helper"
)

func TestAddAndRemoveThingClass(t *testing.T) {
	t.Parallel()

	randomThingClassName := "YellowCars"

	// Ensure that this name is not in the schema yet.
	t.Log("Asserting that this class does not exist yet")
	assert.NotContains(t, GetThingClassNames(t), randomThingClassName)

	tc := &models.SemanticSchemaClass{
		Class: randomThingClassName,
	}

	t.Log("Creating class")
	params := schema.NewWeaviateSchemaThingsCreateParams().WithThingClass(tc)
	resp, err := helper.Client(t).Schema.WeaviateSchemaThingsCreate(params, helper.RootAuth)
	helper.AssertRequestOk(t, resp, err, nil)

	t.Log("Asserting that this class is now created")
	assert.Contains(t, GetThingClassNames(t), randomThingClassName)

	// Now clean up this class.
	t.Log("Remove the class")
	delParams := schema.NewWeaviateSchemaThingsDeleteParams().WithClassName(randomThingClassName)
	delResp, err := helper.Client(t).Schema.WeaviateSchemaThingsDelete(delParams, helper.RootAuth)
	helper.AssertRequestOk(t, delResp, err, nil)

	// And verify that the class does not exist anymore.
	assert.NotContains(t, GetThingClassNames(t), randomThingClassName)
}
