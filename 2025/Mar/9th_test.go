package mar

import (
	"reflect"
	"testing"
)

func TestNumberOfAlternatingGroups(t *testing.T) {
	inputs := []struct {
		colors []int
		k      int
		want   int
	}{
		{[]int{0, 1, 0, 1, 0}, 3, 3},
		{[]int{0, 1, 0, 0, 1, 0, 1}, 6, 2},
		{[]int{1, 1, 0, 1}, 4, 0},
	}

	for _, input := range inputs {
		res := NumberOfAlternatingGroups(input.colors, input.k)
		if !reflect.DeepEqual(res, input.want) {
			t.Errorf("colors:%v,k:%v,want:%v,got:%v", input.colors, input.k, input.want, res)
		}
	}
}
