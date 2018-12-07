package local

import (
	"fmt"

	"github.com/creativesoftwarefdn/weaviate/contextionary"
	"github.com/creativesoftwarefdn/weaviate/database/schema"
	"github.com/creativesoftwarefdn/weaviate/database/schema/kind"
	"github.com/creativesoftwarefdn/weaviate/models"
	"github.com/go-openapi/strfmt"
)

func (l *localSchemaManager) GetSchema() schema.Schema {
	return schema.Schema{
		Actions: l.schemaState.ActionSchema,
		Things:  l.schemaState.ThingSchema,
	}
}

func (l *localSchemaManager) UpdateMeta(kind kind.Kind, atContext strfmt.URI, maintainer strfmt.Email, name string) error {
	semanticSchema := l.schemaState.SchemaFor(kind)
	semanticSchema.AtContext = atContext
	semanticSchema.Maintainer = maintainer
	semanticSchema.Name = name

	return l.saveToDisk()
}

func (l *localSchemaManager) SetContextionary(context contextionary.Contextionary) {
	l.contextionary = context
}

func (l *localSchemaManager) AddClass(kind kind.Kind, class *models.SemanticSchemaClass) error {
	err := l.validateCanAddClass(kind, class)
	if err != nil {
		return err
	} else {
		semanticSchema := l.schemaState.SchemaFor(kind)
		semanticSchema.Classes = append(semanticSchema.Classes, class)
		err := l.saveToDisk()
		if err != nil {
			return err
		}

		return l.connectorMigrator.AddClass(kind, class)
	}
}

func (l *localSchemaManager) DropClass(kind kind.Kind, className string) error {
	semanticSchema := l.schemaState.SchemaFor(kind)

	var classIdx int = -1
	for idx, class := range semanticSchema.Classes {
		if class.Class == className {
			classIdx = idx
			break
		}
	}

	if classIdx == -1 {
		return fmt.Errorf("Could not find class '%s'", className)
	}

	semanticSchema.Classes[classIdx] = semanticSchema.Classes[len(semanticSchema.Classes)-1]
	semanticSchema.Classes[len(semanticSchema.Classes)-1] = nil // to prevent leaking this pointer.
	semanticSchema.Classes = semanticSchema.Classes[:len(semanticSchema.Classes)-1]

	err := l.saveToDisk()
	if err != nil {
		return err
	}

	return l.connectorMigrator.DropClass(kind, className)
}

func (l *localSchemaManager) UpdateClass(kind kind.Kind, className string, newClassName *string, newKeywords *models.SemanticSchemaKeywords) error {
	semanticSchema := l.schemaState.SchemaFor(kind)

	class, err := schema.GetClassByName(semanticSchema, className)
	if err != nil {
		return err
	}

	classNameAfterUpdate := className
	keywordsAfterUpdate := class.Keywords

	// First validate the request
	if newClassName != nil {
		err = l.validateClassNameUniqueness(*newClassName)
		classNameAfterUpdate = *newClassName
		if err != nil {
			return err
		}
	}

	if newKeywords != nil {
		keywordsAfterUpdate = *newKeywords
	}

	// Validate name / keywords in contextionary
	l.validateClassNameOrKeywordsCorrect(kind, classNameAfterUpdate, keywordsAfterUpdate)

	// Validated! Now apply the changes.
	class.Class = classNameAfterUpdate
	class.Keywords = keywordsAfterUpdate

	err = l.saveToDisk()

	if err != nil {
		return nil
	}

	return l.connectorMigrator.UpdateClass(kind, className, newClassName, newKeywords)
}

func (l *localSchemaManager) AddProperty(kind kind.Kind, className string, prop *models.SemanticSchemaClassProperty) error {
	semanticSchema := l.schemaState.SchemaFor(kind)
	class, err := schema.GetClassByName(semanticSchema, className)
	if err != nil {
		return err
	}

	err = l.validateCanAddProperty(prop, class)
	if err != nil {
		return err
	}

	class.Properties = append(class.Properties, prop)

	err = l.saveToDisk()

	if err != nil {
		return nil
	}

	return l.connectorMigrator.AddProperty(kind, className, prop)
}

func (l *localSchemaManager) UpdateProperty(kind kind.Kind, className string, propName string, newName *string, newKeywords *models.SemanticSchemaKeywords) error {
	semanticSchema := l.schemaState.SchemaFor(kind)
	class, err := schema.GetClassByName(semanticSchema, className)

	if err != nil {
		return err
	}

	prop, err := schema.GetPropertyByName(class, propName)
	if err != nil {
		return err
	}

	propNameAfterUpdate := propName
	keywordsAfterUpdate := prop.Keywords

	if newName != nil {
		// verify uniqueness
		err = validatePropertyNameUniqueness(*newName, class)
		propNameAfterUpdate = *newName
		if err != nil {
			return err
		}
	}

	if newKeywords != nil {
		keywordsAfterUpdate = *newKeywords
	}

	// Validate name / keywords in contextionary
	l.validatePropertyNameOrKeywordsCorrect(className, propNameAfterUpdate, keywordsAfterUpdate)

	// Validated! Now apply the changes.
	prop.Name = propNameAfterUpdate
	prop.Keywords = keywordsAfterUpdate

	err = l.saveToDisk()

	if err != nil {
		return nil
	}

	return l.connectorMigrator.UpdateProperty(kind, className, propName, newName, newKeywords)
}

func (l *localSchemaManager) DropProperty(kind kind.Kind, className string, propName string) error {
	semanticSchema := l.schemaState.SchemaFor(kind)
	class, err := schema.GetClassByName(semanticSchema, className)
	if err != nil {
		return err
	}

	var propIdx int = -1
	for idx, prop := range class.Properties {
		if prop.Name == propName {
			propIdx = idx
			break
		}
	}

	if propIdx == -1 {
		return fmt.Errorf("Could not find property '%s'", propName)
	}

	class.Properties[propIdx] = class.Properties[len(class.Properties)-1]
	class.Properties[len(class.Properties)-1] = nil // to prevent leaking this pointer.
	class.Properties = class.Properties[:len(class.Properties)-1]

	err = l.saveToDisk()

	if err != nil {
		return nil
	}

	return l.connectorMigrator.DropProperty(kind, className, propName)
}
