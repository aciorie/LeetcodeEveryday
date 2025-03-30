package Contest

import (
	"reflect"
	"testing"
)

func TestMinCosts(t *testing.T) {
	inputs := []struct {
		cost []int
		want []int
	}{
		{[]int{5, 3, 4, 1, 3, 2}, []int{5, 3, 3, 1, 1, 1}},
		{[]int{1, 2, 4, 6, 7}, []int{1, 1, 1, 1, 1}},
	}

	for _, input := range inputs {
		res := MinCosts(input.cost)
		if !reflect.DeepEqual(res, input.want) {
			t.Errorf("cost:%v,want:%v,got:%v", input.cost, input.want, res)
		}
	}
}
