package hot2nd

func insert(intervals [][]int, newInterval []int) [][]int {
	var res [][]int

	i := 0
	// Add non-overlapping intervals
	for i < len(intervals) && intervals[i][1] < newInterval[0] {
		res = append(res, intervals[i])
		i++
	}

	// Merge overlapping intervals
	for i < len(intervals) && intervals[i][0] <= newInterval[1] {
		newInterval[0] = min(intervals[i][0], newInterval[0])
		newInterval[1] = max(intervals[i][1], newInterval[1])
		i++
	}

	// Add the new merged interval
	res = append(res, newInterval)

	// Add non-overlapping intervals after the new interval
	for i < len(intervals) {
		res = append(res, intervals[i])
		i++
	}

	return res
}
