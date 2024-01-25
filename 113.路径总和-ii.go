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
	if root == nil {
		return nil
	}

	var paths [][]int
	var dfs func(node *TreeNode, path []int, sum int)
	dfs = func(node *TreeNode, path []int, sum int) {
		if node == nil {
			return
		}

		sum += node.Val
		path = append(path, node.Val)

		if sum == targetSum && node.Left == node.Right {
			paths = append(paths, append([]int{}, path...))
		}

		dfs(node.Left, path, sum)
		dfs(node.Right, path, sum)
	}

	dfs(root, []int{}, 0)

	return paths
}

// @lc code=end

