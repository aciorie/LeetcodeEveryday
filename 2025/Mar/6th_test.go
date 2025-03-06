package mar

import (
	"reflect"
	"testing"
)

func TestFindMissingAndRepeatedValues(t *testing.T) {
	inputs := []struct {
		grid [][]int
		want []int
	}{
		{[][]int{{1, 3}, {2, 2}}, []int{2, 4}},
		{[][]int{{9, 1, 7}, {8, 9, 2}, {3, 4, 6}}, []int{9, 5}},
	}

	for _, input := range inputs {
		res := FindMissingAndRepeatedValues(input.grid)
		if !reflect.DeepEqual(res, input.want) {
			t.Errorf("grid:%v,want:%v,got:%v", input.grid, input.want, res)
		}
	}
}
