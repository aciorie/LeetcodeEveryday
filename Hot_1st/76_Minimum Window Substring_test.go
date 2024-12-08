package hot

import "testing"

func TestMinWindow(t *testing.T) {
	inputs := []struct {
		s        string
		t        string
		expected string
	}{
		{s: "ADOBECODEBANC", t: "ABC", expected: "BANC"},
		{s: "a", t: "a", expected: "a"},
		{s: "a", t: "aa", expected: ""},
	}

	for _, input := range inputs {
		res := MinWindow(input.s, input.t)
		if res != input.expected {
			t.Errorf("s:%v,t:%v,expected:%v,got:%v", input.s, input.t, input.expected, res)
		}
	}
}
