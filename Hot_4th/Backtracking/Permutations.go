package backtracking

func permute(nums []int) [][]int {
	res, cur, used := [][]int{}, []int{}, make([]bool, len(nums))
	var dfs_46 func()
	dfs_46 = func() {
		if len(cur) == len(nums) {
			tmp := make([]int, len(cur))
			copy(tmp, cur)
			res = append(res, tmp)
			return
		}

		for i := 0; i < len(nums); i++ {
			if !used[i] {
				cur = append(cur, nums[i])
				used[i] = true
				dfs_46()
				cur = cur[:len(cur)-1]
				used[i] = false
			}
		}
	}
	dfs_46()
	return res
}
