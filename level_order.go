package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) []int {
	var results []int

	if root == nil {
		return results
	}

	var queue []*TreeNode

	queue = append(queue, root)

	for step := 0; step < len(queue); step++ {
		node := queue[step]

		results = append(results, node.Val)

		if node.Left != nil {
			queue = append(queue, node.Left)
		}

		if node.Right != nil {
			queue = append(queue, node.Right)
		}
	}

	return results
}
