package Dec

import "testing"

func TestFindScore(t *testing.T) {
	inputs := []struct {
		nums     []int
		expected int64
	}{
		{nums: []int{2, 1, 3, 4, 5, 2}, expected: 7},
		{nums: []int{2, 3, 5, 1, 3, 2}, expected: 5},
	}

	for _, input := range inputs {
		res := FindScore(input.nums)
		if res != input.expected {
			t.Errorf("Nums:%v,expected:%v,got:%v", input.nums, input.expected, res)
		}
	}
}
