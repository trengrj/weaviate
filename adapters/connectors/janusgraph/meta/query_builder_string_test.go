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

package meta

import (
	"testing"

	"github.com/semi-technologies/weaviate/usecases/traverser"
)

func Test_QueryBuilder_StringProps(t *testing.T) {
	tests := testCases{
		testCase{
			name: "with only a string, with only topOccurrences.value",
			inputProps: []traverser.MetaProperty{
				traverser.MetaProperty{
					Name:                "name",
					StatisticalAnalyses: []traverser.StatisticalAnalysis{traverser.TopOccurrencesValue},
				},
			},
			expectedQuery: `
				.union(
						has("name").groupCount().by("name")
							.order(local).by(values, decr).limit(local, 3).project("topOccurrences").project("name")
				)
				.group().by(select(keys).unfold()).by(
					select(values).unfold().group()
					.by( select(keys).unfold())
					.by( select(values).unfold())
				)
			`,
		},

		testCase{
			name: "with only a string, with only both topOccurrences.value and .occurs",
			inputProps: []traverser.MetaProperty{
				traverser.MetaProperty{
					Name:                "name",
					StatisticalAnalyses: []traverser.StatisticalAnalysis{traverser.TopOccurrencesValue, traverser.TopOccurrencesOccurs},
				},
			},
			expectedQuery: `
				.union(
						has("name").groupCount().by("name")
							.order(local).by(values, decr).limit(local, 3).project("topOccurrences").project("name")
				)
				.group().by(select(keys).unfold()).by(
					select(values).unfold().group()
					.by( select(keys).unfold())
					.by( select(values).unfold())
				)
			`,
		},

		testCase{
			name: "with only a string, with all possible props",
			inputProps: []traverser.MetaProperty{
				traverser.MetaProperty{
					Name:                "name",
					StatisticalAnalyses: []traverser.StatisticalAnalysis{traverser.Type, traverser.Count, traverser.TopOccurrencesValue, traverser.TopOccurrencesOccurs},
				},
			},
			expectedQuery: `
				.union(
				    has("name").count().project("count").project("name"),
						has("name").groupCount().by("name")
							.order(local).by(values, decr).limit(local, 3).project("topOccurrences").project("name")
				)
				.group().by(select(keys).unfold()).by(
					select(values).unfold().group()
					.by( select(keys).unfold())
					.by( select(values).unfold())
				)
			`,
		},
	}

	tests.AssertQuery(t, nil)

}
