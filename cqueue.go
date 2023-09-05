package main

// CQueue
// 使用两个栈实现队列
// https://leetcode.cn/problems/yong-liang-ge-zhan-shi-xian-dui-lie-lcof
type CQueue struct {
	inStack  []int
	outStack []int
}

func Constructor() CQueue {
	return CQueue{}
}

func (q *CQueue) AppendTail(value int) {
	q.inStack = append(q.inStack, value)
}

func (q *CQueue) DeleteHead() int {
	if len(q.outStack) <= 0 {
		if len(q.inStack) <= 0 {
			return -1
		}

		q.outStack = make([]int, 0, len(q.inStack))

		for len(q.inStack) > 0 {
			q.outStack = append(q.outStack, q.inStack[len(q.inStack)-1])
			q.inStack = q.inStack[:len(q.inStack)-1]
		}
	}

	result := q.outStack[len(q.outStack)-1]
	q.outStack = q.outStack[:len(q.outStack)-1]
	return result
}

/**
 * Your CQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AppendTail(value);
 * param_2 := obj.DeleteHead();
 */
