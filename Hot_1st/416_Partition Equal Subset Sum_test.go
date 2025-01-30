package hot

import "testing"

func TestCanPartition(t *testing.T) {
	inputs := []struct {
		nums []int
		want bool
	}{
		{[]int{2, 2, 1, 1}, true},
		{[]int{1, 5, 11, 5}, true},
		{[]int{1, 2, 3, 5}, false},
	}

	for _, input := range inputs {
		res := CanPartition(input.nums)
		if res != input.want {
			t.Errorf("nums:%v,want:%v,got:%v", input.nums, input.want, res)
		}
	}
}
