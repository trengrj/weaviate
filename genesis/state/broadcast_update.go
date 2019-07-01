/*                          _       _
 *__      _____  __ ___   ___  __ _| |_ ___
 *\ \ /\ / / _ \/ _` \ \ / / |/ _` | __/ _ \
 * \ V  V /  __/ (_| |\ V /| | (_| | ||  __/
 *  \_/\_/ \___|\__,_| \_/ |_|\__,_|\__\___|
 *
 * Copyright © 2016 - 2019 Weaviate. All rights reserved.
 * LICENSE WEAVIATE OPEN SOURCE: https://www.semi.technology/playbook/playbook/contract-weaviate-OSS.html
 * LICENSE WEAVIATE ENTERPRISE: https://www.semi.technology/playbook/contract-weaviate-enterprise.html
 * CONCEPT: Bob van Luijt (@bobvanluijt)
 * CONTACT: hello@semi.technology
 */
package state

import (
	weaviate_client "github.com/semi-technologies/weaviate/client"
	weaviate_p2p "github.com/semi-technologies/weaviate/client/p2_p"
	weaviate_models "github.com/semi-technologies/weaviate/entities/models"

	"net/url"

	log "github.com/sirupsen/logrus"
)

func broadcast_update(peer Peer, peers []Peer) {
	log.Debugf("Broadcasting peer update to %v", peer.Id)
	peer_uri, err := url.Parse(string(peer.URI()))

	if err != nil {
		log.Infof("Could not broadcast to peer %v; Peer URI is invalid (%v)", peer.Id, peer.URI())
		return
	}

	transport_config := weaviate_client.TransportConfig{
		Host:     peer_uri.Host,
		BasePath: peer_uri.Path,
		Schemes:  []string{peer_uri.Scheme},
	}

	peer_updates := make(weaviate_models.PeerUpdateList, 0)

	for _, peer := range peers {
		peer_update := weaviate_models.PeerUpdate{
			URI:        peer.URI(),
			ID:         peer.Id,
			Name:       peer.Name(),
			SchemaHash: peer.SchemaHash,
		}

		peer_updates = append(peer_updates, &peer_update)
	}

	client := weaviate_client.NewHTTPClientWithConfig(nil, &transport_config)
	params := weaviate_p2p.NewWeaviateP2pGenesisUpdateParams()
	params.Peers = peer_updates
	_, err = client.P2P.WeaviateP2pGenesisUpdate(params)
	if err != nil {
		log.Debugf("Failed to update %v, because %v", peer.Id, err)
	}
}
