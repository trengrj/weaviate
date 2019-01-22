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
	"github.com/creativesoftwarefdn/weaviate/graphqlapi/local/getmeta"
	"github.com/creativesoftwarefdn/weaviate/gremlin"
)

func (b *Query) crefProp(prop getmeta.MetaProperty) (*gremlin.Query, error) {
	for _, analysis := range prop.StatisticalAnalyses {
		if analysis != getmeta.Count {
			continue
		}

		q := gremlin.New().
			Union(b.crefCountQuery(prop)).
			AsProjectBy(string(prop.Name))

		return q, nil
	}
	return nil, nil
}

func (b *Query) crefCountQuery(prop getmeta.MetaProperty) *gremlin.Query {
	return gremlin.New().
		OutEWithLabel(b.mappedPropertyName(b.params.ClassName, untitle(prop.Name))).Count().
		AsProjectBy("refcount", "count")
}
