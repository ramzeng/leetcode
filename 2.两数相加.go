/*
 * @lc app=leetcode.cn id=2 lang=golang
 *
 * [2] 两数相加
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	// 链表
	var sum, carry int
	// 虚拟头节点
	dummyNode := &ListNode{}
	current := dummyNode

	// 累加
	for l1 != nil || l2 != nil {
		var x, y int
		if l1 != nil {
			x = l1.Val
			l1 = l1.Next
		}

		if l2 != nil {
			y = l2.Val
			l2 = l2.Next
		}

		// 计算当前和余数
		sum = x + y + carry
		sum, carry = sum%10, sum/10
		// 迭代向前
		current.Next = &ListNode{Val: sum}
		current = current.Next
	}

	// 如果余数大于 0
	if carry > 0 {
		current.Next = &ListNode{Val: carry}
	}

	return dummyNode.Next
}

// @lc code=end

