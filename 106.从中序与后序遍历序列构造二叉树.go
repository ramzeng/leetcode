/*
 * @lc app=leetcode.cn id=106 lang=golang
 *
 * [106] 从中序与后序遍历序列构造二叉树
 */

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func buildTree(inorder []int, postorder []int) *TreeNode {
	if len(postorder) == 0 {
		return nil
	}

	root := &TreeNode{Val: postorder[len(postorder)-1]}
	rootIndex := 0

	for k, v := range inorder {
		if v == root.Val {
			rootIndex = k
			break
		}
	}

	root.Left = buildTree(inorder[:rootIndex], postorder[:len(inorder[:rootIndex])])
	root.Right = buildTree(inorder[rootIndex+1:], postorder[len(inorder[:rootIndex]):len(postorder)-1])

	return root
}

// @lc code=end

