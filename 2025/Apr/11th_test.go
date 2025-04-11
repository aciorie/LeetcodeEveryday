package apr

import (
	"reflect"
	"testing"
)

func TestCountSymmetricIntegers(t *testing.T) {
	inputs := []struct {
		low  int
		high int
		want int
	}{
		{1, 100, 9},
		{1200, 1230, 4},
	}

	for _, input := range inputs {
		res := CountSymmetricIntegers(input.low, input.high)
		if !reflect.DeepEqual(res, input.want) {
			t.Errorf("low:%v,high:%v,want:%v,got:%v", input.low, input.high, input.want, res)
		}
	}
}
