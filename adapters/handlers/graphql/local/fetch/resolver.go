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

package fetch

import (
	"fmt"

	"github.com/creativesoftwarefdn/weaviate/adapters/handlers/graphql/local/common_filters"
	"github.com/creativesoftwarefdn/weaviate/entities/schema/kind"
	contextionary "github.com/creativesoftwarefdn/weaviate/database/schema_contextionary"
	"github.com/creativesoftwarefdn/weaviate/usecases/telemetry"
	"github.com/graphql-go/graphql"
)

// Resolver is a local interface that can be composed with other interfaces to
// form the overall GraphQL API main interface. All data-base connectors that
// want to support the GetMeta feature must implement this interface.
type Resolver interface {
	LocalFetchKindClass(info *Params) (interface{}, error)
	LocalFetchFuzzy(info []string) (interface{}, error)
}

// Contextionary is a local abstraction on the contextionary that needs to be
// provided to the graphQL API in order to resolve Local.Fetch queries.
type Contextionary interface {
	SchemaSearch(p contextionary.SearchParams) (contextionary.SearchResults, error)
	SafeGetSimilarWordsWithCertainty(word string, certainty float32) []string
}

// RequestsLog is a local abstraction on the RequestsLog that needs to be
// provided to the graphQL API in order to log Local.Fetch queries.
type RequestsLog interface {
	Register(requestType string, identifier string)
}

// Params to describe the Local->GetMeta->Kind->Class query. Will be passed to
// the individual connector methods responsible for resolving the GetMeta
// query.
type Params struct {
	Kind               kind.Kind
	PossibleClassNames contextionary.SearchResults
	Properties         []Property
}

// Property is a combination of possible names to use for the property as well
// as a match object to perform filtering actions in the db connector based on
// this property
type Property struct {
	PossibleNames contextionary.SearchResults
	Match         PropertyMatch
}

// PropertyMatch defines how in the db connector this property should be used
// as a filter
type PropertyMatch struct {
	Operator common_filters.Operator
	Value    *common_filters.Value
}

func makeResolveClass(kind kind.Kind) graphql.FieldResolveFn {
	return func(p graphql.ResolveParams) (interface{}, error) {
		resources, err := newResources(p.Source)
		if err != nil {
			return nil, err
		}

		where, err := parseWhere(p.Args, kind)
		if err != nil {
			return nil, fmt.Errorf("invalid where filter: %s", err)
		}

		possibleClasses, err := resources.contextionary.SchemaSearch(where.class)
		if err != nil {
			return nil, err
		}

		properties, err := addPossibleNamesToProperties(where.properties, resources.contextionary)
		if err != nil {
			return nil, err
		}

		params := &Params{
			Kind:               kind,
			PossibleClassNames: possibleClasses,
			Properties:         properties,
		}

		if len(possibleClasses.Results) == 0 {
			return nil, fmt.Errorf("the contextionary contains no close matches to " +
				"the provided class name. Try using different search terms or lowering the " +
				"desired certainty")
		}

		if len(properties) == 0 {
			return nil, fmt.Errorf("the contextionary contains no close matches to " +
				"the provided property name. Try using different search terms or lowering " +
				"the desired certainty")
		}
		go func() {
			resources.requestsLog.Register(telemetry.TypeGQL, telemetry.LocalQuery)
		}()

		return func() (interface{}, error) {
			return resources.resolver.LocalFetchKindClass(params)
		}, nil
	}
}

type resources struct {
	resolver      Resolver
	contextionary Contextionary
	requestsLog   RequestsLog
}

func newResources(s interface{}) (*resources, error) {
	source, ok := s.(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("expected source to be a map, but was %T", source)
	}

	resolver, ok := source["Resolver"].(Resolver)
	if !ok {
		return nil, fmt.Errorf("expected source to contain a usable Resolver, but was %#v", source)
	}

	contextionary, ok := source["Contextionary"].(Contextionary)
	if !ok {
		return nil, fmt.Errorf("expected source to contain a usable Contextionary, but was %#v", source)
	}

	requestsLog, ok := source["RequestsLog"].(RequestsLog)
	if !ok {
		return nil, fmt.Errorf("expected source to contain a usable RequestsLog, but was %#v", source)
	}

	return &resources{
		resolver:      resolver,
		contextionary: contextionary,
		requestsLog:   requestsLog,
	}, nil
}

func addPossibleNamesToProperties(whereProperties []whereProperty,
	contextionary Contextionary) ([]Property, error) {
	properties := make([]Property, len(whereProperties), len(whereProperties))
	for i, whereProp := range whereProperties {
		possibleNames, err := contextionary.SchemaSearch(whereProp.search)
		if err != nil {
			return nil, err
		}
		properties[i] = Property{
			PossibleNames: possibleNames,
			Match:         whereProp.match,
		}
	}

	return properties, nil
}

func resolveFuzzy(p graphql.ResolveParams) (interface{}, error) {
	resources, err := newResources(p.Source)
	if err != nil {
		return nil, err
	}

	args := extractFuzzyArgs(p)

	words := resources.contextionary.SafeGetSimilarWordsWithCertainty(args.value, args.certainty)

	res, err := resources.resolver.LocalFetchFuzzy(words)
	if err != nil {
		return nil, fmt.Errorf("could not perform fuzzy search in connector: %v", err)
	}

	return res, nil
}

type fuzzyArgs struct {
	value     string
	certainty float32
}

func extractFuzzyArgs(p graphql.ResolveParams) fuzzyArgs {
	var args fuzzyArgs

	// all args are required, so we don't need to check their existance
	args.value = p.Args["value"].(string)
	args.certainty = float32(p.Args["certainty"].(float64))

	return args
}
