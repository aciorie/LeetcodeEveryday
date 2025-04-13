package hot2nd

func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}

	root, rootIndex := &TreeNode{Val: preorder[0]}, 0
	for i, v := range inorder {
		if v == root.Val {
			rootIndex = inorder[i]
			break
		}
	}

	root.Left = buildTree(preorder[1:rootIndex+1], inorder[:rootIndex])
	root.Right = buildTree(preorder[rootIndex+1:], inorder[rootIndex+1:])
	return root
}
