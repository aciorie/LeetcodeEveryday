package review

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	if root == p || root == q {
		return root
	}

	leftLCA, rightLCA := lowestCommonAncestor(root.Left, p, q), lowestCommonAncestor(root.Right, p, q)
	if leftLCA != nil && rightLCA != nil {
		return root
	}
	if leftLCA != nil && rightLCA == nil {
		return leftLCA
	}
	if leftLCA == nil && rightLCA != nil {
		return rightLCA
	}
	return nil
}
