package leetcode75

import "sort"

func findMinArrowShots(points [][]int) int {
	if len(points) < 2 {
		return 1
	}
	sort.Slice(points, func(i, j int) bool {
		return points[i][1] < points[j][1]
	})

	res := 1
	lastXEnd := points[0][1]

	for i := 1; i < len(points); i++ {
		curBalloon := points[i]
		curXStart := curBalloon[0]
		curXEnd := curBalloon[1]

		if curXStart > lastXEnd {
			res++
			lastXEnd = curXEnd
		}
	}

	return res
}
