package binarysearch

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxPathSum(root *TreeNode) int {
	res := math.MinInt32
	var dfs_124 func(node *TreeNode) int
	dfs_124 = func(node *TreeNode) int {
		if node == nil {
			return 0
		}

		leftSum := max(0, dfs_124(root.Left))
		rightSum := max(0, dfs_124(root.Right))
		curSum := leftSum + rightSum + node.Val

		res = max(res, curSum)
		return node.Val + max(leftSum, rightSum)
	}
	dfs_124(root)
	return res
}
