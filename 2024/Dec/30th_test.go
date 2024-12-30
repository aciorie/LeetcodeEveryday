package Dec

import "testing"

func TestCountGoodStrings(t *testing.T) {
	inputs := []struct {
		low  int
		high int
		zero int
		one  int
		want int
	}{
		{3, 3, 1, 1, 8},
		{2, 3, 1, 2, 5},
	}

	for _, input := range inputs {
		res := CountGoodStrings(input.low, input.high, input.zero, input.one)
		if res != input.want {
			t.Errorf("low:%v,high:%v,zero:%v,one:%v,want:%v,got:%v", input.low, input.high, input.zero, input.one, input.want, res)
		}
	}
}
