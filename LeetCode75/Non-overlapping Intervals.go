package leetcode75

import "sort"

func eraseOverlapIntervals(intervals [][]int) int {
	if len(intervals) < 2 {
		return 0
	}

	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][1] < intervals[j][1]
	})

	res := 0
	lastEndTime := intervals[0][1]

	for i := 1; i < len(intervals); i++ {
		currentInterval := intervals[i]
		currentStartTime := currentInterval[0]

		if currentStartTime < lastEndTime {
			res++
		} else {
			lastEndTime = currentInterval[1]
		}
	}
	return res
}
