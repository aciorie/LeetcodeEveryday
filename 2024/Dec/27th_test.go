package Dec

import "testing"

func TestMaxScoreSightseeingPair(t *testing.T) {
	inputs := []struct {
		values   []int
		expected int
	}{
		{values: []int{8, 1, 5, 2, 6}, expected: 11},
		{values: []int{1, 2}, expected: 2},
	}

	for _, input := range inputs {
		res := MaxScoreSightseeingPair(input.values)
		if res != input.expected {
			t.Errorf("values:%v,expected:%v,got:%v", input.values, input.expected, res)
		}
	}
}
