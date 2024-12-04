package hot

import "testing"

func TestClimbStairs(t *testing.T) {
	tests := []struct {
		input    int
		expected int
	}{
		{input: 2, expected: 2},
		{input: 3, expected: 3},
		{input: 45, expected: 1836311903},
	}

	for _, test := range tests {
		res := ClimbStairs(test.input)
		if !equal70(res, test.expected) {
			t.Errorf("Input:%v,expected:%v,got:%v", test.input, test.expected, res)
		}
	}

	for _, test := range tests {
		res := ClimbStairs2(test.input)
		if !equal70(res, test.expected) {
			t.Errorf("Input:%v,expected:%v,got:%v", test.input, test.expected, res)
		}
	}
}
func equal70(a, b int) bool {
	return a == b
}
