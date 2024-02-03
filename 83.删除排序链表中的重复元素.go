/*
 * @lc app=leetcode.cn id=83 lang=golang
 *
 * [83] 删除排序链表中的重复元素
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteDuplicates(head *ListNode) *ListNode {
	dummyNode := &ListNode{Next: head}
	current := dummyNode

	for current != nil && current.Next != nil {
		for current.Next.Next != nil && current.Next.Val == current.Next.Next.Val {
			current.Next = current.Next.Next
		}

		current = current.Next
	}

	return dummyNode.Next
}

// @lc code=end

