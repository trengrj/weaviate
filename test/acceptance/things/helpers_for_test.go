/*                          _       _
 *__      _____  __ ___   ___  __ _| |_ ___
 *\ \ /\ / / _ \/ _` \ \ / / |/ _` | __/ _ \
 * \ V  V /  __/ (_| |\ V /| | (_| | ||  __/
 *  \_/\_/ \___|\__,_| \_/ |_|\__,_|\__\___|
 *
 * Copyright © 2016 - 2019 Weaviate. All rights reserved.
 * LICENSE: https://github.com/creativesoftwarefdn/weaviate/blob/develop/LICENSE.md
 * DESIGN & CONCEPT: Bob van Luijt (@bobvanluijt)
 * CONTACT: hello@creativesoftwarefdn.org
 */
package test

import (
	"testing"

	"github.com/creativesoftwarefdn/weaviate/client/schema"
	"github.com/creativesoftwarefdn/weaviate/client/things"
	"github.com/creativesoftwarefdn/weaviate/models"
	"github.com/creativesoftwarefdn/weaviate/test/acceptance/helper"
	"github.com/go-openapi/strfmt"
)

func assertCreateThing(t *testing.T, className string, schema map[string]interface{}) strfmt.UUID {
	params := things.NewWeaviateThingsCreateParams().WithBody(things.WeaviateThingsCreateBody{
		Thing: &models.ThingCreate{
			AtContext: "http://example.org",
			AtClass:   className,
			Schema:    schema,
		},
		Async: false,
	})

	resp, _, err := helper.Client(t).Things.WeaviateThingsCreate(params)

	var thingID strfmt.UUID

	// Ensure that the response is OK
	helper.AssertRequestOk(t, resp, err, func() {
		thingID = resp.Payload.ThingID
	})

	return thingID
}

func assertGetThing(t *testing.T, uuid strfmt.UUID) *models.ThingGetResponse {
	getResp, err := helper.Client(t).Things.WeaviateThingsGet(things.NewWeaviateThingsGetParams().WithThingID(uuid))

	var thing *models.ThingGetResponse

	helper.AssertRequestOk(t, getResp, err, func() {
		thing = getResp.Payload
	})

	return thing
}

func assertGetSchema(t *testing.T) *schema.WeaviateSchemaDumpOKBody {
	getResp, err := helper.Client(t).Schema.WeaviateSchemaDump(schema.NewWeaviateSchemaDumpParams())
	var schema *schema.WeaviateSchemaDumpOKBody
	helper.AssertRequestOk(t, getResp, err, func() {
		schema = getResp.Payload
	})

	return schema
}

func assertClassInSchema(t *testing.T, schema *models.SemanticSchema, className string) *models.SemanticSchemaClass {
	for _, class := range schema.Classes {
		if class.Class == className {
			return class
		}
	}

	t.Fatalf("class %s not found in schema", className)
	return nil
}

func assertPropertyInClass(t *testing.T, class *models.SemanticSchemaClass, propertyName string) *models.SemanticSchemaClassProperty {
	for _, prop := range class.Properties {
		if prop.Name == propertyName {
			return prop
		}
	}

	t.Fatalf("property %s not found in class", propertyName)
	return nil
}
