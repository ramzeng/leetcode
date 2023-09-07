package main

type Node struct {
	Val      int
	Children []*Node
}

// N 叉树层序遍历
// 广度优先算法（队列）
// https://leetcode.cn/problems/n-ary-tree-level-order-traversal
func levelOrder(root *Node) [][]int {
	var result [][]int

	if root == nil {
		return result
	}

	queue := []*Node{root}

	for level := 0; len(queue) > 0; level++ {
		result = append(result, []int{})
		steps := len(queue)

		for step := 0; step < steps; step++ {
			root = queue[0]
			queue = queue[1:]
			result[level] = append(result[level], root.Val)
			queue = append(queue, root.Children...)
		}
	}

	return result
}
