package dynamicprogramming

func longestValidParentheses(s string) int {
	if len(s) == 0 {
		return 0
	}
	res := 0
	stack := []int{-1} 
	
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			stack = append(stack, i)
		} else {
			stack = stack[:len(stack)-1]
			if len(stack) == 0 {
				stack = append(stack, i)
			} else {
				cur := i - stack[len(stack)-1]
				if cur > res {
					res = cur
				}
			}
		}
	}
	return res
}
