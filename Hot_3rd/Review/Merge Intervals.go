package review

import "sort"

func merge(intervals [][]int) [][]int {
	if len(intervals) == 1 {
		return intervals
	}
	res := [][]int{}

	sort.Slice(intervals, func(i, j int) bool {
		return (intervals[i][0] < intervals[j][0]) || ((intervals[i][0] == intervals[j][0]) && (intervals[i][1] < intervals[j][1]))
	})

	// Merge
	for _, cur := range intervals {
		if len(res) == 0 || res[len(res)-1][1] < cur[0] {
			res = append(res, cur)
		} else {
			res[len(res)-1][1] = max(res[len(res)-1][1], cur[1])
		}
	}
	return res
}
