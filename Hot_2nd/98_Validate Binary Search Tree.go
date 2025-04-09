package hot2nd

func isValidBST(root *TreeNode) bool {
	if root == nil {
		return true
	}
	if root.Left == nil && root.Right == nil {
		return true
	}
	return inorder_98(root, nil, nil)
}

// func inorder_98(root *TreeNode, min *int, max *int) bool {
// 	if root == nil {
// 		return true
// 	}
// 	if (min != nil && root.Val <= *min) && (max != nil && root.Val >= *max) {
// 		return false
// 	}
// 	return inorder_98(root.Left, min, &root.Val) && inorder_98(root.Right, &root.Val, max)
// }

func inorder_98(root *TreeNode, min *int, max *int) bool {
	if root == nil {
		return true
	}

	if min != nil && root.Val <= *min {
		return false
	}
	if max != nil && root.Val >= *max {
		return false
	}

	return inorder_98(root.Left, min, &root.Val) && inorder_98(root.Right, &root.Val, max)
}
