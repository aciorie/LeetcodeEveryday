package mar

import (
	"reflect"
	"testing"
)

func TestApplyOperations(t *testing.T) {
	inputs := []struct {
		nums []int
		want []int
	}{
		{[]int{1, 2, 2, 1, 1, 0}, []int{1, 4, 2, 0, 0, 0}},
		{[]int{0, 1}, []int{1, 0}},
	}

	for _, input := range inputs {
		res := ApplyOperations(input.nums)
		if !reflect.DeepEqual(res, input.want) {
			t.Errorf("nums:%v,want:%v,got:%v", input.nums, input.want, res)
		}
	}
}
