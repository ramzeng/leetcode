package main

// N 叉树前序遍历
// 深度优先算法
// https://leetcode.cn/problems/n-ary-tree-preorder-traversal
func preorder(root *Node) []int {
	var result []int

	if root == nil {
		return result
	}

	stack := []*Node{root}

	for len(stack) > 0 {
		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		result = append(result, root.Val)

		for step := len(root.Children) - 1; step >= 0; step-- {
			stack = append(stack, root.Children[step])
		}
	}

	return result
}
