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
package schema

import (
	"github.com/creativesoftwarefdn/weaviate/models"
	"testing"
)

func TestDetectPrimitiveTypes(t *testing.T) {
	s := &Schema{}

	for _, type_ := range PrimitiveDataTypes {
		pdt, err := s.FindPropertyDataType([]string{string(type_)})
		if err != nil {
			t.Fatal(err)
		}

		if !pdt.IsPrimitive() {
			t.Fatal("not primitive")
		}

		if pdt.AsPrimitive() != type_ {
			t.Fatal("wrong value")
		}
	}
}

func TestNonExistingClassSingleRef(t *testing.T) {
	s := Empty()

	pdt, err := s.FindPropertyDataType([]string{"NonExistingClass"})

	if err == nil {
		t.Fatal("Should have error")
	}

	if pdt != nil {
		t.Fatal("Should return nil result")
	}
}

func TestExistingClassSingleRef(t *testing.T) {
	s := Empty()

	s.Actions.Classes = append(s.Actions.Classes, &models.SemanticSchemaClass{
		Class: "ExistingClass",
	})

	pdt, err := s.FindPropertyDataType([]string{"ExistingClass"})

	if err != nil {
		t.Fatal(err)
	}

	if !pdt.IsReference() {
		t.Fatal("not single ref")
	}
}
