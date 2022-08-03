//                           _       _
// __      _____  __ ___   ___  __ _| |_ ___
// \ \ /\ / / _ \/ _` \ \ / / |/ _` | __/ _ \
//  \ V  V /  __/ (_| |\ V /| | (_| | ||  __/
//   \_/\_/ \___|\__,_| \_/ |_|\__,_|\__\___|
//
//  Copyright © 2016 - 2022 SeMI Technologies B.V. All rights reserved.
//
//  CONTACT: hello@semi.technology
//

package vectorizer

// TODO: This entire package should be part of the text2vec-contextionary
// module, if methods/objects in here are used from non-modular code, they
// probably shouldn't be in here

import (
	"context"
	"fmt"
	"strings"

	"github.com/fatih/camelcase"
	"github.com/semi-technologies/weaviate/entities/models"
	txt2vecmodels "github.com/semi-technologies/weaviate/modules/text2vec-contextionary/additional/models"
)

// Vectorizer turns objects into vectors
type Vectorizer struct {
	client client
}

type ErrNoUsableWords struct {
	Err error
}

func (e ErrNoUsableWords) Error() string {
	return e.Err.Error()
}

func NewErrNoUsableWordsf(pattern string, args ...interface{}) ErrNoUsableWords {
	return ErrNoUsableWords{Err: fmt.Errorf(pattern, args...)}
}

type client interface {
	VectorForCorpi(ctx context.Context, corpi []string,
		overrides map[string]string) ([]float32, []txt2vecmodels.InterpretationSource, error)
}

// IndexCheck returns whether a property of a class should be indexed
type ClassIndexCheck interface {
	PropertyIndexed(property string) bool
	VectorizeClassName() bool
	VectorizePropertyName(propertyName string) bool
}

// New from c11y client
func New(client client) *Vectorizer {
	return &Vectorizer{client}
}

// Object object to vector
func (v *Vectorizer) Object(ctx context.Context, object *models.Object,
	icheck ClassIndexCheck,
) error {
	var overrides map[string]string
	if object.VectorWeights != nil {
		overrides = object.VectorWeights.(map[string]string)
	}

	vec, sources, err := v.object(ctx, object.Class, object.Properties, overrides,
		icheck)
	if err != nil {
		return err
	}

	object.Vector = vec

	if object.Additional == nil {
		object.Additional = models.AdditionalProperties{}
	}

	object.Additional["interpretation"] = &txt2vecmodels.Interpretation{
		Source: sourceFromInputElements(sources),
	}

	return nil
}

func appendPropIfText(icheck ClassIndexCheck, list *[]string, propName string,
	value interface{},
) {
	valueString, ok := value.(string)
	if ok {
		if icheck.VectorizePropertyName(propName) {
			// use prop and value
			*list = append(*list, strings.ToLower(
				fmt.Sprintf("%s %s", camelCaseToLower(propName), valueString)))
		} else {
			*list = append(*list, strings.ToLower(valueString))
		}
	}
}

func (v *Vectorizer) object(ctx context.Context, className string,
	schema interface{}, overrides map[string]string,
	icheck ClassIndexCheck,
) ([]float32, []txt2vecmodels.InterpretationSource, error) {
	var corpi []string

	if icheck.VectorizeClassName() {
		corpi = append(corpi, camelCaseToLower(className))
	}

	if schema != nil {
		for prop, value := range schema.(map[string]interface{}) {
			if !icheck.PropertyIndexed(prop) {
				continue
			}

			if asSlice, ok := value.([]interface{}); ok {
				for _, elem := range asSlice {
					appendPropIfText(icheck, &corpi, prop, elem)
				}
			} else {
				appendPropIfText(icheck, &corpi, prop, value)
			}
		}
	}

	if len(corpi) == 0 {
		// fall back to using the class name
		corpi = append(corpi, camelCaseToLower(className))
	}

	vector, ie, err := v.client.VectorForCorpi(ctx, []string{strings.Join(corpi, " ")}, overrides)
	if err != nil {
		switch err.(type) {
		case ErrNoUsableWords:
			return nil, nil, fmt.Errorf("The object is invalid, as weaviate could not extract "+
				"any contextionary-valid words from it. This is the case when you have "+
				"set the options 'vectorizeClassName: false' and 'vectorizePropertyName: false' in this class' schema definition "+
				"and not a single property's value "+
				"contains at least one contextionary-valid word. To fix this, you have several "+
				"options:\n\n1.) Make sure that the schema class name or the set properties are "+
				"a contextionary-valid term and include them in vectorization using the "+
				"'vectorizeClassName' or 'vectorizePropertyName' setting. In this case the vector position "+
				"will be composed of both the class/property names and the values for those fields. "+
				"Even if no property values are contextionary-valid, the overall word corpus is still valid "+
				"due to the contextionary-valid class/property names."+
				"\n\n2.) Alternatively, if you do not want to include schema class/property names "+
				"in vectorization, you must make sure that at least one text/string property contains "+
				"at least one contextionary-valid word."+
				"\n\n3.) If the word corpus weaviate extracted from your object "+
				"(see below) does contain enough meaning to build a vector position, but the contextionary "+
				"did not recognize the words, you can extend the contextionary using the "+
				"REST API. This is the case	when you use mostly industry-specific terms which are "+
				"not known to the common language contextionary. Once extended, simply reimport this object."+
				"\n\nThe following words were extracted from your object: %v"+
				"\n\nTo learn more about the contextionary and how it behaves, check out: https://www.semi.technology/documentation/weaviate/current/contextionary.html"+
				"\n\nOriginal error: %v", corpi, err)
		default:
			return nil, nil, fmt.Errorf("vectorizing object with corpus '%+v': %v", corpi, err)
		}
	}

	return vector, ie, nil
}

// Corpi takes any list of strings and builds a common vector for all of them
func (v *Vectorizer) Corpi(ctx context.Context, corpi []string,
) ([]float32, error) {
	for i, corpus := range corpi {
		corpi[i] = camelCaseToLower(corpus)
	}

	vector, _, err := v.client.VectorForCorpi(ctx, corpi, nil)
	if err != nil {
		return nil, fmt.Errorf("vectorizing corpus '%+v': %v", corpi, err)
	}

	return vector, nil
}

func camelCaseToLower(in string) string {
	parts := camelcase.Split(in)
	var sb strings.Builder
	for i, part := range parts {
		if part == " " {
			continue
		}

		if i > 0 {
			sb.WriteString(" ")
		}

		sb.WriteString(strings.ToLower(part))
	}

	return sb.String()
}

func sourceFromInputElements(in []txt2vecmodels.InterpretationSource) []*txt2vecmodels.InterpretationSource {
	out := make([]*txt2vecmodels.InterpretationSource, len(in))
	for i, elem := range in {
		out[i] = &txt2vecmodels.InterpretationSource{
			Concept:    elem.Concept,
			Occurrence: elem.Occurrence,
			Weight:     float64(elem.Weight),
		}
	}

	return out
}
