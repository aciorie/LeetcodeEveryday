package hot

import "testing"

func TestFindKthLargest(t *testing.T) {
	inputs := []struct {
		nums []int
		k    int
		want int
	}{
		{[]int{3, 2, 1, 5, 6, 4}, 2, 5},
		{[]int{3, 2, 3, 1, 2, 4, 5, 5, 6}, 4, 4},
	}

	for _, input := range inputs {
		res := FindKthLargest(input.nums, input.k)
		if res != input.want {
			t.Errorf("nums:%v,k:%v,want:%v,got:%v", input.nums, input.k, input.want, res)
		}
	}
}
