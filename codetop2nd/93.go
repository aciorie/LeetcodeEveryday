package codetop2nd

import (
	"strconv"
	"strings"
)

func restoreIpAddresses(s string) []string {
	// 最终结果切片
	var result []string

	// 存储当前路径的 IP 段
	var segments []int

	// 声明一个匿名函数作为回溯函数
	var backtrack func(start int)

	// 定义回溯函数
	backtrack = func(start int) {
		// 终止条件：如果已经找到4个段
		if len(segments) == 4 {
			// 如果同时用完了字符串 s，说明找到一个有效解
			if start == len(s) {
				// 将 IP 段转换成字符串并添加到结果中
				var ipBuilder strings.Builder
				for i, seg := range segments {
					ipBuilder.WriteString(strconv.Itoa(seg))
					if i < 3 {
						ipBuilder.WriteByte('.')
					}
				}
				result = append(result, ipBuilder.String())
			}
			return
		}

		// 遍历所有可能的段长度（1, 2, 3）
		for length := 1; length <= 3; length++ {
			// 计算当前段的结束索引
			end := start + length

			// 剪枝：如果超出字符串边界，跳出循环
			if end > len(s) {
				break
			}

			// 提取当前段的子串
			currentSegStr := s[start:end]

			// 剪枝：处理前导零
			// 只有当段长度大于1且第一个字符是'0'时，这个段无效
			if length > 1 && currentSegStr[0] == '0' {
				continue
			}

			// 将子串转换为整数
			val, _ := strconv.Atoi(currentSegStr)

			// 剪枝：检查 IP 段的数值是否有效（0-255）
			if val <= 255 {
				// 做选择：将当前段添加到 segments
				segments = append(segments, val)

				// 递归：探索下一个段
				backtrack(end)

				// 撤销选择（回溯）：移除当前段，探索其他可能性
				segments = segments[:len(segments)-1]
			}
		}
	}

	// 启动回溯
	backtrack(0)

	return result
}
