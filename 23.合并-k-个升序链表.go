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
	h := make(MaxHeap, 0)
	heap.Init(&h)

	for _, list := range lists {
		for list != nil {
			heap.Push(&h, list.Val)
			list = list.Next
		}
	}

	dummyNode := &ListNode{}
	current := dummyNode

	for h.Len() > 0 {
		current.Next = &ListNode{Val: heap.Pop(&h).(int)}
		current = current.Next
	}

	return dummyNode.Next
}

type MaxHeap []int

func (m *MaxHeap) Len() int {
	return len(*m)
}

func (m *MaxHeap) Less(i, j int) bool {
	return (*m)[i] < (*m)[j]
}

func (m *MaxHeap) Swap(i, j int) {
	(*m)[i], (*m)[j] = (*m)[j], (*m)[i]
}

func (m *MaxHeap) Push(x interface{}) {
	*m = append(*m, x.(int))
}

func (m *MaxHeap) Pop() interface{} {
	x := (*m)[m.Len()-1]
	*m = (*m)[:m.Len()-1]
	return x
}

// @lc code=end

