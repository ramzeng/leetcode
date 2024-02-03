/*
 * @lc app=leetcode.cn id=234 lang=golang
 *
 * [234] 回文链表
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func isPalindrome(head *ListNode) bool {
	mid := findMiddleNode(head)
	p := head
	q := mid.Next
	mid.Next = nil

	q = reverseNodeList(q)

	for p != nil && q != nil {
		if p.Val != q.Val {
			return false
		}

		p = p.Next
		q = q.Next
	}

	return true
}

func reverseNodeList(head *ListNode) *ListNode {
	var prev, current *ListNode
	current = head

	for current != nil {
		next := current.Next
		current.Next = prev
		prev = current
		current = next
	}

	return prev
}

func findMiddleNode(head *ListNode) *ListNode {
	slow, fast := head, head.Next
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow
}

// @lc code=end

