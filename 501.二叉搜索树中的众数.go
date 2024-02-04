/*
 * @lc app=leetcode.cn id=501 lang=golang
 *
 * [501] 二叉搜索树中的众数
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
func findMode(root *TreeNode) []int {
	counter := make(map[int]int)
	maxCount := 0
	var dfs func(node *TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}

		counter[node.Val]++

		if counter[node.Val] > maxCount {
			maxCount = counter[node.Val]
		}

		dfs(node.Left)
		dfs(node.Right)
	}
	dfs(root)

	var answers []int

	for k, v := range counter {
		if v == maxCount {
			answers = append(answers, k)
		}
	}

	return answers
}

// @lc code=end

