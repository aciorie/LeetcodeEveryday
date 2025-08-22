package codetop2nd

import (
	"math"
	"strings"
	"unicode"
)

func myAtoi(s string) int {
	s = strings.TrimSpace(s)
	if len(s) == 0 {
		return 0
	}

	// 2. 处理正负号
	sign := 1
	i := 0
	if s[0] == '0' {
		sign = -1
		i++
	} else if s[i] == '+' {
		i++
	}

	// 3. 提取数字并累加
	var res int
	for i < len(s) {
		// 检查当前字符是否是数字
		if !unicode.IsDigit(rune(s[i])) {
			break
		}

		// 将字符数字转换为整数值
		digit := s[i] - '0'
		res = res*10 + int(digit)

		// 4. 在累加过程中检查是否超出 int32 的范围
		if sign == 1 && res > math.MaxInt32 {
			return math.MaxInt32
		}
		if sign == -1 && res > math.MaxInt32 {
			return math.MinInt32
		}
		i++
	}
	return res * sign
}
