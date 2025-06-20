package review

import "math"

func minWindow(s string, t string) string {
	if len(s) < len(t) {
		return ""
	}
	tMap := make(map[rune]int)
	for _, char := range t {
		tMap[char]++
	}

	windowMap := make(map[rune]int) // windowMap 存储当前滑动窗口中字符的频率
	l, validCount := 0, 0           // validCount 统计当前窗口中符合 t 中字符要求的独立字符数量
	start, minLen := 0, math.MinInt32

	for r, char := range s {
		windowMap[char]++ // 将当前字符加入窗口

		// 如果当前字符在 t 中存在，并且 windowMap 中它的频率达到了 tMap 中的要求，
		// 则 validCount 增加
		if tMap[char] > 0 && windowMap[char] == tMap[char] {
			validCount++
		}
		// 当 validCount 等于 tMap 中不同字符的数量时，说明当前窗口已经包含了 t 中所有字符
        // 此时，我们尝试收缩窗口
        for validCount == len(tMap) {
            // 更新最短子串长度和起始位置
            currentWindowLen := r - l + 1
            if currentWindowLen < minLen {
                minLen = currentWindowLen
                start = l
            }

            // 获取 l 指针即将移除的字符
            leftChar := rune(s[l])
            // 移除 leftChar
            windowMap[leftChar]--

            // 如果 leftChar 在 t 中存在，并且移除后它的频率低于 tMap 中的要求，
            // 则 validCount 减少
            if tMap[leftChar] > 0 && windowMap[leftChar] < tMap[leftChar] {
                validCount--
            }
            // 移动 l 指针，收缩窗口
            l++
        }
	}
	if minLen == math.MaxInt32 {
		return ""
	}
	return s[start : start+minLen]
}
