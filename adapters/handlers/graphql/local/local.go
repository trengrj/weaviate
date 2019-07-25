//                           _       _
// __      _____  __ ___   ___  __ _| |_ ___
// \ \ /\ / / _ \/ _` \ \ / / |/ _` | __/ _ \
//  \ V  V /  __/ (_| |\ V /| | (_| | ||  __/
//   \_/\_/ \___|\__,_| \_/ |_|\__,_|\__\___|
//
//  Copyright © 2016 - 2019 Weaviate. All rights reserved.
//  LICENSE: https://github.com/semi-technologies/weaviate/blob/develop/LICENSE.md
//  DESIGN & CONCEPT: Bob van Luijt (@bobvanluijt)
//  CONTACT: hello@semi.technology
//

package local

import (
	"github.com/graphql-go/graphql"
	"github.com/semi-technologies/weaviate/adapters/handlers/graphql/local/aggregate"
	"github.com/semi-technologies/weaviate/adapters/handlers/graphql/local/explore"
	"github.com/semi-technologies/weaviate/adapters/handlers/graphql/local/fetch"
	"github.com/semi-technologies/weaviate/adapters/handlers/graphql/local/get"
	"github.com/semi-technologies/weaviate/adapters/handlers/graphql/local/getmeta"
	"github.com/semi-technologies/weaviate/entities/schema"
	"github.com/semi-technologies/weaviate/usecases/config"
	"github.com/semi-technologies/weaviate/usecases/network/common/peers"
	"github.com/sirupsen/logrus"
)

// Build the local queries from the database schema.
func Build(dbSchema *schema.Schema, peers peers.Peers, logger logrus.FieldLogger,
	config config.Config) (graphql.Fields, error) {
	getField, err := get.Build(dbSchema, peers, logger)
	if err != nil {
		return nil, err
	}
	getMetaField, err := getmeta.Build(dbSchema, config)
	if err != nil {
		return nil, err
	}
	aggregateField, err := aggregate.Build(dbSchema, config)
	if err != nil {
		return nil, err
	}
	fetchField := fetch.Build()
	exploreField := explore.Build()

	localFields := graphql.Fields{
		"Get":       getField,
		"GetMeta":   getMetaField,
		"Aggregate": aggregateField,
		"Fetch":     fetchField,
		"Explore":   exploreField,
	}

	return localFields, nil
}
