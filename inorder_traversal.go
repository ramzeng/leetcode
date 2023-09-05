package main

// 二叉树中序遍历：左根右
// 深度优先算法（栈）
// https://leetcode.cn/problems/binary-tree-inorder-traversal
func inorderTraversal(root *TreeNode) []int {
	var results []int

	if root == nil {
		return results
	}

	var stack []*TreeNode
	node := root

	for node != nil || len(stack) > 0 {
		for node != nil {
			stack = append(stack, node)
			node = node.Left
		}

		node = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		results = append(results, node.Val)
		node = node.Right
	}

	return results
}
