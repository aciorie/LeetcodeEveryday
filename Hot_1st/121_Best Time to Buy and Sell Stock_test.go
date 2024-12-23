package hot

import "testing"

func TestMaxProfit(t *testing.T) {
	inputs := []struct {
		prices   []int
		expected int
	}{
		{prices: []int{7, 1, 5, 3, 6, 4}, expected: 5},
		{prices: []int{7, 6, 4, 3, 1}, expected: 0},
	}

	for _, input := range inputs {
		res := MaxProfit(input.prices)
		if res != input.expected {
			t.Errorf("Prices:%v,expected:%v,got:%v", input.prices, input.expected, res)
		}
	}
}
