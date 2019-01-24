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
package aggregate

import (
	"testing"

	"github.com/creativesoftwarefdn/weaviate/database/schema"
	ag "github.com/creativesoftwarefdn/weaviate/graphqlapi/local/aggregate"
	cf "github.com/creativesoftwarefdn/weaviate/graphqlapi/local/common_filters"
)

// These tests only assert that multiple props work together correctly. See the
// tests for the individual property types for more detailed tests.

func Test_QueryBuilder_MultipleProps(t *testing.T) {
	tests := testCases{
		testCase{
			name: "counting a stirng prop, grouped by a primitive prop",
			inputProps: []ag.Property{
				ag.Property{
					Name:        "name",
					Aggregators: []ag.Aggregator{ag.Count},
				},
				ag.Property{
					Name:        "population",
					Aggregators: []ag.Aggregator{ag.Count},
				},
			},
			inputGroupBy: &cf.Path{
				Class:    schema.ClassName("City"),
				Property: schema.PropertyName("isCapital"),
			},
			expectedQuery: `
				.group().by("isCapital").by(
					fold()
						.match(
							__.as("a").unfold().values("name").count().as("name__count"),
							__.as("a").unfold().values("population").count().as("population__count")
						)
						.select("name__count").by(project("name__count")).as("name")
						.select("population__count").by(project("population__count")).as("population")
						.select("name", "population")
					)
				`,
		},
	}

	tests.AssertQuery(t, nil)

}
