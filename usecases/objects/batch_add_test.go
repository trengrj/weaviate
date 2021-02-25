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

package objects

import (
	"context"
	"testing"

	"github.com/go-openapi/strfmt"
	"github.com/semi-technologies/weaviate/entities/models"
	"github.com/semi-technologies/weaviate/entities/schema"
	"github.com/semi-technologies/weaviate/usecases/config"
	"github.com/sirupsen/logrus/hooks/test"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func Test_BatchManager_AddObjects_WithNoVectorizerModule(t *testing.T) {
	var (
		vectorRepo *fakeVectorRepo
		manager    *BatchManager
	)

	schema := schema.Schema{
		Objects: &models.Schema{
			Classes: []*models.Class{
				{
					Vectorizer: config.VectorizerModuleNone,
					Class:      "Foo",
				},
			},
		},
	}

	reset := func() {
		vectorRepo = &fakeVectorRepo{}
		config := &config.WeaviateConfig{}
		locks := &fakeLocks{}
		schemaManager := &fakeSchemaManager{
			GetSchemaResponse: schema,
		}
		logger, _ := test.NewNullLogger()
		authorizer := &fakeAuthorizer{}
		vectorizer := &fakeVectorizer{}
		vecProvider := &fakeVectorizerProvider{vectorizer}
		manager = NewBatchManager(vectorRepo, vecProvider, locks,
			schemaManager, config, logger, authorizer)
	}

	ctx := context.Background()

	t.Run("without any objects", func(t *testing.T) {
		reset()
		expectedErr := NewErrInvalidUserInput("invalid param 'objects': cannot be empty, need at least" +
			" one object for batching")

		_, err := manager.AddObjects(ctx, nil, []*models.Object{}, []*string{})

		assert.Equal(t, expectedErr, err)
	})

	t.Run("with objects without IDs", func(t *testing.T) {
		reset()
		vectorRepo.On("BatchPutObjects", mock.Anything).Return(nil).Once()
		objects := []*models.Object{
			{
				Class:  "Foo",
				Vector: []float32{0.1, 0.1, 0.1111},
			},
			{
				Class:  "Foo",
				Vector: []float32{0.2, 0.2, 0.2222},
			},
		}

		_, err := manager.AddObjects(ctx, nil, objects, []*string{})
		repoCalledWithObjects := vectorRepo.Calls[0].Arguments[0].(BatchObjects)

		assert.Nil(t, err)
		require.Len(t, repoCalledWithObjects, 2)
		assert.Len(t, repoCalledWithObjects[0].UUID, 36,
			"a uuid was set for the first object")
		assert.Len(t, repoCalledWithObjects[1].UUID, 36,
			"a uuid was set for the second object")
		assert.Nil(t, repoCalledWithObjects[0].Err)
		assert.Nil(t, repoCalledWithObjects[1].Err)
		assert.Equal(t, []float32{0.1, 0.1, 0.1111}, repoCalledWithObjects[0].Vector,
			"the correct vector was used")
		assert.Equal(t, []float32{0.2, 0.2, 0.2222}, repoCalledWithObjects[1].Vector,
			"the correct vector was used")
	})

	t.Run("with user-specified IDs", func(t *testing.T) {
		reset()
		vectorRepo.On("BatchPutObjects", mock.Anything).Return(nil).Once()
		id1 := strfmt.UUID("2d3942c3-b412-4d80-9dfa-99a646629cd2")
		id2 := strfmt.UUID("cf918366-3d3b-4b90-9bc6-bc5ea8762ff6")
		objects := []*models.Object{
			{
				ID:     id1,
				Class:  "Foo",
				Vector: []float32{0.1, 0.1, 0.1111},
			},
			{
				ID:     id2,
				Class:  "Foo",
				Vector: []float32{0.2, 0.2, 0.2222},
			},
		}

		_, err := manager.AddObjects(ctx, nil, objects, []*string{})
		repoCalledWithObjects := vectorRepo.Calls[0].Arguments[0].(BatchObjects)

		assert.Nil(t, err)
		require.Len(t, repoCalledWithObjects, 2)
		assert.Equal(t, id1, repoCalledWithObjects[0].UUID, "the user-specified uuid was used")
		assert.Equal(t, id2, repoCalledWithObjects[1].UUID, "the user-specified uuid was used")
		assert.Nil(t, repoCalledWithObjects[0].Err)
		assert.Nil(t, repoCalledWithObjects[1].Err)
		assert.Equal(t, []float32{0.1, 0.1, 0.1111}, repoCalledWithObjects[0].Vector,
			"the correct vector was used")
		assert.Equal(t, []float32{0.2, 0.2, 0.2222}, repoCalledWithObjects[1].Vector,
			"the correct vector was used")
	})

	t.Run("with an invalid user-specified IDs", func(t *testing.T) {
		reset()
		vectorRepo.On("BatchPutObjects", mock.Anything).Return(nil).Once()
		id1 := strfmt.UUID("invalid")
		id2 := strfmt.UUID("cf918366-3d3b-4b90-9bc6-bc5ea8762ff6")
		objects := []*models.Object{
			{
				ID:     id1,
				Class:  "Foo",
				Vector: []float32{0.1, 0.1, 0.1111},
			},
			{
				ID:     id2,
				Class:  "Foo",
				Vector: []float32{0.2, 0.2, 0.2222},
			},
		}

		_, err := manager.AddObjects(ctx, nil, objects, []*string{})
		repoCalledWithObjects := vectorRepo.Calls[0].Arguments[0].(BatchObjects)

		assert.Nil(t, err)
		require.Len(t, repoCalledWithObjects, 2)
		assert.Equal(t, repoCalledWithObjects[0].Err.Error(), "uuid: incorrect UUID length: invalid")
		assert.Equal(t, id2, repoCalledWithObjects[1].UUID, "the user-specified uuid was used")
	})
}

func Test_BatchManager_AddObjects_WithExternalVectorizerModule(t *testing.T) {
	var (
		vectorRepo *fakeVectorRepo
		manager    *BatchManager
	)

	schema := schema.Schema{
		Objects: &models.Schema{
			Classes: []*models.Class{
				{
					Vectorizer: config.VectorizerModuleText2VecContextionary,
					Class:      "Foo",
				},
			},
		},
	}

	reset := func() {
		vectorRepo = &fakeVectorRepo{}
		config := &config.WeaviateConfig{}
		locks := &fakeLocks{}
		schemaManager := &fakeSchemaManager{
			GetSchemaResponse: schema,
		}
		logger, _ := test.NewNullLogger()
		authorizer := &fakeAuthorizer{}
		vectorizer := &fakeVectorizer{}
		vecProvider := &fakeVectorizerProvider{vectorizer}
		vectorizer.On("UpdateObject", mock.Anything).Return([]float32{0, 1, 2}, nil)
		manager = NewBatchManager(vectorRepo, vecProvider, locks,
			schemaManager, config, logger, authorizer)
	}

	ctx := context.Background()

	t.Run("without any objects", func(t *testing.T) {
		reset()
		expectedErr := NewErrInvalidUserInput("invalid param 'objects': cannot be empty, need at least" +
			" one object for batching")

		_, err := manager.AddObjects(ctx, nil, []*models.Object{}, []*string{})

		assert.Equal(t, expectedErr, err)
	})

	t.Run("with objects without IDs", func(t *testing.T) {
		reset()
		vectorRepo.On("BatchPutObjects", mock.Anything).Return(nil).Once()
		objects := []*models.Object{
			{
				Class: "Foo",
			},
			{
				Class: "Foo",
			},
		}

		_, err := manager.AddObjects(ctx, nil, objects, []*string{})
		repoCalledWithObjects := vectorRepo.Calls[0].Arguments[0].(BatchObjects)

		assert.Nil(t, err)
		require.Len(t, repoCalledWithObjects, 2)
		assert.Len(t, repoCalledWithObjects[0].UUID, 36, "a uuid was set for the first object")
		assert.Len(t, repoCalledWithObjects[1].UUID, 36, "a uuid was set for the second object")
		assert.Nil(t, repoCalledWithObjects[0].Err)
		assert.Nil(t, repoCalledWithObjects[1].Err)
		assert.Equal(t, []float32{0, 1, 2}, repoCalledWithObjects[0].Vector,
			"the correct vector was used")
		assert.Equal(t, []float32{0, 1, 2}, repoCalledWithObjects[1].Vector,
			"the correct vector was used")
	})

	t.Run("with user-specified IDs", func(t *testing.T) {
		reset()
		vectorRepo.On("BatchPutObjects", mock.Anything).Return(nil).Once()
		id1 := strfmt.UUID("2d3942c3-b412-4d80-9dfa-99a646629cd2")
		id2 := strfmt.UUID("cf918366-3d3b-4b90-9bc6-bc5ea8762ff6")
		objects := []*models.Object{
			{
				ID:    id1,
				Class: "Foo",
			},
			{
				ID:    id2,
				Class: "Foo",
			},
		}

		_, err := manager.AddObjects(ctx, nil, objects, []*string{})
		repoCalledWithObjects := vectorRepo.Calls[0].Arguments[0].(BatchObjects)

		assert.Nil(t, err)
		require.Len(t, repoCalledWithObjects, 2)
		assert.Equal(t, id1, repoCalledWithObjects[0].UUID, "the user-specified uuid was used")
		assert.Equal(t, id2, repoCalledWithObjects[1].UUID, "the user-specified uuid was used")
	})

	t.Run("with an invalid user-specified IDs", func(t *testing.T) {
		reset()
		vectorRepo.On("BatchPutObjects", mock.Anything).Return(nil).Once()
		id1 := strfmt.UUID("invalid")
		id2 := strfmt.UUID("cf918366-3d3b-4b90-9bc6-bc5ea8762ff6")
		objects := []*models.Object{
			{
				ID:    id1,
				Class: "Foo",
			},
			{
				ID:    id2,
				Class: "Foo",
			},
		}

		_, err := manager.AddObjects(ctx, nil, objects, []*string{})
		repoCalledWithObjects := vectorRepo.Calls[0].Arguments[0].(BatchObjects)

		assert.Nil(t, err)
		require.Len(t, repoCalledWithObjects, 2)
		assert.Equal(t, repoCalledWithObjects[0].Err.Error(), "uuid: incorrect UUID length: invalid")
		assert.Equal(t, id2, repoCalledWithObjects[1].UUID, "the user-specified uuid was used")
	})
}
