/*
 * @lc app=leetcode.cn id=103 lang=golang
 *
 * [103] 二叉树的锯齿形层序遍历
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
func zigzagLevelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}

	levels := [][]int{}
	currentLevel := []*TreeNode{root}

	for even := false; len(currentLevel) > 0; even = !even {
		nextLevel := []*TreeNode{}
		n := len(currentLevel)
		level := make([]int, n)
		for index, node := range currentLevel {
			if !even {
				level[index] = node.Val
			} else {
				level[n-1-index] = node.Val
			}

			if node.Left != nil {
				nextLevel = append(nextLevel, node.Left)
			}

			if node.Right != nil {
				nextLevel = append(nextLevel, node.Right)
			}
		}
		currentLevel = nextLevel
		levels = append(levels, level)
	}

	return levels
}

// @lc code=end

