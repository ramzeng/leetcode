/*
 * @lc app=leetcode.cn id=24 lang=golang
 *
 * [24] 两两交换链表中的节点
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func swapPairs(head *ListNode) *ListNode {
	// 链表问题
	if head == nil {
		return nil
	}

	// 虚拟头节点
	dummyNode := &ListNode{Next: head}
	current := dummyNode

	// current->1->2->3
	// current->2 => current.Next = current.Next.Next
	// 2->1 => current.Next.Next.Next = current.Next
	// 1->3 => current.Next.Next = current.Next.Next.Next
	// current move to 1 => current = current.Next
	for current != nil && current.Next != nil && current.Next.Next != nil {
		current, current.Next, current.Next.Next, current.Next.Next.Next = current.Next, current.Next.Next, current.Next.Next.Next, current.Next
	}

	return dummyNode.Next
}

// @lc code=end

