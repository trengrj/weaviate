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

package answer

import (
	"context"
	"testing"

	"github.com/semi-technologies/weaviate/entities/search"
	qnamodels "github.com/semi-technologies/weaviate/modules/qna-transformers/additional/models"
	"github.com/semi-technologies/weaviate/modules/qna-transformers/ent"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAdditionalAnswerProvider(t *testing.T) {
	t.Run("should fail with empty content", func(t *testing.T) {
		// given
		qnaClient := &fakeQnAClient{}
		fakeHelper := &fakeParamsHelper{}
		answerProvider := New(qnaClient, fakeHelper)
		in := []search.Result{
			{
				ID: "some-uuid",
			},
		}
		fakeParams := &Params{}
		limit := 1
		argumentModuleParams := map[string]interface{}{}

		// when
		out, err := answerProvider.AdditionalPropertyFn(context.Background(), in, fakeParams, &limit, argumentModuleParams)

		// then
		require.NotNil(t, err)
		require.NotEmpty(t, out)
		assert.Error(t, err, "empty content")
	})

	t.Run("should fail with empty question", func(t *testing.T) {
		// given
		qnaClient := &fakeQnAClient{}
		fakeHelper := &fakeParamsHelper{}
		answerProvider := New(qnaClient, fakeHelper)
		in := []search.Result{
			{
				ID: "some-uuid",
				Schema: map[string]interface{}{
					"content": "content",
				},
			},
		}
		fakeParams := &Params{}
		limit := 1
		argumentModuleParams := map[string]interface{}{}

		// when
		out, err := answerProvider.AdditionalPropertyFn(context.Background(), in, fakeParams, &limit, argumentModuleParams)

		// then
		require.NotNil(t, err)
		require.NotEmpty(t, out)
		assert.Error(t, err, "empty content")
	})

	t.Run("should answer", func(t *testing.T) {
		// given
		qnaClient := &fakeQnAClient{}
		fakeHelper := &fakeParamsHelper{}
		answerProvider := New(qnaClient, fakeHelper)
		in := []search.Result{
			{
				ID: "some-uuid",
				Schema: map[string]interface{}{
					"content": "content",
				},
			},
		}
		fakeParams := &Params{}
		limit := 1
		argumentModuleParams := map[string]interface{}{
			"ask": map[string]interface{}{
				"question": "question",
			},
		}

		// when
		out, err := answerProvider.AdditionalPropertyFn(context.Background(), in, fakeParams, &limit, argumentModuleParams)

		// then
		require.Nil(t, err)
		require.NotEmpty(t, out)
		assert.Equal(t, 1, len(in))
		answer, answerOK := in[0].AdditionalProperties["answer"]
		assert.True(t, answerOK)
		assert.NotNil(t, answer)
		answerAdditional, answerAdditionalOK := answer.(*qnamodels.Answer)
		assert.True(t, answerAdditionalOK)
		assert.Equal(t, "answer", *answerAdditional.Result)
	})

	t.Run("should answer with property", func(t *testing.T) {
		// given
		qnaClient := &fakeQnAClient{}
		fakeHelper := &fakeParamsHelper{}
		answerProvider := New(qnaClient, fakeHelper)
		in := []search.Result{
			{
				ID: "some-uuid",
				Schema: map[string]interface{}{
					"content":  "content with answer",
					"content2": "this one is just a title",
				},
			},
		}
		fakeParams := &Params{}
		limit := 1
		argumentModuleParams := map[string]interface{}{
			"ask": map[string]interface{}{
				"question":   "question",
				"properties": []string{"content", "content2"},
			},
		}

		// when
		out, err := answerProvider.AdditionalPropertyFn(context.Background(), in, fakeParams, &limit, argumentModuleParams)

		// then
		require.Nil(t, err)
		require.NotEmpty(t, out)
		assert.Equal(t, 1, len(in))
		answer, answerOK := in[0].AdditionalProperties["answer"]
		assert.True(t, answerOK)
		assert.NotNil(t, answer)
		answerAdditional, answerAdditionalOK := answer.(*qnamodels.Answer)
		assert.True(t, answerAdditionalOK)
		assert.Equal(t, "answer", *answerAdditional.Result)
		assert.Equal(t, "content", *answerAdditional.Property)
		assert.Equal(t, 0.8, *answerAdditional.Certainty)
		assert.Equal(t, 13, answerAdditional.StartPosition)
		assert.Equal(t, 19, answerAdditional.EndPosition)
		assert.Equal(t, true, answerAdditional.HasAnswer)
	})

	t.Run("should answer with certainty set above ask certainty", func(t *testing.T) {
		// given
		qnaClient := &fakeQnAClient{}
		fakeHelper := &fakeParamsHelper{}
		answerProvider := New(qnaClient, fakeHelper)
		in := []search.Result{
			{
				ID: "some-uuid",
				Schema: map[string]interface{}{
					"content":  "content with answer",
					"content2": "this one is just a title",
				},
			},
		}
		fakeParams := &Params{}
		limit := 1
		argumentModuleParams := map[string]interface{}{
			"ask": map[string]interface{}{
				"question":   "question",
				"properties": []string{"content", "content2"},
				"certainty":  float64(0.8),
			},
		}

		// when
		out, err := answerProvider.AdditionalPropertyFn(context.Background(), in, fakeParams, &limit, argumentModuleParams)

		// then
		require.Nil(t, err)
		require.NotEmpty(t, out)
		assert.Equal(t, 1, len(in))
		answer, answerOK := in[0].AdditionalProperties["answer"]
		assert.True(t, answerOK)
		assert.NotNil(t, answer)
		answerAdditional, answerAdditionalOK := answer.(*qnamodels.Answer)
		assert.True(t, answerAdditionalOK)
		assert.Equal(t, "answer", *answerAdditional.Result)
		assert.Equal(t, "content", *answerAdditional.Property)
		assert.Equal(t, 0.8, *answerAdditional.Certainty)
		assert.Equal(t, 13, answerAdditional.StartPosition)
		assert.Equal(t, 19, answerAdditional.EndPosition)
		assert.Equal(t, true, answerAdditional.HasAnswer)
	})

	t.Run("should not answer with certainty set below ask certainty", func(t *testing.T) {
		// given
		qnaClient := &fakeQnAClient{}
		fakeHelper := &fakeParamsHelper{}
		answerProvider := New(qnaClient, fakeHelper)
		in := []search.Result{
			{
				ID: "some-uuid",
				Schema: map[string]interface{}{
					"content":  "content with answer",
					"content2": "this one is just a title",
				},
			},
		}
		fakeParams := &Params{}
		limit := 1
		argumentModuleParams := map[string]interface{}{
			"ask": map[string]interface{}{
				"question":   "question",
				"properties": []string{"content", "content2"},
				"certainty":  float64(0.81),
			},
		}

		// when
		out, err := answerProvider.AdditionalPropertyFn(context.Background(), in, fakeParams, &limit, argumentModuleParams)

		// then
		require.Nil(t, err)
		require.NotEmpty(t, out)
		assert.Equal(t, 1, len(in))
		answer, answerOK := in[0].AdditionalProperties["answer"]
		assert.True(t, answerOK)
		assert.NotNil(t, answer)
		answerAdditional, answerAdditionalOK := answer.(*qnamodels.Answer)
		assert.True(t, answerAdditionalOK)
		assert.True(t, answerAdditional.Result == nil)
		assert.True(t, answerAdditional.Property == nil)
		assert.True(t, answerAdditional.Certainty == nil)
		assert.Equal(t, 0, answerAdditional.StartPosition)
		assert.Equal(t, 0, answerAdditional.EndPosition)
		assert.Equal(t, false, answerAdditional.HasAnswer)
	})
}

type fakeQnAClient struct{}

func (c *fakeQnAClient) Answer(ctx context.Context,
	text, question string) (*ent.AnswerResult, error) {
	answerString := "answer"
	var certainty float64 = 0.8
	answer := &ent.AnswerResult{
		Text:      question,
		Question:  question,
		Answer:    &answerString,
		Certainty: &certainty,
	}
	return answer, nil
}

type fakeParamsHelper struct{}

func (h *fakeParamsHelper) GetQuestion(params interface{}) string {
	if fakeParamsMap, ok := params.(map[string]interface{}); ok {
		if question, ok := fakeParamsMap["question"].(string); ok {
			return question
		}
	}
	return ""
}

func (h *fakeParamsHelper) GetProperties(params interface{}) []string {
	if fakeParamsMap, ok := params.(map[string]interface{}); ok {
		if properties, ok := fakeParamsMap["properties"].([]string); ok {
			return properties
		}
	}
	return nil
}

func (h *fakeParamsHelper) GetCertainty(params interface{}) float64 {
	if fakeParamsMap, ok := params.(map[string]interface{}); ok {
		if certainty, ok := fakeParamsMap["certainty"].(float64); ok {
			return certainty
		}
	}
	return 0
}
