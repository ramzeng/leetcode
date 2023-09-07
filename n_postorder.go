package main

// N 叉树的后序遍历
// 深度优先算法
// https://leetcode.cn/problems/n-ary-tree-postorder-traversal/submissions
func postorder(root *Node) []int {
	var result []int

	if root == nil {
		return result
	}

	stack := []*Node{root}
	nodes := make(map[*Node]bool)

	for len(stack) > 0 {
		root = stack[len(stack)-1]

		if root.Children == nil || nodes[root] {
			result = append(result, root.Val)
			stack = stack[:len(stack)-1]
			continue
		}

		for step := len(root.Children) - 1; step >= 0; step-- {
			stack = append(stack, root.Children[step])
		}

		nodes[root] = true
	}

	return result
}
