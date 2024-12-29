package contest

import "testing"

func TestMinimumOperations(t *testing.T) {
	inputs := []struct {
		grid     [][]int
		expected int
	}{
		{[][]int{{3, 2}, {1, 3}, {3, 4}, {0, 1}}, 15},
		{[][]int{{3, 2, 1}, {2, 1, 0}, {1, 2, 3}}, 12},
	}

	for _, input := range inputs {
		res := MinimumOperations(input.grid)
		if res != input.expected {
			t.Errorf("grid:%v,expected:%v,got:%v", input.grid, input.expected, res)
		}
	}
}
