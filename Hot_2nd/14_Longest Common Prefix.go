package hot2nd

func longestCommonPrefix(strs []string) string {
	res := strs[0]
	for i := 1; i < len(strs); i++ {
		for j := 0; j < len(res); j++ {
			if len(strs[i]) <= j || strs[i][j] != res[j] {
				res = res[0:j]
				break
			}
		}
	}
	return res
}