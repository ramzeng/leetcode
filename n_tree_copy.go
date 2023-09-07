package main

type Node struct {
	Val      int
	Children []*Node
}

func cloneTree(root *Node) *Node {
	if root == nil {
		return nil
	}

	clonedRoot := &Node{Val: root.Val}

	originalQueue := []*Node{root}
	clonedQueue := []*Node{clonedRoot}

	for step := 0; step < len(originalQueue); step++ {
		originalNode := originalQueue[step]
		clonedNode := clonedQueue[step]

		for _, child := range originalNode.Children {
			clonedChild := &Node{Val: child.Val}
			clonedNode.Children = append(clonedNode.Children, clonedChild)
			originalQueue = append(originalQueue, child)
			clonedQueue = append(clonedQueue, clonedChild)
		}
	}

	return clonedRoot
}
