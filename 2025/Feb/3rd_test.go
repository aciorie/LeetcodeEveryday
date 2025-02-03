package feb

import "testing"

func TestLongestMonotonicSubarray2(t *testing.T) {
	inputs := []struct {
		nums     []int
		expected int
	}{
		{nums: []int{100, 4, 200, 1, 3, 2}, expected: 4},
		{nums: []int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1}, expected: 9},
		{nums: []int{1, 4, 3, 3, 2}, expected: 2},
	}

	for _, input := range inputs {
		res := LongestMonotonicSubarray2(input.nums)
		if res != input.expected {
			t.Errorf("Nums:%v,expected:%v,got:%v", input.nums, input.expected, res)
		}
	}
}
