package main

// 二叉树后序遍历：左右根
// 深度优先算法（栈）
// https://leetcode.cn/problems/binary-tree-postorder-traversal
func postorderTraversal(root *TreeNode) []int {
	var results []int

	if root == nil {
		return results
	}

	node := root
	var preNode *TreeNode
	var stack []*TreeNode

	for node != nil || len(stack) > 0 {
		for node != nil {
			stack = append(stack, node)
			node = node.Left
		}

		node = stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if node.Right == nil || node.Right == preNode {
			results = append(results, node.Val)
			preNode = node
			node = nil
		} else {
			stack = append(stack, node)
			node = node.Right
		}
	}

	return results
}
