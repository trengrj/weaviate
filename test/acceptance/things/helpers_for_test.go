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

package test

import (
	"testing"

	"github.com/go-openapi/strfmt"
	"github.com/semi-technologies/weaviate/client/things"
	"github.com/semi-technologies/weaviate/entities/models"
	"github.com/semi-technologies/weaviate/test/acceptance/helper"
	testhelper "github.com/semi-technologies/weaviate/test/helper"
)

func assertCreateThing(t *testing.T, className string, schema map[string]interface{}) strfmt.UUID {
	params := things.NewThingsCreateParams().WithBody(
		&models.Thing{
			Class:  className,
			Schema: schema,
		})

	resp, err := helper.Client(t).Things.ThingsCreate(params, nil)

	var thingID strfmt.UUID

	// Ensure that the response is OK
	helper.AssertRequestOk(t, resp, err, func() {
		thingID = resp.Payload.ID
	})

	return thingID
}

func assertCreateThingWithID(t *testing.T, className string, id strfmt.UUID, schema map[string]interface{}) {
	params := things.NewThingsCreateParams().WithBody(
		&models.Thing{
			ID:     id,
			Class:  className,
			Schema: schema,
		})

	resp, err := helper.Client(t).Things.ThingsCreate(params, nil)

	// Ensure that the response is OK
	helper.AssertRequestOk(t, resp, err, nil)
}

func assertGetThing(t *testing.T, uuid strfmt.UUID) *models.Thing {
	getResp, err := helper.Client(t).Things.ThingsGet(things.NewThingsGetParams().WithID(uuid), nil)

	var thing *models.Thing

	helper.AssertRequestOk(t, getResp, err, func() {
		thing = getResp.Payload
	})

	return thing
}

func assertGetThingEventually(t *testing.T, uuid strfmt.UUID) *models.Thing {
	var (
		resp *things.ThingsGetOK
		err  error
	)

	checkThunk := func() interface{} {
		resp, err = helper.Client(t).Things.ThingsGet(things.NewThingsGetParams().WithID(uuid), nil)
		return err == nil
	}

	testhelper.AssertEventuallyEqual(t, true, checkThunk)

	var thing *models.Thing

	helper.AssertRequestOk(t, resp, err, func() {
		thing = resp.Payload
	})

	return thing
}
