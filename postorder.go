package main

// N 叉树的后序遍历
// 深度优先算法
// https://leetcode.cn/problems/n-ary-tree-postorder-traversal/submissions
func postorder(root *Node) []int {
	var results []int

	if root == nil {
		return results
	}

	stack := []*Node{root}
	nodes := make(map[*Node]bool)

	for len(stack) > 0 {
		node := stack[len(stack)-1]

		if len(node.Children) == 0 || nodes[node] {
			results = append(results, node.Val)
			stack = stack[:len(stack)-1]
			continue
		}

		for i := len(node.Children) - 1; i >= 0; i-- {
			stack = append(stack, node.Children[i])
		}

		nodes[node] = true
	}

	return results
}
