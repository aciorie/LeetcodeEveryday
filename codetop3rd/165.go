package codetop3rd

import (
	"strconv"
	"strings"
)

func compareVersion(version1 string, version2 string) int {
	v1Parts := strings.Split(version1, ".")
	v2Parts := strings.Split(version2, ".")

	m := len(v1Parts)
	n := len(v2Parts)

	maxLength := m
	if n > m {
		maxLength = n
	}

	for i := 0; i < maxLength; i++ {
		num1 := 0
		if i < m {
			// strconv.Atoi 函数会返回一个错误，这里我们忽略了
			num1, _ = strconv.Atoi(v1Parts[i])
		}

		num2 := 0
		if i < n {
			num2, _ = strconv.Atoi(v2Parts[i])
		}

		if num1 > num2 {
			return 1
		}
		if num1 < num2 {
			return -1
		}
	}

	return 0
}
