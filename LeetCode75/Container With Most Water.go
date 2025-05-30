package leetcode75

func maxArea(height []int) int {
	l, r := 0, len(height)-1
	res := (r - l) * min(height[l], height[r])
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
