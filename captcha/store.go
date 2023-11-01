package captcha

import (
	"errors"
	"fmt"
	"sync"
	"time"
)

type Store interface {
	Get(key string) (string, error)
	Set(key string, value string, expire time.Duration) error
	Delete(key string) error
}

type Item struct {
	value     string
	expiredAt time.Time
}

type MemoryStore struct {
	items map[string]*Item
	lock  sync.RWMutex
}

func (m *MemoryStore) Get(key string) (string, error) {
	m.lock.RLock()
	defer m.lock.RUnlock()

	if item, ok := m.items[key]; ok {
		if item.expiredAt.Before(time.Now()) {
			return "", errors.New("key is expired")
		}

		return item.value, nil
	} else {
		return "", errors.New("key is not found")
	}
}

func (m *MemoryStore) Set(key string, value string, expire time.Duration) error {
	m.lock.Lock()
	defer m.lock.Unlock()

	m.items[key] = &Item{
		value:     value,
		expiredAt: time.Now().Add(expire),
	}

	return nil
}

func (m *MemoryStore) Delete(key string) error {
	m.lock.Lock()
	defer m.lock.Unlock()

	delete(m.items, key)

	return nil
}

func (m *MemoryStore) clear() {
	for {
		select {
		case <-time.Tick(time.Second):
			m.lock.Lock()
			for key, item := range m.items {
				if item.expiredAt.Before(time.Now()) {
					delete(m.items, key)
					fmt.Printf("key %s is expired\n", key)
				}
			}
			m.lock.Unlock()
		}
	}
}

func NewMemoryStore() *MemoryStore {
	ms := &MemoryStore{
		items: make(map[string]*Item),
	}

	go ms.clear()

	return ms
}
