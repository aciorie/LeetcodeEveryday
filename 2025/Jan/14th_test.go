package jan

import (
	"reflect"
	"testing"
)

func TestFindThePrefixCommonArray(t *testing.T) {
	inputs := []struct {
		A    []int
		B    []int
		want []int
	}{
		{[]int{1, 3, 2, 4}, []int{3, 1, 2, 4}, []int{0, 2, 3, 4}},
		{[]int{2, 3, 1}, []int{3, 1, 2}, []int{0, 1, 3}},
	}

	for _, input := range inputs {
		res := FindThePrefixCommonArray(input.A, input.B)
		if !reflect.DeepEqual(input.want, res) {
			t.Errorf("A:%v,B:%v,want:%v,got:%v", input.A, input.B, input.want, res)
		}
	}
}
