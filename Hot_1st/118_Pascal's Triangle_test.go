package hot

import "testing"

func TestGenerate(t *testing.T) {
	inputs := []struct {
		numRows  int
		expected [][]int
	}{
		{numRows: 5, expected: [][]int{{1}, {1, 1}, {1, 2, 1}, {1, 3, 3, 1}, {1, 4, 6, 4, 1}}},
		{numRows: 1, expected: [][]int{{1}}},
	}

	for _, input := range inputs {
		res := Generate(input.numRows)
		if !equalTwoDArray(res, input.expected) {
			t.Errorf("numrows:%v,expected:%v,got:%v", input.numRows, input.expected, res)
		}
	}
}
