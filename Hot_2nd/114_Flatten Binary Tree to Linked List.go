package hot2nd

func flatten(root *TreeNode) {
	if root == nil {
		return
	}
	stack := []*TreeNode{root}
	for len(stack) > 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		// Because the stack is LIFO (Last In, First Out), when pushing the left and right child nodes onto
		// the stack, you need to push the Right child first and then the Left child. This way,
		// when popping from the stack, the Left child will be processed first, followed by the Right child.
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
