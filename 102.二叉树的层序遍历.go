/*
 * @lc app=leetcode.cn id=102 lang=golang
 *
 * [102] 二叉树的层序遍历
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
func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}

	currentLevel := []*TreeNode{root}
	levels := [][]int{}

	for len(currentLevel) > 0 {
		nextLevel := []*TreeNode{}
		level := []int{}

		for _, node := range currentLevel {
			level = append(level, node.Val)
			if node.Left != nil {
				nextLevel = append(nextLevel, node.Left)
			}
			if node.Right != nil {
				nextLevel = append(nextLevel, node.Right)
			}
		}

		levels = append(levels, level)
		currentLevel = nextLevel
	}

	return levels
}

// @lc code=end

