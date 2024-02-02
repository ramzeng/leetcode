/*
 * @lc app=leetcode.cn id=129 lang=golang
 *
 * [129] 求根节点到叶节点数字之和
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
func sumNumbers(root *TreeNode) int {
	var dfs func(node *TreeNode, sum int) int
	dfs = func(node *TreeNode, sum int) int {
		if node == nil {
			return 0
		}

		sum = 10*sum + node.Val

		if node.Left == node.Right {
			return sum
		}

		return dfs(node.Left, sum) + dfs(node.Right, sum)
	}

	return dfs(root, 0)
}

// @lc code=end

