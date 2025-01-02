package hot

import "testing"

func TestFindMin(t *testing.T) {
	inputs := []struct {
		nums []int
		want int
	}{
		{[]int{3, 4, 5, 1, 2}, 1},
		{[]int{4, 5, 6, 7, 0, 1, 2}, 0},
		{[]int{11, 13, 15, 17}, 11},
	}

	for _, input := range inputs {
		res := FindMin(input.nums)
		if res != input.want {
			t.Errorf("nums:%v,got:%v,want:%v", input.nums, res, input.want)
		}
	}
}
