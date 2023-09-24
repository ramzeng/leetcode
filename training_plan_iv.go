package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func trainingPlanIV(l1 *ListNode, l2 *ListNode) *ListNode {
	head := &ListNode{}

	current := head

	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			current.Next = l1
			l1 = l1.Next
		} else {
			current.Next = l2
			l2 = l2.Next
		}

		current = current.Next
	}

	return head.Next
}
