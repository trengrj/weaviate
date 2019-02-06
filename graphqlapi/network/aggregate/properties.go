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
package aggregate

import (
	"fmt"

	"github.com/creativesoftwarefdn/weaviate/graphqlapi/descriptions"
	"github.com/creativesoftwarefdn/weaviate/models"
	"github.com/graphql-go/graphql"
)

func numericPropertyFields(class *models.SemanticSchemaClass, property *models.SemanticSchemaClassProperty, prefix string) *graphql.Object {
	getMetaIntFields := graphql.Fields{
		"sum": &graphql.Field{
			Name:        fmt.Sprintf("%s%s%sSum", prefix, class.Class, property.Name),
			Description: descriptions.NetworkAggregateSum,
			Type:        graphql.Float,
		},
		"minimum": &graphql.Field{
			Name:        fmt.Sprintf("%s%s%sMinimum", prefix, class.Class, property.Name),
			Description: descriptions.NetworkAggregateMin,
			Type:        graphql.Float,
		},
		"maximum": &graphql.Field{
			Name:        fmt.Sprintf("%s%s%sMaximum", prefix, class.Class, property.Name),
			Description: descriptions.NetworkAggregateMax,
			Type:        graphql.Float,
		},
		"mean": &graphql.Field{
			Name:        fmt.Sprintf("%s%s%sMean", prefix, class.Class, property.Name),
			Description: descriptions.NetworkAggregateMean,
			Type:        graphql.Float,
		},
		"mode": &graphql.Field{
			Name:        fmt.Sprintf("%s%s%sMode", prefix, class.Class, property.Name),
			Description: descriptions.NetworkAggregateMode,
			Type:        graphql.Float,
		},
		"median": &graphql.Field{
			Name:        fmt.Sprintf("%s%s%sMedian", prefix, class.Class, property.Name),
			Description: descriptions.NetworkAggregateMedian,
			Type:        graphql.Float,
		},
		"count": &graphql.Field{
			Name:        fmt.Sprintf("%s%s%sCount", prefix, class.Class, property.Name),
			Description: descriptions.NetworkAggregateCount,
			Type:        graphql.Int,
		},
	}

	return graphql.NewObject(graphql.ObjectConfig{
		Name:        fmt.Sprintf("%s%s%sObj", prefix, class.Class, property.Name),
		Fields:      getMetaIntFields,
		Description: descriptions.NetworkAggregatePropertyObject,
	})
}

func nonNumericPropertyFields(class *models.SemanticSchemaClass,
	property *models.SemanticSchemaClassProperty, prefix string) *graphql.Object {
	getMetaPointingFields := graphql.Fields{
		"count": &graphql.Field{
			Name:        fmt.Sprintf("%s%sCount", prefix, class.Class),
			Description: descriptions.NetworkAggregateCount,
			Type:        graphql.Int,
		},
	}

	return graphql.NewObject(graphql.ObjectConfig{
		Name:        fmt.Sprintf("%s%s%sObj", prefix, class.Class, property.Name),
		Fields:      getMetaPointingFields,
		Description: descriptions.NetworkAggregatePropertyObject,
	})
}

func groupedByProperty(class *models.SemanticSchemaClass, peerName string) *graphql.Object {
	classProperties := graphql.Fields{
		"path": &graphql.Field{
			Description: descriptions.NetworkAggregateGroupedByGroupedByPath,
			Type:        graphql.NewList(graphql.String),
		},
		"value": &graphql.Field{
			Description: descriptions.NetworkAggregateGroupedByGroupedByValue,
			Type:        graphql.String,
		},
	}

	classPropertiesObj := graphql.NewObject(graphql.ObjectConfig{
		Name:        fmt.Sprintf("Aggregate%s%sGroupedByObj", peerName, class.Class),
		Fields:      classProperties,
		Description: descriptions.NetworkAggregateGroupedByObj,
	})

	return classPropertiesObj
}
