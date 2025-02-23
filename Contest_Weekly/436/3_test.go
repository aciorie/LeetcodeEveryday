package contest

import "testing"

func TestCountSubstrings(t *testing.T) {
	inputs := []struct {
		s    string
		want int64
	}{
		{"12936", 11},
		{"5701283", 18},
		{"1010101010", 25},
	}

	for _, input := range inputs {
		res := CountSubstrings(input.s)
		if res != input.want {
			t.Errorf("s:%v,want:%v,got:%v", input.s, input.want, res)
		}
	}
}
