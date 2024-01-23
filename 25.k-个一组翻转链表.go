/*
 * @lc app=leetcode.cn id=25 lang=golang
 *
 * [25] K 个一组翻转链表
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseKGroup(head *ListNode, k int) *ListNode {
	var size int

	for current := head; current != nil; current = current.Next {
		size++
	}

	dummyNode := &ListNode{Next: head}
	p := dummyNode

	for size >= k {
		size -= k

		var prev, current *ListNode
		current = p.Next

		for i := 0; i < k; i++ {
			next := current.Next
			current.Next = prev
			prev = current
			current = next
		}

		p, p.Next, p.Next.Next = p.Next, prev, current
	}

	return dummyNode.Next
}
// @lc code=end

