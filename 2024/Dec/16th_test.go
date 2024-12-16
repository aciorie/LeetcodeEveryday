package Dec

import (
	"testing"
)

func TestGetFinalState(t *testing.T) {
	inputs := []struct {
		nums       []int
		k          int
		multiplier int
		expected   []int
	}{
		{nums: []int{2, 1, 3, 5, 6}, k: 5, multiplier: 2, expected: []int{8, 4, 6, 5, 6}},
		{nums: []int{1, 2}, k: 3, multiplier: 4, expected: []int{16, 8}},
	}

	for _, input := range inputs {
		res := GetFinalState(input.nums, input.k, input.multiplier)
		if !equal1DIntArray(res, input.expected) {
			t.Errorf("nums:%v,k:%v,multiplier:%v,expected:%v,got:%v", input.nums, input.k, input.multiplier, input.expected, res)
		}
	}
}

func equal1DIntArray(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
