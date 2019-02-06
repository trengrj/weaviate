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

// Package network provides the network graphql endpoint for Weaviate
package network

import (
	"github.com/creativesoftwarefdn/weaviate/graphqlapi/descriptions"
	network_introspect "github.com/creativesoftwarefdn/weaviate/graphqlapi/network/introspect"
	"github.com/creativesoftwarefdn/weaviate/graphqlapi/utils"
	"github.com/graphql-go/graphql"
)

func genNetworkWhereOperatorEnum() *graphql.Enum {
	enumConf := graphql.EnumConfig{
		Name: "NetworkWhereOperatorEnum",
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
	}

	return graphql.NewEnum(enumConf)
}

// This is a translation of the Prototype from JS to Go. In the prototype some filter elements are declared as global variables, this is recreated here.
func genGlobalNetworkFilterElements(filterContainer *utils.FilterContainer) {
	filterContainer.WeaviateNetworkWhereKeywordsInpObj = genWeaviateNetworkWhereNameKeywordsInpObj()
	filterContainer.WeaviateNetworkIntrospectPropertiesObjField = network_introspect.GenWeaviateNetworkIntrospectPropertiesObjField()
}

func genWeaviateNetworkWhereNameKeywordsInpObj() *graphql.InputObject {
	weaviateNetworkWhereNameKeywordsInpObj := graphql.NewInputObject(
		graphql.InputObjectConfig{
			Name: "WeaviateNetworkWhereNameKeywordsInpObj",
			Fields: graphql.InputObjectConfigFieldMap{
				"value": &graphql.InputObjectFieldConfig{
					Type:        graphql.String,
					Description: descriptions.WhereKeywordsValue,
				},
				"weight": &graphql.InputObjectFieldConfig{
					Type:        graphql.Float,
					Description: descriptions.WhereKeywordsWeight,
				},
			},
			Description: descriptions.WhereKeywordsInpObj,
		},
	)
	return weaviateNetworkWhereNameKeywordsInpObj
}
