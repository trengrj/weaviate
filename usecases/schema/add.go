/*                          _       _
 *__      _____  __ ___   ___  __ _| |_ ___
 *\ \ /\ / / _ \/ _` \ \ / / |/ _` | __/ _ \
 * \ V  V /  __/ (_| |\ V /| | (_| | ||  __/
 *  \_/\_/ \___|\__,_| \_/ |_|\__,_|\__\___|
 *
 * Copyright © 2016 - 2019 Weaviate. All rights reserved.
 * LICENSE: https://github.com/semi-technologies/weaviate/blob/develop/LICENSE.md
 * DESIGN & CONCEPT: Bob van Luijt (@bobvanluijt)
 * CONTACT: hello@semi.technology
 */package schema

import (
	"context"
	"fmt"
	"strings"

	"github.com/semi-technologies/weaviate/entities/models"
	"github.com/semi-technologies/weaviate/entities/schema/kind"
)

// AddAction Class to the schema
func (m *Manager) AddAction(ctx context.Context, class *models.SemanticSchemaClass) error {
	return m.addClass(ctx, class, kind.Action)
}

// AddThing Class to the schema
func (m *Manager) AddThing(ctx context.Context, class *models.SemanticSchemaClass) error {
	return m.addClass(ctx, class, kind.Thing)
}

func (m *Manager) addClass(ctx context.Context, class *models.SemanticSchemaClass, k kind.Kind) error {
	unlock, err := m.locks.LockSchema()
	if err != nil {
		return err
	}
	defer unlock()

	class.Class = upperCaseClassName(class.Class)
	class.Properties = lowerCaseAllPropertyNames(class.Properties)

	err = m.validateCanAddClass(k, class)
	if err != nil {
		return err
	}

	semanticSchema := m.state.SchemaFor(k)
	semanticSchema.Classes = append(semanticSchema.Classes, class)
	err = m.saveSchema(ctx)
	if err != nil {
		return err
	}

	return m.migrator.AddClass(ctx, k, class)
	// TODO gh-846: Rollback state upate if migration fails
}

func (m *Manager) validateCanAddClass(knd kind.Kind, class *models.SemanticSchemaClass) error {
	// First check if there is a name clash.
	err := m.validateClassNameUniqueness(class.Class)
	if err != nil {
		return err
	}

	err = m.validateClassNameAndKeywords(knd, class.Class, class.Keywords)
	if err != nil {
		return err
	}

	// Check properties
	foundNames := map[string]bool{}
	for _, property := range class.Properties {
		err = m.validatePropertyNameAndKeywords(class.Class, property.Name, property.Keywords)
		if err != nil {
			return err
		}

		if foundNames[property.Name] == true {
			return fmt.Errorf("Name '%s' already in use as a property name for class '%s'", property.Name, class.Class)
		}

		foundNames[property.Name] = true

		// Validate data type of property.
		schema := m.GetSchema()
		_, err := (&schema).FindPropertyDataType(property.DataType)
		if err != nil {
			return fmt.Errorf("Data type fo property '%s' is invalid; %v", property.Name, err)
		}
	}

	// all is fine!
	return nil
}

func upperCaseClassName(name string) string {
	if len(name) < 1 {
		return name
	}

	if len(name) == 1 {
		return strings.ToUpper(name)
	}

	return strings.ToUpper(string(name[0])) + name[1:]
}

func lowerCaseAllPropertyNames(props []*models.SemanticSchemaClassProperty) []*models.SemanticSchemaClassProperty {
	for i, prop := range props {
		props[i].Name = lowerCaseFirstLetter(prop.Name)
	}

	return props
}

func lowerCaseFirstLetter(name string) string {
	if len(name) < 1 {
		return name
	}

	if len(name) == 1 {
		return strings.ToLower(name)
	}

	return strings.ToLower(string(name[0])) + name[1:]
}
