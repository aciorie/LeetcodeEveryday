package hot

/*
Given the root of a binary tree, flatten the tree into a "linked list":

The "linked list" should use the same TreeNode class where the right child pointer points to the next node in the list and the left child pointer is always null.
The "linked list" should be in the same order as a pre-order traversal of the binary tree.


Example 1:
Input: root = [1,2,5,3,4,null,6]
Output: [1,null,2,null,3,null,4,null,5,null,6]

Example 2:
Input: root = []
Output: []

Example 3:
Input: root = [0]
Output: [0]


Constraints:

The number of nodes in the tree is in the range [0, 2000].
-100 <= Node.val <= 100


Follow up: Can you flatten the tree in-place (with O(1) extra space)?
*/

func Flatten(root *TreeNode) {
	if root == nil {
		return
	}
	stack := []*TreeNode{root}

	// While there are nodes in the stack
	for len(stack) > 0 {
		// Pop the top node from the stack
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		// If the current node has a right child, push it onto the stack
		// This is because the data structure is a stack so, to make sure it is lifo
		if node.Right != nil {
			stack = append(stack, node.Right)
		}

		if node.Left != nil {
			stack = append(stack, node.Left)
		}

		// If there are still nodes in the stack, set the right pointer of the current node
		if len(stack) > 0 {
			node.Right = stack[len(stack)-1] // Point to the next node in pre-order
		}
		// Set the left pointer to nil to maintain the linked list structure
		node.Left = nil
	}
}
