package hot

import "testing"

func TestSingleNumber(t *testing.T) {
	inputs := []struct {
		nums     []int
		expected int
	}{
		{nums: []int{2, 2, 1}, expected: 1},
		{nums: []int{4, 1, 2, 1, 2}, expected: 4},
		{nums: []int{1}, expected: 1},
	}

	for _, input := range inputs {
		res := SingleNumber(input.nums)
		if res != input.expected {
			t.Errorf("nums:%v,expected:%v,got:%v", input.nums, input.expected, res)
		}
	}
}
