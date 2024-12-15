package contest

import "testing"

func TestButtonWithLongestTime(t *testing.T) {
	inputs := []struct {
		events   [][]int
		expected int
	}{
		{
			events:   [][]int{{1, 2}, {2, 5}, {3, 9}, {1, 15}},
			expected: 1,
		},
		{
			events:   [][]int{{10, 5}, {1, 7}},
			expected: 10,
		},
	}

	for _, input := range inputs {
		res := ButtonWithLongestTime(input.events)
		if res != input.expected {
			t.Errorf("Events:%v,expected:%v,got:%v", input.events, input.events, res)
		}
	}
}
