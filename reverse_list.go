package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	var prev *ListNode

	for head != nil {
		// 保存下一个节点
		next := head.Next
		// 当前节点指向前一个节点
		head.Next = prev
		// 前一个节点指向当前节点
		prev = head
		// 当前节点指向下一个节点
		head = next
	}

	return prev
}
