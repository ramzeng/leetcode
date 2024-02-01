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

type Selector struct {
	peers []*Peer
	mu    sync.Mutex
}

func NewSelector() *Selector {
	return &Selector{
		peers: make([]*Peer, 0),
	}
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
	selector.AddPeer(&Peer{Weight: 10, EffctiveWeight: 10, Name: "Peer One"})
	selector.AddPeer(&Peer{Weight: 20, EffctiveWeight: 20, Name: "Peer Two"})
	selector.AddPeer(&Peer{Weight: 30, EffctiveWeight: 30, Name: "Peer Three"})

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
