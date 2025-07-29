package greedy

func partitionLabels(s string) []int {
	occurLast := make(map[byte]int, 0)
	for i := 0; i < len(s); i++ {
		occurLast[s[i]] = i
	}

	res := make([]int, 0)
	start, last := 0, 0
	for i := 0; i < len(s); i++ {
		last = occurLast[s[i]]
		if i == last {
			res = append(res, last-start+1)
			start = i + 1
		}
	}
	return res
}
