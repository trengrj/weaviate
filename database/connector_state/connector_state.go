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
package connector_state

import (
	"encoding/json"
)

// Internal connector state
type StateManager interface {
	// Called by a connector when it has updated it's internal state that needs to be shared across all connectors in other Weaviate instances.
	SetState(state json.RawMessage) error

	// Used by a connector to get the initial state.
	GetInitialConnectorState() json.RawMessage

	// Link a connector to this state manager.
	// When the internal state of some connector is updated, this state connector will call SetState on the provided conn.
	SetStateConnector(conn Connector)
}

// Implemented by a connector
type Connector interface {
	// Called by the state manager, when an external state update is happening.
	SetState(state json.RawMessage)

	// Link a StateManager to this connector, so that when a connector is updating it's own state, it can propagate these changes to othr Weaviate instances.
	SetStateManager(manager StateManager)
}
