package contest

import (
	"reflect"
	"testing"
)

func TestSolveQueries(t *testing.T) {
	inputs := []struct {
		nums    []int
		queries []int
		want    []int
	}{
		{[]int{1, 3, 1, 4, 1, 3, 2}, []int{0, 3, 5}, []int{2, -1, 3}},
		{[]int{1, 2, 3, 4}, []int{0, 1, 2, 3}, []int{-1, -1, -1, -1}},
	}

	for _, input := range inputs {
		res := SolveQueries(input.nums, input.queries)
		if !reflect.DeepEqual(res, input.want) {
			t.Errorf("nums:%v,queries:%v,want:%v,got:%v", input.nums, input.queries, input.want, res)
		}
	}
}
