package hot

import (
	"reflect"
	"testing"
)

func TestDailyTemperatures(t *testing.T) {
	inputs := []struct {
		temperatures []int
		want         []int
	}{
		{[]int{73, 74, 75, 71, 69, 72, 76, 73}, []int{1, 1, 4, 2, 1, 1, 0, 0}},
		{[]int{30, 40, 50, 60}, []int{1, 1, 1, 0}},
		{[]int{30, 60, 90}, []int{1, 1, 0}},
	}

	for _, input := range inputs {
		res := DailyTemperatures(input.temperatures)
		if !reflect.DeepEqual(res, input.want) {
			t.Errorf("temperatures:%v,want:%v,got:%v", input.temperatures, input.want, res)
		}
	}
}
