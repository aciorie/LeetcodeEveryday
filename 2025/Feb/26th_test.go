package feb

import "testing"

func TestMaxAbsoluteSum(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{
			name: "Example 1",
			nums: []int{1, -3, 2, 3, -4},
			want: 5,
		},
		{
			name: "Example 2",
			nums: []int{2, -5, 1, -4, 3, -2},
			want: 8,
		},
		{
			name: "Test Case 3",
			nums: []int{-1, -1, -1, -1, -1},
			want: 5, // abs(-5)
		},
		{
			name: "Test Case 4",
			nums: []int{1, 1, 1, 1, 1},
			want: 5,
		},
		{
			name: "Test Case 5",
			nums: []int{0, 0, 0, 0, 0},
			want: 0,
		},
		{
			name: "Test Case 6",
			nums: []int{-100, 1, 2, -3, 4},
			want: 100, // abs(-100)
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MaxAbsoluteSum(tt.nums); got != tt.want {
				t.Errorf("MaxAbsoluteSum() = %v, want %v", got, tt.want)
			}
		})
	}
}