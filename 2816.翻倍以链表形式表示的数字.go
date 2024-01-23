/*
 * @lc app=leetcode.cn id=2816 lang=golang
 *
 * [2816] 翻倍以链表形式表示的数字
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func doubleIt(head *ListNode) *ListNode {
	if head.Val > 4 {
		head = &ListNode{Next: head}
	}

	for current := head; current != nil; current = current.Next {
		current.Val = current.Val*2%10

		if current.Next != nil && current.Next.Val > 4 {
			current.Val++
		}
	}

	return head
}
// @lc code=end

