//                           _       _
// __      _____  __ ___   ___  __ _| |_ ___
// \ \ /\ / / _ \/ _` \ \ / / |/ _` | __/ _ \
//  \ V  V /  __/ (_| |\ V /| | (_| | ||  __/
//   \_/\_/ \___|\__,_| \_/ |_|\__,_|\__\___|
//
//  Copyright © 2016 - 2020 SeMI Technologies B.V. All rights reserved.
//
//  CONTACT: hello@semi.technology
//

package get

import (
	"fmt"

	"github.com/graphql-go/graphql"
	"github.com/semi-technologies/weaviate/adapters/handlers/graphql/descriptions"
	"github.com/semi-technologies/weaviate/entities/schema"
	"github.com/sirupsen/logrus"
)

// Build the Local.Get part of the graphql tree
func Build(schema *schema.Schema, logger logrus.FieldLogger) (*graphql.Field, error) {
	getKinds := graphql.Fields{}

	if len(schema.Actions.Classes) == 0 && len(schema.Things.Classes) == 0 {
		return nil, fmt.Errorf("there are no Actions or Things classes defined yet")
	}

	cb := newClassBuilder(schema, logger)

	if len(schema.Actions.Classes) > 0 {
		actions, err := cb.actions()
		if err != nil {
			return nil, err
		}

		getKinds["Actions"] = &graphql.Field{
			Name:        "GetActions",
			Description: descriptions.GetActions,
			Type:        actions,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				// Does nothing; pass through the filters
				return p.Source, nil
			},
		}
	}

	if len(schema.Things.Classes) > 0 {
		things, err := cb.things()
		if err != nil {
			return nil, err
		}

		getKinds["Things"] = &graphql.Field{
			Name:        "GetThings",
			Description: descriptions.GetThings,
			Type:        things,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				// Does nothing; pass through the filters
				return p.Source, nil
			},
		}
	}

	return &graphql.Field{
		Name:        "Get",
		Description: descriptions.Get,
		Type: graphql.NewObject(graphql.ObjectConfig{
			Name:        "GetObj",
			Fields:      getKinds,
			Description: descriptions.GetObj,
		}),
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			return p.Source, nil
		},
	}, nil
}
