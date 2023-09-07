package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// https://leetcode.cn/problems/cong-shang-dao-xia-da-yin-er-cha-shu-lcof
func levelOrder(root *TreeNode) []int {
	var result []int

	if root == nil {
		return result
	}

	queue := []*TreeNode{root}

	for step := 0; step < len(queue); step++ {
		root = queue[step]

		result = append(result, root.Val)

		if root.Left != nil {
			queue = append(queue, root.Left)
		}

		if root.Right != nil {
			queue = append(queue, root.Right)
		}
	}

	return result
}
