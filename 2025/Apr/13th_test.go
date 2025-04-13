package apr

import (
	"reflect"
	"testing"
)

func TestCountGoodNumbers(t *testing.T) {
	inputs := []struct {
		n    int64
		want int
	}{
		{1, 5},
		{4, 400},
		{50, 564908303},
	}

	for _, input := range inputs {
		res := CountGoodNumbers(input.n)
		if !reflect.DeepEqual(res, input.want) {
			t.Errorf("n:%v,want:%v,got:%v", input.n, input.want, res)
		}
	}
}
