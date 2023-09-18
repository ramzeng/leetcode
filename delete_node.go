package main

func deleteNode(head *ListNode, val int) *ListNode {
	if head == nil {
		return head
	}

	if head.Val == val {
		return head.Next
	}

	prev, current := head, head.Next

	for current != nil {
		if current.Val == val {
			prev.Next = current.Next
			break
		}

		prev = current
		current = current.Next
	}

	return head
}
