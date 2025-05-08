package leetcode75

import "math"

var maxLength int

func longestZigZag(root *TreeNode) int {
	maxLength = 0

	if root.Left != nil {
		dfs_1372(root.Left, 1, false)
	}
	if root.Right != nil {
		dfs_1372(root.Right, 1, true)
	}

	return maxLength
}

func dfs_1372(node *TreeNode, currentLength int, lastStepRight bool) {
	if node == nil {
		return
	}

	maxLength = int(math.Max(float64(maxLength), float64(currentLength)))

	if lastStepRight {
		if node.Left != nil {
			dfs_1372(node.Left, currentLength+1, false)
		}
		if node.Right != nil {
			dfs_1372(node.Right, 1, true)
		}
	} else {
		if node.Right != nil {
			dfs_1372(node.Right, currentLength+1, true)
		}
		if node.Left != nil {
			dfs_1372(node.Left, 1, false)
		}
	}
}
