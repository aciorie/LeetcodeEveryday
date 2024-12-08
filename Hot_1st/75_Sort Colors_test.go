package hot

import "testing"

func TestSortColors(t *testing.T) {
	inputs := []struct {
		nums     []int
		expected []int
	}{
		{nums: []int{2, 0, 2, 1, 1, 0}, expected: []int{0, 0, 1, 1, 2, 2}},
		{nums: []int{2, 0, 1}, expected: []int{0, 1, 2}},
	}

	for _, input := range inputs {
		tmp := input.nums
		SortColors(input.nums)
		if !equal75(input.nums, input.expected) {
			t.Errorf("Input:%v,expected:%v,got:%v", tmp, input.expected, input.nums)
		}
	}
}

func equal75(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := 0; i < len(a)-1; i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
