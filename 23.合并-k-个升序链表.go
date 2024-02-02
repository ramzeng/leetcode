/*
 * @lc app=leetcode.cn id=23 lang=golang
 *
 * [23] 合并 K 个升序链表
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
import "container/heap"

func mergeKLists(lists []*ListNode) *ListNode {
	m := make(MinHeap, 0)
	heap.Init(&m)

	for i := 0; i < len(lists); i++ {
		for current := lists[i]; current != nil; current = current.Next {
			heap.Push(&m, current.Val)
		}
	}

	dummyNode := &ListNode{}
	current := dummyNode

	for m.Len() > 0 {
		current.Next = &ListNode{Val: heap.Pop(&m).(int)}
		current = current.Next
	}

	return dummyNode.Next
}

type MinHeap []int

func (m *MinHeap) Len() int {
	return len(*m)
}

func (m *MinHeap) Less(i, j int) bool {
	return (*m)[i] < (*m)[j]
}

func (m *MinHeap) Swap(i, j int) {
	(*m)[i], (*m)[j] = (*m)[j], (*m)[i]
}

func (m *MinHeap) Push(x interface{}) {
	*m = append(*m, x.(int))
}

func (m *MinHeap) Pop() interface{} {
	tail := (*m)[m.Len()-1]
	*m = (*m)[:m.Len()-1]
	return tail
}

// @lc code=end

