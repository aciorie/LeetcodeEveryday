package feb

import "testing"

func TestMaximumSum(t *testing.T) {
	inputs := []struct {
		nums []int
		want int
	}{
		{[]int{18, 43, 36, 13, 7}, 54},
		{[]int{10, 12, 19, 14}, -1},
	}
	for _, input := range inputs {
		res := MaximumSum(input.nums)
		if res != input.want {
			t.Errorf("nums:%v,want:%v,got:%v", input.nums, input.want, res)
		}
	}
}
