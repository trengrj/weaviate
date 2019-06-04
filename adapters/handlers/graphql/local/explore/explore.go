/*                          _       _
 *__      _____  __ ___   ___  __ _| |_ ___
 *\ \ /\ / / _ \/ _` \ \ / / |/ _` | __/ _ \
 * \ V  V /  __/ (_| |\ V /| | (_| | ||  __/
 *  \_/\_/ \___|\__,_| \_/ |_|\__,_|\__\___|
 *
 * Copyright © 2016 - 2019 Weaviate. All rights reserved.
 * LICENSE: https://github.com/semi-technologies/weaviate/blob/develop/LICENSE.md
 * DESIGN & CONCEPT: Bob van Luijt (@bobvanluijt)
 * CONTACT: hello@semi.technology
 */

// Package explore provides the Local->Explore GraphQL API
package explore

import (
	"github.com/graphql-go/graphql"
)

// Build builds the object containing the Local->Explore Fields, such as Things/Actions
func Build() *graphql.Field {
	return &graphql.Field{
		Name: "WeaviateLocalExplore",
		// Description: descriptions.LocalExplore,
		Type:    exploreObj(),
		Resolve: bubbleUpResolver,
	}
}

func exploreObj() *graphql.Object {

	fields := graphql.Fields{
		"Concepts": &graphql.Field{
			Name: "WeaviateLocalExplore",
			// Description: descriptions.NetworkExploreConcepts,
			Type: graphql.NewList(conceptsFieldsObj()),
			Args: graphql.FieldConfigArgument{
				"values": &graphql.ArgumentConfig{
					// Description: descriptions.ExploreConceptsValue,
					Type: graphql.NewNonNull(graphql.NewList(graphql.String)),
				},
			},
			Resolve: resolveConcepts,
		},
	}

	return graphql.NewObject(graphql.ObjectConfig{
		Name:   "WeaviateLocalExploreObj",
		Fields: fields,
		// Description: descriptions.LocalExploreObj,
	})
}

func bubbleUpResolver(p graphql.ResolveParams) (interface{}, error) {
	return p.Source, nil
}
