package hot2nd

func isValid(s string) bool {
	n := len(s) - 1
	if n%2 == 1 {
		return false
	}
	stack := []rune{}
	for _, cha := range s {
		switch cha {
		case '(':
			stack = append(stack, ')')
		case '[':
			stack = append(stack, ']')
		case '{':
			stack = append(stack, '}')
		default:
			if len(stack) == 0 || cha != stack[len(stack)-1] {
				return false
			}
			stack = stack[:len(stack)-1]
		}
	}
	return len(stack) == 0
}
