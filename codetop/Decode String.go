package codetop

func decodeString(s string) string {
	numstack, strstack := make([]int, 0), make([]string, 0)
	res, multi := "", 0

	for _, ch := range s {
		// 如果是数字
		if ch >= '0' && ch <= '9' {
			multi = multi*10 + int(ch-'0')
		} else if ch == '[' {
			// 如果是左括号，表示进入一个新的嵌套
			// 将当前的重复次数和已解码字符串压入栈
			numstack = append(numstack, multi)
			strstack = append(strstack, res)

			multi = 0
			res = ""
		} else if ch == ']' {
			// 如果是右括号，表示一个嵌套层结束

			k := numstack[len(numstack)-1]
			numstack = numstack[:len(numstack)-1]

			prevstr := strstack[len(strstack)-1]
			strstack = strstack[:len(strstack)-1]

			// 将当前字符串重复k次，并拼接到上一层字符串后面
			repeatedstr := ""
			for range k {
				repeatedstr += res
			}
			res = prevstr + repeatedstr
		} else {
			// 如果是字母，直接拼接到当前字符串
			res += string(ch)
		}
	}
	return res
}
