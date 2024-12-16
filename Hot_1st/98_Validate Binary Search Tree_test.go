package hot

import "testing"

func TestIsValidBST(t *testing.T) {
	inputs := []struct {
		root     *TreeNode
		expected bool
	}{
		{
			root:     &TreeNode{Val: 2, Left: &TreeNode{Val: 1}, Right: &TreeNode{Val: 3}},
			expected: true,
		},
		{
			root:     &TreeNode{Val: 5, Left: &TreeNode{Val: 1}, Right: &TreeNode{Val: 4, Left: &TreeNode{Val: 3}, Right: &TreeNode{Val: 6}}},
			expected: false,
		},
	}

	for _, input := range inputs {
		result := IsValidBST(input.root)
		if result != input.expected {
			t.Errorf("For input %v, expected %v, but got %v", input.root, input.expected, result)
		}
	}
}
