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

	for level := 0; len(queue) > 0; level++ {
		var tempQueue []*TreeNode
		results = append(results, []int{})

		for i := 0; i < len(queue); i++ {
			node := queue[i]

			results[level] = append(results[level], node.Val)

			if node.Left != nil {
				tempQueue = append(tempQueue, node.Left)
			}

			if node.Right != nil {
				tempQueue = append(tempQueue, node.Right)
			}
		}

		queue = tempQueue
	}

	return results
}
