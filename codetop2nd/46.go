package codetop2nd

func permute(nums []int) [][]int {
	res, cur, used, n := [][]int{}, []int{}, make([]bool, 0), len(nums)
	var dfs_46 func()
	dfs_46 = func() {
		if len(cur) == n {
			tmp := make([]int, len(cur))
			copy(tmp, cur)
			res = append(res, tmp)
			return
		}

		for i := 0; i < n; i++ {
			if !used[nums[i]] {
				used[nums[i]] = true
				cur = append(cur, nums[i])
				dfs_46()
				cur = cur[:len(cur)-1]
				used[nums[i]] = false
			}
		}
	}
	dfs_46()
	return res
}
