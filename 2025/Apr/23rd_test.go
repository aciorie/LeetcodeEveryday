package apr

import "testing"

func TestCountLargestGroup(t *testing.T) {
	inputs := []struct {
		n    int
		want int
	}{
		{13, 4},
		{2, 2},
	}
	for _, input := range inputs {
		res := CountLargestGroup(input.n)
		if res != input.want {
			t.Errorf("n:%v,got:%v,want:%v", input.n, res, input.want)
		}
	}
}
