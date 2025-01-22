package hot

import "testing"

func TestSearchMatrix_240_2(t *testing.T) {
	inputs := []struct {
		matrix [][]int
		target int
		want   bool
	}{
		{[][]int{{1, 4, 7, 11, 15}, {2, 5, 8, 12, 19}, {3, 6, 9, 16, 22}, {10, 13, 14, 17, 24}, {18, 21, 23, 26, 30}}, 5, true},
		{[][]int{{1, 4, 7, 11, 15}, {2, 5, 8, 12, 19}, {3, 6, 9, 16, 22}, {10, 13, 14, 17, 24}, {18, 21, 23, 26, 30}}, 20, false},
	}

	for _, input := range inputs {
		res := SearchMatrix_240(input.matrix, input.target)
		if res != input.want {
			t.Errorf("matrix:%v,target:%v,want:%v,got:%v", input.matrix, input.target, input.want, res)
		}
	}
}
