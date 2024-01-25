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
	// 只有一个数字不用排了
	if head == nil || head.Next == nil {
		return head
	}

	mid := findMiddleNode(head)
	left := head
	right := mid.Next
	mid.Next = nil

	p := sortList(left)
	q := sortList(right)

	return mergeSort(p, q)
}

func findMiddleNode(head *ListNode) *ListNode {
	// fast 必须等于 head.Next
	// 偶数情况，第一个节点为中间节点
	slow, fast := head, head.Next

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	return slow
}

func mergeSort(p, q *ListNode) *ListNode {
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

