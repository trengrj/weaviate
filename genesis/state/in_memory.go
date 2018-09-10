package state

import (
	"fmt"
	"github.com/go-openapi/strfmt"
	"github.com/satori/go.uuid"
	log "github.com/sirupsen/logrus"
	"sync"
	"time"
)

type inMemoryState struct {
	sync.Mutex
	peers map[strfmt.UUID]Peer
}

func NewInMemoryState() State {
	state := inMemoryState{
		peers: make(map[strfmt.UUID]Peer),
	}
	go state.garbage_collect()
	return State(&state)
}

func (im *inMemoryState) RegisterPeer(name string, uri strfmt.URI) (*Peer, error) {
	im.Lock()
	defer im.Unlock()

	uuid, err := uuid.NewV4()
	if err != nil {
		panic(err)
	}

	id := strfmt.UUID(uuid.String())

	log.Debugf("Registering peer '%v' with id '%v'", name, id)
	peer := Peer{
		PeerInfo: PeerInfo{
			Id:            id,
			LastContactAt: time.Now(),
		},
		name: name,
		uri:  uri,
	}

	im.peers[id] = peer
	go im.broadcast_update()
	return &peer, nil
}

func (im *inMemoryState) ListPeers() ([]Peer, error) {
	im.Lock()
	defer im.Unlock()

	peers := make([]Peer, 0)

	for _, v := range im.peers {
		peers = append(peers, v)
	}

	return peers, nil
}

func (im *inMemoryState) RemovePeer(id strfmt.UUID) error {
	im.Lock()
	defer im.Unlock()

	_, ok := im.peers[id]

	if ok {
		delete(im.peers, id)
	}

	go im.broadcast_update()

	return nil
}

func (im *inMemoryState) UpdateLastContact(id strfmt.UUID, contact_at time.Time) error {
	log.Debugf("Updating last contact for %v", id)

	im.Lock()
	defer im.Unlock()

	peer, ok := im.peers[id]

	if ok {
		peer.LastContactAt = contact_at
		im.peers[id] = peer
		return nil
	} else {
		return fmt.Errorf("No such peer exists")
	}
}

func (im *inMemoryState) garbage_collect() {
	for {
		time.Sleep(1 * time.Second)
		deleted_some := false

		im.Lock()
		for key, peer := range im.peers {
			peer_times_out_at := peer.PeerInfo.LastContactAt.Add(time.Second * 60)
			if time.Now().After(peer_times_out_at) {
				log.Infof("Garbage collecting peer %v", peer.Id)
				delete(im.peers, key)
				deleted_some = true
			}
		}
		im.Unlock()

		if deleted_some {
			im.broadcast_update()
		}
	}
}

func (im *inMemoryState) broadcast_update() {
	log.Info("Broadcasting peer update")
	im.Lock()
	defer im.Unlock()

	peers := make([]Peer, 0)

	for _, peer := range im.peers {
		peers = append(peers, peer)
	}

	for _, peer := range peers {
		go broadcast_update(peer, peers)
	}
}
