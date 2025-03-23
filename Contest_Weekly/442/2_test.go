package contest

import (
	"reflect"
	"testing"
)

func TestNumberOfComponents(t *testing.T) {
	inputs := []struct {
		properties [][]int
		k          int
		want       int
	}{
		{[][]int{{1, 2, 3}, {2, 3, 4}, {4, 3, 5}}, 2, 1},
		{[][]int{{1, 1}, {1, 1}}, 2, 2},
	}

	for _, input := range inputs {
		res := NumberOfComponents(input.properties, input.k)
		if !reflect.DeepEqual(res, input.want) {
			t.Errorf("properties:%v,k:%v,want:%v,got:%v", input.properties, input.k, input.want, res)
		}
	}
}
