package contest

import "testing"

func TestCountPartitions(t *testing.T) {
	inputs := []struct {
		nums []int
		want int
	}{
		{[]int{10, 10, 3, 7, 6}, 4},
		{[]int{1, 2, 2}, 0},
		{[]int{2, 4, 6, 8}, 3},
	}

	for _, input := range inputs {
		res := CountPartitions(input.nums)
		if res != input.want {
			t.Errorf("nums:%v,want:%v,got:%v", input.nums, input.want, res)
		}
	}
}
