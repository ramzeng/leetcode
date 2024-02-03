/*
 * @lc app=leetcode.cn id=257 lang=golang
 *
 * [257] 二叉树的所有路径
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
func binaryTreePaths(root *TreeNode) []string {
	var paths []string
	var path []string
	var dfs func(node *TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}

		path = append(path, strconv.Itoa(node.Val))

		if node.Left == node.Right {
			paths = append(paths, strings.Join(path, "->"))
		}

		dfs(node.Left)
		dfs(node.Right)
		path = path[:len(path)-1]
	}
	dfs(root)

	return paths
}

// @lc code=end

