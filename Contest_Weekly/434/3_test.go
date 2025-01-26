package contest

import "testing"

func TestMaxFrequency(t *testing.T) {
	inputs := []struct {
		nums []int
		k    int
		want int
	}{
		{[]int{1, 2, 3, 4, 5, 6}, 1, 2},
		{[]int{10, 2, 3, 4, 5, 5, 4, 3, 2, 2}, 10, 4},
	}

	for _, input := range inputs {
		res := MaxFrequency2(input.nums, input.k)
		if res != input.want {
			t.Errorf("nums:%v,k:%v,want:%v,got:%v", input.nums, input.k, input.want, res)
		}
	}
}
