package apr

import (
	"reflect"
	"testing"
)

func TestSubsetXORSum(t *testing.T) {
	inputs := []struct {
		nums []int
		want int
	}{
		{[]int{1, 3}, 6},
		{[]int{5, 1, 6}, 28},
		{[]int{3, 4, 5, 6, 7, 8}, 480},
	}

	for _, input := range inputs {
		res := SubsetXORSum(input.nums)
		if !reflect.DeepEqual(res, input.want) {
			t.Errorf("nums:%v,want:%v,got:%v", input.nums, input.want, res)
		}
	}
}
