package kinds

import (
	"context"

	"github.com/creativesoftwarefdn/weaviate/usecases/config"
	"github.com/creativesoftwarefdn/weaviate/usecases/network"
)

// BatchManager manages kind changes in batch at a use-case level , i.e.
// agnostic of underlying databases or storage providers
type BatchManager struct {
	network       network.Network
	config        *config.WeaviateConfig
	repo          batchAndGetRepo
	locks         locks
	schemaManager schemaManager
}

// Repo describes the requirements the kinds UC has to the connected database
type batchRepo interface {
	AddThingsBatch(ctx context.Context, things BatchThings) error
	AddActionsBatch(ctx context.Context, actions BatchActions) error
	AddBatchReferences(ctx context.Context, references BatchReferences) error
}

type batchAndGetRepo interface {
	batchRepo
	getRepo
}

// NewBatchManager creates a new manager
func NewBatchManager(repo Repo, locks locks, schemaManager schemaManager, network network.Network, config *config.WeaviateConfig) *BatchManager {
	return &BatchManager{
		network:       network,
		config:        config,
		repo:          repo,
		locks:         locks,
		schemaManager: schemaManager,
	}
}
