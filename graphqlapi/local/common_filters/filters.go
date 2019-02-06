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

// Package common_filters provides the filters for the graphql endpoint for Weaviate
package common_filters

import (
	"fmt"

	"github.com/creativesoftwarefdn/weaviate/graphqlapi/descriptions"
	"github.com/graphql-go/graphql"
)

// The filters common to Local->Get and Local->GetMeta queries.
func BuildNew(path string) graphql.InputObjectConfigFieldMap {
	commonFilters := graphql.InputObjectConfigFieldMap{
		"operator": &graphql.InputObjectFieldConfig{
			Type: graphql.NewEnum(graphql.EnumConfig{
				Name: fmt.Sprintf("%sWhereOperatorEnum", path),
				Values: graphql.EnumValueConfigMap{
					"And":              &graphql.EnumValueConfig{},
					"Or":               &graphql.EnumValueConfig{},
					"Equal":            &graphql.EnumValueConfig{},
					"Not":              &graphql.EnumValueConfig{},
					"NotEqual":         &graphql.EnumValueConfig{},
					"GreaterThan":      &graphql.EnumValueConfig{},
					"GreaterThanEqual": &graphql.EnumValueConfig{},
					"LessThan":         &graphql.EnumValueConfig{},
					"LessThanEqual":    &graphql.EnumValueConfig{},
				},
				Description: descriptions.WhereOperatorEnum,
			}),
			Description: descriptions.WhereOperator,
		},
		"path": &graphql.InputObjectFieldConfig{
			Type:        graphql.NewList(graphql.String),
			Description: descriptions.WherePath,
		},
		"valueInt": &graphql.InputObjectFieldConfig{
			Type:        graphql.Int,
			Description: descriptions.WhereValueInt,
		},
		"valueNumber": &graphql.InputObjectFieldConfig{
			Type:        graphql.Float,
			Description: descriptions.WhereValueNumber,
		},
		"valueBoolean": &graphql.InputObjectFieldConfig{
			Type:        graphql.Boolean,
			Description: descriptions.WhereValueBoolean,
		},
		"valueString": &graphql.InputObjectFieldConfig{
			Type:        graphql.String,
			Description: descriptions.WhereValueString,
		},
		"valueText": &graphql.InputObjectFieldConfig{
			Type:        graphql.String,
			Description: descriptions.WhereValueText,
		},
		"valueDate": &graphql.InputObjectFieldConfig{
			Type:        graphql.String,
			Description: descriptions.WhereValueString,
		},
	}

	// Recurse into the same time.
	commonFilters["operands"] = &graphql.InputObjectFieldConfig{
		Description: descriptions.WhereOperands,
		Type: graphql.NewList(graphql.NewInputObject(
			graphql.InputObjectConfig{
				Name:        fmt.Sprintf("%sWhereOperandsInpObj", path),
				Description: descriptions.WhereOperandsInpObj,
				Fields: (graphql.InputObjectConfigFieldMapThunk)(func() graphql.InputObjectConfigFieldMap {
					return commonFilters
				}),
			},
		)),
	}

	return commonFilters
}
