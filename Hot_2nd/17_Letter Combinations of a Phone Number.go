package hot2nd

func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}
	var res []string
	dict := []string{"", "", "abc", "def", "ghi", "jkl", "mno", "pqrs", "tuv", "wxyz"}
	dfs(digits, dict, 0, "", &res)
	return res
}

func dfs(digits string, dict []string, pos int, cur string, res *[]string) {
	if pos == len(digits) {
		*res = append(*res, cur)
		return
	}
	str := dict[digits[pos]-'0']
	for i := 0; i < len(str); i++ {
		dfs(digits, dict, pos+1, cur+string(str[i]), res)
	}
}