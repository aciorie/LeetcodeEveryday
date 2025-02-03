package hot

import (
	"reflect"
	"testing"
)

func TestSearch(t *testing.T) {
	inputs := []struct {
		nums   []int
		target int
		want   int
	}{
		{[]int{-1, 0, 3, 5, 9, 12}, 9, 4},
		{[]int{-1, 0, 3, 5, 9, 12}, 2, -1},
	}

	for _, input := range inputs {
		res := Search(input.nums, input.target)
		if !reflect.DeepEqual(res, input.want) {
			t.Errorf("nums:%v,target:%v,want:%v,got:%v", input.nums, input.target, input.want, res)
		}
	}
}
