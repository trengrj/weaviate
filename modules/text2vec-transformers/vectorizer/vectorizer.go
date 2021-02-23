package vectorizer

import (
	"context"
	"fmt"
	"strings"

	"github.com/fatih/camelcase"
	"github.com/semi-technologies/weaviate/entities/models"
	"github.com/semi-technologies/weaviate/modules/text2vec-transformers/ent"
)

type Vectorizer struct {
	client Client
}

func New(client Client) *Vectorizer {
	return &Vectorizer{
		client: client,
	}
}

type Client interface {
	Vectorize(ctx context.Context, input string,
		cfg ent.VectorizationConfig) (*ent.VectorizationResult, error)
}

// IndexCheck returns whether a property of a class should be indexed
type ClassSettings interface {
	PropertyIndexed(property string) bool
	VectorizeClassName() bool
	VectorizePropertyName(propertyName string) bool
	PoolingStrategy() string
}

func (v *Vectorizer) Object(ctx context.Context, object *models.Object,
	icheck ClassSettings) error {
	vec, err := v.object(ctx, object.Class, object.Properties, icheck)
	if err != nil {
		return err
	}

	object.Vector = vec
	return nil
}

func (v *Vectorizer) object(ctx context.Context, className string,
	schema interface{}, icheck ClassSettings) ([]float32, error) {
	var corpi []string

	if icheck.VectorizeClassName() {
		corpi = append(corpi, camelCaseToLower(className))
	}

	if schema != nil {
		for prop, value := range schema.(map[string]interface{}) {
			if !icheck.PropertyIndexed(prop) {
				continue
			}

			valueString, ok := value.(string)
			if ok {
				if icheck.VectorizePropertyName(prop) {
					// use prop and value
					corpi = append(corpi, strings.ToLower(
						fmt.Sprintf("%s %s", camelCaseToLower(prop), valueString)))
				} else {
					corpi = append(corpi, strings.ToLower(valueString))
				}
			}
		}
	}

	if len(corpi) == 0 {
		// fall back to using the class name
		corpi = append(corpi, camelCaseToLower(className))
	}

	text := strings.Join(corpi, " ")
	res, err := v.client.Vectorize(ctx, text, ent.VectorizationConfig{
		PoolingStrategy: icheck.PoolingStrategy(),
	})
	if err != nil {
		return nil, err
	}

	return res.Vector, nil
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
