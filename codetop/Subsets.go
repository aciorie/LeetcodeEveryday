package codetop

func subsets(nums []int) [][]int {
	var res [][]int
	n, cur := len(nums), make([]int, 0)
	var dfs_78 func(int)
	dfs_78 = func(start int) {
		tmp := make([]int, len(cur))
		copy(tmp, cur)
		res = append(res, cur)

		for i := start; i < n; i++ {
			cur = append(cur, nums[i])
			dfs_78(i + 1)
			cur = cur[:len(cur)-1]
		}
	}
	dfs_78(0)
	return res
}
