package codetop2nd

func generateParenthesis(n int) []string {
	if n == 0 {
		return []string{}
	}
	res := make([]string, 0)
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
