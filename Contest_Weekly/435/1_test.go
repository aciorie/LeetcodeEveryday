package contest

import "testing"

func TestMaxDifference(t *testing.T) {
	inputs := []struct {
		s    string
		want int
	}{
		{"aaaaabbc", 3},
		{"abcabcab", 1},
		{"mmsmsym", -1},
	}

	for _, input := range inputs {
		res := MaxDifference2(input.s)
		if res != input.want {
			t.Errorf("s:%v,want:%v,got:%v", input.s, input.want, res)
		}
	}
}
