package main

type Node struct {
	Val      int
	Children []*Node
}

// N 叉树层序遍历
// 广度优先算法（队列）
// https://leetcode.cn/problems/n-ary-tree-level-order-traversal
func levelOrder(root *Node) [][]int {
	var results [][]int

	if root == nil {
		return results
	}

	if root.Children == nil || len(root.Children) == 0 {
		results = append(results, []int{root.Val})
	} else {
		// 广度优先算法
		queue := []*Node{root}

		for len(queue) > 0 {
			var values []int
			var children []*Node

			// 清空队列（处理完当前层）
			for _, node := range queue {
				values = append(values, node.Val)
				// 将 children 节点（下一层）压入队列
				children = append(children, node.Children...)
			}

			queue = children
			results = append(results, values)
		}
	}

	return results
}
