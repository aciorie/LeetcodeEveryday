package contest

import "testing"

func TestMaxContainers(t *testing.T) {
	testCases := []struct {
		n         int
		w         int
		maxWeight int
		expected  int
	}{
		{2, 3, 15, 4},
		{3, 5, 20, 4},
		{1, 1, 1, 1},
		{10, 2, 5, 2},
		{1000, 1, 1000000, 1000000}, // Test large n
		{2, 10, 5, 0},               // Test maxWeight smaller than w
		{5, 2, 51, 25},
		{2, 2, 7, 3},
	}

	for _, tc := range testCases {
		result := MaxContainers(tc.n, tc.w, tc.maxWeight)
		if result != tc.expected {
			t.Errorf("For n=%d, w=%d, maxWeight=%d, expected %d, but got %d", tc.n, tc.w, tc.maxWeight, tc.expected, result)
		}
	}
}
