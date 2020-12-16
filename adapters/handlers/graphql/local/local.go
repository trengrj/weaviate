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

package local

import (
	"github.com/graphql-go/graphql"
	"github.com/semi-technologies/weaviate/adapters/handlers/graphql/local/aggregate"
	"github.com/semi-technologies/weaviate/adapters/handlers/graphql/local/explore"
	"github.com/semi-technologies/weaviate/adapters/handlers/graphql/local/get"
	"github.com/semi-technologies/weaviate/entities/schema"
	"github.com/semi-technologies/weaviate/usecases/config"
	"github.com/sirupsen/logrus"
)

// Build the local queries from the database schema.
func Build(dbSchema *schema.Schema, logger logrus.FieldLogger,
	config config.Config) (graphql.Fields, error) {
	getField, err := get.Build(dbSchema, logger)
	if err != nil {
		return nil, err
	}

	aggregateField, err := aggregate.Build(dbSchema, config)
	if err != nil {
		return nil, err
	}
	exploreField := explore.Build()

	localFields := graphql.Fields{
		"Get":       getField,
		"Aggregate": aggregateField,
		"Explore":   exploreField,
	}

	return localFields, nil
}
