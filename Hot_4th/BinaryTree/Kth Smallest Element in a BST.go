package binarytree

func kthSmallest(root *TreeNode, k int) int {
	res := inorderTraversal_230(root)
	return res[k-1]
}

func inorderTraversal_230(root *TreeNode) []int {
	if root == nil {
		return []int{}
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
