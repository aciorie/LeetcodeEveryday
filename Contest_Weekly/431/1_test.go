package contest

import "testing"

func TestMaxLength(t *testing.T) {
	inputs := []struct {
		nums []int
		want int
	}{
		{[]int{1, 2, 1, 2, 1, 1, 1}, 5},
		{[]int{2, 3, 4, 5, 6}, 3},
		{[]int{1, 2, 3, 1, 4, 5, 1}, 5},
	}

	for _, input := range inputs {
		res := MaxLength(input.nums)
		if res != input.want {
			t.Errorf("nums:%v,want:%v,got:%v", input.nums, input.want, res)
		}
	}
}
