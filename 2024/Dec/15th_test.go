package Dec

import (
	"math"
	"testing"
)

func TestMaxAverageRatio(t *testing.T) {
	inputs := []struct {
		classes       [][]int
		extraStudents int
		expected      float64
	}{
		{
			classes:       [][]int{{1, 2}, {3, 5}, {2, 2}},
			extraStudents: 2,
			expected:      0.78333,
		},
		{
			classes:       [][]int{{2, 4}, {3, 9}, {4, 5}, {2, 10}},
			extraStudents: 4,
			expected:      0.53485,
		},
	}

	for _, input := range inputs {
		res := MaxAverageRatio(input.classes, input.extraStudents)
		if math.Abs(res-input.expected) > 1e-5 {
			t.Errorf("classes:%v,extraStudents:%v,expected:%v,got:%v", input.classes, input.extraStudents, input.expected, res)
		}
	}
}
