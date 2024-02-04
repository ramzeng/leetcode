/*
 * @lc app=leetcode.cn id=530 lang=golang
 *
 * [530] 二叉搜索树的最小绝对差
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
func getMinimumDifference(root *TreeNode) int {
	var dfs func(node *TreeNode)
	var preNode *TreeNode
	minDiff := math.MaxInt

	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}

		dfs(node.Left)

		if preNode != nil {
			minDiff = min(minDiff, abs(node.Val-preNode.Val))
		}

		preNode = node

		dfs(node.Right)
	}
	dfs(root)
	return minDiff
}

func abs(x int) int {
	if x < 0 {
		return -x
	} else {
		return x
	}
}

// @lc code=end

