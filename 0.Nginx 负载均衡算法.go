package main

import (
	"fmt"
	"sync"
	"time"
)

type Peer struct {
	Weight         int
	CurrentWeight  int
	EffctiveWeight int
	Name           string
}

func NewSelector() *Selector {
	return &Selector{
		peers: make([]*Peer, 0),
	}
}

type Selector struct {
	peers []*Peer
	mu    sync.Mutex
}

func (s *Selector) AddPeer(peer *Peer) {
	s.peers = append(s.peers, peer)
}

func (s *Selector) SelectPeer() *Peer {
	s.mu.Lock()
	defer s.mu.Unlock()

	if len(s.peers) == 0 {
		return nil
	}

	selectedPeer := s.peers[0]
	totalWeight := 0

	for _, peer := range s.peers {
		totalWeight += peer.EffctiveWeight
		peer.CurrentWeight += peer.EffctiveWeight
		if peer.CurrentWeight > selectedPeer.CurrentWeight {
			selectedPeer = peer
		}
	}

	selectedPeer.CurrentWeight -= totalWeight

	return selectedPeer
}

func main() {
	selector := NewSelector()
	selector.AddPeer(&Peer{Name: "peer one", Weight: 10, EffctiveWeight: 10})
	selector.AddPeer(&Peer{Name: "peer two", Weight: 20, EffctiveWeight: 20})
	selector.AddPeer(&Peer{Name: "peer three", Weight: 30, EffctiveWeight: 30})

	for i := 0; i < 10; i++ {
		go func() {
			for {
				fmt.Printf("this is %s\n", selector.SelectPeer().Name)
				time.Sleep(time.Second)
			}
		}()
	}

	select {}
}
