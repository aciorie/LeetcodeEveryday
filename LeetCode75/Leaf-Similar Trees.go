package leetcode75

import "reflect"

func leafSimilar(root1 *TreeNode, root2 *TreeNode) bool {
	leftLeaves, rightLeaves := preOrderTraverse(root1), preOrderTraverse(root2)
	return reflect.DeepEqual(leftLeaves, rightLeaves)
}

func preOrderTraverse(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	if root.Left == nil && root.Right == nil {
		return []int{root.Val}
	}
	leftLeaves := preOrderTraverse(root.Left)
	rightLeaves := preOrderTraverse(root.Right)
	return append(leftLeaves, rightLeaves...)
}
