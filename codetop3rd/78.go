package codetop3rd

func subsets(nums []int) [][]int {
	n, res, cur := len(nums), make([][]int, 0), make([]int, 0)
	var dfs_78 func(int)
	dfs_78 = func(start int) {
		tmp := make([]int, len(cur))
		copy(tmp, cur)
		res = append(res, tmp)

		for i := start; i < n; i++ {
			cur = append(cur, nums[i])
			dfs_78(i + 1)
			cur = cur[:len(cur)-1]
		}
	}
	dfs_78(0)
	return res
}
