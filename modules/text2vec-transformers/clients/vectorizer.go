package clients

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/pkg/errors"
	"github.com/semi-technologies/weaviate/modules/text2vec-transformers/ent"
)

type vectorizer struct {
	origin     string
	httpClient *http.Client
}

func New(origin string) *vectorizer {
	return &vectorizer{
		origin:     origin,
		httpClient: &http.Client{},
	}
}

func (v *vectorizer) Vectorize(ctx context.Context, input string,
	config ent.VectorizationConfig) (*ent.VectorizationResult, error) {
	body, err := json.Marshal(vecRequest{
		Text: input,
		Config: vecRequestConfig{
			PoolingStrategy: config.PoolingStrategy,
		},
	})
	if err != nil {
		return nil, errors.Wrapf(err, "marshal body")
	}

	req, err := http.NewRequestWithContext(ctx, "POST", v.url("/vectors"),
		bytes.NewReader(body))
	if err != nil {
		return nil, errors.Wrap(err, "create POST request")
	}

	res, err := v.httpClient.Do(req)
	if err != nil {
		return nil, errors.Wrap(err, "send POST request")
	}
	defer res.Body.Close()

	bodyBytes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, errors.Wrap(err, "read response body")
	}

	var resBody vecRequest
	if err := json.Unmarshal(bodyBytes, &resBody); err != nil {
		return nil, errors.Wrap(err, "unmarshal response body")
	}

	if res.StatusCode > 399 {
		return nil, errors.Errorf("fail with status %d: %s", res.StatusCode,
			resBody.Error)
	}

	return &ent.VectorizationResult{
		Text:       resBody.Text,
		Dimensions: resBody.Dims,
		Vector:     resBody.Vector,
	}, nil
}

func (v *vectorizer) url(path string) string {
	return fmt.Sprintf("%s%s", v.origin, path)
}

type vecRequest struct {
	Text   string           `json:"text"`
	Dims   int              `json:"dims"`
	Vector []float32        `json:"vector"`
	Error  string           `json:"error"`
	Config vecRequestConfig `json:"config"`
}

type vecRequestConfig struct {
	PoolingStrategy string `json:"pooling_strategy"`
}
