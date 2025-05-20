package hot3rd

func subsets(nums []int) [][]int {
	n, res, cur := len(nums), make([][]int, 0), make([]int, 0)
	var backtrack func(int)
	backtrack = func(index int) {
		tmp := make([]int, len(cur))
		copy(tmp, cur)
		res = append(res, tmp)

		for i := index; i < n; i++ {
			cur = append(cur, nums[i])
			backtrack(i + 1)
			cur = cur[:len(cur)-1]
		}
	}
	backtrack(0)
	return res
}
