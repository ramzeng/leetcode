/*
 * @lc app=leetcode.cn id=92 lang=golang
 *
 * [92] 反转链表 II
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseBetween(head *ListNode, left int, right int) *ListNode {
	dummyNode := &ListNode{Next: head}
	p := dummyNode

	for i := 0; i < left-1; i++ {
		p = p.Next
	}

	var prev, current *ListNode
	current = p.Next

	for i := left; i <= right; i++ {
		next := current.Next
		current.Next = prev
		prev = current
		current = next
	}

	p.Next, p.Next.Next = prev, current

	return dummyNode.Next
}
// @lc code=end

