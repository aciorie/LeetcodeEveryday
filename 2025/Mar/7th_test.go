package mar

import (
	"reflect"
	"testing"
)

func TestClosestPrimes(t *testing.T) {
	inputs := []struct {
		left  int
		right int
		want  []int
	}{
		{10, 19, []int{11, 13}},
		{4, 6, []int{-1, -1}},
		{19, 31, []int{29, 31}},
	}

	for _, input := range inputs {
		res := ClosestPrimes2(input.left, input.right)
		if !reflect.DeepEqual(res, input.want) {
			t.Errorf("left:%v,right:%v,want:%v,got:%v", input.left, input.right, input.want, res)
		}
	}
}
