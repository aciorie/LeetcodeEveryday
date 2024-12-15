package hot

import (
	"fmt"
	"testing"
)

// Helper function to create a binary tree from a slice
func createTree(nodes []interface{}, index int) *TreeNode {
	if index >= len(nodes) || nodes[index] == nil {
		return nil
	}
	node := &TreeNode{Val: nodes[index].(int)}
	node.Left = createTree(nodes, 2*index+1)
	node.Right = createTree(nodes, 2*index+2)
	return node
}
func TestInorderTraversal(t *testing.T) {
	testCases := []struct {
		input    []interface{}
		expected []int
	}{
		{[]interface{}{1, nil, 2, 3}, []int{1, 3, 2}},
		{[]interface{}{1, 2, 3, 4, 5, nil, 8, nil, nil, 6, 7, 9}, []int{4, 2, 6, 5, 7, 1, 3, 9, 8}},
		{[]interface{}{}, []int{}},
		{[]interface{}{1}, []int{1}},
	}

	for _, tc := range testCases {
		root := createTree(tc.input, 0)
		result := InorderTraversal(root)
		fmt.Printf("Input: %v, Output: %v, Expected: %v\n", tc.input, result, tc.expected)
	}
}
