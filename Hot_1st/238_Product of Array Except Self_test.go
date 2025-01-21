package hot

import (
	"reflect"
	"testing"
)

func TestProductExceptSelf(t *testing.T) {
	inputs := []struct {
		nums []int
		want []int
	}{
		{[]int{1, 2, 3, 4}, []int{24, 12, 8, 6}},
		{[]int{-1, 1, 0, -3, 3}, []int{0, 0, 9, 0, 0}},
	}

	for _, input := range inputs {
		res := ProductExceptSelf(input.nums)
		if !reflect.DeepEqual(res, input.want) {
			t.Errorf("nums:%v,want:%v,got:%v", input.nums, input.want, res)
		}
	}
}
