package hot

import (
	"reflect"
	"testing"
)

func TestPartition(t *testing.T) {
	inputs := []struct {
		s        string
		expected [][]string
	}{
		{s: "aab", expected: [][]string{{"a", "a", "b"}, {"aa", "b"}}},
		{s: "a", expected: [][]string{{"a"}}},
	}

	for _, input := range inputs {
		res := Partition(input.s)
		if !reflect.DeepEqual(res, input.expected) {
			t.Errorf("s:%v,expected:%v,got:%v", input.s, input.expected, res)
		}
	}
}
