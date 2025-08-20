package codetop2nd

func addStrings(num1 string, num2 string) string {
	i, j := len(num1)-1, len(num2)-1
	carry := 0
	res := []byte{}

	for i >= 0 || j >= 0 || carry > 0 {
		d1, d2 := 0, 0
		if i >= 0 {
			d1 = int(num1[i] - '0')
			i--
		}
		if j >= 0 {
			d2 = int(num2[j] - '0')
			j--
		}

		sum := d1 + d2 + carry
		carry = sum / 10
		res = append(res, byte(sum%10+'0'))
	}

	for i, j := 0, len(res)-1; i < j; i, j = i+1, j-1 {
		res[i], res[j] = res[j], res[i]
	}

	return string(res)
}
