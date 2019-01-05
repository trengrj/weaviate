/*                          _       _
 *__      _____  __ ___   ___  __ _| |_ ___
 *\ \ /\ / / _ \/ _` \ \ / / |/ _` | __/ _ \
 * \ V  V /  __/ (_| |\ V /| | (_| | ||  __/
 *  \_/\_/ \___|\__,_| \_/ |_|\__,_|\__\___|
 *
 * Copyright © 2016 - 2019 Weaviate. All rights reserved.
 * LICENSE: https://github.com/creativesoftwarefdn/weaviate/blob/develop/LICENSE.md
 * DESIGN: Bob van Luijt (bob@k10y.co)
 */
package schema_migrator

import (
	"github.com/creativesoftwarefdn/weaviate/database/schema/kind"
	"github.com/creativesoftwarefdn/weaviate/models"
)

type Migrator interface {
	// Add a class to the Thing or Action schema, depending on the kind parameter.
	AddClass(kind kind.Kind, class *models.SemanticSchemaClass) error

	// Drop a class from the schema.
	DropClass(kind kind.Kind, className string) error

	// Update a given class. If newClassName is not nil, update the class name, if newKeywords is not nil, update the keywords.
	// If both updates are specified, either both updates succeed, or none do.
	UpdateClass(kind kind.Kind, className string, newClassName *string, newKeywords *models.SemanticSchemaKeywords) error

	// Add a property to a given class.
	AddProperty(kind kind.Kind, className string, prop *models.SemanticSchemaClassProperty) error

	// Update a given property. If newName is not nil, update the property name, if newKeywords is not nil, update the keywords.
	// If both updates are specified, either both updates succeed, or none do.
	UpdateProperty(kind kind.Kind, className string, propName string, newName *string, newKeywords *models.SemanticSchemaKeywords) error

	// Update a given property. Idempotently add a dataType to the list of dataTypes
	UpdatePropertyAddDataType(kind kind.Kind, className string, propName string, newDataType string) error

	// Drop the given property from the schema
	DropProperty(kind kind.Kind, className string, propName string) error
}
