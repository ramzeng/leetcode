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

	return abs(height(root.Left)-height(root.Right)) <= 1 && isBalanced(root.Left) && isBalanced(root.Right)
}

func height(node *TreeNode) int {
	var maxDepth int
	var dfs func(node *TreeNode, depth int)
	dfs = func(node *TreeNode, depth int) {
		if node == nil {
			return
		}

		depth++
		maxDepth = max(maxDepth, depth)
		dfs(node.Left, depth)
		dfs(node.Right, depth)
	}
	dfs(node, 0)
	return maxDepth
}

func abs(x int) int {
	if x < 0 {
		return -x
	} else {
		return x
	}
}

// @lc code=end

