package feb

import "testing"

func TestNumOfSubarrays(t *testing.T) {
	inputs := []struct {
		arr  []int
		want int
	}{
		{[]int{1, 3, 5}, 4},
		{[]int{2, 4, 6}, 0},
		{[]int{1, 2, 3, 4, 5, 6, 7}, 16},
	}

	for _, input := range inputs {
		res := NumOfSubarrays2(input.arr)
		if res != input.want {
			t.Errorf("arr:%v,want:%v,got:%v", input.arr, input.want, res)
		}
	}
}
