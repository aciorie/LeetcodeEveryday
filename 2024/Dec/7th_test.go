package Dec

import "testing"

func TestMinimumSize(t *testing.T) {
	inputs := []struct {
		nums          []int
		maxOperations int
		expected      int
	}{
		{nums: []int{9}, maxOperations: 2, expected: 3},
		{nums: []int{2, 4, 8, 2}, maxOperations: 4, expected: 2},
	}

	for _, input := range inputs {
		res := MinimumSize(input.nums, input.maxOperations)
		if res != input.expected {
			t.Errorf("nums:%v,maxOperations:%d,expected:%d,got:%d", input.nums, input.maxOperations, input.expected, res)
		}
	}
}
