package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func getKthFromEnd(head *ListNode, k int) *ListNode {
	if head == nil {
		return head
	}

	prev, current := head, head.Next
	distance := 1

	for current != nil {
		if distance == k {
			prev = prev.Next
		} else {
			distance++
		}

		current = current.Next
	}

	return prev
}
