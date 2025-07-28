package codetop

func addStrings(num1 string, num2 string) string {
	i, j := len(num1), len(num2)
	carry := 0
	res := []byte{}
	for i >= 0 || j >= 0 || carry > 0 {
		digit1, digit2 := 0, 0

		if i >= 0 {
			digit1 = int(num1[i] - '0')
			i--
		}
		if j >= 0 {
			digit2 = int(num2[j] - '0')
			j--
		}

		sum := digit1 + digit2 + carry
		carry = sum / 10
		curDigit := sum % 10

		res = append(res, byte(curDigit+'0'))
	}
	for k, l := 0, len(res)-1; k < l; k, l = k+1, l-1 {
		res[k], res[l] = res[l], res[k]
	}

	return string(res)
}
