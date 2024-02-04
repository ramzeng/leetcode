/*
 * @lc app=leetcode.cn id=589 lang=golang
 *
 * [589] N 叉树的前序遍历
 */

// @lc code=start
/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */

func preorder(root *Node) []int {
	var nodes []int
	var dfs func(node *Node)
	dfs = func(node *Node) {
		if node == nil {
			return
		}

		nodes = append(nodes, node.Val)

		for _, child := range node.Children {
			dfs(child)
		}
	}

	dfs(root)

	return nodes
}

// @lc code=end

