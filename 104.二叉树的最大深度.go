/*
 * @lc app=leetcode.cn id=104 lang=golang
 *
 * [104] 二叉树的最大深度
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
func maxDepth(root *TreeNode) int {
	// if root == nil {
	// 	return 0
	// }

	// return max(maxDepth(root.Left), maxDepth(root.Right)) + 1

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

	dfs(root, 0)

	return maxDepth
}

// @lc code=end

