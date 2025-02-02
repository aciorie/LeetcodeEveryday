package hot

import (
	"reflect"
	"testing"
)

func TestSubarraySum(t *testing.T) {
	inputs := []struct {
		nums []int
		k    int
		want int
	}{
		{[]int{1, 1, 1}, 2, 2},
		{[]int{1, 2, 3}, 3, 2},
	}

	for _, input := range inputs {
		res := SubarraySum(input.nums, input.k)
		if !reflect.DeepEqual(res, input.want) {
			t.Errorf("nums:%v,k:%v,want:%v,got:%v", input.nums, input.k, input.want, res)
		}
	}
}
