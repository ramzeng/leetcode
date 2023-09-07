package main

// 二叉树后序遍历：左右根
// 深度优先算法（栈）
// https://leetcode.cn/problems/binary-tree-postorder-traversal
func postorderTraversal(root *TreeNode) []int {
	var result []int

	var preNode *TreeNode
	var stack []*TreeNode

	for root != nil || len(stack) > 0 {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}

		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if root.Right == nil || root.Right == preNode {
			result = append(result, root.Val)
			preNode = root
			root = nil
		} else {
			stack = append(stack, root)
			root = root.Right
		}
	}

	return result
}
