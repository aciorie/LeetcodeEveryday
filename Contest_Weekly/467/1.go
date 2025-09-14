package Contest

import "math"

func earliestTime(tasks [][]int) int {
	n := len(tasks)
	res := math.MaxInt32
	for i := 0; i < n; i++ {
		res = min(res, tasks[i][1]+tasks[i][0])
	}
	return res
}
