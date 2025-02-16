package hot2nd

func maxArea(height []int) int {
	l, r := 0, len(height)-1
	res := min(height[l], height[r]) * (r - l)
	for l < r {
		if height[l] < height[r] {
			res = max(res, height[l]*(r-l))
			l++
		} else {
			res = max(res, height[r]*(r-l))
			r--
		}
	}
	return res
}
