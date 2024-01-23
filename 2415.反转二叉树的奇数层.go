/*
 * @lc app=leetcode.cn id=2415 lang=golang
 *
 * [2415] 反转二叉树的奇数层
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
func reverseOddLevels(root *TreeNode) *TreeNode {
	var dfs func(p, q *TreeNode, isOddLevel bool)
	dfs = func(p, q *TreeNode, isOddLevel bool) {
		if p == nil {
			return
		}

		if isOddLevel {
			p.Val, q.Val = q.Val, p.Val
		}

		dfs(p.Left, q.Right, !isOddLevel)
		dfs(p.Right, q.Left, !isOddLevel)
	}

	dfs(root.Left, root.Right, true)
	return root
}
// @lc code=end

