package review

func partitionLabels(s string) []int {
	occurLastLetter := make(map[byte]int, 0)
	for i := 0; i < len(s); i++ {
		occurLastLetter[s[i]] = i
	}

	res := make([]int, 0) // 存储每个片段的长度
	start, end := 0, 0    // start: 当前片段的起始索引；end: 当前片段必须延伸到的最远索引

	// 遍历字符串，划分片段
	for i := 0; i < len(s); i++ {
		// 更新当前片段必须延伸到的最远位置
		// `end` 总是当前片段中所有字符的最后出现位置的最大值
		end = max(end, occurLastLetter[s[i]])

		// 如果当前遍历的索引 `i` 达到了当前片段必须延伸到的最远位置 `end`
		// 这意味着当前片段内的所有字符，它们的最后一次出现都在 `i` 或 `i` 之前
		// 此时，可以安全地进行切分
		if end == i {
			res = append(res, end - start + 1) // 计算当前片段的长度 (结束索引 - 起始索引 + 1)
			
			// 更新下一个片段的起始索引
			start = i + 1
		}
	}
	return res
}
