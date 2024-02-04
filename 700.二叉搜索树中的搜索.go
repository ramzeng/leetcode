/*
 * @lc app=leetcode.cn id=700 lang=golang
 *
 * [700] 二叉搜索树中的搜索
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
func searchBST(root *TreeNode, val int) *TreeNode {
	var dfs func(node *TreeNode) *TreeNode
	dfs = func(node *TreeNode) *TreeNode {
		if node == nil {
			return nil
		}

		if val == node.Val {
			return node
		}

		if val < node.Val {
			return dfs(node.Left)
		}

		if val > node.Val {
			return dfs(node.Right)
		}

		return nil
	}

	return dfs(root)
}

// @lc code=end

