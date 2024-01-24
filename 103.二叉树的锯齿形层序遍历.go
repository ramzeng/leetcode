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

	var answers [][]int

	currentLevel := []*TreeNode{root}

	for even := false; len(currentLevel) > 0; even = !even {

		var nextLevel []*TreeNode
		size := len(currentLevel)
		answer := make([]int, size)

		for i := 0; i < size; i++ {
			node := currentLevel[0]
			currentLevel = currentLevel[1:]

			if !even {
				answer[i] = node.Val
			} else {
				answer[size-1-i] = node.Val
			}

			if node.Left != nil {
				nextLevel = append(nextLevel, node.Left)
			}

			if node.Right != nil {
				nextLevel = append(nextLevel, node.Right)
			}
		}

		answers = append(answers, answer)
		currentLevel = nextLevel
	}

	return answers
}

// @lc code=end

