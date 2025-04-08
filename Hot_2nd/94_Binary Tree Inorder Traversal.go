package hot2nd

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	if root.Left == nil && root.Right == nil {
		return []int{root.Val}
	}
	return inorder(root)
}

func inorder(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	left := inorder(root.Left)
	current := []int{root.Val}
	right := inorder(root.Right)

	return append(append(left, current...), right...)
}
