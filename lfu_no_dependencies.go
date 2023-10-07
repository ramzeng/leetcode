package main

// https://leetcode.cn/problems/lfu-cache
func NewLinkedList() *LinkedList {
	head := NewNode()
	tail := NewNode()

	head.next = tail
	tail.prev = head

	return &LinkedList{
		head: head,
		tail: tail,
		size: 0,
	}
}

type LinkedList struct {
	head, tail *Node
	size       int
}

func NewNode() *Node {
	return &Node{
		key:   0,
		value: 0,
		level: 0,
		prev:  nil,
		next:  nil,
	}
}

type Node struct {
	key        int
	value      int
	level      int
	prev, next *Node
}

func (l *LinkedList) AddToHead(node *Node) {
	node.prev = l.head
	node.next = l.head.next
	l.head.next.prev = node
	l.head.next = node
	l.size++
}

func (l *LinkedList) MoveToHead(node *Node) {
	l.Remove(node)
	l.AddToHead(node)
}

func (l *LinkedList) Remove(node *Node) {
	node.prev.next = node.next
	node.next.prev = node.prev
	l.size--
}

func (l *LinkedList) RemoveTail() *Node {
	node := l.tail.prev
	l.Remove(node)
	return node
}

type LFUCache struct {
	values         map[int]*Node
	linkedLists    map[int]*LinkedList
	size, capacity int
	minLevel       int
}

func Constructor(capacity int) LFUCache {
	return LFUCache{
		values:      make(map[int]*Node),
		linkedLists: make(map[int]*LinkedList),
		size:        0,
		capacity:    capacity,
	}
}

func (this *LFUCache) Get(key int) int {
	if node, ok := this.values[key]; ok {
		this.linkedLists[node.level].Remove(node)

		if this.linkedLists[node.level+1] == nil {
			this.linkedLists[node.level+1] = NewLinkedList()
		}

		this.linkedLists[node.level+1].AddToHead(node)

		if this.linkedLists[node.level].size == 0 {
			if this.minLevel == node.level {
				this.minLevel++
			}
		}

		node.level++

		return node.value
	} else {
		return -1
	}
}

func (this *LFUCache) Put(key int, value int) {
	if node, ok := this.values[key]; ok {
		node.value = value
		this.Get(key)
	} else {
		node := &Node{
			key:   key,
			value: value,
			level: 0,
		}

		if this.linkedLists[node.level] == nil {
			this.linkedLists[node.level] = NewLinkedList()
		}

		this.size++

		for this.size > this.capacity {
			tail := this.linkedLists[this.minLevel].RemoveTail()
			delete(this.values, tail.key)
			this.size--
		}

		this.linkedLists[node.level].AddToHead(node)

		this.minLevel = node.level
		this.values[key] = node
	}
}
