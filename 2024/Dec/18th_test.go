package Dec

import "testing"

func TestFinalPrices(t *testing.T) {
	inputs := []struct {
		prices   []int
		expected []int
	}{
		{prices: []int{8, 4, 6, 2, 3}, expected: []int{4, 2, 4, 2, 3}},
		{prices: []int{1, 2, 3, 4, 5}, expected: []int{1, 2, 3, 4, 5}},
		{prices: []int{10, 1, 1, 6}, expected: []int{9, 0, 1, 6}},
	}

	for _, input := range inputs {
		res := FinalPrices(input.prices)
		if !equal1DIntArray(res, input.expected) {
			t.Errorf("prices:%v,expected:%v,got:%v", input.prices, input.expected, res)
		}
	}
}
