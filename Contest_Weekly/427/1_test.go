package contest

import "testing"

func TestConstructTransformedArray(t *testing.T) {
	inputs := []struct {
		nums     []int
		expected []int
	}{
		{nums: []int{3, -2, 1, 1}, expected: []int{1, 1, 1, 3}},
		{nums: []int{-1, 4, -1}, expected: []int{-1, -1, 4}},
	}

	for _, input := range inputs {
		res := ConstructTransformedArray(input.nums)
		if equalArray(res, input.expected) {
			t.Errorf("input:%v,expect:%v,got:%v", input.nums, input.expected, res)
		}
	}
}

func equalArray(a, b []int) bool {
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
