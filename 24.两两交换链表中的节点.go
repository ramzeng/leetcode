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
	dummyNode := &ListNode{Next: head}
	current := dummyNode

	for current != nil && current.Next != nil && current.Next.Next != nil {
		// current -> current.Next
		// current.Next -> current.Next.Next
		// current.Next.Next -> current.Next.Next.Next
		current, current.Next, current.Next.Next, current.Next.Next.Next = current.Next, current.Next.Next, current.Next.Next.Next, current.Next
	}

	return dummyNode.Next
}

// @lc code=end

