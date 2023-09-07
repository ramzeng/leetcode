package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// https://leetcode.cn/problems/cong-shang-dao-xia-da-yin-er-cha-shu-iii-lcof
func levelOrder(root *TreeNode) [][]int {
	var result [][]int

	if root == nil {
		return result
	}

	queue := []*TreeNode{root}

	for level := 0; len(queue) > 0; level++ {

		result = append(result, []int{})
		steps := len(queue)

		for step := 0; step < steps; step++ {
			root = queue[0]
			queue = queue[1:]
			result[level] = append(result[level], root.Val)

			if root.Left != nil {
				queue = append(queue, root.Left)
			}

			if root.Right != nil {
				queue = append(queue, root.Right)
			}
		}

		if level%2 == 1 {
			for i, j := 0, len(result[level])-1; i < j; i, j = i+1, j-1 {
				result[level][i], result[level][j] = result[level][j], result[level][i]
			}
		}
	}

	return result
}
