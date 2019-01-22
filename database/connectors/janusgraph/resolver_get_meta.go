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
package janusgraph

import (
	"github.com/creativesoftwarefdn/weaviate/database/connectors/janusgraph/filters"
	"github.com/creativesoftwarefdn/weaviate/database/connectors/janusgraph/meta"
	"github.com/creativesoftwarefdn/weaviate/graphqlapi/local/getmeta"
	"github.com/creativesoftwarefdn/weaviate/gremlin"
)

// LocalGetMeta based on GraphQL Query params
func (j *Janusgraph) LocalGetMeta(params *getmeta.Params) (interface{}, error) {
	className := j.state.GetMappedClassName(params.ClassName)
	q := gremlin.New().Raw(`g.V()`).HasString("classId", string(className))

	filterProvider := filters.New(params.Filters, &j.state)

	metaQuery, err := meta.NewQuery(params, &j.state, &j.schema, filterProvider).String()
	if err != nil {
		return nil, err
	}

	q = q.Raw(metaQuery)

	typeInfo, err := meta.NewTypeInspector(&j.schema).Process(params)
	if err != nil {
		return nil, err
	}

	return meta.NewProcessor(j.client).Process(q, typeInfo)
}
