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
package p2p

import (
	"reflect"
	"testing"

	"github.com/creativesoftwarefdn/weaviate/network/common/peers"
)

func TestPeerUpdateWithNewPeers(t *testing.T) {
	oldPeers := []peers.Peer{}
	newPeers := []peers.Peer{{
		Name: "best-weaviate",
		ID:   "uuid",
		URI:  "does-not-matter",
	}}

	subject := network{
		peers:           oldPeers,
		downloadChanged: downloadChangedFake(newPeers),
	}

	callbackCalled := false
	callbackCalledWith := []peers.Peer{}
	callbackSpy := func(peers peers.Peers) {
		callbackCalled = true
		callbackCalledWith = peers
	}

	subject.RegisterUpdatePeerCallback(callbackSpy)
	subject.UpdatePeers(newPeers)

	if callbackCalled != true {
		t.Error("expect PeerUpdateCallback to be called, but was never called")
	}

	if !reflect.DeepEqual(callbackCalledWith, newPeers) {
		t.Errorf("expect PeerUpdateCallback to be called with new peers, but was called with %#v",
			callbackCalledWith)
	}
}

func TestPeerUpdateWithoutAnyChange(t *testing.T) {
	unchangedPeers := []peers.Peer{{
		Name: "best-weaviate",
		ID:   "uuid",
		URI:  "does-not-matter",
	}}

	subject := network{
		peers:           unchangedPeers,
		downloadChanged: downloadChangedFake(unchangedPeers),
	}

	callbackCalled := false
	callbackSpy := func(peers peers.Peers) {
		callbackCalled = true
	}

	subject.RegisterUpdatePeerCallback(callbackSpy)
	subject.UpdatePeers(unchangedPeers)

	if callbackCalled != false {
		t.Error("expect PeerUpdateCallback not to be called, but it was called")
	}
}

func downloadChangedFake(peerList peers.Peers) func(peers.Peers) peers.Peers {
	return func(peers.Peers) peers.Peers {
		return peerList
	}
}
