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

package traverser

import (
	"fmt"
	"regexp"

	"github.com/semi-technologies/weaviate/entities/filters"
	"github.com/semi-technologies/weaviate/entities/schema"
	"github.com/semi-technologies/weaviate/entities/schema/kind"
	libprojector "github.com/semi-technologies/weaviate/usecases/projector"
	"github.com/semi-technologies/weaviate/usecases/sempath"
)

type GetParams struct {
	Kind       kind.Kind
	Filters    *filters.LocalFilter
	ClassName  string
	Pagination *filters.Pagination
	Properties SelectProperties

	// TODO: The searcher params must be some sort of an interface and accept a
	// list, since NearText will be provided from the `text2vec-contextionary`
	// module when we have actual modularization
	NearText             *NearTextParams
	NearVector           *NearVectorParams
	SearchVector         []float32
	Group                *GroupParams
	AdditionalProperties AdditionalProperties
}

type SelectProperty struct {
	Name string

	IsPrimitive bool

	// Include the __typename in all the Refs below.
	IncludeTypeName bool

	// Not a primitive type? Then select these properties.
	Refs []SelectClass
}

type SelectClass struct {
	ClassName     string
	RefProperties SelectProperties
}

type GroupParams struct {
	Strategy string
	Force    float32
}

// FindSelectClass by specifying the exact class name
func (sp SelectProperty) FindSelectClass(className schema.ClassName) *SelectClass {
	for _, selectClass := range sp.Refs {
		if selectClass.ClassName == string(className) {
			return &selectClass
		}
	}

	return nil
}

// HasPeer returns true if any of the referenced classes are from the specified
// peer
func (sp SelectProperty) HasPeer(peerName string) bool {
	r := regexp.MustCompile(fmt.Sprintf("^%s__", peerName))
	for _, selectClass := range sp.Refs {
		if r.MatchString(selectClass.ClassName) {
			return true
		}
	}

	return false
}

type SelectProperties []SelectProperty

func (sp SelectProperties) HasRefs() bool {
	for _, p := range sp {
		if !p.IsPrimitive {
			return true
		}
	}

	return false
}

func (sp SelectProperties) ShouldResolve(path []string) (bool, error) {
	if len(path)%2 != 0 || len(path) == 0 {
		return false, fmt.Errorf("used incorrectly: path must have even number of segments in the form of " +
			"refProp, className, refProp, className, etc.")
	}

	// the above gives us the guarantuee that path contains at least two elements
	property := path[0]
	class := schema.ClassName(path[1])

	for _, p := range sp {
		if p.IsPrimitive {
			continue
		}

		if p.Name != property {
			continue
		}

		selectClass := p.FindSelectClass(class)
		if selectClass == nil {
			continue
		}

		if len(path) > 2 {
			// we're not done yet, this one's nested
			return selectClass.RefProperties.ShouldResolve(path[2:])
		}

		// we are done and found the path
		return true, nil
	}

	return false, nil
}

func (sp SelectProperties) FindProperty(propName string) *SelectProperty {
	for _, prop := range sp {
		if prop.Name == propName {
			return &prop
		}
	}

	return nil
}

type AdditionalProperties struct {
	Classification    bool
	RefMeta           bool
	Vector            bool
	Interpretation    bool
	NearestNeighbors  bool
	SemanticPath      *sempath.Params
	FeatureProjection *libprojector.Params
	Certainty         bool
	ID                bool
}
