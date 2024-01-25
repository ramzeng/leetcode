/*
 * @lc app=leetcode.cn id=21 lang=golang
 *
 * [21] 合并两个有序链表
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	// 链表题
	// 排序：从小到大排序
	// 虚拟头节点
	dummyNode := &ListNode{}
	current := dummyNode

	// 合并两个有序链表
	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			current.Next = list1
			list1 = list1.Next
		} else {
			current.Next = list2
			list2 = list2.Next
		}

		current = current.Next
	}

	// 如果 list1 还没搞定，接上即可
	if list1 != nil {
		current.Next = list1
	}

	// 同理
	if list2 != nil {
		current.Next = list2
	}

	return dummyNode.Next
}

// @lc code=end

