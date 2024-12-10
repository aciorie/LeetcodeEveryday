package hot

import "testing"

func TestSubsets(t *testing.T) {
	inputs := []struct {
		nums     []int
		expected [][]int
	}{
		{nums: []int{1, 2, 3}, expected: [][]int{{}, {1}, {2}, {1, 2}, {3}, {1, 3}, {2, 3}, {1, 2, 3}}},
		{nums: []int{0}, expected: [][]int{{}, {0}}},
	}

	for _, input := range inputs {
		res := Subsets(input.nums)
		if !equalTwoDArray(res, input.expected) {
			t.Errorf("nums:%v,expected:%v,got:%v", input.nums, input.expected, res)
		}
	}
}
