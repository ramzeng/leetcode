/*
 * @lc app=leetcode.cn id=1110 lang=golang
 *
 * [1110] 删点成林
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
func delNodes(root *TreeNode, to_delete []int) []*TreeNode {
	var nodes []*TreeNode
    var dfs func(node *TreeNode) *TreeNode
	dfs = func(node *TreeNode) *TreeNode {
		if node == nil {
			return nil
		}

		node.Left = dfs(node.Left)
		node.Right = dfs(node.Right)

		for _, v := range to_delete {
			if node.Val == v {
				if node.Left != nil {
					nodes = append(nodes, node.Left)
				}

				if node.Right != nil {
					nodes = append(nodes, node.Right)
				}

				return nil
			}
		}

		return node
	}

	if dfs(root) != nil {
		nodes = append(nodes, root)
	}

	return nodes
}
// @lc code=end

