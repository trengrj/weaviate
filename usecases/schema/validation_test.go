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
	"testing"

	"github.com/semi-technologies/weaviate/entities/models"
	"github.com/semi-technologies/weaviate/entities/schema/kind"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_Validation_ClassNames(t *testing.T) {
	type testCase struct {
		input     string
		valid     bool
		storedAs  string
		name      string
		vectorize bool
	}

	// TODO: most or all of these contain text2vec-contextionary-specific logic
	// for all test cases keep in mind that the word "carrot" is not present in
	// the fake c11y, but every other word is.
	//
	// Additionally, the word "the" is a stopword
	//
	// all inputs represent class names (!)
	tests := []testCase{
		// valid names
		testCase{
			name:      "Single uppercase word present in the c11y",
			input:     "Car",
			valid:     true,
			storedAs:  "Car",
			vectorize: true,
		},
		testCase{
			name:      "Single lowercase word present in the c11y, stored as uppercase",
			input:     "car",
			valid:     true,
			storedAs:  "Car",
			vectorize: true,
		},
		testCase{
			name:      "combination of valid words starting with uppercase letter",
			input:     "CarGarage",
			valid:     true,
			storedAs:  "CarGarage",
			vectorize: true,
		},
		testCase{
			name:      "combination of valid words starting with lowercase letter, stored as uppercase",
			input:     "carGarage",
			valid:     true,
			storedAs:  "CarGarage",
			vectorize: true,
		},
		testCase{
			name:      "combination of valid words and stopwords, starting with uppercase",
			input:     "TheCarGarage",
			valid:     true,
			storedAs:  "TheCarGarage",
			vectorize: true,
		},
		testCase{
			name:      "combination of valid words and stopwords starting with lowercase letter, stored as uppercase",
			input:     "carTheGarage",
			valid:     true,
			storedAs:  "CarTheGarage",
			vectorize: true,
		},

		// inavlid names
		testCase{
			name:      "Single uppercase word NOT present in the c11y",
			input:     "Carrot",
			valid:     false,
			vectorize: true,
		},
		testCase{
			name:      "Single lowercase word NOT present in the c11y",
			input:     "carrot",
			valid:     false,
			vectorize: true,
		},
		testCase{
			name:      "Single uppercase stopword",
			input:     "The",
			valid:     false,
			vectorize: true,
		},
		testCase{
			name:      "Single lowercase stopword",
			input:     "the",
			valid:     false,
			vectorize: true,
		},
		testCase{
			name:      "combination of valid and invalid words, valid word first lowercased",
			input:     "potatoCarrot",
			valid:     false,
			vectorize: true,
		},
		testCase{
			name:      "combination of valid and invalid words, valid word first uppercased",
			input:     "PotatoCarrot",
			valid:     false,
			vectorize: true,
		},
		testCase{
			name:      "combination of valid and invalid words, invalid word first lowercased",
			input:     "carrotPotato",
			valid:     false,
			vectorize: true,
		},
		testCase{
			name:      "combination of valid and invalid words, invalid word first uppercased",
			input:     "CarrotPotato",
			valid:     false,
			vectorize: true,
		},
		testCase{
			name:      "combination of only stopwords, starting with lowercase",
			input:     "theThe",
			valid:     false,
			vectorize: true,
		},
		testCase{
			name:      "combination of only stopwords, starting with uppercase",
			input:     "TheThe",
			valid:     false,
			vectorize: true,
		},

		// vectorize turned off
		testCase{
			name:      "non-vectorized: combination of only stopwords, starting with uppercase",
			input:     "TheThe",
			valid:     true,
			vectorize: false,
			storedAs:  "TheThe",
		},
		testCase{
			name:      "non-vectorized: excluded word",
			input:     "carrot",
			valid:     true,
			vectorize: false,
			storedAs:  "Carrot",
		},
		testCase{
			name:      "empty class",
			input:     "",
			valid:     false,
			vectorize: false,
		},
	}

	t.Run("adding a class", func(t *testing.T) {
		t.Run("different class names without keywords or properties", func(t *testing.T) {
			for _, test := range tests {
				t.Run(test.name+" as thing class", func(t *testing.T) {
					moduleConfig := map[string]interface{}{
						"text2vec-contextionary": map[string]interface{}{
							"vectorizeClassName": test.vectorize,
						},
					}

					class := &models.Class{
						Vectorizer:   "text2vec-contextionary",
						Class:        test.input,
						ModuleConfig: moduleConfig,
						Properties: []*models.Property{
							&models.Property{
								Name:     "dummyPropSoWeDontRunIntoAllNoindexedError",
								DataType: []string{"string"},
							},
						},
					}

					m := newSchemaManager()
					err := m.AddObject(context.Background(), nil, class)
					t.Log(err)
					assert.Equal(t, test.valid, err == nil)

					// only proceed if input was supposed to be valid
					if test.valid == false {
						return
					}

					classNames := testGetClassNames(m, kind.Object)
					assert.Contains(t, classNames, test.storedAs, "class should be stored correctly")
				})
			}
		})

		t.Run("different class names with valid keywords", func(t *testing.T) {
			for _, test := range tests {
				t.Run(test.name+" as thing class", func(t *testing.T) {
					moduleConfig := map[string]interface{}{
						"text2vec-contextionary": map[string]interface{}{
							"vectorizeClassName": test.vectorize,
						},
					}
					class := &models.Class{
						Vectorizer:   "text2vec-contextionary",
						Class:        test.input,
						ModuleConfig: moduleConfig,
						Properties: []*models.Property{
							&models.Property{
								Name:     "dummyPropSoWeDontRunIntoAllNoindexedError",
								DataType: []string{"string"},
							},
						},
					}

					m := newSchemaManager()
					err := m.AddObject(context.Background(), nil, class)
					t.Log(err)
					assert.Equal(t, test.valid, err == nil)

					// only proceed if input was supposed to be valid
					if test.valid == false {
						return
					}

					classNames := testGetClassNames(m, kind.Object)
					assert.Contains(t, classNames, test.storedAs, "class should be stored correctly")
				})
			}
		})
	})

	t.Run("updating an existing class", func(t *testing.T) {
		t.Run("different class names without keywords or properties", func(t *testing.T) {
			for _, test := range tests {
				originalName := "ValidOriginalName"
				t.Run(test.name+" as thing class", func(t *testing.T) {
					moduleConfig := map[string]interface{}{
						"text2vec-contextionary": map[string]interface{}{
							"vectorizeClassName": test.vectorize,
						},
					}
					class := &models.Class{
						Vectorizer:   "text2vec-contextionary",
						Class:        originalName,
						ModuleConfig: moduleConfig,
						Properties: []*models.Property{
							&models.Property{
								Name:     "dummyPropSoWeDontRunIntoAllNoindexedError",
								DataType: []string{"string"},
							},
						},
					}

					m := newSchemaManager()
					err := m.AddObject(context.Background(), nil, class)
					require.Nil(t, err)

					// now try to update
					updatedClass := &models.Class{
						Vectorizer: "text2vec-contextionary",
						Class:      test.input,
					}

					err = m.UpdateObject(context.Background(), nil, originalName, updatedClass)
					assert.Equal(t, test.valid, err == nil)

					// only proceed if input was supposed to be valid
					if test.valid == false {
						return
					}

					classNames := testGetClassNames(m, kind.Object)
					assert.Contains(t, classNames, test.storedAs, "class should be stored correctly")
				})

			}
		})

		t.Run("different class names with valid keywords", func(t *testing.T) {
			for _, test := range tests {
				originalName := "ValidOriginalName"

				t.Run(test.name+" as thing class", func(t *testing.T) {
					moduleConfig := map[string]interface{}{
						"text2vec-contextionary": map[string]interface{}{
							"vectorizeClassName": test.vectorize,
						},
					}
					class := &models.Class{
						Vectorizer:   "text2vec-contextionary",
						Class:        originalName,
						ModuleConfig: moduleConfig,
						Properties: []*models.Property{
							&models.Property{
								Name:     "dummyPropSoWeDontRunIntoAllNoindexedError",
								DataType: []string{"string"},
							},
						},
					}

					m := newSchemaManager()
					err := m.AddObject(context.Background(), nil, class)
					require.Nil(t, err)

					// now update
					updatedClass := &models.Class{
						Vectorizer: "text2vec-contextionary",
						Class:      test.input,
					}
					err = m.UpdateObject(context.Background(), nil, originalName, updatedClass)
					assert.Equal(t, test.valid, err == nil)

					// only proceed if input was supposed to be valid
					if test.valid == false {
						return
					}

					classNames := testGetClassNames(m, kind.Object)
					assert.Contains(t, classNames, test.storedAs, "class should be stored correctly")
				})
			}
		})
	})
}

