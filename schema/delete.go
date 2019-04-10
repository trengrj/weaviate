package schema

import (
	"context"

	"github.com/creativesoftwarefdn/weaviate/database/schema/kind"
)

// DeleteAction Class to the schema
func (m *Manager) DeleteAction(ctx context.Context, class string) error {
	return m.deleteClass(ctx, class, kind.ACTION_KIND)
}

// DeleteThing Class to the schema
func (m *Manager) DeleteThing(ctx context.Context, class string) error {
	return m.deleteClass(ctx, class, kind.THING_KIND)
}

func (m *Manager) deleteClass(ctx context.Context, class string, k kind.Kind) error {
	schemaLock, err := m.db.SchemaLock()
	if err != nil {
		return err
	}
	defer unlock(schemaLock)

	schemaManager := schemaLock.SchemaManager()
	err = schemaManager.DropClass(ctx, k, class)
	if err != nil {
		return err
	}

	return nil
}
