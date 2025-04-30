package leetcode75

func largestAltitude(gain []int) int {
	cur, res := 0, 0
	for i := 0; i < len(gain); i++ {
		cur += gain[i]
		res = max(res, cur)
	}
	return res
}
