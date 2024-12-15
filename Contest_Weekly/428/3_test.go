package contest

import "testing"

func TestBeautifulSplits(t *testing.T) {
	inputs := []struct {
		nums     []int
		expected int
	}{
		{nums: []int{1, 1, 2, 1}, expected: 2},
		{nums: []int{1, 2, 3, 4}, expected: 0},
	}

	for _, input := range inputs {
		res := BeautifulSplits(input.nums)
		if res != input.expected {
			t.Errorf("Nums:%v,expected:%v,got:%v", input.nums, input.expected, res)
		}
	}
}
