package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isValidBST(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return dfs(root, nil, nil)
}

func dfs(root *TreeNode, min *int, max *int) bool {
	if root == nil {
		return true
	}
	if (min != nil && root.Val <= *min) || (max != nil && root.Val >= *max) {
		return false
	}
	return dfs(root.Left, min, &root.Val) && dfs(root.Right, &root.Val, max)
}
