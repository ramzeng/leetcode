package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 根左右，左根右
// 输入: preorder = [3,9,20,15,7], inorder = [9,3,15,20,7]
// 输出: [3,9,20,null,null,15,7]
// https://leetcode.cn/problems/zhong-jian-er-cha-shu-lcof
func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}

	// 前序遍历的第一个节点就是根节点
	root := &TreeNode{Val: preorder[0]}

	// 在中序遍历中找到根节点的位置
	var i int
	for i = 0; i < len(inorder); i++ {
		if inorder[i] == preorder[0] {
			break
		}
	}

	// 递归构建左右子树
	root.Left = buildTree(preorder[1:len(inorder[:i+1])], inorder[:i])
	root.Right = buildTree(preorder[len(inorder[:i+1]):], inorder[i+1:])

	return root
}
