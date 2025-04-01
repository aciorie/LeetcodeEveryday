package apr

import (
	"reflect"
	"testing"
)

func TestMostPoints(t *testing.T) {
	inputs := []struct {
		questions [][]int
		want      int64
	}{
		{[][]int{{3, 2}, {4, 3}, {4, 4}, {2, 5}}, 5},
		{[][]int{{1, 1}, {2, 2}, {3, 3}, {4, 4}, {5, 5}}, 7},
	}

	for _, input := range inputs {
		res := MostPoints(input.questions)
		if !reflect.DeepEqual(res, input.want) {
			t.Errorf("questions:%v,want:%v,got:%v", input.questions, input.want, res)
		}
	}
}
