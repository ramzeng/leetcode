package main

import (
	"container/list"
	"sync"
)

type Entity struct {
	Key   string
	Value Value
}

func (e *Entity) Size() uint64 {
	return uint64(len(e.Key)) + e.Value.Size()
}

type Value interface {
	Size() uint64
}

func NewLRU(maxBytesSize uint64) *LRU {
	return &LRU{
		maxBytesSize:     maxBytesSize,
		doubleLinkedList: list.New(),
		values:           make(map[string]*list.Element),
	}
}

type LRU struct {
	maxBytesSize     uint64
	usedBytesSize    uint64
	doubleLinkedList *list.List
	values           map[string]*list.Element
	lock             sync.RWMutex
}

func (l *LRU) Get(key string) (Value, bool) {
	l.lock.RLock()
	defer l.lock.RUnlock()

	if element, ok := l.values[key]; ok {
		l.doubleLinkedList.MoveToFront(element)

		return element.Value.(*Entity).Value, true
	}

	return nil, false
}

func (l *LRU) Set(key string, value Value) {
	l.lock.Lock()
	defer l.lock.Unlock()

	if element, ok := l.values[key]; !ok {
		entity := &Entity{
			Key:   key,
			Value: value,
		}

		l.doubleLinkedList.PushFront(entity)

		l.usedBytesSize += entity.Size()
	} else {
		l.doubleLinkedList.MoveToFront(element)

		entity := element.Value.(*Entity)
		l.usedBytesSize += value.Size() - entity.Value.Size()
		entity.Value = value
	}

	for l.maxBytesSize > 0 && l.doubleLinkedList.Len() > 0 && l.usedBytesSize > l.maxBytesSize {
		element := l.doubleLinkedList.Back()
		l.doubleLinkedList.Remove(element)

		entity := element.Value.(*Entity)
		delete(l.values, entity.Key)
		l.usedBytesSize -= entity.Size()
	}
}
