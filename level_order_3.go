package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	var results [][]int

	if root == nil {
		return results
	}

	queue := []*TreeNode{root}

	for level := 1; len(queue) > 0; level++ {
		var tempQueue []*TreeNode
		var values []int

		for i := 0; i < len(queue); i++ {
			node := queue[i]

			values = append(values, node.Val)

			if node.Left != nil {
				tempQueue = append(tempQueue, node.Left)
			}

			if node.Right != nil {
				tempQueue = append(tempQueue, node.Right)
			}
		}

		queue = tempQueue

		if level%2 == 0 {
			for i, j := 0, len(values)-1; i < j; i, j = i+1, j-1 {
				values[i], values[j] = values[j], values[i]
			}
		}

		results = append(results, values)
	}

	return results
}
