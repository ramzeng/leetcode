/*
 * @lc app=leetcode.cn id=101 lang=golang
 *
 * [101] 对称二叉树
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
func isSymmetric(root *TreeNode) bool {
	var f func(p, q *TreeNode) bool
	f = func(p, q *TreeNode) bool {
		if p == nil || q == nil {
			return p == q
		}

		return p.Val == q.Val && f(p.Left, q.Right) && f(p.Right, q.Left)
	}

	return f(root, root)
}
// @lc code=end

