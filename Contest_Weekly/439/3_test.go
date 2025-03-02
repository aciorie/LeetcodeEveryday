package contest

import "testing"

func TestMaxSum(t *testing.T) {
	inputs := []struct {
		nums []int
		k    int
		m    int
		want int
	}{
		{[]int{1, 2, -1, 3, 3, 4}, 2, 2, 13},
		{[]int{-10, 3, -1, -2}, 4, 1, -10},
	}

	for _, input := range inputs {
		res := MaxSum(input.nums, input.k, input.m)
		if res != input.want {
			t.Errorf("nums:%v,k:%v,m:%v,want:%v,got:%v", input.nums, input.k, input.m, input.want, res)
		}
	}
}
