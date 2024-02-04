/*
 * @lc app=leetcode.cn id=559 lang=golang
 *
 * [559] N 叉树的最大深度
 */

// @lc code=start
/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */

func maxDepth(root *Node) int {
	var answer int
	var dfs func(node *Node, depth int)
	dfs = func(node *Node, depth int) {
		if node == nil {
			return
		}

		depth++

		if len(node.Children) == 0 {
			answer = max(answer, depth)
		}

		for _, child := range node.Children {
			dfs(child, depth)
		}
	}
	dfs(root, 0)

	return answer
}

// @lc code=end

