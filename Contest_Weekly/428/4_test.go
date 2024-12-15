package contest

import "testing"

func TestMakeStringGood(t *testing.T) {
	inputs := []struct {
		s        string
		expected int
	}{
		{s: "acab", expected: 1},
		{s: "wddw", expected: 0},
		{s: "aaabc", expected: 2},
	}

	for _, input := range inputs {
		res := MakeStringGood(input.s)
		if res != input.expected {
			t.Errorf("s:%v,expected:%v,got:%v", input.s, input.expected, res)
		}
	}
}
