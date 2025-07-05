package jul

func FindLucky(arr []int) int {
	if len(arr) == 0 {
		return 0
	}
	freq := make(map[int]int)
	for _, num := range arr {
		freq[num]++
	}
	res := -1
	for k, v := range freq {
		if k == v {
			if k > res {
				res = k
			}
		}
	}
	return res
}
