/*
 * @lc app=leetcode.cn id=199 lang=golang
 *
 * [199] 二叉树的右视图
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
func rightSideView(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	var nodes []int

	currentLevel := []*TreeNode{root}

	for len(currentLevel) > 0 {
		nodes = append(nodes, currentLevel[len(currentLevel)-1].Val)
		var nextLevel []*TreeNode
		for _, node := range currentLevel {
			if node.Left != nil {
				nextLevel = append(nextLevel, node.Left)
			}

			if node.Right != nil {
				nextLevel = append(nextLevel, node.Right)
			}
		}

		currentLevel = nextLevel
	}

	return nodes
}

// @lc code=end

