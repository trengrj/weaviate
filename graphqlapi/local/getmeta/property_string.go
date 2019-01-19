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
package getmeta

import (
	"fmt"

	"github.com/creativesoftwarefdn/weaviate/graphqlapi/descriptions"
	"github.com/creativesoftwarefdn/weaviate/models"
	"github.com/graphql-go/graphql"
)

func stringPropertyFields(class *models.SemanticSchemaClass,
	property *models.SemanticSchemaClassProperty) *graphql.Object {
	getMetaPointingFields := graphql.Fields{
		"type": &graphql.Field{
			Name:        fmt.Sprintf("Meta%sType", class.Class),
			Description: descriptions.GetMetaPropertyTypeDesc,
			Type:        graphql.String,
		},

		"count": &graphql.Field{
			Name:        fmt.Sprintf("Meta%sCount", class.Class),
			Description: descriptions.GetMetaPropertyCountDesc,
			Type:        graphql.Int,
		},

		"topOccurrences": &graphql.Field{
			Name:        fmt.Sprintf("Meta%sTopOccurrences", class.Class),
			Description: descriptions.GetMetaPropertyTopOccurrencesDesc,
			Type:        graphql.NewList(stringTopOccurrences(class, property)),
			Args: graphql.FieldConfigArgument{
				"first": &graphql.ArgumentConfig{
					Description: descriptions.FirstDesc,
					Type:        graphql.Int,
				},
				"after": &graphql.ArgumentConfig{
					Description: descriptions.AfterDesc,
					Type:        graphql.Int,
				},
			},
		},
	}

	getMetaStringProperty := graphql.ObjectConfig{
		Name:        fmt.Sprintf("Meta%s%sObj", class.Class, property.Name),
		Fields:      getMetaPointingFields,
		Description: descriptions.GetMetaPropertyObjectDesc,
	}

	return graphql.NewObject(getMetaStringProperty)
}

func stringTopOccurrences(class *models.SemanticSchemaClass,
	property *models.SemanticSchemaClassProperty) *graphql.Object {
	getMetaPointingFields := graphql.Fields{
		"value": &graphql.Field{
			Name:        fmt.Sprintf("Meta%s%sTopOccurrencesValue", class.Class, property.Name),
			Description: descriptions.GetMetaPropertyTopOccurrencesValueDesc,
			Type:        graphql.String,
		},

		"occurs": &graphql.Field{
			Name:        fmt.Sprintf("Meta%s%sTopOccurrencesOccurs", class.Class, property.Name),
			Description: descriptions.GetMetaPropertyTopOccurrencesOccursDesc,
			Type:        graphql.Int,
		},
	}

	getMetaPointing := graphql.ObjectConfig{
		Name:        fmt.Sprintf("Meta%s%sTopOccurrencesObj", class.Class, property.Name),
		Fields:      getMetaPointingFields,
		Description: descriptions.GetMetaPropertyTopOccurrencesDesc,
	}

	return graphql.NewObject(getMetaPointing)
}
