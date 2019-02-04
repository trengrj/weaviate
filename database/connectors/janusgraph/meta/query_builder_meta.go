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
package meta

import (
	"fmt"

	"github.com/creativesoftwarefdn/weaviate/database/schema"
	"github.com/creativesoftwarefdn/weaviate/graphqlapi/local/getmeta"
	"github.com/creativesoftwarefdn/weaviate/gremlin"
)

const (
	// MetaProp ("meta") is special in that it can be used an any class
	// regardless of the the actual properties, to retrieve generic information
	// such as the count of all class instances.
	MetaProp schema.PropertyName = "meta"
)

func (b *Query) metaProp(prop getmeta.MetaProperty) (*gremlin.Query, error) {
	if len(prop.StatisticalAnalyses) != 1 {
		return nil, fmt.Errorf(
			"meta prop only supports exactly one statistical analysis prop 'count', but have: %#v",
			prop.StatisticalAnalyses)
	}

	analysis := prop.StatisticalAnalyses[0]
	if analysis != getmeta.Count {
		return nil, fmt.Errorf(
			"meta prop only supports statistical analysis prop 'count', but have '%s'", analysis)
	}

	q := gremlin.New().
		Union(b.metaCountQuery()).
		AsProjectBy(string(MetaProp))

	return q, nil
}

func (b *Query) metaCountQuery() *gremlin.Query {
	return gremlin.New().
		Count().
		AsProjectBy("count")
}
