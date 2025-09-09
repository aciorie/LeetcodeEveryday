package codetop3rd

import "strconv"

func addStrings(num1 string, num2 string) string {
	i, j := len(num1)-1, len(num2)-1
	carry := 0
	res := ""

	for i >= 0 || j >= 0 || carry != 0 {
		var n1, n2 int
		if i >= 0 {
			n1 = int(num1[i] - '0')
		}
		if j >= 0 {
			n2 = int(num2[j] - '0')
		}

		sum := n1 + n2 + carry
		res = strconv.Itoa(sum%10) + res
		carry = sum / 10

		i--
		j--
	}
	return res
}
