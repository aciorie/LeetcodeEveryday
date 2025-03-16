package mar

import (
	"reflect"
	"testing"
)

func TestRepairCars(t *testing.T) {
	inputs := []struct {
		ranks []int
		cars  int
		want  int64
	}{
		{[]int{4, 2, 3, 1}, 10, 16},
		{[]int{5, 1, 8}, 6, 16},
	}

	for _, input := range inputs {
		res := RepairCars(input.ranks, input.cars)
		if !reflect.DeepEqual(res, input.want) {
			t.Errorf("ranks:%v,cars:%v,want:%v,got:%v", input.ranks, input.cars, input.want, res)
		}
	}
}
