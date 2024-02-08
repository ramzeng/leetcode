package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	store := &Store{
		items: make(map[int]*Item),
	}
	store.Set(0, 1, time.Now().Add(time.Second*10))

	for {
		fmt.Println("get key 0 from store, result is: ", store.Get(0))
		time.Sleep(time.Second)
	}
}

type Item struct {
	value     int
	expiredAt time.Time
}

type Store struct {
	items map[int]*Item
	mu    sync.Mutex
}

func (s *Store) Get(key int) int {
	s.mu.Lock()
	defer s.mu.Unlock()

	if item, ok := s.items[key]; ok {
		if item.expiredAt.Before(time.Now()) {
			delete(s.items, key)
			return -1
		}
		return item.value
	} else {
		return -1
	}
}

func (s *Store) Set(key, value int, expiredAt time.Time) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if item, ok := s.items[key]; ok {
		item.value = value
		item.expiredAt = expiredAt
	} else {
		item := &Item{value: value, expiredAt: expiredAt}
		s.items[key] = item
	}
}
