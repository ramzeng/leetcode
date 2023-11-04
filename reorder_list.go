package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

// https://leetcode.cn/problems/reorder-list
func reorderList(head *ListNode) {
	if head == nil {
		return
	}

	var nodes []*ListNode
	
	// 借助外部存储，实现下标访问
	for node := head; node != nil; node = node.Next {
		nodes = append(nodes, node)
	}

	// 左右指针
	left, right := 0, len(nodes)-1

	for left < right {
		// 交替前进
		nodes[left].Next = nodes[right]
		left++

		if left == right {
			break
		}

		nodes[right].Next = nodes[left]
		right--
	}

	// 不能忘记结束
	nodes[right].Next = nil
}
