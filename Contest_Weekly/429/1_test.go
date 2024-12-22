package contest

import "testing"

func TestMinimumOperations(t *testing.T) {
	inputs := []struct {
		nums     []int
		expected int
	}{
		{nums: []int{1, 2, 3, 4, 2, 3, 3, 5, 7}, expected: 2},
		{nums: []int{4, 5, 6, 4, 4}, expected: 2},
		{nums: []int{6, 7, 8, 9}, expected: 0},
	}
	for _, input := range inputs {
		res := MinimumOperations(input.nums)
		if res != input.expected {
			t.Errorf("nums:%v,expected:%v,got:%v", input.nums, input.expected, res)
		}
	}
}
