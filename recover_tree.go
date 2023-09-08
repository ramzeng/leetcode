package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// https://leetcode.cn/problems/recover-binary-search-tree/submissions
func recoverTree(root *TreeNode) {
	if root == nil {
		return
	}

	var stack []*TreeNode
	var x, y, pre *TreeNode

	for root != nil || len(stack) > 0 {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}

		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if pre != nil && pre.Val > root.Val {
			y = root
			if x == nil {
				x = pre
			} else {
				break
			}
		}

		pre = root
		root = root.Right
	}

	x.Val, y.Val = y.Val, x.Val
}
