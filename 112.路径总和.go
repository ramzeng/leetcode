/*
 * @lc app=leetcode.cn id=112 lang=golang
 *
 * [112] 路径总和
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
func hasPathSum(root *TreeNode, targetSum int) bool {
	var dfs func(node *TreeNode, sum int) bool
	dfs = func(node *TreeNode, sum int) bool {
		if node == nil {
			return false
		}

		sum += node.Val

		if sum == targetSum && node.Left == node.Right {
			return true
		}

		return dfs(node.Left, sum) || dfs(node.Right, sum)
	}

	return dfs(root, 0)
}

// @lc code=end

