package main

type Node struct {
	key, value int
	prev, next *Node
}

// LRUCache https://leetcode.cn/problems/lru-cache/description
type LRUCache struct {
	head, tail     *Node
	size, capacity int
	nodes          map[int]*Node
	//lock           sync.RWMutex
}

func Constructor(capacity int) LRUCache {
	head := &Node{}
	tail := &Node{}
	head.next = tail
	tail.prev = head

	return LRUCache{
		head:     head,
		tail:     tail,
		size:     0,
		capacity: capacity,
		nodes:    make(map[int]*Node),
	}
}

func (this *LRUCache) Get(key int) int {
	//this.lock.RLock()
	//defer this.lock.RUnlock()

	if node, ok := this.nodes[key]; ok {
		this.moveToHead(node)
		return node.value
	}

	return -1
}

func (this *LRUCache) Put(key int, value int) {
	//this.lock.Lock()
	//defer this.lock.Unlock()

	if node, ok := this.nodes[key]; ok {
		node.value = value
		this.moveToHead(node)
	} else {
		node = &Node{key: key, value: value}
		this.nodes[key] = node
		this.addToHead(node)
		this.size++
	}

	for this.size > this.capacity {
		node := this.removeTail()
		delete(this.nodes, node.key)
		this.size--
	}
}

func (this *LRUCache) addToHead(node *Node) {
	node.prev = this.head
	node.next = this.head.next
	this.head.next.prev = node
	this.head.next = node
}

func (this *LRUCache) moveToHead(node *Node) {
	this.removeNode(node)
	this.addToHead(node)
}

func (this *LRUCache) removeNode(node *Node) {
	node.prev.next = node.next
	node.next.prev = node.prev
}

func (this *LRUCache) removeTail() *Node {
	node := this.tail.prev
	this.removeNode(node)
	return node
}
