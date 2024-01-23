/*
 * @lc app=leetcode.cn id=116 lang=golang
 *
 * [116] 填充每个节点的下一个右侧节点指针
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

		preNode := queue[0]
		for i := 0; i < size; i++ {
			p := queue[0]
			queue = queue[1:]

			if i > 0 {
				preNode.Next = p
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

