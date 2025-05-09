package hot2nd

func trap(height []int) int {
	l, r, lMax, rMax, res := 0, len(height)-1, 0, 0, 0

	for l < r {
		if height[l] < height[r] {
			if height[l] >= lMax {
				lMax = height[l]
			} else {
				res += lMax - height[l]
			}
			l++
		} else {
			if height[r] >= rMax {
				rMax = height[r]
			} else {
				res += rMax - height[r]
			}
			r--
		}
	}
	return res
}
