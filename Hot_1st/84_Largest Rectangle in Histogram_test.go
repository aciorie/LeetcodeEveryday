package hot

import "testing"

//failed at solutions so does the test file
func TestLargestRectangleArea(t *testing.T) {
	inputs := []struct {
		heights  []int
		expected int
	}{
		{heights: []int{2, 1, 5, 6, 2, 3}, expected: 10},
		{heights: []int{2, 4}, expected: 4},
	}
	for _, input := range inputs {
		// res := LargestRectangleArea(input.heights)
		res := 0
		if res != input.expected {
			t.Errorf("heights:%v,expected:%v,got:%v", input.heights, input.expected, res)
		}
	}
}
