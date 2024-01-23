/*
 * @lc app=leetcode.cn id=82 lang=golang
 *
 * [82] 删除排序链表中的重复元素 II
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
		p := current.Next
		q := current.Next.Next

		if q != nil && p.Val == q.Val {
			for p != nil && p.Val == q.Val {
				p = p.Next
			}
			current.Next = p
		} else {
			current = p
		}
	}

	return dummyNode.Next
}
// @lc code=end

