/*
 * @lc app=leetcode.cn id=1123 lang=golang
 *
 * [1123] 最深叶节点的最近公共祖先
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
func lcaDeepestLeaves(root *TreeNode) *TreeNode {
	var maxDepth int
	var answer *TreeNode

	var dfs func(node *TreeNode, depth int) int
	dfs = func(node *TreeNode, depth int) int {
		if node == nil {
			maxDepth = max(maxDepth, depth)
			return depth
		}

		leftMaxDepth := dfs(node.Left, depth+1)
		rightMaxDepth := dfs(node.Right, depth+1)

		if leftMaxDepth == maxDepth && rightMaxDepth == maxDepth {
			answer = node
		}

		return max(leftMaxDepth, rightMaxDepth)
	}

	dfs(root, 0)

	return answer
}
// @lc code=end

