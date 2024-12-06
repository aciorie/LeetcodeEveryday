package hot

import "testing"

func TestSetZeroes(t *testing.T) {
	inputs := []struct {
		matrix   [][]int
		expected [][]int
	}{
		{
			matrix: [][]int{
				{1, 1, 1},
				{1, 0, 1},
				{1, 1, 1}},
			expected: [][]int{
				{1, 0, 1},
				{0, 0, 0},
				{1, 0, 1}},
		},
		{
			matrix: [][]int{
				{0, 1, 2, 0},
				{3, 4, 5, 2},
				{1, 3, 1, 5}},
			expected: [][]int{
				{0, 0, 0, 0},
				{0, 4, 5, 0},
				{0, 3, 1, 0}},
		},
	}

	for _, input := range inputs {
		SetZeroes2(input.matrix)
		if !equal73(input.matrix, input.expected) {
			t.Errorf("Expected:%v,got:%v", input.expected, input.matrix)
		}
	}
}

func equal73(a, b [][]int) bool {
	if len(a) != len(b) || len(a[0]) != len(b[0]) {
		return false
	}
	for i := 0; i < len(a); i++ {
		for j := 0; j < len(a[0]); j++ {
			if a[i][j] != b[i][j] {
				return false
			}
		}
	}
	return true
}
