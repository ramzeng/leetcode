/*
 * @lc app=leetcode.cn id=148 lang=golang
 *
 * [148] 排序链表
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	p := head
	mid := findMiddleNode(head)
	q := mid.Next
	mid.Next = nil

	p = sortList(p)
	q = sortList(q)
	return mergeNodeLists(p, q)
}

func findMiddleNode(head *ListNode) *ListNode {
	slow, fast := head, head.Next

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	return slow
}

func mergeNodeLists(p, q *ListNode) *ListNode {
	dummyNode := &ListNode{}
	current := dummyNode

	for p != nil && q != nil {
		if p.Val < q.Val {
			current.Next = p
			p = p.Next
		} else {
			current.Next = q
			q = q.Next
		}
		current = current.Next
	}

	if p != nil {
		current.Next = p
	}

	if q != nil {
		current.Next = q
	}

	return dummyNode.Next
}

// @lc code=end

