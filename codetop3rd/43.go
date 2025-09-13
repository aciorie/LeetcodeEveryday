package codetop3rd

import (
	"strconv"
	"strings"
)

func multiply(num1 string, num2 string) string {
	if num1 == "0" || num2 == "0" {
		return "0"
	}
	m, n := len(num1), len(num2)
	res := make([]int, m+n)
	for i := m - 1; i >= 0; i-- {
		digit1 := int(num1[i] - '0')
		for j := n - 1; j >= 0; j-- {
			digit2 := int(num2[j] - '0')

			product := digit1 * digit2

			// 确定乘积在结果数组中的位置
			// pos1是进位的位置，pos2是当前位的位置
			pos1, pos2 := i+j, i+j+1

			// 将乘积与当前位置已有的值相加
			sum := product + res[pos2]

			// 更新当前位的值
			res[pos2] = sum % 10

			// 将进位加到前一位
			res[pos1] += sum / 10
		}
	}
	// 将结果切片转换为字符串
	var resultBuilder strings.Builder

	// 跳过可能的前导零
	startIndex := 0
	if res[0] == 0 {
		startIndex = 1
	}

	for i := startIndex; i < len(res); i++ {
		resultBuilder.WriteString(strconv.Itoa(res[i]))
	}

	return resultBuilder.String()
}
