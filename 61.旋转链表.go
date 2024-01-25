/*
 * @lc app=leetcode.cn id=61 lang=golang
 *
 * [61] 旋转链表
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil {
		return nil
	}

	// 求链表长度
	length := 0

	for current := head; current != nil; current = current.Next {
		length++
	}

	// 取余数
	k = k % length

	if k == 0 {
		return head
	}

	// 找到倒数第 k 个节点
	slow, fast := head, head

	for i := 0; i < k; i++ {
		fast = fast.Next
	}

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next
	}

	// 此时 fast 在尾部，slow 在倒数第 K 个节点
	head, fast.Next, slow.Next = slow.Next, head, nil

	return head
}

// @lc code=end

