package Dec

import "testing"

func TestMaxChunksToSorted(t *testing.T) {
	inputs := []struct {
		arr      []int
		expected int
	}{
		{arr: []int{4, 3, 2, 1, 0}, expected: 1},
		{arr: []int{1, 0, 2, 3, 4}, expected: 4},
	}

	for _, input := range inputs {
		res := MaxChunksToSorted(input.arr)
		if res != input.expected {
			t.Errorf("arr:%v,expected:%v,got:%v", input.arr, input.expected, res)
		}
	}
}
