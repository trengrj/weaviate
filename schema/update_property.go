package schema

import (
	"context"

	"github.com/creativesoftwarefdn/weaviate/database/schema/kind"
	"github.com/creativesoftwarefdn/weaviate/models"
)

// UpdateActionProperty of an existing Action Property
func (m *Manager) UpdateActionProperty(ctx context.Context, class string, name string,
	property *models.SemanticSchemaClassProperty) error {
	return m.updateClassProperty(ctx, class, name, property, kind.ACTION_KIND)
}

// UpdateThingProperty of an existing Thing Property
func (m *Manager) UpdateThingProperty(ctx context.Context, class string, name string,
	property *models.SemanticSchemaClassProperty) error {
	return m.updateClassProperty(ctx, class, name, property, kind.THING_KIND)
}

// TODO: gh-832: Implement full capabilities, not just keywords/naming
func (m *Manager) updateClassProperty(ctx context.Context, class string, name string,
	property *models.SemanticSchemaClassProperty, k kind.Kind) error {
	schemaLock, err := m.db.SchemaLock()
	if err != nil {
		return err
	}
	defer unlock(schemaLock)

	schemaManager := schemaLock.SchemaManager()
	var newName *string
	var newKeywords *models.SemanticSchemaKeywords

	if property.Name != name {
		// the name in the URI and body don't match, so we assume the user wants to rename
		newName = &property.Name
	}

	// TODO gh-619: This implies that we can't undo setting keywords, because we can't detect if keywords is not present, or empty.
	if len(property.Keywords) > 0 {
		newKeywords = &property.Keywords
	}
	err = schemaManager.UpdateProperty(ctx, k, class, name, newName, newKeywords)
	if err != nil {
		return err
	}

	return nil
}
