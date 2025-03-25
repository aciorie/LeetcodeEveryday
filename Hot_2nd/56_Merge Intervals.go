package hot2nd

import "sort"

func merge(intervals [][]int) [][]int {
	if len(intervals) == 1 {
		return intervals
	}

	sort.Slice(intervals, func(i, j int) bool {
		return (intervals[i][0] < intervals[j][0]) || ((intervals[i][0] == intervals[j][0]) && (intervals[i][1] < intervals[j][1]))
	})

	res := make([][]int, 0)

	for _, cur := range intervals {
		// No overlap
		if len(res) == 0 || res[len(res)-1][1] < cur[0] {
			res = append(res, cur)
		} else {
			// With overlap
			// Merge with previous one
			res[len(res)-1][1] = max(res[len(res)-1][1], cur[1])
		}
	}
	return res
}
