package jan

import "testing"

func TestWaysToSplitArray(t *testing.T) {
	inputs := []struct {
		nums []int
		want int
	}{
		{[]int{10, 4, -8, 7}, 2},
		{[]int{2, 3, 1, 0}, 2},
	}

	for _, input := range inputs {
		res := WaysToSplitArray(input.nums)
		if res != input.want {
			t.Errorf("nums:%v,want:%v,got:%v", input.nums, input.want, res)
		}
	}
}
