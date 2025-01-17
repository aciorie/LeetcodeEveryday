package hot

/*
Given the root of a binary tree, invert the tree, and return its root.



Example 1:
Input: root = [4,2,7,1,3,6,9]
Output: [4,7,2,9,6,3,1]

Example 2:
Input: root = [2,1,3]
Output: [2,3,1]

Example 3:
Input: root = []
Output: []


Constraints:

The number of nodes in the tree is in the range [0, 100].
-100 <= Node.val <= 100
*/
func InvertTree(root *TreeNode) *TreeNode {
	traverse_226(root)
	return root
}

func traverse_226(root *TreeNode) {
	if root == nil {
		return
	}

	root.Left, root.Right = root.Right, root.Left
	traverse_226(root.Left)
	traverse_226(root.Right)
}
