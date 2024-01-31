/*
 * @lc app=leetcode.cn id=117 lang=golang
 *
 * [117] 填充每个节点的下一个右侧节点指针 II
 */

// @lc code=start
/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Left *Node
 *     Right *Node
 *     Next *Node
 * }
 */

func connect(root *Node) *Node {
	if root == nil {
		return nil
	}

	queue := []*Node{root}

	for len(queue) > 0 {
		size := len(queue)

		var preNode *Node
		for i := 0; i < size; i++ {
			p := queue[0]
			queue = queue[1:]

			if preNode != nil {
				preNode.Next = p
				preNode = p
			} else {
				preNode = p
			}

			if p.Left != nil {
				queue = append(queue, p.Left)
			}

			if p.Right != nil {
				queue = append(queue, p.Right)
			}
		}
	}

	return root
}
// @lc code=end

