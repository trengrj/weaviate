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

package local_get

import (
	"fmt"
	"strings"

	"github.com/creativesoftwarefdn/weaviate/database/schema"
	"github.com/creativesoftwarefdn/weaviate/database/schema/kind"
	"github.com/creativesoftwarefdn/weaviate/graphqlapi/descriptions"
	"github.com/creativesoftwarefdn/weaviate/graphqlapi/local/common_filters"
	"github.com/creativesoftwarefdn/weaviate/models"
	"github.com/graphql-go/graphql"
)

// Build the Local.Get part of the graphql tree
func Build(dbSchema *schema.Schema) (*graphql.Field, error) {
	getKinds := graphql.Fields{}

	if len(dbSchema.Actions.Classes) == 0 && len(dbSchema.Things.Classes) == 0 {
		return nil, fmt.Errorf("there are no Actions or Things classes defined yet")
	}

	knownClasses := map[string]*graphql.Object{}

	if len(dbSchema.Actions.Classes) > 0 {
		localGetActions, err := buildGetClasses(dbSchema, kind.ACTION_KIND, dbSchema.Actions, &knownClasses)
		if err != nil {
			return nil, err
		}

		getKinds["Actions"] = &graphql.Field{
			Name:        "WeaviateLocalGetActions",
			Description: descriptions.LocalGetActionsDesc,
			Type:        localGetActions,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				fmt.Printf("- LocalGetActions (pass on Source)\n")
				// Does nothing; pass through the filters
				return p.Source, nil
			},
		}
	}

	if len(dbSchema.Things.Classes) > 0 {
		localGetThings, err := buildGetClasses(dbSchema, kind.THING_KIND, dbSchema.Things, &knownClasses)
		if err != nil {
			return nil, err
		}

		getKinds["Things"] = &graphql.Field{
			Name:        "WeaviateLocalGetThings",
			Description: descriptions.LocalGetThingsDesc,
			Type:        localGetThings,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				fmt.Printf("- LocalGetThings (pass on Source)\n")
				// Does nothing; pass through the filters
				return p.Source, nil
			},
		}
	}

	field := graphql.Field{
		Name:        "WeaviateLocalGet",
		Description: descriptions.LocalGetDesc,
		Type: graphql.NewObject(graphql.ObjectConfig{
			Name:        "WeaviateLocalGetObj",
			Fields:      getKinds,
			Description: descriptions.LocalGetObjDesc,
		}),
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			fmt.Printf("- LocalGet (extract resolver from source, parse filters )\n")
			resolver := p.Source.(map[string]interface{})["Resolver"].(Resolver)
			filters, err := common_filters.ExtractFilters(p.Args)

			if err != nil {
				return nil, err
			}

			return &filtersAndResolver{
				filters:  filters,
				resolver: resolver,
			}, nil
		},
	}

	return &field, nil
}

// Builds the classes below a Local -> Get -> (k kind.Kind)
func buildGetClasses(dbSchema *schema.Schema, k kind.Kind, semanticSchema *models.SemanticSchema, knownClasses *map[string]*graphql.Object) (*graphql.Object, error) {
	classFields := graphql.Fields{}
	kindName := strings.Title(k.Name())

	for _, class := range semanticSchema.Classes {
		classField, err := buildGetClass(dbSchema, k, class, knownClasses)
		if err != nil {
			return nil, fmt.Errorf("Could not build class for %s", class.Class)
		}
		classFields[class.Class] = classField
	}

	classes := graphql.NewObject(graphql.ObjectConfig{
		Name:        fmt.Sprintf("WeaviateLocalGet%ssObj", kindName),
		Fields:      classFields,
		Description: fmt.Sprintf(descriptions.LocalGetThingsActionsObjDesc, kindName),
	})

	return classes, nil
}
