package contest

import (
	"testing"
)

func TestLargestInteger(t *testing.T) {
	inputs := []struct {
		nums []int
		k    int
		want int
	}{
		{[]int{3, 9, 2, 1, 7}, 3, 7},
		{[]int{3, 9, 7, 2, 1, 7}, 4, 3},
	}

	for _, input := range inputs {
		res := LargestInteger(input.nums, input.k)
		if res != input.want {
			t.Errorf("nums:%v,k:%v,want:%v,got:%v", input.nums, input.k, input.want, res)
		}
	}
}
