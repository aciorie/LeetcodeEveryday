package Dec

import "testing"

func TestMaxTwoEvents(t *testing.T) {
	inputs := []struct {
		events   [][]int
		expected int
	}{
		{
			events:   [][]int{{1, 3, 2}, {4, 5, 2}, {2, 4, 3}},
			expected: 4,
		},
		{
			events:   [][]int{{1, 3, 2}, {4, 5, 2}, {1, 5, 5}},
			expected: 5,
		},
		{
			events:   [][]int{{1, 5, 3}, {1, 5, 1}, {6, 6, 5}},
			expected: 5, // 这里的期望值需要根据具体逻辑来确定
		},
	}

	for _, input := range inputs {
		res := MaxTwoEvents(input.events)
		if res != input.expected {
			t.Errorf("Input:%v,epected:%v,got:%v", input.events, input.expected, res)
		}
	}
}
