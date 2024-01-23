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
	p := dummyNode

	for p != nil && p.Next != nil && p.Next.Next != nil {
		p, p.Next, p.Next.Next, p.Next.Next.Next = p.Next, p.Next.Next, p.Next.Next.Next, p.Next
	}

	return dummyNode.Next
}
// @lc code=end

