/*
 * @lc app=leetcode.cn id=113 lang=golang
 *
 * [113] 路径总和 II
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
func pathSum(root *TreeNode, targetSum int) [][]int {
	var paths [][]int
	var path []int
	var dfs func(node *TreeNode, sum int, path []int)

	dfs = func(node *TreeNode, sum int, path []int) {
		if node == nil {
			return
		}

		sum += node.Val
		path = append(path, node.Val)

		if node.Left == node.Right && sum == targetSum {
			paths = append(paths, append([]int{}, path...))
		}

		dfs(node.Left, sum, path)
		dfs(node.Right, sum, path)
		path = path[:len(path)-1]
	}
	dfs(root, 0, path)
	return paths
}

// @lc code=end

