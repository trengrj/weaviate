package vectorizer

import (
	"testing"

	"github.com/semi-technologies/weaviate/entities/models"
	"github.com/semi-technologies/weaviate/usecases/modules"
	"github.com/stretchr/testify/assert"
)

func TestIndexChecker(t *testing.T) {
	t.Run("with all defaults", func(t *testing.T) {
		class := &models.Class{
			Class: "MyClass",
			Properties: []*models.Property{{
				Name: "someProp",
			}},
		}

		cfg := modules.NewClassBasedModuleConfig(class, "my-module")
		ic := NewIndexChecker(cfg)

		assert.True(t, ic.PropertyIndexed("someProp"))
		assert.False(t, ic.VectorizePropertyName("someProp"))
		assert.True(t, ic.VectorizeClassName())
	})

	t.Run("with all explicit config matching the defaults", func(t *testing.T) {
		class := &models.Class{
			Class: "MyClass",
			ModuleConfig: map[string]interface{}{
				"my-module": map[string]interface{}{
					"vectorizeClassName": true,
				},
			},
			Properties: []*models.Property{{
				Name: "someProp",
				ModuleConfig: map[string]interface{}{
					"my-module": map[string]interface{}{
						"skip":                  false,
						"vectorizePropertyName": false,
					},
				},
			}},
		}

		cfg := modules.NewClassBasedModuleConfig(class, "my-module")
		ic := NewIndexChecker(cfg)

		assert.True(t, ic.PropertyIndexed("someProp"))
		assert.False(t, ic.VectorizePropertyName("someProp"))
		assert.True(t, ic.VectorizeClassName())
	})

	t.Run("with all explicit config using non-default values", func(t *testing.T) {
		class := &models.Class{
			Class: "MyClass",
			ModuleConfig: map[string]interface{}{
				"my-module": map[string]interface{}{
					"vectorizeClassName": false,
				},
			},
			Properties: []*models.Property{{
				Name: "someProp",
				ModuleConfig: map[string]interface{}{
					"my-module": map[string]interface{}{
						"skip":                  true,
						"vectorizePropertyName": true,
					},
				},
			}},
		}

		cfg := modules.NewClassBasedModuleConfig(class, "my-module")
		ic := NewIndexChecker(cfg)

		assert.False(t, ic.PropertyIndexed("someProp"))
		assert.True(t, ic.VectorizePropertyName("someProp"))
		assert.False(t, ic.VectorizeClassName())
	})
}
