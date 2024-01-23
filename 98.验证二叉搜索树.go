/*
 * @lc app=leetcode.cn id=98 lang=golang
 *
 * [98] 验证二叉搜索树
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
// func isValidBST(root *TreeNode) bool {
// 	// 根左右
// 	var dfs func(node *TreeNode, left, right int) bool
// 	dfs = func(node *TreeNode, left, right int) bool {
// 		if node == nil {
// 			return true
// 		}

// 		return node.Val > left && node.Val < right && dfs(node.Left, left, node.Val) && dfs(node.Right, node.Val, right)
// 	}

// 	return dfs(root, math.MinInt, math.MaxInt)
// }

// func isValidBST(root *TreeNode) bool {
// 	// 左根右
// 	preNode := &TreeNode{Val: math.MinInt}
// 	var dfs func(node *TreeNode) bool
// 	dfs = func(node *TreeNode) bool {
// 		if node == nil {
// 			return true
// 		}

// 		if !dfs(node.Left) {
// 			return false
// 		}

// 		if node.Val <= preNode.Val {
// 			return false
// 		}

// 		preNode = node

// 		return dfs(node.Right)
// 	}

// 	return dfs(root)
// }

func isValidBST(root *TreeNode) bool {
	// 左右根
	var dfs func(node *TreeNode) (int, int)
	dfs = func(node *TreeNode) (int, int) {
		if node == nil {
			return math.MaxInt, math.MinInt
		}

		leftMin, leftMax := dfs(node.Left)
		rightMin, rightMax := dfs(node.Right)

		if node.Val <= leftMax || node.Val >= rightMin {
			return math.MinInt, math.MaxInt
		}

		return min(leftMin, node.Val), max(rightMax, node.Val)
	}

	_, max := dfs(root)

	return max != math.MaxInt
}
// @lc code=end

