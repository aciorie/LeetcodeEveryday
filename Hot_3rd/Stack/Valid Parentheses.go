package stack

func isValid(s string) bool {
	n := len(s)
	if n < 2 {
		return false
	}
	stack := []rune{}
	for _, c := range s {
		switch c {
		case '(':
			stack = append(stack, ')')
		case '[':
			stack = append(stack, ']')
		case '{':
			stack = append(stack, '}')
		default:
			if len(stack) == 0 || c != stack[len(stack)-1] {
				return false
			}
		}
	}
	return len(stack) == 0
}
