package leetcode75

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxDepth(root *TreeNode) int {
	return traverse(root, 0)
}

func traverse(root *TreeNode, depth int) int {
	if root == nil {
		return depth
	}
	leftDepth, rightDepth := traverse(root.Left, depth+1), traverse(root.Right, depth+1)
	return max(leftDepth, rightDepth)
}
