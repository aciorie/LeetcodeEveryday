package jan

import "testing"

func TestFirstCompleteIndex(t *testing.T) {
	inputs := []struct {
		arr  []int
		mat  [][]int
		want int
	}{
		{[]int{1, 3, 4, 2}, [][]int{{1, 4}, {2, 3}}, 2},
		{[]int{2, 8, 7, 4, 1, 3, 5, 6, 9}, [][]int{{3, 2, 5}, {1, 4, 6}, {8, 7, 9}}, 3},
	}

	for _, input := range inputs {
		res := FirstCompleteIndex(input.arr, input.mat)
		if res != input.want {
			t.Errorf("arr:%v,mat:%v,want:%v,got:%v", input.arr, input.mat, input.want, res)
		}

	}
}
