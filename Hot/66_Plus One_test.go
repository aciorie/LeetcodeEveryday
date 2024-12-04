package hot

import (
	"testing"
)

func TestPlusOne(t *testing.T) {
	tests := []struct {
		input    []int
		expected []int
	}{
		{input: []int{1, 2, 3}, expected: []int{1, 2, 4}},
		{input: []int{4, 3, 2, 1}, expected: []int{4, 3, 2, 2}},
		{input: []int{9}, expected: []int{1, 0}},
		{input: []int{8, 9, 9}, expected: []int{9, 0, 0}},
	}

	for _, test := range tests {
		res := PlusOne(test.input)
		if !equal66(res, test.expected) {
			t.Errorf("Input:%v,expected:%v,got:%v", test.input, test.expected, res)
		}
	}
}

func equal66(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
