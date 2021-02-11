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

package objects

import (
	"context"
	"errors"
	"fmt"
	"reflect"
	"testing"

	"github.com/go-openapi/strfmt"
	"github.com/semi-technologies/weaviate/entities/models"
	"github.com/semi-technologies/weaviate/usecases/config"
	"github.com/semi-technologies/weaviate/usecases/traverser"
	"github.com/sirupsen/logrus/hooks/test"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// A component-test like test suite that makes sure that every available UC is
// potentially protected with the Authorization plugin

func Test_Kinds_Authorization(t *testing.T) {
	type testCase struct {
		methodName       string
		additionalArgs   []interface{}
		expectedVerb     string
		expectedResource string
	}

	tests := []testCase{
		// single kind
		testCase{
			methodName:       "AddObject",
			additionalArgs:   []interface{}{(*models.Object)(nil)},
			expectedVerb:     "create",
			expectedResource: "objects",
		},
		testCase{
			methodName:       "ValidateObject",
			additionalArgs:   []interface{}{(*models.Object)(nil)},
			expectedVerb:     "validate",
			expectedResource: "objects",
		},
		testCase{
			methodName:       "GetObject",
			additionalArgs:   []interface{}{strfmt.UUID("foo"), traverser.AdditionalProperties{}},
			expectedVerb:     "get",
			expectedResource: "objects/foo",
		},
		testCase{
			methodName:       "DeleteObject",
			additionalArgs:   []interface{}{strfmt.UUID("foo")},
			expectedVerb:     "delete",
			expectedResource: "objects/foo",
		},
		testCase{
			methodName:       "UpdateObject",
			additionalArgs:   []interface{}{strfmt.UUID("foo"), (*models.Object)(nil)},
			expectedVerb:     "update",
			expectedResource: "objects/foo",
		},
		testCase{
			methodName:       "MergeObject",
			additionalArgs:   []interface{}{strfmt.UUID("foo"), (*models.Object)(nil)},
			expectedVerb:     "update",
			expectedResource: "objects/foo",
		},

		// list kinds
		testCase{
			methodName:       "GetObjects",
			additionalArgs:   []interface{}{(*int64)(nil), traverser.AdditionalProperties{}},
			expectedVerb:     "list",
			expectedResource: "objects",
		},

		// reference on kinds
		testCase{
			methodName:       "AddObjectReference",
			additionalArgs:   []interface{}{strfmt.UUID("foo"), "some prop", (*models.SingleRef)(nil)},
			expectedVerb:     "update",
			expectedResource: "objects/foo",
		},
		testCase{
			methodName:       "DeleteObjectReference",
			additionalArgs:   []interface{}{strfmt.UUID("foo"), "some prop", (*models.SingleRef)(nil)},
			expectedVerb:     "update",
			expectedResource: "objects/foo",
		},
		testCase{
			methodName:       "UpdateObjectReferences",
			additionalArgs:   []interface{}{strfmt.UUID("foo"), "some prop", (models.MultipleRef)(nil)},
			expectedVerb:     "update",
			expectedResource: "objects/foo",
		},
	}

	t.Run("verify that a test for every public method exists", func(t *testing.T) {
		testedMethods := make([]string, len(tests))
		for i, test := range tests {
			testedMethods[i] = test.methodName
		}

		for _, method := range allExportedMethods(&Manager{}) {
			assert.Contains(t, testedMethods, method)
		}
	})

	t.Run("verify the tested methods require correct permissions from the authorizer", func(t *testing.T) {
		principal := &models.Principal{}
		logger, _ := test.NewNullLogger()
		for _, test := range tests {
			t.Run(test.methodName, func(t *testing.T) {
				schemaManager := &fakeSchemaManager{}
				locks := &fakeLocks{}
				cfg := &config.WeaviateConfig{}
				authorizer := &authDenier{}
				extender := &fakeExtender{}
				projector := &fakeProjector{}
				vectorizer := &fakeVectorizer{}
				vecProvider := &fakeVectorizerProvider{vectorizer}
				vectorRepo := &fakeVectorRepo{}
				manager := NewManager(locks, schemaManager,
					cfg, logger, authorizer, vecProvider, vectorRepo, extender, projector)

				args := append([]interface{}{context.Background(), principal}, test.additionalArgs...)
				out, _ := callFuncByName(manager, test.methodName, args...)

				require.Len(t, authorizer.calls, 1, "authorizer must be called")
				assert.Equal(t, errors.New("just a test fake"), out[len(out)-1].Interface(),
					"execution must abort with authorizer error")
				assert.Equal(t, authorizeCall{principal, test.expectedVerb, test.expectedResource},
					authorizer.calls[0], "correct paramteres must have been used on authorizer")
			})
		}
	})
}

func Test_BatchKinds_Authorization(t *testing.T) {
	type testCase struct {
		methodName       string
		additionalArgs   []interface{}
		expectedVerb     string
		expectedResource string
	}

	tests := []testCase{
		// testCase{
		// 	methodName:       "AddActions",
		// 	additionalArgs:   []interface{}{[]*models.Object{}, []*string{}},
		// 	expectedVerb:     "create",
		// 	expectedResource: "batch/actions",
		// },

		testCase{
			methodName:       "AddObjects",
			additionalArgs:   []interface{}{[]*models.Object{}, []*string{}},
			expectedVerb:     "create",
			expectedResource: "batch/objects",
		},

		testCase{
			methodName:       "AddReferences",
			additionalArgs:   []interface{}{[]*models.BatchReference{}},
			expectedVerb:     "update",
			expectedResource: "batch/*",
		},
	}

	t.Run("verify that a test for every public method exists", func(t *testing.T) {
		testedMethods := make([]string, len(tests))
		for i, test := range tests {
			testedMethods[i] = test.methodName
		}

		for _, method := range allExportedMethods(&BatchManager{}) {
			assert.Contains(t, testedMethods, method)
		}
	})

	t.Run("verify the tested methods require correct permissions from the authorizer", func(t *testing.T) {
		principal := &models.Principal{}
		logger, _ := test.NewNullLogger()
		for _, test := range tests {
			schemaManager := &fakeSchemaManager{}
			locks := &fakeLocks{}
			cfg := &config.WeaviateConfig{}
			authorizer := &authDenier{}
			vectorRepo := &fakeVectorRepo{}
			vectorizer := &fakeVectorizer{}
			vecProvider := &fakeVectorizerProvider{vectorizer}
			manager := NewBatchManager(vectorRepo, vecProvider, locks, schemaManager, cfg, logger, authorizer)

			args := append([]interface{}{context.Background(), principal}, test.additionalArgs...)
			out, _ := callFuncByName(manager, test.methodName, args...)

			require.Len(t, authorizer.calls, 1, "authorizer must be called")
			assert.Equal(t, errors.New("just a test fake"), out[len(out)-1].Interface(),
				"execution must abort with authorizer error")
			assert.Equal(t, authorizeCall{principal, test.expectedVerb, test.expectedResource},
				authorizer.calls[0], "correct paramteres must have been used on authorizer")
		}
	})
}

type authorizeCall struct {
	principal *models.Principal
	verb      string
	resource  string
}

type authDenier struct {
	calls []authorizeCall
}

func (a *authDenier) Authorize(principal *models.Principal, verb, resource string) error {
	a.calls = append(a.calls, authorizeCall{principal, verb, resource})
	return errors.New("just a test fake")
}

// inspired by https://stackoverflow.com/a/33008200
func callFuncByName(manager interface{}, funcName string, params ...interface{}) (out []reflect.Value, err error) {
	managerValue := reflect.ValueOf(manager)
	m := managerValue.MethodByName(funcName)
	if !m.IsValid() {
		return make([]reflect.Value, 0), fmt.Errorf("Method not found \"%s\"", funcName)
	}
	in := make([]reflect.Value, len(params))
	for i, param := range params {
		in[i] = reflect.ValueOf(param)
	}
	out = m.Call(in)
	return
}

func allExportedMethods(subject interface{}) []string {
	var methods []string
	subjectType := reflect.TypeOf(subject)
	for i := 0; i < subjectType.NumMethod(); i++ {
		name := subjectType.Method(i).Name
		if name[0] >= 'A' && name[0] <= 'Z' {
			methods = append(methods, name)
		}
	}

	return methods
}
