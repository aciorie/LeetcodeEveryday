package codetop2nd

func isValid(s string) bool {
	n := len(s)
	if n%2 == 1 {
		return false
	}
	stack := make([]rune, 0)
	for _, ch := range s {
		switch ch {
		case '(':
			stack = append(stack, ')')
		case '[':
			stack = append(stack, ']')
		case '{':
			stack = append(stack, '}')
		default:
			if len(stack) == 0 || stack[len(stack)-1] != ch {
				return false
			}
			stack = stack[:len(stack)-1]
		}
	}
	return len(stack) == 0
}