func Test_Validation_PropertyNames(t *testing.T) {
	type testCase struct {
		input     string
		valid     bool
		storedAs  string
		name      string
		vectorize bool
	}

	// for all test cases keep in mind that the word "carrot" is not present in
	// the fake c11y, but every other word is
	//
	// all inputs represent property names (!)
	tests := []testCase{
		// valid names
		testCase{
			name:      "Single uppercase word present in the c11y, stored as lowercase",
			input:     "Brand",
			valid:     true,
			storedAs:  "brand",
			vectorize: true,
		},
		testCase{
			name:      "Single lowercase word present in the c11y",
			input:     "brand",
			valid:     true,
			storedAs:  "brand",
			vectorize: true,
		},
		testCase{
			name:      "combination of valid words starting with uppercase letter, stored as lowercase",
			input:     "BrandGarage",
			valid:     true,
			storedAs:  "brandGarage",
			vectorize: true,
		},
		testCase{
			name:      "combination of valid words starting with lowercase letter",
			input:     "brandGarage",
			valid:     true,
			storedAs:  "brandGarage",
			vectorize: true,
		},
		testCase{
			name:      "combination of valid words and stop words starting with uppercase letter, stored as lowercase",
			input:     "TheGarage",
			valid:     true,
			storedAs:  "theGarage",
			vectorize: true,
		},
		testCase{
			name:      "combination of valid words and stop words starting with lowercase letter",
			input:     "theGarage",
			valid:     true,
			storedAs:  "theGarage",
			vectorize: true,
		},

		// inavlid names
		testCase{
			name:      "Single uppercase word NOT present in the c11y",
			input:     "Carrot",
			valid:     false,
			vectorize: true,
		},
		testCase{
			name:      "Single lowercase word NOT present in the c11y",
			input:     "carrot",
			valid:     false,
			vectorize: true,
		},
		testCase{
			name:      "Single lowercase stop word",
			input:     "the",
			valid:     false,
			vectorize: true,
		},
		testCase{
			name:      "combination of valid and invalid words, valid word first lowercased",
			input:     "potatoCarrot",
			valid:     false,
			vectorize: true,
		},
		testCase{
			name:      "combination of valid and invalid words, valid word first uppercased",
			input:     "PotatoCarrot",
			valid:     false,
			vectorize: true,
		},
		testCase{
			name:      "combination of valid and invalid words, invalid word first lowercased",
			input:     "carrotPotato",
			valid:     false,
			vectorize: true,
		},
		testCase{
			name:      "combination of valid and invalid words, invalid word first uppercased",
			input:     "CarrotPotato",
			valid:     false,
			vectorize: true,
		},
		testCase{
			name:      "combination of only stop words,  first lowercased",
			input:     "theThe",
			valid:     false,
			vectorize: true,
		},
		testCase{
			name:      "combination of only stop words, first uppercased",
			input:     "TheThe",
			valid:     false,
			vectorize: true,
		},

		// without vectorizing
		testCase{
			name:      "non-vectorizing: combination of only stop words, first uppercased",
			input:     "TheThe",
			valid:     true,
			vectorize: false,
			storedAs:  "theThe",
		},
		testCase{
			name:      "non-vectorizing: combination of only stop words, first uppercased",
			input:     "carrot",
			valid:     true,
			vectorize: false,
			storedAs:  "carrot",
		},
		testCase{
			name:      "non-vectorizing: empty",
			input:     "",
			valid:     false,
			vectorize: false,
		},
	}

	t.Run("when adding a new class", func(t *testing.T) {
		t.Run("different property names without keywords for the prop", func(t *testing.T) {
			for _, test := range tests {
				t.Run(test.name+" as thing class", func(t *testing.T) {
					class := &models.Class{
						Vectorizer: "text2vec-contextionary",
						Class:      "ValidName",
						Properties: []*models.Property{{
							DataType: []string{"string"},
							Name:     test.input,
							ModuleConfig: map[string]interface{}{
								"text2vec-contextionary": map[string]interface{}{
									"vectorizePropertyName": test.vectorize,
								},
							},
						}},
					}

					m := newSchemaManager()
					err := m.AddObject(context.Background(), nil, class)
					t.Log(err)
					assert.Equal(t, test.valid, err == nil)

					// only proceed if input was supposed to be valid
					if test.valid == false {
						return
					}

					schema, _ := m.GetSchema(nil)
					propName := schema.Objects.Classes[0].Properties[0].Name
					assert.Equal(t, propName, test.storedAs, "class should be stored correctly")
				})
			}
		})

		t.Run("different property names  with valid keywords for the prop", func(t *testing.T) {
			for _, test := range tests {
				t.Run(test.name+" as thing class", func(t *testing.T) {
					class := &models.Class{
						Vectorizer: "text2vec-contextionary",
						Class:      "ValidName",
						Properties: []*models.Property{{
							DataType: []string{"string"},
							Name:     test.input,
							ModuleConfig: map[string]interface{}{
								"text2vec-contextionary": map[string]interface{}{
									"vectorizePropertyName": test.vectorize,
								},
							},
						}},
					}

					m := newSchemaManager()
					err := m.AddObject(context.Background(), nil, class)
					t.Log(err)
					assert.Equal(t, test.valid, err == nil)

					// only proceed if input was supposed to be valid
					if test.valid == false {
						return
					}

					schema, _ := m.GetSchema(nil)
					propName := schema.Objects.Classes[0].Properties[0].Name
					assert.Equal(t, propName, test.storedAs, "class should be stored correctly")
				})
			}
		})
	})

	t.Run("when updating an existing class with a new property", func(t *testing.T) {
		t.Run("different property names without keywords for the prop", func(t *testing.T) {
			for _, test := range tests {
				t.Run(test.name+" as thing class", func(t *testing.T) {
					class := &models.Class{
						Vectorizer: "text2vec-contextionary",
						Class:      "ValidName",
						Properties: []*models.Property{
							&models.Property{
								Name:     "dummyPropSoWeDontRunIntoAllNoindexedError",
								DataType: []string{"string"},
							},
						},
					}

					m := newSchemaManager()
					err := m.AddObject(context.Background(), nil, class)
					require.Nil(t, err)

					property := &models.Property{
						DataType: []string{"string"},
						Name:     test.input,
						ModuleConfig: map[string]interface{}{
							"text2vec-contextionary": map[string]interface{}{
								"vectorizePropertyName": test.vectorize,
							},
						},
					}
					err = m.AddObjectProperty(context.Background(), nil, "ValidName", property)
					t.Log(err)
					require.Equal(t, test.valid, err == nil)

					// only proceed if input was supposed to be valid
					if test.valid == false {
						return
					}

					schema, _ := m.GetSchema(nil)
					propName := schema.Objects.Classes[0].Properties[1].Name
					assert.Equal(t, propName, test.storedAs, "class should be stored correctly")
				})
			}
		})

		t.Run("different property names  with valid keywords for the prop", func(t *testing.T) {
			for _, test := range tests {
				t.Run(test.name+" as thing class", func(t *testing.T) {
					class := &models.Class{
						Vectorizer: "text2vec-contextionary",
						Class:      "ValidName",
						Properties: []*models.Property{{
							DataType: []string{"string"},
							Name:     test.input,
							ModuleConfig: map[string]interface{}{
								"text2vec-contextionary": map[string]interface{}{
									"vectorizePropertyName": test.vectorize,
								},
							},
						}},
					}

					m := newSchemaManager()
					err := m.AddObject(context.Background(), nil, class)
					t.Log(err)
					assert.Equal(t, test.valid, err == nil)

					// only proceed if input was supposed to be valid
					if test.valid == false {
						return
					}

					schema, _ := m.GetSchema(nil)
					propName := schema.Objects.Classes[0].Properties[0].Name
					assert.Equal(t, propName, test.storedAs, "class should be stored correctly")
				})
			}
		})
	})

	t.Run("when updating an existing property with a new prop name", func(t *testing.T) {
		t.Run("different property names without keywords for the prop", func(t *testing.T) {
			for _, test := range tests {
				originalName := "validPropertyName"

				t.Run(test.name+" as thing class", func(t *testing.T) {
					class := &models.Class{
						Vectorizer: "text2vec-contextionary",
						Class:      "ValidName",
						Properties: []*models.Property{
							&models.Property{
								DataType: []string{"string"},
								ModuleConfig: map[string]interface{}{
									"text2vec-contextionary": map[string]interface{}{
										"vectorizePropertyName": test.vectorize,
									},
								},
								Name: originalName,
							},
						},
					}

					m := newSchemaManager()
					err := m.AddObject(context.Background(), nil, class)
					require.Nil(t, err)

					updatedProperty := &models.Property{
						DataType: []string{"string"},
						Name:     test.input,
					}
					err = m.UpdateObjectProperty(context.Background(), nil, "ValidName", originalName, updatedProperty)
					t.Log(err)
					require.Equal(t, test.valid, err == nil)

					// only proceed if input was supposed to be valid
					if test.valid == false {
						return
					}

					schema, _ := m.GetSchema(nil)
					propName := schema.Objects.Classes[0].Properties[0].Name
					assert.Equal(t, propName, test.storedAs, "class should be stored correctly")
				})
			}
		})

		t.Run("different property names  with valid keywords for the prop", func(t *testing.T) {
			for _, test := range tests {
				t.Run(test.name+" as thing class", func(t *testing.T) {
					class := &models.Class{
						Vectorizer: "text2vec-contextionary",
						Class:      "ValidName",
						Properties: []*models.Property{{
							DataType: []string{"string"},
							Name:     test.input,
							ModuleConfig: map[string]interface{}{
								"text2vec-contextionary": map[string]interface{}{
									"vectorizePropertyName": test.vectorize,
								},
							},
						}},
					}

					m := newSchemaManager()
					err := m.AddObject(context.Background(), nil, class)
					t.Log(err)
					assert.Equal(t, test.valid, err == nil)

					// only proceed if input was supposed to be valid
					if test.valid == false {
						return
					}

					schema, _ := m.GetSchema(nil)
					propName := schema.Objects.Classes[0].Properties[0].Name
					assert.Equal(t, propName, test.storedAs, "class should be stored correctly")
				})
			}
		})
	})
}

