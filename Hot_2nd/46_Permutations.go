package hot2nd

func permute(nums []int) [][]int {
	res, cur, used := [][]int{}, []int{}, make([]bool, len(nums))

	var backtrack func()
	backtrack = func() {
		// Base Case: Permutation is complete
		if len(cur) == len(nums) {
			// cur is modified in place
			tmp := make([]int, len(cur))
			copy(tmp, cur)
			res = append(res, tmp)
			return
		}

		// Explore Choices
		for i := 0; i < len(nums); i++ {
			if !used[i] { // If the number at index 'i' hasn't been used yet

				// 1. Make the choice
				cur = append(cur, nums[i])
				used[i] = true

				// 2. Explore (recursive call)
				backtrack()

				// 3. Undo the choice (backtrack)
				cur = cur[:len(cur)-1] // Remove the last element
				used[i] = false
			}
		}
	}
	backtrack()
	return res
}
