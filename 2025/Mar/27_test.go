package mar

import (
	"reflect"
	"testing"
)

func TestMinimumIndex(t *testing.T) {
	inputs := []struct {
		nums []int
		want int
	}{
		{[]int{1, 2, 2, 2}, 2},
		{[]int{2, 1, 3, 1, 1, 1, 7, 1, 2, 1}, 4},
		{[]int{3, 3, 3, 3, 7, 2, 2}, -1},
	}

	for _, input := range inputs {
		res := MinimumIndex(input.nums)
		if !reflect.DeepEqual(res, input.want) {
			t.Errorf("nums:%v,want:%v,got:%v", input.nums, input.want, res)
		}
	}
}
