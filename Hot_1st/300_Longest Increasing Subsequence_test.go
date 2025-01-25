package hot

import "testing"

func TestLengthOfLIS2(t *testing.T) {
	inputs := []struct {
		nums []int
		want int
	}{
		{[]int{10, 9, 2, 5, 3, 7, 101, 18}, 4},
		{[]int{0, 1, 0, 3, 2, 3}, 4},
		{[]int{7, 7, 7, 7, 7, 7, 7}, 1},
		{[]int{1, 3, 6, 7, 9, 4, 10, 5, 6}, 5},
	}

	for _, input := range inputs {
		res := LengthOfLIS2(input.nums)
		if res != input.want {
			t.Errorf("nums:%v,want:%v,got:%v", input.nums, input.want, res)
		}
	}
}
