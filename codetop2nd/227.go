package codetop2nd

func calculate(s string) int {
	n := len(s)
	if n == 0 {
		return 0
	}

	numstack := make([]int, 0)

	// num 用于构建多位数字
	num := 0
	// op 用于记录上一个遇到的运算符，初始为 '+' 以处理第一个数
	op := byte('+')

	for i := 0; i < n; i++ {
		// 1. 忽略空格
		if s[i] == ' ' {
			continue
		}

		// 2. 解析多位数字
		if s[i] >= '0' && s[i] <= '9' {
			num = num*10 + int(s[i]-'0')
		}

		// 3. 遇到运算符或到达字符串末尾时，处理上一个数字
		// 这里的逻辑处理了两种情况：
		// a) 遇到了新的运算符
		// b) 遍历到了字符串的最后一个字符
		if s[i] < '0' || s[i] > '9' || i == n-1 {
			switch op {
			case '+':
				numstack = append(numstack, num)
			case '-':
				numstack = append(numstack, -num)
			case '*':
				top := numstack[len(numstack)-1]
				numstack = numstack[:len(numstack)-1] // 弹出栈顶元素
				numstack = append(numstack, top*num)
			case '/':
				top := numstack[len(numstack)-1]
				numstack = numstack[:len(numstack)-1] // 弹出栈顶元素
				numstack = append(numstack, top/num)
			}
			// 更新运算符，并重置数字
			op = s[i]
			num = 0
		}
	}

	// 4. 将栈中所有数字相加，得到最终结果
	result := 0
	for _, v := range numstack {
		result += v
	}
	return result
}
