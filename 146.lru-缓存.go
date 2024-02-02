/*
 * @lc app=leetcode.cn id=146 lang=golang
 *
 * [146] LRU 缓存
 */

// @lc code=start
type LRUCache struct {
	size, capacity int
	head, tail     *LinkedNode
	values         map[int]*LinkedNode
}

type LinkedNode struct {
	prev, next *LinkedNode
	key, value int
}

func Constructor(capacity int) LRUCache {
	head, tail := &LinkedNode{}, &LinkedNode{}
	head.next, tail.prev = tail, head
	return LRUCache{
		head:     head,
		tail:     tail,
		capacity: capacity,
		values:   make(map[int]*LinkedNode),
	}
}

func (this *LRUCache) addToHead(node *LinkedNode) {
	node.next = this.head.next
	node.prev = this.head
	this.head.next.prev = node
	this.head.next = node
}

func (this *LRUCache) moveToHead(node *LinkedNode) {
	this.removeNode(node)
	this.addToHead(node)
}

func (this *LRUCache) removeNode(node *LinkedNode) {
	node.prev.next = node.next
	node.next.prev = node.prev
}

func (this *LRUCache) removeTail() *LinkedNode {
	node := this.tail.prev
	this.removeNode(node)
	return node
}

func (this *LRUCache) Get(key int) int {
	if node, ok := this.values[key]; ok {
		this.moveToHead(node)
		return node.value
	} else {
		return -1
	}
}

func (this *LRUCache) Put(key int, value int) {
	if node, ok := this.values[key]; ok {
		node.value = value
		this.moveToHead(node)
	} else {
		node := &LinkedNode{key: key, value: value}
		this.values[key] = node
		this.addToHead(node)
		this.size++
	}

	for this.size > this.capacity {
		tail := this.removeTail()
		delete(this.values, tail.key)
		this.size--
	}
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
// @lc code=end

