package codetop3rd

func permute(nums []int) [][]int {
	res, cur, used := [][]int{}, []int{}, make([]bool, len(nums))

	var backtrack func()
	backtrack = func() {
		if len(cur) == len(nums) {
			tmp := make([]int, len(cur))
			copy(tmp, cur)
			res = append(res, tmp)
			return
		}

		for i := 0; i < len(nums); i++ {
			if !used[i] {
				used[i] = true
				cur = append(cur, nums[i])
				backtrack()
				used[i] = false
				cur = cur[:len(cur)-1]
			}
		}
	}
	backtrack()
	return res
}
