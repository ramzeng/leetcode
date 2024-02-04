/*
 * @lc app=leetcode.cn id=116 lang=golang
 *
 * [116] 填充每个节点的下一个右侧节点指针
 */

// @lc code=start
/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Left *Node
 *     Right *Node
 *     Next *Node
 * }
 */

func connect(root *Node) *Node {
	if root == nil {
		return nil
	}

	currentLevel := []*Node{root}
	for len(currentLevel) > 0 {
		var nextLevel []*Node
		var preNode *Node
		for _, node := range currentLevel {
			if preNode != nil {
				preNode.Next = node
			}

			preNode = node

			if node.Left != nil {
				nextLevel = append(nextLevel, node.Left)
			}

			if node.Right != nil {
				nextLevel = append(nextLevel, node.Right)
			}
		}

		currentLevel = nextLevel
	}

	return root
}

// @lc code=end

