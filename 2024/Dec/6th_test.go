package Dec

import "testing"

func TestMaxCount(t *testing.T) {
	inputs := []struct {
		banned   []int
		n        int
		maxSum   int
		expected int
	}{
		{banned: []int{1, 6, 5}, n: 5, maxSum: 6, expected: 2},
		{banned: []int{1, 2, 3, 4, 5, 6, 7}, n: 8, maxSum: 1, expected: 0},
		{banned: []int{11}, n: 7, maxSum: 50, expected: 7},
	}

	for _, input := range inputs {
		res := MaxCount(input.banned, input.n, input.maxSum)
		if res != input.expected {
			t.Errorf("banned:%v,n:%v,maxSum:%v,expected:%v,got:%v", input.banned, input.n, input.maxSum, input.expected, res)
		}
	}
}
