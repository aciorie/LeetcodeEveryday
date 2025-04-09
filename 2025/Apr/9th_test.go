package apr

import (
	"reflect"
	"testing"
)

func TestMinOperations(t *testing.T) {
	inputs := []struct {
		nums []int
		k    int
		want int
	}{
		{[]int{5, 2, 5, 4, 5}, 2, 2},
		{[]int{2, 1, 2}, 2, -1},
		{[]int{9, 7, 5, 3}, 1, 4},
	}

	for _, input := range inputs {
		res := MinOperations(input.nums, input.k)
		if !reflect.DeepEqual(res, input.want) {
			t.Errorf("nums:%v,k:%v,want:%v,got:%v", input.nums, input.k, input.want, res)
		}
	}
}
