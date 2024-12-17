package Dec

import "testing"

func TestRepeatLimitedString(t *testing.T) {
	inputs := []struct {
		s           string
		repeatLimit int
		expected    string
	}{
		{s: "cczazcc", repeatLimit: 3, expected: "zzcccac"},
		{s: "aababab", repeatLimit: 2, expected: "bbabaa"},
	}

	for _, input := range inputs {
		res := RepeatLimitedString(input.s, input.repeatLimit)
		if res != input.expected {
			t.Errorf("s:%v,repeatLimit:%v,expected:%v,got:%v", input.s, input.repeatLimit, input.expected, res)
		}
	}
}
