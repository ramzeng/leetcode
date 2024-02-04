/*
 * @lc app=leetcode.cn id=1302 lang=golang
 *
 * [1302] 层数最深叶子节点的和
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
func deepestLeavesSum(root *TreeNode) int {
	var maxDepth int
	var sum int
	var dfs func(node *TreeNode, depth int)

	dfs = func(node *TreeNode, depth int) {
		if node == nil {
			return
		}

		depth++

		if depth > maxDepth {
			maxDepth = depth
			sum = 0
		}

		if depth == maxDepth && node.Left == node.Right {
			sum += node.Val
		}

		dfs(node.Left, depth)
		dfs(node.Right, depth)
	}
	dfs(root, 0)

	return sum
}

// @lc code=end

