package jan

import (
	"reflect"
	"testing"
)

func TestMinOperations(t *testing.T) {
	inputs := []struct {
		boxes string
		want  []int
	}{
		{"110", []int{1, 1, 3}},
		{"001011", []int{11, 8, 5, 4, 3, 4}},
	}

	for _, input := range inputs {
		res := MinOperations(input.boxes)
		if !reflect.DeepEqual(res, input.want) {
			t.Errorf("boxes:%v,want:%v,got:%v", input.boxes, input.want, res)
		}
	}
}
