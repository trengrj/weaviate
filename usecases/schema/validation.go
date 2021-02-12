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

package schema

import (
	"context"
	"fmt"

	"github.com/pkg/errors"
	"github.com/semi-technologies/weaviate/entities/models"
	"github.com/semi-technologies/weaviate/entities/schema"
	"github.com/semi-technologies/weaviate/usecases/config"
)

func (m *Manager) validateClassNameUniqueness(className string) error {
	for _, otherClass := range m.state.SchemaFor().Classes {
		if className == otherClass.Class {
			return fmt.Errorf("Name '%s' already used as a name for an Object class", className)
		}
	}

	return nil
}

// Check that the format of the name is correct
func (m *Manager) validateClassName(ctx context.Context, className string, vectorizeClass bool) error {
	_, err := schema.ValidateClassName(className)
	return err
}

func validatePropertyNameUniqueness(propertyName string, class *models.Class) error {
	for _, otherProperty := range class.Properties {
		if propertyName == otherProperty.Name {
			return fmt.Errorf("Name '%s' already in use as a property name for class '%s'", propertyName, class.Class)
		}
	}

	return nil
}

// Check that the format of the name is correct
func (m *Manager) validatePropertyName(ctx context.Context, className string,
	propertyName string, moduleConfig interface{}) error {
	_, err := schema.ValidatePropertyName(propertyName)
	return err
}

func (m *Manager) validateVectorSettings(ctx context.Context, class *models.Class) error {
	if err := m.validateVectorizer(ctx, class); err != nil {
		return err
	}

	if err := m.validateVectorIndex(ctx, class); err != nil {
		return err
	}

	return nil
}

func (m *Manager) validateVectorizer(ctx context.Context, class *models.Class) error {
	if class.Vectorizer == config.VectorizerModuleNone {
		return nil
	}

	if err := m.vectorizerValidator.ValidateVectorizer(class.Vectorizer); err != nil {
		return errors.Wrap(err, "vectorizer")
	}

	return nil
}

func (m *Manager) validateVectorIndex(ctx context.Context, class *models.Class) error {
	switch class.VectorIndexType {
	case "hnsw":
		return nil
	default:
		return errors.Errorf("unrecognized or unsupported vectorIndexType %q",
			class.VectorIndexType)
	}
}
