//                           _       _
// __      _____  __ ___   ___  __ _| |_ ___
// \ \ /\ / / _ \/ _` \ \ / / |/ _` | __/ _ \
//  \ V  V /  __/ (_| |\ V /| | (_| | ||  __/
//   \_/\_/ \___|\__,_| \_/ |_|\__,_|\__\___|
//
//  Copyright © 2016 - 2019 SeMI Holding B.V. (registered @ Dutch Chamber of Commerce no 75221632). All rights reserved.
//  LICENSE WEAVIATE OPEN SOURCE: https://www.semi.technology/playbook/playbook/contract-weaviate-OSS.html
//  LICENSE WEAVIATE ENTERPRISE: https://www.semi.technology/playbook/contract-weaviate-enterprise.html
//  CONCEPT: Bob van Luijt (@bobvanluijt)
//  CONTACT: hello@semi.technology
//

package test

import (
	"testing"

	"github.com/semi-technologies/weaviate/test/acceptance/helper"
	"github.com/stretchr/testify/assert"
)

func TestGetWithComplexFilter(t *testing.T) {
	t.Run("without filters <- this is the control", func(t *testing.T) {
		query := `
		{
				Get {
					Things {
						Airport {
							code
						}
					}
				}
		}
		`
		result := AssertGraphQL(t, helper.RootAuth, query)
		airports := result.Get("Get", "Things", "Airport").AsSlice()

		expected := []interface{}{
			map[string]interface{}{"code": "10000"},
			map[string]interface{}{"code": "20000"},
			map[string]interface{}{"code": "30000"},
			map[string]interface{}{"code": "40000"},
		}

		assert.ElementsMatch(t, expected, airports)
	})

	t.Run("with filters applied", func(t *testing.T) {
		query := `
		{
				Get {
					Things {
						Airport(where:{
							operator:And
							operands: [
								{
									operator: GreaterThan,
									valueInt: 600000,
									path:["inCity", "City", "population"]
								}
								{
									operator: Equal,
									valueString:"Germany"
									path:["inCity", "City", "inCountry", "Country", "name"]
								}
							]
						}){
							code
						}
					}
				}
		}
		`
		result := AssertGraphQL(t, helper.RootAuth, query)
		airports := result.Get("Get", "Things", "Airport").AsSlice()

		expected := []interface{}{
			map[string]interface{}{"code": "40000"},
		}

		assert.ElementsMatch(t, expected, airports)
	})

	// TODO: https://github.com/semi-technologies/weaviate/issues/949
	// t.Run("with or filters applied", func(t *testing.T) {
	// 	// this test was added to prevent a regression on the bugfix for gh-758

	// 	query := `
	// 		{
	// 				Meta {
	// 					Things {
	// 						City(where:{
	// 							operator:Or
	// 							operands:[{
	// 								valueString:"Amsterdam",
	// 								operator:Equal,
	// 								path:["name"]
	// 							}, {
	// 								valueString:"Berlin",
	// 								operator:Equal,
	// 								path:["name"]
	// 							}]
	// 						}) {
	// 							__typename
	// 							name {
	// 								__typename
	// 								count
	// 							}
	// 						}
	// 					}
	// 				}
	// 		}
	// 	`
	// 	result := AssertGraphQL(t, helper.RootAuth, query)
	// 	cityMeta := result.Get("Meta", "Things", "City").Result

	// 	expected := map[string]interface{}{
	// 		"__typename": "MetaCity",
	// 		"name": map[string]interface{}{
	// 			"__typename": "MetaCitynameObj",
	// 			"count":      json.Number("2"),
	// 		},
	// 	}

	// 	assert.Equal(t, expected, cityMeta)
	// })
}
