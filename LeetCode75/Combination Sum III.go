package leetcode75

func combinationSum3(k int, n int) [][]int {
	var res [][]int
	dfs_216([]int{}, 1, n, k, &res)
	return res
}

func dfs_216(cur []int, index, n, k int, res *[][]int) {
	if len(cur) == k && sum(cur) == n {
		*res = append(*res, append([]int{}, cur...))
		return
	}
	if len(cur) > k || sum(cur) > n {
		return
	}
	for i := index; i < 10; i++ {
		cur = append(cur, i)
		dfs_216(cur, i+1, n, k, res)
		cur = cur[:len(cur)-1]
	}
}

func sum(cur []int) int {
	var res int
	for _, num := range cur {
		res += num
	}
	return res
}
