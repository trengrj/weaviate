//                           _       _
// __      _____  __ ___   ___  __ _| |_ ___
// \ \ /\ / / _ \/ _` \ \ / / |/ _` | __/ _ \
//  \ V  V /  __/ (_| |\ V /| | (_| | ||  __/
//   \_/\_/ \___|\__,_| \_/ |_|\__,_|\__\___|
//
//  Copyright © 2016 - 2019 Weaviate. All rights reserved.
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

// This file contains only a single test to verify that combining multiple
// props together works as intended and helpers for other tests. Each
// individual property type however, is also extensively tested, please see the
// test files for individual props with the format query_builder_<type>_test.go

func Test_QueryBuilder_MultipleProps(t *testing.T) {
	tests := testCases{
		testCase{
			name: "with multiple props",
			inputProps: []traverser.MetaProperty{
				traverser.MetaProperty{
					Name: "isCapital",
					StatisticalAnalyses: []traverser.StatisticalAnalysis{
						traverser.Count, traverser.TotalTrue, traverser.TotalFalse, traverser.PercentageTrue, traverser.PercentageFalse,
					},
				},
				traverser.MetaProperty{
					Name: "population",
					StatisticalAnalyses: []traverser.StatisticalAnalysis{
						traverser.Mean, traverser.Sum, traverser.Maximum, traverser.Minimum, traverser.Count,
					},
				},
			},
			expectedQuery: `
				.union(
					values("isCapital").union(
						count().project("count").project("isCapital"),
						groupCount().unfold().project("isCapital")
					),
				  values("population").union(
					  mean().project("mean").project("population"),
					  sum().project("sum").project("population"),
					  max().project("maximum").project("population"),
					  min().project("minimum").project("population"),
					  count().project("count").project("population")
					)
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

func Test_QueryBuilder_MultiplePropsWithFilter(t *testing.T) {
	tests := testCases{
		testCase{
			name: "with multiple props",
			inputProps: []traverser.MetaProperty{
				traverser.MetaProperty{
					Name: "isCapital",
					StatisticalAnalyses: []traverser.StatisticalAnalysis{
						traverser.Count, traverser.TotalTrue, traverser.TotalFalse, traverser.PercentageTrue, traverser.PercentageFalse,
					},
				},
				traverser.MetaProperty{
					Name: "population",
					StatisticalAnalyses: []traverser.StatisticalAnalysis{
						traverser.Mean, traverser.Sum, traverser.Maximum, traverser.Minimum, traverser.Count,
					},
				},
			},
			expectedQuery: `
			  .has("foo", eq("bar"))
				.union(
					values("isCapital").union(
						count().project("count").project("isCapital"),
						groupCount().unfold().project("isCapital")
					),
				  values("population").union(
					  mean().project("mean").project("population"),
					  sum().project("sum").project("population"),
					  max().project("maximum").project("population"),
					  min().project("minimum").project("population"),
					  count().project("count").project("population")
					)
				)
				.group().by(select(keys).unfold()).by(
					select(values).unfold().group()
					.by( select(keys).unfold())
					.by( select(values).unfold())
				)
			`,
		},
	}

	filter := &fakeFilterSource{
		queryToReturn: `.has("foo", eq("bar"))`,
	}

	tests.AssertQueryWithFilterSource(t, nil, filter)
}
