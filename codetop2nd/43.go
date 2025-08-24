package codetop2nd

import "strings"

func multiply(num1 string, num2 string) string {
	if num1 == "0" || num2 == "0" {
		return "0"
	}

	m, n := len(num1), len(num2)
	res := make([]int, m+n)

	// 从右向左遍历 num1 和 num2
	for i := m - 1; i >= 0; i-- {
		d1 := num1[i] - '0'
		for j := n - 1; j >= 0; j-- {
			d2 := num2[j] - '0'

			// 当前位的乘积
			product := int(d1) * int(d2)
			// 乘积将影响 res[i+j+1] 和 res[i+j] 这两个位置
			// sum = 当前乘积 + res[i+j+1] 的原有值
			sum := product + res[i+j+1]
			// 将 sum 的个位放到 res[i+j+1]
			res[i+j+1] = sum % 10
			// 将 sum 的十位（进位）加到 res[i+j]
			res[i+j] += sum / 10
		}
	}

	// 找到第一个非零的数字，处理前导零
	startIdx := 0
	if res[0] == 0 {
		startIdx = 1
	}

	// 构建最终结果字符串
	var builder strings.Builder
	for i := startIdx; i < len(res); i++ {
		builder.WriteByte(byte(res[i]) + '0')
	}

	return builder.String()
}
