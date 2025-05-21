package hot3rd

func generateParenthesis(n int) []string {
	if n == 0 {
		return []string{}
	}
	if n == 1 {
		return []string{"()"}
	}
	var res []string
	var dfs func(int, int, string)
	dfs = func(start, end int, t string) {
		if start > n || end > n || start < end {
			return
		}
		if start == n || end == n {
			res = append(res, t)
			return
		}
		dfs(start+1, end, t+"(")
		dfs(start, end+1, t+")")
	}
	dfs(0, 0, "")
	return res
}
