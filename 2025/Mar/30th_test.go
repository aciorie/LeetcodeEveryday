package mar

import (
	"reflect"
	"testing"
)

func TestPartitionLabels(t *testing.T) {
	inputs := []struct {
		s    string
		want []int
	}{
		{"ababcbacadefegdehijhklij", []int{9, 7, 8}},
		{"eccbbbbdec", []int{10}},
	}

	for _, input := range inputs {
		res := PartitionLabels(input.s)
		if !reflect.DeepEqual(input.want, res) {
			t.Errorf("s:%v,want:%v,got:%v", input.s, input.want, res)
		}
	}
}
