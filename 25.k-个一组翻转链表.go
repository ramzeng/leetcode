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
	// 链表问题
	// 先拿到链表总长度
	length := 0
	for current := head; current != nil; current = current.Next {
		length++
	}

	// 虚拟头节点
	dummyNode := &ListNode{Next: head}
	p := dummyNode

	for length >= k {
		length -= k

		// p->1->2->3->4->5->6
		// p->1<-2<-3 4->5->6
		// prev 在 3, current 在 4
		// 1->4 p.Next.Next = current
		// p->3 p.Next= prev
		// p 移动到 1 p = p.Next
		var prev, current *ListNode

		// 从 p 的下一个开始
		current = p.Next

		// 翻转 k 个
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

