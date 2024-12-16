package hot

import (
	"math"
)

/*
Given the root of a binary tree, determine if it is a valid binary search tree (BST).

A valid BST is defined as follows:

The left subtree of a node contains only nodes with keys less than the node's key.
The right subtree of a node contains only nodes with keys greater than the node's key.
Both the left and right subtrees must also be binary search trees.


Example 1:
Input: root = [2,1,3]
Output: true

Example 2:
Input: root = [5,1,4,null,null,3,6]
Output: false
Explanation: The root node's value is 5 but its right child's value is 4.


Constraints:
The number of nodes in the tree is in the range [1, 104].
-231 <= Node.val <= 231 - 1
*/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// a failed test
func IsValidBST1(root *TreeNode) bool {
	if root.Left == nil && root.Right == nil {
		return true
	}
	return helper(root, math.MinInt32, math.MaxInt32)
}

func helper(root *TreeNode, left, right int32) bool {
	if root == nil {
		return false
	}
	if root.Left.Val >= root.Val || root.Right.Val <= root.Val {
		return false
	}
	return helper(root.Left, left, int32(root.Val)) && helper(root.Right, int32(root.Val), right)
}

func IsValidBST(root *TreeNode) bool {
	var prev *TreeNode
	var inorder func(root *TreeNode) bool

	inorder = func(root *TreeNode) bool {
		if root == nil {
			return true
		}

		if !inorder(root.Left) {
			return false
		}

		if prev != nil && root.Val <= prev.Val {
			return false
		}
		prev = root

		return inorder(root.Right)
	}

	return inorder(root)
}
