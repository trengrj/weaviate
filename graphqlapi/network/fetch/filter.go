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
package network_fetch

import (
	"github.com/creativesoftwarefdn/weaviate/graphqlapi/descriptions"
	"github.com/creativesoftwarefdn/weaviate/graphqlapi/local/common_filters"
	"github.com/creativesoftwarefdn/weaviate/graphqlapi/utils"
	"github.com/graphql-go/graphql"
)

func thingsAndActionsFilterFields(filterContainer *utils.FilterContainer) graphql.InputObjectConfigFieldMap {
	networkFetchWhereInpObjPropertiesObj := whereInpObjPropertiesObj(filterContainer)
	networkFetchWhereInpObjClassInpObj := whereInpObjClassInpObj(filterContainer)

	networkFetchThingsAndActionsFilterFields := graphql.InputObjectConfigFieldMap{
		"class": &graphql.InputObjectFieldConfig{
			Type:        graphql.NewList(networkFetchWhereInpObjClassInpObj),
			Description: descriptions.WhereClass,
		},
		"properties": &graphql.InputObjectFieldConfig{
			Type:        graphql.NewList(networkFetchWhereInpObjPropertiesObj),
			Description: descriptions.WhereProperties,
		},
		"first": &graphql.InputObjectFieldConfig{
			Type:        graphql.Int,
			Description: descriptions.First,
		},
	}

	return networkFetchThingsAndActionsFilterFields
}

func whereInpObjPropertiesObj(filterContainer *utils.FilterContainer) *graphql.InputObject {
	filterPropertiesElements := common_filters.BuildNew("WeaviateNetworkFetch")

	// Remove path and operands fields as they are not required here
	delete(filterPropertiesElements, "path")
	delete(filterPropertiesElements, "operands")

	filterPropertiesElements["certainty"] = &graphql.InputObjectFieldConfig{
		Type:        graphql.Float,
		Description: descriptions.WhereCertainty,
	}
	filterPropertiesElements["name"] = &graphql.InputObjectFieldConfig{
		Type:        graphql.String,
		Description: descriptions.WhereName,
	}
	filterPropertiesElements["keywords"] = &graphql.InputObjectFieldConfig{
		Type:        graphql.NewList(wherePropertyWhereKeywordsInpObj()),
		Description: descriptions.WhereKeywords,
	}

	networkFetchWhereInpObjPropertiesObj := graphql.NewInputObject(
		graphql.InputObjectConfig{
			Name:        "WeaviateNetworkFetchWhereInpObjProperties",
			Fields:      filterPropertiesElements,
			Description: descriptions.WhereProperties,
		},
	)

	return networkFetchWhereInpObjPropertiesObj
}

func wherePropertyWhereKeywordsInpObj() *graphql.InputObject {
	outputObject := graphql.NewInputObject(
		graphql.InputObjectConfig{
			Name: "NetworkFetchWherePropertyWhereKeywordsInpObj",
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
	return outputObject
}

func whereInpObjClassInpObj(filterContainer *utils.FilterContainer) *graphql.InputObject {
	filterClassElements := graphql.InputObjectConfigFieldMap{
		"name": &graphql.InputObjectFieldConfig{
			Type:        graphql.String,
			Description: descriptions.WhereName,
		},
		"certainty": &graphql.InputObjectFieldConfig{
			Type:        graphql.Float,
			Description: descriptions.WhereCertainty,
		},
		"keywords": &graphql.InputObjectFieldConfig{
			Type:        graphql.NewList(filterContainer.WeaviateNetworkWhereKeywordsInpObj),
			Description: descriptions.WhereKeywords,
		},
		"first": &graphql.InputObjectFieldConfig{
			Type:        graphql.Int,
			Description: descriptions.First,
		},
	}

	networkFetchWhereInpObjClassInpObj := graphql.NewInputObject(
		graphql.InputObjectConfig{
			Name:        "WeaviateNetworkFetchWhereInpObjClassInpObj",
			Fields:      filterClassElements,
			Description: descriptions.WhereClass,
		},
	)
	return networkFetchWhereInpObjClassInpObj
}
