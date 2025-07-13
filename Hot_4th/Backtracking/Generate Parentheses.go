package backtracking

func generateParenthesis(n int) []string {
	var res []string
	var dfs_22 func(int, int, string)
	dfs_22 = func(l, r int, cur string) {
		if l > n || r > n || l < r {
			return
		}
		if l == n && r == n {
			res = append(res, cur)
			return
		}
		dfs_22(l+1, r, cur+"(")
		dfs_22(l, r+1, cur+")")
	}
	dfs_22(0, 0, "")
	return res
}
