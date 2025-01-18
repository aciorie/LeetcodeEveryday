package hot

/*
Given the root of a binary search tree, and an integer k, return the kth smallest value (1-indexed) of all the values of the nodes in the tree.



Example 1:
Input: root = [3,1,4,null,2], k = 1
Output: 1

Example 2:
Input: root = [5,3,6,2,4,null,null,1], k = 3
Output: 3


Constraints:

The number of nodes in the tree is n.
1 <= k <= n <= 104
0 <= Node.val <= 104


Follow up: If the BST is modified often (i.e., we can do insert and delete operations) and you need to find the kth smallest frequently, how would you optimize?
*/

func KthSmallest(root *TreeNode, k int) int {
	res := inorderTraversal(root)
	return res[k-1]
}

func inorderTraversal(root *TreeNode) []int {
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
