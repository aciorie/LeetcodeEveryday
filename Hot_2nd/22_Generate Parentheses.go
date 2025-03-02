package hot2nd

func generateParenthesis(n int) []string {
	var res []string
	var dfs func(int, int, string)
	// t represents the current bracket sequence
	dfs = func(l, r int, t string) {
		if l > n || r > n || l < r {
			return
		}
		if l == n && r == n {
			res = append(res, t)
			return
		}
		dfs(l+1, r, t+"(")
		dfs(l, r+1, t+")")
	}
	dfs(0, 0, "")
	return res
}
