package main

// N 叉树前序遍历
// 深度优先算法
// https://leetcode.cn/problems/n-ary-tree-preorder-traversal
func preorder(root *Node) []int {
	var results []int

	if root == nil {
		return results
	}

	stack := []*Node{root}

	for len(stack) > 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		results = append(results, node.Val)

		for i := len(node.Children) - 1; i >= 0; i-- {
			stack = append(stack, node.Children[i])
		}
	}

	return results
}
