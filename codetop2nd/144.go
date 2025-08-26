package codetop2nd

func preorderTraversal(root *TreeNode) []int {
	var res []int
	var dfs_144 func(node *TreeNode)
	dfs_144 = func(node *TreeNode) {
		if node == nil {
			return
		}
		res = append(res, node.Val)
		dfs_144(node.Left)
		dfs_144(node.Right)
	}
	dfs_144(root)
	return res
}
