package Dec

import "testing"

func TestIsArraySpecial(t *testing.T) {
	inputs := []struct {
		nums     []int
		queries  [][]int
		expected []bool
	}{
		{nums: []int{3, 4, 1, 2, 6}, queries: [][]int{{0, 4}}, expected: []bool{false}},
		{nums: []int{4, 3, 1, 6}, queries: [][]int{{0, 2}, {2, 3}}, expected: []bool{false, true}},
	}

	for _, input := range inputs {
		res := IsArraySpecial(input.nums, input.queries)
		if !equalBoolArray(res, input.expected) {
			t.Errorf("nums:%v,queries:%v,expected:%v,got:%v", input.nums, input.queries, input.expected, res)
		}
	}
}

func equalBoolArray(a, b []bool) bool {
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
