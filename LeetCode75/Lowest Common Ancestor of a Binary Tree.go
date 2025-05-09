package leetcode75

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	if root == p || root == q {
		return root
	}

	left_lca, right_lca := lowestCommonAncestor(root.Left, p, q), lowestCommonAncestor(root.Right, p, q)
	if left_lca != nil && right_lca != nil {
		return root
	}
	if left_lca != nil && right_lca == nil {
		return left_lca
	}
	if left_lca == nil && right_lca != nil {
		return right_lca
	}
	return nil
}
