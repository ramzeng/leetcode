/*
 * @lc app=leetcode.cn id=1026 lang=golang
 *
 * [1026] 节点与其祖先之间的最大差值
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
func maxAncestorDiff(root *TreeNode) int {
	var maxDiff int
    var f func(node *TreeNode, minValue, maxValue int)
	f = func(node *TreeNode, minValue, maxValue int) {
		if node == nil {
			return
		}

		minValue = min(minValue, node.Val)
		maxValue = max(maxValue, node.Val)
		diff := max(abs(node.Val-minValue), abs(node.Val-maxValue))
		maxDiff = max(maxDiff, diff)

		f(node.Left, minValue, maxValue)
		f(node.Right, minValue, maxValue)
	}

	f(root, root.Val, root.Val)

	return maxDiff
}

func abs(x int) int {
	if x < 0 {
		return -x
	} else {
		return x
	}
}
// @lc code=end

