package binarytree

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	if root == p || root == q {
		return root
	}

	leftAncestor, rightAncestor := lowestCommonAncestor(root.Left, p, q), lowestCommonAncestor(root.Right, p, q)

	if leftAncestor != nil && rightAncestor != nil {
		return root
	}
	if leftAncestor != nil && rightAncestor == nil {
		return leftAncestor
	}
	if leftAncestor == nil && rightAncestor != nil {
		return rightAncestor
	}

	return nil
}
