package codetop2nd

import (
	"strconv"
	"strings"
)

func compareVersion(version1 string, version2 string) int {
	// 使用 strings.Split 将版本字符串按 "." 分割成切片
	v1Revisions := strings.Split(version1, ".")
	v2Revisions := strings.Split(version2, ".")

	// 获取两个版本号的修订位数
	n1, n2 := len(v1Revisions), len(v2Revisions)

	// 使用一个 for 循环，只要任何一个版本号还有修订位，就继续
	// 这个循环巧妙地处理了两个版本号长度不一致的情况
	for i := 0; i < n1 || i < n2; i++ {
		// 初始化当前修订号的数值
		var num1, num2 int

		// 如果当前索引 i 小于 v1 的长度，将该修订字符串转换为整数
		if i < n1 {
			// strconv.Atoi 将字符串转换为整数，忽略潜在的转换错误
			num1, _ = strconv.Atoi(v1Revisions[i])
		}

		// 如果当前索引 i 小于 v2 的长度，将该修订字符串转换为整数
		if i < n2 {
			num2, _ = strconv.Atoi(v2Revisions[i])
		}

		// 比较两个修订号
		if num1 > num2 {
			return 1 // v1 更大，直接返回 1
		}
		if num1 < num2 {
			return -1 // v2 更大，直接返回 -1
		}
		// 如果相等，继续循环比较下一个修订号
	}

	// 如果循环结束，说明所有修订号都相等
	return 0
}
