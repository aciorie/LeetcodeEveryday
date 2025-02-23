package feb

/*
889. Construct Binary Tree from Preorder and Postorder Traversal
Medium
Topics
Companies
Given two integer arrays, preorder and postorder where preorder is the preorder traversal of a binary tree of distinct values and postorder is the postorder traversal of the same tree, reconstruct and return the binary tree.

If there exist multiple answers, you can return any of them.



Example 1:


Input: preorder = [1,2,4,5,3,6,7], postorder = [4,5,2,6,7,3,1]
Output: [1,2,3,4,5,6,7]
Example 2:

Input: preorder = [1], postorder = [1]
Output: [1]


Constraints:

1 <= preorder.length <= 30
1 <= preorder[i] <= preorder.length
All the values of preorder are unique.
postorder.length == preorder.length
1 <= postorder[i] <= postorder.length
All the values of postorder are unique.
It is guaranteed that preorder and postorder are the preorder traversal and postorder traversal of the same binary tree.
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func ConstructFromPrePost(preorder []int, postorder []int) *TreeNode {
	if len(preorder) == 0 || len(postorder) == 0 {
		return nil
	}
	root := &TreeNode{Val: postorder[0]}
	if len(postorder) == 1 {
		return root
	}

	leftRootVal := preorder[1]
	leftSubtreeSize := 0

	// Find the size of the left subtree
	for i := 0; i < len(postorder); i++ {
		if postorder[i] == leftRootVal {
			leftSubtreeSize = i + 1
			break
		}
	}

	root.Left = ConstructFromPrePost(preorder[1:leftSubtreeSize+1], postorder[:leftSubtreeSize])
	root.Right = ConstructFromPrePost(preorder[leftSubtreeSize+1:], postorder[leftSubtreeSize:len(postorder)-1])

	return root
}
