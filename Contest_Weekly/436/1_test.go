package contest

import (
	"reflect"
	"testing"
)

func TestSortMatrix(t *testing.T) {
	inputs := []struct {
		grid [][]int
		want [][]int
	}{
		{[][]int{{1, 7, 3}, {9, 8, 2}, {4, 5, 6}}, [][]int{{8, 2, 3}, {9, 6, 7}, {4, 5, 1}}},
		{[][]int{{0, 1}, {1, 2}}, [][]int{{2, 1}, {1, 0}}},
		{[][]int{{1}}, [][]int{{1}}},
	}

	for _, input := range inputs {
		res := SortMatrix(input.grid)
		if !reflect.DeepEqual(res, input.want) {
			t.Errorf("grid:%v,want:%v,got:%v", input.grid, input.want, res)
		}
	}
}
