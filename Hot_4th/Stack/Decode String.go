package stack

func decodeString(s string) string {
	numStack, strStack := []int{}, []string{}
	res, multi := "", 0 // multi 存放当前正在解析的数字

	for _, ch := range s {
		// 如果是数字
		if ch >= '0' && ch <= '9' {
			multi = multi*10 + int(ch-'0')
		} else if ch == '[' {
			numStack = append(numStack, multi)
			strStack = append(strStack, res)

			multi = 0
			res = ""
		} else if ch == ']' {
			// 如果是右括号，表示一个嵌套层结束
			// 弹出数字栈顶的重复次数
			k := numStack[len(numStack)-1]
			numStack = numStack[:len(numStack)-1]

			// 弹出字符串栈顶的上一层字符串
			prevStr := strStack[len(strStack)-1]
			strStack = strStack[:len(strStack)-1]

			// 将当前字符串重复k次，并拼接到上一层字符串后面
			repeatedStr := ""
			for i := 0; i < k; i++ {
				repeatedStr += res
			}
			res = prevStr + repeatedStr
		} else {
			// 如果是字母，直接拼接到当前字符串
			res += string(ch)
		}
	}
	return res
}
