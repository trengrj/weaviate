//                           _       _
// __      _____  __ ___   ___  __ _| |_ ___
// \ \ /\ / / _ \/ _` \ \ / / |/ _` | __/ _ \
//  \ V  V /  __/ (_| |\ V /| | (_| | ||  __/
//   \_/\_/ \___|\__,_| \_/ |_|\__,_|\__\___|
//
//  Copyright © 2016 - 2022 SeMI Technologies B.V. All rights reserved.
//
//  CONTACT: hello@semi.technology
//

package sharding

import (
	"crypto/rand"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestState(t *testing.T) {
	size := 1000

	cfg, err := ParseConfig(map[string]interface{}{"desiredCount": float64(4)}, 14)
	require.Nil(t, err)

	nodes := fakeNodes{[]string{"node1", "node2"}}
	state, err := InitState("my-index", cfg, nodes)
	require.Nil(t, err)

	physicalCount := map[string]int{}
	var names [][]byte

	for i := 0; i < size; i++ {
		name := make([]byte, 16)
		rand.Read(name)
		names = append(names, name)

		phid := state.PhysicalShard(name)
		physicalCount[phid]++
	}

	// verify each shard contains at least 15% of data. The expected value would
	// be 25%, but since this is random, we should take a lower value to reduce
	// flakyness

	for name, count := range physicalCount {
		if owns := float64(count) / float64(size); owns < 0.15 {
			t.Errorf("expected shard %q to own at least 15%%, but it only owns %f", name, owns)
		}
	}

	// Marshal and recreate, verify results
	bytes, err := state.JSON()
	require.Nil(t, err)

	// destroy old version
	state = nil

	stateReloaded, err := StateFromJSON(bytes, nodes)
	require.Nil(t, err)

	physicalCountReloaded := map[string]int{}

	// hash the same values again and verify the counts are exactly the same
	for _, name := range names {
		phid := stateReloaded.PhysicalShard(name)
		physicalCountReloaded[phid]++
	}

	assert.Equal(t, physicalCount, physicalCountReloaded)
}

type fakeNodes struct {
	nodes []string
}

func (f fakeNodes) AllNames() []string {
	return f.nodes
}

func (f fakeNodes) LocalName() string {
	return f.nodes[0]
}

func TestAdjustReplicas(t *testing.T) {
	t.Run("scaling up from 1 to 3", func(t *testing.T) {
		nodes := fakeNodes{nodes: []string{"node1", "node2", "node3"}}
		shard := Physical{BelongsToNodes: []string{"node1"}}
		expected := Physical{BelongsToNodes: []string{"node1", "node2", "node3"}}

		require.Nil(t, shard.AdjustReplicas(3, nodes))
		assert.Equal(t, expected, shard)
	})

	t.Run("scaling up from 2 to 3", func(t *testing.T) {
		nodes := fakeNodes{nodes: []string{"node1", "node2", "node3"}}
		shard := Physical{BelongsToNodes: []string{"node2", "node3"}}
		expected := Physical{BelongsToNodes: []string{"node2", "node3", "node1"}}

		require.Nil(t, shard.AdjustReplicas(3, nodes))
		assert.Equal(t, expected, shard)
	})

	t.Run("scaling from 3 to 3 - no change required", func(t *testing.T) {
		nodes := fakeNodes{nodes: []string{"node1", "node2", "node3"}}
		shard := Physical{BelongsToNodes: []string{"node1", "node2", "node3"}}
		expected := Physical{BelongsToNodes: []string{"node1", "node2", "node3"}}

		require.Nil(t, shard.AdjustReplicas(3, nodes))
		assert.Equal(t, expected, shard)
	})

	t.Run("scaling down from 3 to 2", func(t *testing.T) {
		nodes := fakeNodes{nodes: []string{"node1", "node2", "node3"}}
		shard := Physical{BelongsToNodes: []string{"node1", "node2", "node3"}}
		expected := Physical{BelongsToNodes: []string{"node1", "node2"}}

		require.Nil(t, shard.AdjustReplicas(2, nodes))
		assert.Equal(t, expected, shard)
	})

	t.Run("attempting to scale to a negative value", func(t *testing.T) {
		nodes := fakeNodes{nodes: []string{"node1", "node2", "node3"}}
		shard := Physical{BelongsToNodes: []string{"node1", "node2", "node3"}}

		require.NotNil(t, shard.AdjustReplicas(-22, nodes))
	})
}
