package schema

import (
	"encoding/json"
	"testing"

	"github.com/semi-technologies/weaviate/entities/models"
	"github.com/semi-technologies/weaviate/usecases/cluster"
	"github.com/semi-technologies/weaviate/usecases/config"
	"github.com/sirupsen/logrus/hooks/test"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestStartupSync(t *testing.T) {
	t.Run("new node joining, other nodes have schema", func(t *testing.T) {
		clusterState := &fakeClusterState{
			hosts: []string{"node1", "node2"},
		}

		txJSON, _ := json.Marshal(ReadSchemaPayload{
			Schema: &State{
				ObjectSchema: &models.Schema{
					Classes: []*models.Class{
						{
							Class:           "Bongourno",
							VectorIndexType: "hnsw",
						},
					},
				},
			},
		})

		txClient := &fakeTxClient{
			openInjectPayload: json.RawMessage(txJSON),
		}

		sm, err := newManagerWithClusterAndTx(t, clusterState, txClient, nil)
		require.Nil(t, err)

		localSchema := sm.GetSchemaSkipAuth()
		assert.Equal(t, "Bongourno", localSchema.FindClassByName("Bongourno").Class)
	})

	t.Run("new node joining, other nodes have no schema", func(t *testing.T) {
		clusterState := &fakeClusterState{
			hosts: []string{"node1", "node2"},
		}

		txJSON, _ := json.Marshal(ReadSchemaPayload{
			Schema: &State{
				ObjectSchema: &models.Schema{
					Classes: []*models.Class{},
				},
			},
		})

		txClient := &fakeTxClient{
			openInjectPayload: json.RawMessage(txJSON),
		}

		sm, err := newManagerWithClusterAndTx(t, clusterState, txClient, nil)
		require.Nil(t, err)

		localSchema := sm.GetSchemaSkipAuth()
		assert.Len(t, localSchema.Objects.Classes, 0)
	})

	t.Run("new node joining, conflict in schema between nodes", func(t *testing.T) {
		clusterState := &fakeClusterState{
			hosts: []string{"node1", "node2"},
		}

		txJSON, _ := json.Marshal(ReadSchemaPayload{
			Schema: &State{
				ObjectSchema: &models.Schema{
					Classes: []*models.Class{
						{
							Class:           "Bongourno",
							VectorIndexType: "hnsw",
						},
					},
				},
			},
		})

		txClient := &fakeTxClient{
			openInjectPayload: json.RawMessage(txJSON),
		}

		_, err := newManagerWithClusterAndTx(t, clusterState, txClient, &State{
			ObjectSchema: &models.Schema{
				Classes: []*models.Class{
					{
						Class:           "Hola",
						VectorIndexType: "hnsw",
					},
				},
			},
		})
		require.NotNil(t, err)
		assert.Contains(t, err.Error(), "corrupt")
	})
}

func newManagerWithClusterAndTx(t *testing.T, clusterState clusterState,
	txClient cluster.Client, initialSchema *State,
) (*Manager, error) {
	logger, _ := test.NewNullLogger()
	repo := newFakeRepo()
	repo.schema = initialSchema
	sm, err := NewManager(&NilMigrator{}, repo, logger, &fakeAuthorizer{},
		config.Config{DefaultVectorizerModule: config.VectorizerModuleNone},
		dummyParseVectorConfig, // only option for now
		&fakeVectorizerValidator{}, dummyValidateInvertedConfig,
		&fakeModuleConfig{}, clusterState, txClient,
	)

	return sm, err
}
