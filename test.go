package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	// 包含额外格式化标签的字符串
	str := ""

	// 编译一个正则表达式来匹配浮点数（包含数字和小数点）
	// MustCompile 会在正则表达式无效时 panic，适合在初始化时使用
	re := regexp.MustCompile(`[0-9\.]+`)

	// 使用 " & " 作为分隔符来分割字符串
	stringParts := strings.Split(str, " & ")

	// 初始化一个变量来存储总和
	var total float64

	// 遍历分割后的字符串切片
	for _, part := range stringParts {
		// 从每个部分中找到匹配正则表达式的第一个子字符串（也就是数字）
		numberStr := re.FindString(part)

		// 如果没有找到数字，就跳过这个部分
		if numberStr == "" {
			continue
		}

		// 将提取出的数字字符串转换为 float64
		// 这里我们假设提取出的已经是干净的数字，所以不再需要 TrimSpace
		num, err := strconv.ParseFloat(numberStr, 64)
		if err != nil {
			// 如果在转换过程中发生错误，打印错误信息
			fmt.Printf("无法转换 '%s' (从 '%s' 提取): %v\n", numberStr, part, err)
			continue
		}

		// 将转换后的数字加到总和中
		total += num
	}

	// 打印最终的总和
	fmt.Printf("字符串中的数字总和是: %f\n，平均数：%f\n", total,total/8)
}
