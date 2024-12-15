package hot

/*
Given the root of a binary tree, return the inorder traversal of its nodes' values.



Example 1:
Input: root = [1,null,2,3]
Output: [1,3,2]
Explanation:



Example 2:
Input: root = [1,2,3,4,5,null,8,null,null,6,7,9]
Output: [4,2,6,5,7,1,3,9,8]
Explanation:



Example 3:
Input: root = []
Output: []

Example 4:
Input: root = [1]
Output: [1]



Constraints:

The number of nodes in the tree is in the range [0, 100].
-100 <= Node.val <= 100


Follow up: Recursive solution is trivial, could you do it iteratively?
*/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func InorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	if root.Right == nil && root.Left == nil {
		return []int{root.Val}
	}

	var res []int
	var stack []*TreeNode
	cur := root

	for cur != nil || len(stack) > 0 {
		//Traverse all nodes in the left subtree
		for cur != nil {
			stack = append(stack, cur)
			cur = cur.Left
		}

		//Pop the top node of the stack and access
		cur = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		res = append(res, cur.Val)

		//Visit the right subtree
		cur = cur.Right
	}
	return res
}
