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

	"github.com/creativesoftwarefdn/weaviate/adapters/connectors/janusgraph/fetch"
	"github.com/creativesoftwarefdn/weaviate/adapters/connectors/janusgraph/fetchfuzzy"
	graphqlfetch "github.com/creativesoftwarefdn/weaviate/adapters/handlers/graphql/local/fetch"
	"github.com/creativesoftwarefdn/weaviate/gremlin"
)

// LocalFetchKindClass based on GraphQL Query params
func (j *Janusgraph) LocalFetchKindClass(params *graphqlfetch.Params) (interface{}, error) {
	q, err := fetch.NewQuery(*params, &j.state, &j.schema).String()
	if err != nil {
		return nil, fmt.Errorf("could not build query: %s", err)
	}

	res, err := fetch.NewProcessor(j.client, params.Kind, "localhost", &j.state).
		Process(gremlin.New().Raw(q))
	return res, err
}

// LocalFetchFuzzy based on GraphQL Query params
func (j *Janusgraph) LocalFetchFuzzy(words []string) (interface{}, error) {
	q, err := fetchfuzzy.NewQuery(words, &j.state, &j.schema).String()
	if err != nil {
		return nil, fmt.Errorf("could not build query: %s", err)
	}

	res, err := fetchfuzzy.NewProcessor(j.client, "localhost", &j.state).
		Process(gremlin.New().Raw(q))
	return res, err
}
