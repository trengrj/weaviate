/*                          _       _
 *__      _____  __ ___   ___  __ _| |_ ___
 *\ \ /\ / / _ \/ _` \ \ / / |/ _` | __/ _ \
 * \ V  V /  __/ (_| |\ V /| | (_| | ||  __/
 *  \_/\_/ \___|\__,_| \_/ |_|\__,_|\__\___|
 *
 * Copyright © 2016 - 2019 Weaviate. All rights reserved.
 * LICENSE: https://github.com/creativesoftwarefdn/weaviate/blob/develop/LICENSE.md
 * DESIGN & CONCEPT: Bob van Luijt (@bobvanluijt)
 * CONTACT: hello@creativesoftwarefdn.org
 */
package schema

import (
	"fmt"
	"time"

	schemaclient "github.com/creativesoftwarefdn/weaviate/client/schema"
	"github.com/creativesoftwarefdn/weaviate/database/schema"
	"github.com/creativesoftwarefdn/weaviate/network/common/peers"
)

func download(peer peers.Peer) (schema.Schema, error) {
	peerClient, err := peer.CreateClient()
	if err != nil {
		return schema.Schema{}, fmt.Errorf(
			"could not create client for %s: %s", peer.Name, err)
	}

	params := &schemaclient.WeaviateSchemaDumpParams{}
	params.WithTimeout(2 * time.Second)
	ok, err := peerClient.Schema.WeaviateSchemaDump(params)
	if err != nil {
		return schema.Schema{}, fmt.Errorf(
			"could not download schema from %s: %s", peer.Name, err)
	}

	return schema.Schema{
		Things:  ok.Payload.Things,
		Actions: ok.Payload.Actions,
	}, nil
}
