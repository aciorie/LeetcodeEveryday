package Dec

import "testing"

func TestMaximumBeauty(t *testing.T) {
	inputs := []struct {
		nums     []int
		k        int
		expected int
	}{
		{nums: []int{4, 6, 1, 2}, k: 2, expected: 3},
		{nums: []int{1, 1, 1, 1}, k: 10, expected: 4},
		{nums: []int{52, 34}, k: 21, expected: 2},
	}
	for _, input := range inputs {
		// res := MaximumBeauty(input.nums, input.k)
		res := MaximumBeauty2(input.nums, input.k)
		if res != input.expected {
			t.Errorf("nums:%v,k:%v,expected:%v,got:%v", input.nums, input.k, input.expected, res)
		}
	}
}
