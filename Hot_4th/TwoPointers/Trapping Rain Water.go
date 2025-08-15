package twopointers

func trap(height []int) int {
	l, r, lMax, rMax, res := 0, len(height)-1, 0, 0, 0

	for l < r {
		if height[l] < height[r] {
			lMax = max(lMax, height[l])
			res += lMax - height[l]
			l++
		} else {
			rMax = max(rMax, height[r])
			res += rMax - height[r]
			r--
		}
	}
	return res
}
