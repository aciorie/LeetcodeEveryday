package apr

import (
	"reflect"
	"testing"
)

func TestMaximumTripletValue(t *testing.T) {
	inputs := []struct {
		nums []int
		want int64
	}{
		{[]int{12, 6, 1, 2, 7}, 77},
		{[]int{1, 10, 3, 4, 19}, 133},
		{[]int{1, 2, 3}, 0},
	}

	for _, input := range inputs {
		res := MaximumTripletValue(input.nums)
		if !reflect.DeepEqual(res, input.want) {
			t.Errorf("nums:%v,want:%v,got:%v", input.nums, input.want, res)
		}
	}
}
