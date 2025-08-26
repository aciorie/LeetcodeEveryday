package codetop2nd

func maxDepth(root *TreeNode) int {
	return traverse_104(root, 0)
}

func traverse_104(root *TreeNode, cur int) int {
	if root == nil {
		return cur
	}
	lef, rig := traverse_104(root.Left, cur), traverse_104(root.Right, cur)
	return max(lef, rig)
}
