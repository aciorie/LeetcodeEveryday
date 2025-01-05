package jan

import "testing"

func TestShiftingLetters(t *testing.T) {
	inputs := []struct {
		s      string
		shifts [][]int
		want   string
	}{
		{"abc", [][]int{{0, 1, 0}, {1, 2, 1}, {0, 2, 1}}, "ace"},
		{"dztz", [][]int{{0, 0, 0}, {1, 1, 1}}, "catz"},
	}

	for _, input := range inputs {
		res := ShiftingLetters(input.s, input.shifts)
		if res != input.want {
			t.Errorf("s:%v,shifts:%v,want:%v,got:%v", input.s, input.shifts, input.want, res)
		}
	}
}
