package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sortedArrayToBST(nums []int) *TreeNode {
	n := len(nums)
	if n == 0 {
		return nil
	}
	return buildTreeHelper(nums, 0, len(nums)-1)
}

func buildTreeHelper(nums []int, left int, right int) *TreeNode {
	if left > right {
		return nil
	}

	mid := left + (right-left)/2
	root := &TreeNode{Val: nums[mid]}

	root.Left = buildTreeHelper(nums, left, mid-1)

	root.Right = buildTreeHelper(nums, mid+1, right)

	return root
}
