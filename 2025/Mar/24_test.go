package mar

import (
	"reflect"
	"testing"
)

func TestCountDays(t *testing.T) {
	inputs := []struct {
		days     int
		meetings [][]int
		want     int
	}{
		{10, [][]int{{5, 7}, {1, 3}, {9, 10}}, 2},
		{5, [][]int{{2, 4}, {1, 3}}, 1},
		{6, [][]int{{1, 6}}, 0},
	}

	for _, input := range inputs {
		res := CountDays(input.days, input.meetings)
		if !reflect.DeepEqual(res, input.want) {
			t.Errorf("days:%v,meetings:%v,want:%v,got:%v", input.days, input.meetings, input.want, res)
		}
	}
}
