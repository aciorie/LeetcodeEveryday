package leetcode75

import "math"

func goodNodes(root *TreeNode) int {
	return dfs(root, math.MinInt32)
}

func dfs(root *TreeNode, maxSoFar int) int {
	if root == nil {
		return 0
	}

	isGood := 0
	if root.Val >= maxSoFar {
		isGood = 1
	}
	newMaxSoFar := max(root.Val, maxSoFar)
	leftGoodNodes := dfs(root.Left, newMaxSoFar)
	rightGoodNodes := dfs(root.Right, newMaxSoFar)

	return isGood + leftGoodNodes + rightGoodNodes
}
