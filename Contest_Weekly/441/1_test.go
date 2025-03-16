package contest

import "testing"

func TestMaxSum(t *testing.T) {
	inputs := []struct {
		nums []int
		want int
	}{
		{[]int{1, 2, 3, 4, 5}, 15},
		{[]int{1, 1, 0, 1, 1}, 1},
		{[]int{1, 2, -1, -2, 1, 0, -1}, 3},
		{[]int{-100}, -100},
	}

	for _, input := range inputs {
		res := MaxSum2(input.nums)
		if res != input.want {
			t.Errorf("nums:%v,want:%v,got:%v", input.nums, input.want, res)
		}
	}
}
