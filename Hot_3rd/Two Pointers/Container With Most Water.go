package main

func maxArea(height []int) int {
	l, r := 0, len(height)-1
	res := (r - l) * min(height[l], height[r])
	for l < r {
		if height[l] < height[r] {
			res = max(res, (r-l)*height[l])
			l++
		} else {
			res = max(res, (r-l)*height[r])
			r--
		}
	}
	return res
}
