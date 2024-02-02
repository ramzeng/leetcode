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
	// 层序遍历
	// 取每层的最右侧节点的值
	if root == nil {
		return nil
	}

	currentLevel := []*TreeNode{root}
	values := []int{}

	for len(currentLevel) > 0 {
		nextLevel := []*TreeNode{}
		values = append(values, currentLevel[len(currentLevel)-1].Val)

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

	return values
}

// @lc code=end

