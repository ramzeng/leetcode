/*
 * @lc app=leetcode.cn id=429 lang=golang
 *
 * [429] N 叉树的层序遍历
 */

// @lc code=start
/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */

func levelOrder(root *Node) [][]int {
	if root == nil {
		return nil
	}

	var levels [][]int
	currentLevel := []*Node{root}

	for len(currentLevel) > 0 {
		var nextLevel []*Node
		var level []int
		for _, node := range currentLevel {
			level = append(level, node.Val)
			nextLevel = append(nextLevel, node.Children...)
		}
		levels = append(levels, level)
		currentLevel = nextLevel
	}

	return levels
}

// @lc code=end

