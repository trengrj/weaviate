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

package janusgraph

import (
	"fmt"

	"github.com/creativesoftwarefdn/weaviate/database/connectors/janusgraph/fetch"
	graphqlfetch "github.com/creativesoftwarefdn/weaviate/graphqlapi/local/fetch"
	"github.com/creativesoftwarefdn/weaviate/gremlin"
	"github.com/davecgh/go-spew/spew"
)

// LocalFetchKindClass based on GraphQL Query params
func (j *Janusgraph) LocalFetchKindClass(params *graphqlfetch.Params) (interface{}, error) {
	q, err := fetch.NewQuery(*params, &j.state, &j.schema).String()
	if err != nil {
		return nil, fmt.Errorf("could not build query: %s", err)
	}

	res, err := fetch.NewProcessor(j.client, params.Kind, "localhost").
		Process(gremlin.New().Raw(q))
	spew.Dump(res)
	return res, err
}
