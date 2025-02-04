package feb

import "testing"

func TestMaxAscendingSum(t *testing.T) {
	inputs := []struct {
		nums     []int
		expected int
	}{
		{nums: []int{10, 20, 30, 5, 10, 50}, expected: 65},
		{nums: []int{10, 20, 30, 40, 50}, expected: 150},
		{nums: []int{12, 17, 15, 13, 10, 11, 12}, expected: 33},
		{nums: []int{100,10,1}, expected: 100},
		{nums: []int{3,6,10,1,8,9,9,8,9}, expected: 19},
	}

	for _, input := range inputs {
		res := MaxAscendingSum3(input.nums)
		if res != input.expected {
			t.Errorf("Nums:%v,expected:%v,got:%v", input.nums, input.expected, res)
		}
	}
}
