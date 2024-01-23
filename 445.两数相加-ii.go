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
	var head, tail *ListNode

	var sum, carry int

	l1 = reverseList(l1)
	l2 = reverseList(l2)

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

		sum = x+y+carry
		sum, carry = sum%10, sum/10

		if head == nil {
			tail = &ListNode{Val: sum}
			head = &ListNode{Next: tail}
		} else {
			tail.Next = &ListNode{Val: sum}
			tail = tail.Next
		}
	}

	if carry > 0 {
		tail.Next = &ListNode{Val: carry}
	}

	return reverseList(head.Next)
}

func reverseList(head *ListNode) *ListNode {
	var prev, current *ListNode
	current = head

	for current != nil {
		next := current.Next
		current.Next = prev
		prev = current
		current =next
	}

	return prev
}
// @lc code=end

