package binarytree

func flatten(root *TreeNode) {
	if root == nil {
		return
	}

	stack := []*TreeNode{root}
	for len(stack) > 0 {
		node := stack[0]
		stack = stack[1:]
		if node.Right != nil {
			stack = append(stack, node.Right)
		}
		if node.Left != nil {
			stack = append(stack, node.Left)
		}

		if len(stack) > 0 {
			node.Right = stack[len(stack)-1]
		}
		node.Left = nil
	}
}
