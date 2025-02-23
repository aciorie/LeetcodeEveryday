package contest

import (
	"reflect"
	"testing"
)

func TestSortMatrix(t *testing.T) {
	inputs := []struct {
		s    string
		want bool
	}{
		{"3902", true},
		{"34789", false},
	}

	for _, input := range inputs {
		res := HasSameDigits(input.s)
		if !reflect.DeepEqual(res, input.want) {
			t.Errorf("s:%v,want:%v,got:%v", input.s, input.want, res)
		}
	}
}
