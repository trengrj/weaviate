//                           _       _
// __      _____  __ ___   ___  __ _| |_ ___
// \ \ /\ / / _ \/ _` \ \ / / |/ _` | __/ _ \
//  \ V  V /  __/ (_| |\ V /| | (_| | ||  __/
//   \_/\_/ \___|\__,_| \_/ |_|\__,_|\__\___|
//
//  Copyright © 2016 - 2021 SeMI Technologies B.V. All rights reserved.
//
//  CONTACT: hello@semi.technology
//

package get

import (
	"fmt"

	"github.com/graphql-go/graphql"
	"github.com/semi-technologies/weaviate/adapters/handlers/graphql/descriptions"
	"github.com/semi-technologies/weaviate/entities/models"
	"github.com/semi-technologies/weaviate/entities/schema"
	"github.com/sirupsen/logrus"
)

type ModulesProvider interface {
	GetArguments(class *models.Class) map[string]*graphql.ArgumentConfig
	ExtractSearchParams(arguments map[string]interface{}) map[string]interface{}
}

// Build the Local.Get part of the graphql tree
func Build(schema *schema.Schema, logger logrus.FieldLogger,
	modulesProvider ModulesProvider) (*graphql.Field, error) {
	if len(schema.Objects.Classes) == 0 {
		return nil, fmt.Errorf("there are no Objects classes defined yet")
	}

	cb := newClassBuilder(schema, logger, modulesProvider)

	var err error
	var objects *graphql.Object
	if len(schema.Objects.Classes) > 0 {
		objects, err = cb.objects()
		if err != nil {
			return nil, err
		}
	}

	return &graphql.Field{
		Name:        "Get",
		Description: descriptions.GetObjects,
		Type:        objects,
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			// Does nothing; pass through the filters
			return p.Source, nil
		},
	}, nil
}
