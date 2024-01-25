/*
 * @lc app=leetcode.cn id=19 lang=golang
 *
 * [19] 删除链表的倒数第 N 个结点
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	// 链表
	// 删除倒数第 N 个节点
	// 要拿到倒数第 N+1 个节点
	// 虚拟节点，防止头节点也要被删除
	dummyNode := &ListNode{Next: head}
	// 快慢指针
	slow, fast := dummyNode, dummyNode
	// 快指针先走 N-1 步，由于是从虚拟节点开始，所以走 N 步即可
	for i := 0; i < n; i++ {
		fast = fast.Next
	}

	// 然后快慢指针一起走，直到快指针到最后一个节点
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next
	}

	slow.Next = slow.Next.Next

	return dummyNode.Next
}

// @lc code=end

