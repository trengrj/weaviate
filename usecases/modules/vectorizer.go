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

package modules

import (
	"context"
	"fmt"

	"github.com/pkg/errors"
	"github.com/semi-technologies/weaviate/entities/models"
	"github.com/semi-technologies/weaviate/entities/modulecapabilities"
	"github.com/semi-technologies/weaviate/entities/schema"
	"github.com/semi-technologies/weaviate/entities/vectorindex/hnsw"
	"github.com/semi-technologies/weaviate/usecases/config"
	"github.com/sirupsen/logrus"
)

const (
	errorVectorizerCapability = "module %q exists, but does not provide the " +
		"Vectorizer or ReferenceVectorizer capability"

	errorVectorIndexType = "vector index config (%T) is not of type HNSW, " +
		"but objects manager is restricted to HNSW"

	warningVectorIgnored = "This vector will be ignored. If you meant to index " +
		"the vector, make sure to set vectorIndexConfig.skip to 'false'. If the previous " +
		"setting is correct, make sure you set vectorizer to 'none' in the schema and " +
		"provide a null-vector (i.e. no vector) at import time."

	warningSkipVectorGenerated = "this class is configured to skip vector indexing, " +
		"but a vector was generated by the %q vectorizer. " + warningVectorIgnored

	warningSkipVectorProvided = "this class is configured to skip vector indexing, " +
		"but a vector was explicitly provided. " + warningVectorIgnored
)

func (m *Provider) ValidateVectorizer(moduleName string) error {
	mod := m.GetByName(moduleName)
	if mod == nil {
		return errors.Errorf("no module with name %q present", moduleName)
	}

	_, okVec := mod.(modulecapabilities.Vectorizer)
	_, okRefVec := mod.(modulecapabilities.ReferenceVectorizer)
	if !okVec && !okRefVec {
		return errors.Errorf(errorVectorizerCapability, moduleName)
	}

	return nil
}

func (m *Provider) UsingRef2Vec(className string) bool {
	class, err := m.getClass(className)
	if err != nil {
		return false
	}

	cfg := class.ModuleConfig
	if cfg == nil {
		return false
	}

	for modName := range cfg.(map[string]interface{}) {
		mod := m.GetByName(modName)
		if _, ok := mod.(modulecapabilities.ReferenceVectorizer); ok {
			return true
		}
	}

	return false
}

func (m *Provider) UpdateVector(ctx context.Context, object *models.Object,
	findObjectFn modulecapabilities.FindObjectFn, logger logrus.FieldLogger,
) error {
	class, err := m.getClass(object.Class)
	if err != nil {
		return err
	}

	vectorizerName, idxCfg, err := m.getClassVectorizer(object.Class)
	if err != nil {
		return err
	}

	hnswConfig, ok := idxCfg.(hnsw.UserConfig)
	if !ok {
		return fmt.Errorf(errorVectorIndexType, idxCfg)
	}

	if vectorizerName == config.VectorizerModuleNone {
		if hnswConfig.Skip && len(object.Vector) > 0 {
			logger.WithField("className", object.Class).
				Warningf(warningSkipVectorProvided)
		}

		return nil
	}

	if hnswConfig.Skip {
		logger.WithField("className", object.Class).
			WithField("vectorizer", vectorizerName).
			Warningf(warningSkipVectorGenerated, vectorizerName)
	}

	var found modulecapabilities.Module
	for modName := range class.ModuleConfig.(map[string]interface{}) {
		if err := m.ValidateVectorizer(modName); err == nil {
			found = m.GetByName(modName)
			break
		}
	}

	if found == nil {
		return fmt.Errorf(
			"no vectorizer found for class %q: %w", object.Class, err)
	}

	cfg := NewClassBasedModuleConfig(class, found.Name())

	if vectorizer, ok := found.(modulecapabilities.Vectorizer); ok {
		if object.Vector == nil {
			if err := vectorizer.VectorizeObject(ctx, object, cfg); err != nil {
				return fmt.Errorf("update vector: %w", err)
			}
		}
	} else {
		refVectorizer := found.(modulecapabilities.ReferenceVectorizer)
		if err := refVectorizer.VectorizeObject(
			ctx, object, cfg, findObjectFn); err != nil {
			return fmt.Errorf("update reference vector: %w", err)
		}
	}

	return nil
}

func (m *Provider) VectorizerName(className string) (string, error) {
	name, _, err := m.getClassVectorizer(className)
	if err != nil {
		return "", err
	}
	return name, nil
}

func (m *Provider) getClassVectorizer(className string) (string, interface{}, error) {
	sch := m.schemaGetter.GetSchemaSkipAuth()

	class := sch.FindClassByName(schema.ClassName(className))
	if class == nil {
		// this should be impossible by the time this method gets called, but let's
		// be 100% certain
		return "", nil, fmt.Errorf("class %s not present", className)
	}

	return class.Vectorizer, class.VectorIndexConfig, nil
}
