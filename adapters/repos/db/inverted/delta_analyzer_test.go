package inverted

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDeltaAnalyzer(t *testing.T) {
	t.Run("without previous indexing", func(t *testing.T) {
		previous := []Property(nil)
		next := []Property{
			Property{
				Name: "prop1",
				Items: []Countable{
					{
						Data:          []byte("value1"),
						TermFrequency: 7,
					},
					{
						Data:          []byte("value2"),
						TermFrequency: 3,
					},
				},
			},
			Property{
				Name: "prop2",
				Items: []Countable{
					{
						Data:          []byte("value3"),
						TermFrequency: 7,
					},
					{
						Data:          []byte("value4"),
						TermFrequency: 3,
					},
				},
			},
		}

		res := Delta(previous, next)
		assert.Equal(t, next, res.ToAdd)
		assert.Len(t, res.ToDelete, 0)
	})

	t.Run("with previous indexing and no changes", func(t *testing.T) {
		previous := []Property{
			Property{
				Name: "prop1",
				Items: []Countable{
					{
						Data:          []byte("value1"),
						TermFrequency: 7,
					},
					{
						Data:          []byte("value2"),
						TermFrequency: 3,
					},
				},
			},
			Property{
				Name: "prop2",
				Items: []Countable{
					{
						Data:          []byte("value3"),
						TermFrequency: 7,
					},
					{
						Data:          []byte("value4"),
						TermFrequency: 3,
					},
				},
			},
		}
		next := []Property{
			Property{
				Name: "prop1",
				Items: []Countable{
					{
						Data:          []byte("value1"),
						TermFrequency: 7,
					},
					{
						Data:          []byte("value2"),
						TermFrequency: 3,
					},
				},
			},
			Property{
				Name: "prop2",
				Items: []Countable{
					{
						Data:          []byte("value3"),
						TermFrequency: 7,
					},
					{
						Data:          []byte("value4"),
						TermFrequency: 3,
					},
				},
			},
		}

		res := Delta(previous, next)
		assert.Len(t, res.ToDelete, 0)
		assert.Len(t, res.ToAdd, 0)
	})

	t.Run("with previous indexing - only additions", func(t *testing.T) {
		previous := []Property{
			Property{
				Name: "prop1",
				Items: []Countable{
					{
						Data:          []byte("value2"),
						TermFrequency: 3,
					},
				},
			},
			Property{
				Name: "prop2",
				Items: []Countable{
					{
						Data:          []byte("value4"),
						TermFrequency: 3,
					},
				},
			},
		}
		next := []Property{
			Property{
				Name: "prop1",
				Items: []Countable{
					{
						Data:          []byte("value1"),
						TermFrequency: 7,
					},
					{
						Data:          []byte("value2"),
						TermFrequency: 3,
					},
				},
			},
			Property{
				Name: "prop2",
				Items: []Countable{
					{
						Data:          []byte("value3"),
						TermFrequency: 7,
					},
					{
						Data:          []byte("value4"),
						TermFrequency: 3,
					},
				},
			},
		}

		expectedAdd := []Property{
			Property{
				Name: "prop1",
				Items: []Countable{
					{
						Data:          []byte("value1"),
						TermFrequency: 7,
					},
				},
			},
			Property{
				Name: "prop2",
				Items: []Countable{
					{
						Data:          []byte("value3"),
						TermFrequency: 7,
					},
				},
			},
		}

		res := Delta(previous, next)
		assert.Equal(t, expectedAdd, res.ToAdd)
		assert.Len(t, res.ToDelete, 0)
	})

	t.Run("with previous indexing - both additions and deletions", func(t *testing.T) {
		previous := []Property{
			Property{
				Name: "prop1",
				Items: []Countable{
					{
						Data:          []byte("value2"),
						TermFrequency: 3,
					},
				},
			},
			Property{
				Name: "prop2",
				Items: []Countable{
					{
						Data:          []byte("value4"),
						TermFrequency: 3,
					},
				},
			},
		}
		next := []Property{
			Property{
				Name: "prop1",
				Items: []Countable{
					{
						Data:          []byte("value1"),
						TermFrequency: 7,
					},
				},
			},
			Property{
				Name: "prop2",
				Items: []Countable{
					{
						Data:          []byte("value3"),
						TermFrequency: 7,
					},
					{
						Data:          []byte("value4"),
						TermFrequency: 3,
					},
				},
			},
		}

		expectedAdd := []Property{
			Property{
				Name: "prop1",
				Items: []Countable{
					{
						Data:          []byte("value1"),
						TermFrequency: 7,
					},
				},
			},
			Property{
				Name: "prop2",
				Items: []Countable{
					{
						Data:          []byte("value3"),
						TermFrequency: 7,
					},
				},
			},
		}

		expectedDelete := []Property{
			{
				Name: "prop1",
				Items: []Countable{
					{
						Data:          []byte("value2"),
						TermFrequency: 3,
					},
				},
			},
		}

		res := Delta(previous, next)
		assert.Equal(t, expectedAdd, res.ToAdd)
		assert.Equal(t, expectedDelete, res.ToDelete)
	})
}
