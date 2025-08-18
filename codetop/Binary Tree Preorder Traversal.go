package codetop

func preorderTraversal(root *TreeNode) []int {
	var res []int

	var dfs_144 func(*TreeNode)
	dfs_144 = func(root *TreeNode) {
		if root == nil {
			return
		}
		res = append(res, root.Val)
		dfs_144(root.Left)
		dfs_144(root.Right)
	}
	dfs_144(root)
	return res
}
