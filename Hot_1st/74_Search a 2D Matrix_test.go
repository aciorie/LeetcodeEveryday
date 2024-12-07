package hot

import "testing"

func TestSearchMatrix(t *testing.T) {
	inputs := []struct {
		matrix   [][]int
		target   int
		expected bool
	}{
		{matrix: [][]int{{1, 3, 4, 5}, {10, 11, 16, 20}, {23, 30, 34, 60}}, target: 3, expected: true},
		{matrix: [][]int{{1, 3, 4, 5}, {10, 11, 16, 20}, {23, 30, 34, 60}}, target: 13, expected: false},
		{matrix: [][]int{{-8, -7, -5, -3, -3, -1, 1}, {2, 2, 2, 3, 3, 5, 7}, {8, 9, 11, 11, 13, 15, 17}, {18, 18, 18, 20, 20, 20, 21}, {23, 24, 26, 26, 26, 27, 27}, {28, 29, 29, 30, 32, 32, 34}}, target: -5, expected: true},
	}

	for _, input := range inputs {
		res := SearchMatrix(input.matrix, input.target)
		if res != input.expected {
			t.Errorf("Matrix:%v,target:%v,expected:%v,got:%v", input.matrix, input.target, input.expected, res)
		}

		res2 := SearchMatrix2(input.matrix, input.target)
		if res2 != input.expected {
			t.Errorf("Matrix:%v,target:%v,expected:%v,got:%v", input.matrix, input.target, input.expected, res)
		}
	}
}
