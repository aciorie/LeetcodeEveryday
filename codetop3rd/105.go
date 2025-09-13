package codetop3rd

func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	rootVal := preorder[0]
	rootIndex := 0
	for k, v := range inorder {
		if v == rootVal {
			rootIndex = k
		}
	}

	root := &TreeNode{Val: rootVal}
	root.Left, root.Right = buildTree(preorder[1:rootIndex+1], inorder[:rootIndex]), buildTree(preorder[rootIndex+1:], inorder[rootIndex+1:])
	return root
}
