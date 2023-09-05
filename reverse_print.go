package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	head := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 3,
			Next: &ListNode{
				Val:  2,
				Next: nil,
			},
		},
	}

	fmt.Println(reversePrint(head))
}

// 输入：head = [1,3,2]
// 输出：[2,3,1]
// https://leetcode.cn/problems/cong-wei-dao-tou-da-yin-lian-biao-lcof
func reversePrint(head *ListNode) []int {
	values := make([]int, 0)

	if head == nil {
		return values
	}

	node := head

	for node != nil {
		values = append(values, node.Val)

		node = node.Next
	}

	length := len(values)

	results := make([]int, length)

	for i := length - 1; i >= 0; i-- {
		results[i] = values[length-1-i]
	}

	return results
}
