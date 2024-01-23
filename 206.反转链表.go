/*
 * @lc app=leetcode.cn id=206 lang=golang
 *
 * [206] 反转链表
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseList(head *ListNode) *ListNode {
	var prev, current *ListNode
	current = head

	for current != nil {
		next := current.Next
		current.Next = prev
		prev = current
		current =next
	}

	return prev
}
// @lc code=end

