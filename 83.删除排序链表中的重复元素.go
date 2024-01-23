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
	current := dummyNode.Next

	for current != nil && current.Next != nil {
		p := current.Next

		for p != nil && current.Val == p.Val {
			p = p.Next
		}

		current.Next = p
		current = p
	}

	return dummyNode.Next
}
// @lc code=end

