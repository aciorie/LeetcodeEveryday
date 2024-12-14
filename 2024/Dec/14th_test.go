package Dec

import "testing"

func TestContinuousSubarrays(t *testing.T) {
	inputs := []struct {
		nums     []int
		expected int64
	}{
		{nums: []int{5, 4, 2, 4}, expected: 8},
		{nums: []int{1, 2, 3}, expected: 6},
	}

	for _, input := range inputs {
		res := ContinuousSubarrays(input.nums)
		if res != input.expected {
			t.Errorf("nums:%v,expected:%v,got:%v", input.nums, input.expected, res)
		}
	}
}
