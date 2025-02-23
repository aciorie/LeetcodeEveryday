package contest

import (
	"reflect"
	"testing"
)

func TestMaxSum(t *testing.T) {
	inputs := []struct {
		grid   [][]int
		limits []int
		k      int
		want   int64
	}{
		{[][]int{{1, 2}, {3, 4}}, []int{1, 2}, 2, 7},
		{[][]int{{5, 3, 7}, {8, 2, 6}}, []int{2, 2}, 3, 21},
	}

	for _, input := range inputs {
		res := MaxSum(input.grid, input.limits, input.k)
		if !reflect.DeepEqual(res, input.want) {
			t.Errorf("grid:%v,limits:%v,k:%v,want:%v,got:%v", input.grid, input.limits, input.k, input.want, res)
		}
	}
}
