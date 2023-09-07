package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// https://leetcode.cn/problems/cong-shang-dao-xia-da-yin-er-cha-shu-ii-lcof
// https://leetcode.cn/problems/binary-tree-level-order-traversal/
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
	}

	return result
}
