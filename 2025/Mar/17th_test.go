package mar

import (
	"reflect"
	"testing"
)

func TestDivideArray(t *testing.T) {
	inputs := []struct {
		nums []int
		want bool
	}{
		{[]int{3, 2, 3, 2, 2, 2}, true},
		{[]int{1, 2, 3, 4}, false},
	}

	for _, input := range inputs {
		res := DivideArray(input.nums)
		if !reflect.DeepEqual(res, input.want) {
			t.Errorf("nums:%v,want:%v,got:%v", input.nums, input.want, res)
		}
	}
}
