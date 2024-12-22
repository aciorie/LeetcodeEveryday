package contest

import "testing"

func TestMaxDistinctElements(t *testing.T) {
	inputs := []struct {
		nums     []int
		k        int
		expected int
	}{
		{nums: []int{1, 2, 2, 3, 3, 4}, k: 2, expected: 6},
		{nums: []int{4, 4, 4, 4}, k: 1, expected: 3},
	}
	for _, input := range inputs {
		res := MaxDistinctElements(input.nums, input.k)
		if res != input.expected {
			t.Errorf("nums:%v,k:%v,expected:%v,got:%v", input.nums, input.k, input.expected, res)
		}
	}
}