func Test_AllUsablePropsNoindexed(t *testing.T) {
	t.Run("all schema vectorization turned off", func(t *testing.T) {
		class := &models.Class{
			Vectorizer: "text2vec-contextionary",
			Class:      "ValidName",
			ModuleConfig: map[string]interface{}{
				"text2vec-contextionary": map[string]interface{}{
					"vectorizeClassName": false,
				},
			},
			Properties: []*models.Property{
				{
					DataType: []string{"text"},
					Name:     "decsription",
					ModuleConfig: map[string]interface{}{
						"text2vec-contextionary": map[string]interface{}{
							"vectorizeClassName": false,
							"skip":               true,
						},
					},
				},
				{
					DataType: []string{"string"},
					Name:     "name",
					ModuleConfig: map[string]interface{}{
						"text2vec-contextionary": map[string]interface{}{
							"vectorizeClassName": false,
							"skip":               true,
						},
					},
				},
				{
					DataType: []string{"int"},
					Name:     "amount",
					ModuleConfig: map[string]interface{}{
						"text2vec-contextionary": map[string]interface{}{
							"vectorizeClassName": false,
							"skip":               true,
						},
					},
				},
			},
		}

		m := newSchemaManager()
		err := m.AddObject(context.Background(), nil, class)
		assert.NotNil(t, err)
	})
}
