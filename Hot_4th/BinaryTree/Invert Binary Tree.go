package binarytree

func invertTree(root *TreeNode) *TreeNode {
	var invert_226 func(node *TreeNode)
	invert_226 = func(node *TreeNode) {
		if node == nil {
			return
		}
		node.Left, node.Right = node.Right, node.Left
		invert_226(node.Left)
		invert_226(node.Right)
	}
	invert_226(root)
	return root
}
