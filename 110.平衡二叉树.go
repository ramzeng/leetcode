/*
 * @lc app=leetcode.cn id=110 lang=golang
 *
 * [110] 平衡二叉树
 */

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}

	diff := abs(height(root.Left) - height(root.Right))

	return diff <= 1 && isBalanced(root.Left) && isBalanced(root.Right)
}

func abs(x int) int {
	if x < 0 {
		return -x
	} else {
		return x
	}
}

func height(node *TreeNode) int {
	if node == nil {
		return 0
	}

	return max(height(node.Left), height(node.Right)) + 1
}

// @lc code=end

