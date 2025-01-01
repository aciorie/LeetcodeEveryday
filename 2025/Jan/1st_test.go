package jan

import "testing"

func TestMaxScore(t *testing.T) {
	inputs := []struct {
		s    string
		want int
	}{
		{"011101", 5},
		{"00111", 5},
		{"1111", 3},
	}

	for _, input := range inputs {
		if input.want != MaxScore(input.s) {
			t.Errorf("s:%v,want:%v,got:%v", input.s, input.want, MaxScore(input.s))
		}
	}
}
