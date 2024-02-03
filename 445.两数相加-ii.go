/*
 * @lc app=leetcode.cn id=445 lang=golang
 *
 * [445] 两数相加 II
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
	l1 = reverseNodeList(l1)
	l2 = reverseNodeList(l2)

	dummyNode := &ListNode{}
	current := dummyNode
	sum, carry := 0, 0

	for l1 != nil || l2 != nil {
		x, y := 0, 0

		if l1 != nil {
			x = l1.Val
			l1 = l1.Next
		}

		if l2 != nil {
			y = l2.Val
			l2 = l2.Next
		}

		sum = x + y + carry
		sum, carry = sum%10, sum/10
		current.Next = &ListNode{Val: sum}
		current = current.Next
	}

	if carry > 0 {
		current.Next = &ListNode{Val: carry}
	}

	return reverseNodeList(dummyNode.Next)
}

func reverseNodeList(head *ListNode) *ListNode {
	var prev, current *ListNode
	current = head

	for current != nil {
		next := current.Next
		current.Next = prev
		prev = current
		current = next
	}

	return prev
}

// @lc code=end

