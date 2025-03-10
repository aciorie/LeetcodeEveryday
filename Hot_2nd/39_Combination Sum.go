package hot2nd

import "sort"

func combinationSum(candidates []int, target int) [][]int {
	sort.Ints(candidates)

	var res [][]int
	var cur []int

	var dfs func(int, int)
	dfs = func(target int, start int) {
		if target < 0 {
			return
		}
		if target == 0 {
			tmp := make([]int, len(cur))
			copy(tmp, cur)
			res = append(res, tmp)
			return
		}

		for i := start; i < len(candidates); i++ {
			cur = append(cur, candidates[i])
			dfs(target-candidates[i], i)
			cur = cur[:len(cur)-1]
		}
	}

	dfs(target, 0)
	return res
}
