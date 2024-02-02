/*
 * @lc app=leetcode.cn id=143 lang=golang
 *
 * [143] 重排链表
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reorderList(head *ListNode) {
	p := head
	mid := findMiddleNode(head)
	q := mid.Next
	mid.Next = nil

	q = reverseNodeList(q)
	mergeNodeLists(p, q)
}

func findMiddleNode(head *ListNode) *ListNode {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow
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

func mergeNodeLists(p, q *ListNode) {
	for p != nil && q != nil {
		p, p.Next, q, q.Next = p.Next, q, q.Next, p.Next
	}
}

// @lc code=end

