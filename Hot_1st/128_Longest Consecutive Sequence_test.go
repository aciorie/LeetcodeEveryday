package hot

import "testing"

func TestLongestConsecutive(t *testing.T) {
	inputs := []struct {
		nums     []int
		expected int
	}{
		{nums: []int{100, 4, 200, 1, 3, 2}, expected: 4},
		{nums: []int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1}, expected: 9},
	}

	for _, input := range inputs {
		res := LongestConsecutive(input.nums)
		if res != input.expected {
			t.Errorf("Nums:%v,expected:%v,got:%v", input.nums, input.expected, res)
		}
	}
}
