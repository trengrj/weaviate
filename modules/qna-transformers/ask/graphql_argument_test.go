//                           _       _
// __      _____  __ ___   ___  __ _| |_ ___
// \ \ /\ / / _ \/ _` \ \ / / |/ _` | __/ _ \
//  \ V  V /  __/ (_| |\ V /| | (_| | ||  __/
//   \_/\_/ \___|\__,_| \_/ |_|\__,_|\__\___|
//
//  Copyright © 2016 - 2021 SeMI Technologies B.V. All rights reserved.
//
//  CONTACT: hello@semi.technology
//

package ask

import (
	"testing"

	"github.com/graphql-go/graphql"
	"github.com/stretchr/testify/assert"
)

func TestAskGraphQLArgument(t *testing.T) {
	t.Run("should generate ask argument properly", func(t *testing.T) {
		// given
		prefix := "Prefix"
		classname := "Class"
		// when
		ask := New(nil).askArgument(prefix, classname)

		// then
		// the built graphQL field needs to support this structure:
		// ask {
		//   question: "question?",
		//   certainty: 0.9
		//   properties: ["prop1", "prop2"]
		// }
		assert.NotNil(t, ask)
		assert.Equal(t, "QnATransformersPrefixClassAskInpObj", ask.Type.Name())
		askFields, ok := ask.Type.(*graphql.InputObject)
		assert.True(t, ok)
		assert.NotNil(t, askFields)
		assert.Equal(t, 3, len(askFields.Fields()))
		fields := askFields.Fields()
		question := fields["question"]
		questionNonNull, questionNonNullOK := question.Type.(*graphql.NonNull)
		assert.True(t, questionNonNullOK)
		assert.Equal(t, "String", questionNonNull.OfType.Name())
		assert.NotNil(t, question)
		assert.NotNil(t, fields["certainty"])
		properties := fields["properties"]
		propertiesList, propertiesListOK := properties.Type.(*graphql.List)
		assert.True(t, propertiesListOK)
		assert.Equal(t, "String", propertiesList.OfType.Name())
	})
}

func TestAskGraphQLArgumentWithAutocorrect(t *testing.T) {
	t.Run("should generate ask argument properly with autocorrect", func(t *testing.T) {
		// given
		prefix := "Prefix"
		classname := "Class"
		// when
		ask := New(&fakeTransformer{}).askArgument(prefix, classname)

		// then
		// the built graphQL field needs to support this structure:
		// ask {
		//   question: "question?",
		//   certainty: 0.9
		//   properties: ["prop1", "prop2"]
		//   autocorrect: true
		// }
		assert.NotNil(t, ask)
		assert.Equal(t, "QnATransformersPrefixClassAskInpObj", ask.Type.Name())
		askFields, ok := ask.Type.(*graphql.InputObject)
		assert.True(t, ok)
		assert.NotNil(t, askFields)
		assert.Equal(t, 4, len(askFields.Fields()))
		fields := askFields.Fields()
		question := fields["question"]
		questionNonNull, questionNonNullOK := question.Type.(*graphql.NonNull)
		assert.True(t, questionNonNullOK)
		assert.Equal(t, "String", questionNonNull.OfType.Name())
		assert.NotNil(t, question)
		assert.NotNil(t, fields["certainty"])
		properties := fields["properties"]
		propertiesList, propertiesListOK := properties.Type.(*graphql.List)
		assert.True(t, propertiesListOK)
		assert.Equal(t, "String", propertiesList.OfType.Name())
		assert.NotNil(t, fields["autocorrect"])
	})
}
