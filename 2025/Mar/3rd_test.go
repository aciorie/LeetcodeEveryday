package mar

import (
	"reflect"
	"testing"
)

func TestPivotArray(t *testing.T) {
	inputs := []struct {
		nums  []int
		pivot int
		want  []int
	}{
		{[]int{9, 12, 5, 10, 14, 3, 10}, 10, []int{9, 5, 3, 10, 10, 12, 14}},
		{[]int{-3, 4, 3, 2}, 2, []int{-3, 2, 4, 3}},
	}

	for _, input := range inputs {
		res := PivotArray(input.nums, input.pivot)
		if !reflect.DeepEqual(res, input.want) {
			t.Errorf("nums:%v,pivot:%v,want:%v,got:%v", input.nums, input.pivot, input.want, res)
		}
	}
}
