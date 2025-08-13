package stack

import "math"

func largestRectangleArea(heights []int) int {
	if len(heights) < 2 {
		return heights[0]
	}

	res := 0
	for i := 0; i < len(heights); i++ {
		minHeight := math.MaxInt32
		for j := i; j < len(heights); j++ {
			// Fine minimum height
			for k := i; k <= j; k++ {
				if heights[k] < minHeight {
					minHeight = heights[k]
				}
			}
			// Now got minimum height
			res = max(minHeight*(j-i+1), res)
		}
	}
	return res
}
