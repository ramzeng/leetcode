/*
 * @lc app=leetcode.cn id=94 lang=golang
 *
 * [94] 二叉树的中序遍历
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
func inorderTraversal(root *TreeNode) []int {
	// var values []int
	// var dfs func(node *TreeNode)
	// dfs = func(node *TreeNode) {
	// 	if node == nil {
	// 		return
	// 	}

	// 	dfs(node.Left)
	// 	values = append(values, node.Val)
	// 	dfs(node.Right)
	// }
	// dfs(root)
	// return values

	var values []int
	stack := []*TreeNode{}

	for root != nil || len(stack) > 0 {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}

		root = stack[len(stack)-1]
		values = append(values, root.Val)
		stack = stack[:len(stack)-1]
		root = root.Right
	}

	return values
}

// @lc code=end

