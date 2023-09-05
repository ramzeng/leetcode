package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 二叉树前序遍历：根左右
// 深度优先算法（栈）
// https://leetcode.cn/problems/binary-tree-preorder-traversal
func preorderTraversal(root *TreeNode) []int {
	var results []int

	if root == nil {
		return results
	}

	node := root

	var stack []*TreeNode

	for node != nil || len(stack) > 0 {
		for node != nil {
			stack = append(stack, node)
			results = append(results, node.Val)
			node = node.Left
		}

		node = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		node = node.Right
	}

	return results
}
