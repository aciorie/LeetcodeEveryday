package apr

/*
1123. Lowest Common Ancestor of Deepest Leaves
Medium
Topics
Companies
Hint
Given the root of a binary tree, return the lowest common ancestor of its deepest leaves.

Recall that:

The node of a binary tree is a leaf if and only if it has no children
The depth of the root of the tree is 0. if the depth of a node is d, the depth of each of its children is d + 1.
The lowest common ancestor of a set S of nodes, is the node A with the largest depth such that every node in S is in the subtree with root A.


Example 1:


Input: root = [3,5,1,6,2,0,8,null,null,7,4]
Output: [2,7,4]
Explanation: We return the node with value 2, colored in yellow in the diagram.
The nodes coloured in blue are the deepest leaf-nodes of the tree.
Note that nodes 6, 0, and 8 are also leaf nodes, but the depth of them is 2, but the depth of nodes 7 and 4 is 3.
Example 2:

Input: root = [1]
Output: [1]
Explanation: The root is the deepest node in the tree, and it's the lca of itself.
Example 3:

Input: root = [0,1,3,null,2]
Output: [2]
Explanation: The deepest leaf node in the tree is 2, the lca of one node is itself.


Constraints:

The number of nodes in the tree will be in the range [1, 1000].
0 <= Node.val <= 1000
The values of the nodes in the tree are unique.
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func lcaDeepestLeaves(root *TreeNode) *TreeNode {
	if root == nil || (root.Left == nil && root.Right == nil) {
		return root
	}

	maxDepth, deepestLeaves := 0, make([]*TreeNode, 0)
	var dfs func(node *TreeNode, depth int)
	var lca func(node *TreeNode) *TreeNode

	dfs = func(node *TreeNode, depth int) {
		if node == nil {
			// Update maxDepth on null node
			maxDepth = max(maxDepth, depth-1)
			return
		}

		if node.Left == nil && node.Right == nil {
			if depth > maxDepth {
				maxDepth = depth
				// Reset the deepest leaf node slice and add the current leaf node to
				deepestLeaves = []*TreeNode{node}
			} else if depth == maxDepth {
				// If the depth of the current leaf node is equal to the maximum depth, add the current leaf node to the slice
				deepestLeaves = append(deepestLeaves, node)
			}
		}
		dfs(node.Left, depth+1)
		dfs(node.Right, depth+1)
	}

	lca = func(node *TreeNode) *TreeNode {
		if node == nil {
			return nil
		}
		// Recursively find the nearest common ancestor of the left and right subtrees
		l, r := lca(node.Left), lca(node.Right)

		// Check if the current node is in deepestLeaves (BEFORE checking children)
		for _, leaf := range deepestLeaves {
			if node == leaf {
				return node
			}
		}

		// If the LCA of both left and right subtrees is not empty, the current node is the LCA
		if l != nil && r != nil {
			return node
		} else if l != nil {
			return l
		} else if r != nil {
			return r
		} else {
			return nil
		}
	}

	// Find the maximum depth and deepest leaf node
	dfs(root, 0)
	return lca(root)
}
