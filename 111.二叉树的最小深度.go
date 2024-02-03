/*
 * @lc app=leetcode.cn id=111 lang=golang
 *
 * [111] 二叉树的最小深度
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
func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	answer := math.MaxInt
	var dfs func(node *TreeNode, depth int)

	dfs = func(node *TreeNode, depth int) {
		if node == nil {
			return
		}

		depth++

		if node.Left == node.Right {
			answer = min(answer, depth)
		}

		dfs(node.Left, depth)
		dfs(node.Right, depth)
	}

	dfs(root, 0)

	return answer
}

// @lc code=end

