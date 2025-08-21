package codetop2nd

import "math"

func maxPathSum(root *TreeNode) int {
	res := math.MinInt64
	var dfs_124 func(root *TreeNode) int
	dfs_124 = func(root *TreeNode) int {
		if root == nil {
			return 0
		}
		leftsum, rightsum := max(0, dfs_124(root.Left)), max(0, dfs_124(root.Right))
		cursum := root.Val + leftsum + rightsum
		res = max(res, cursum)
		return max(leftsum, rightsum) + root.Val
	}
	dfs_124(root)
	return res
}
