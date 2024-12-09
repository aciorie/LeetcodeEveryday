package hot

import "testing"

func TestCombine(t *testing.T) {
	inputs := []struct {
		n        int
		k        int
		expected [][]int
	}{
		{n: 4, k: 2, expected: [][]int{{1, 2}, {1, 3}, {1, 4}, {2, 3}, {2, 4}, {3, 4}}},
		{n: 1, k: 1, expected: [][]int{{1}}},
	}

	for _, input := range inputs {
		res := Combine(input.n, input.k)
		if !equalTwoDArray(res, input.expected) {
			t.Errorf("n:%v,k:%v,expected:%v,got:%v", input.n, input.k, input.expected, res)
		}
	}
}

func equalTwoDArray(a, b [][]int) bool {
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
