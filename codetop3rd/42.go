package codetop3rd

func trap(height []int) int {
	l, r, lmax, rmax, res := 0, len(height)-1, 0, 0, 0
	for l < r {
		if height[l] < height[r] {
			lmax = max(lmax, height[l])
			res += lmax - height[l]
		} else {
			rmax = max(rmax, height[r])
			res += rmax - height[r]
		}
	}
	return res
}
