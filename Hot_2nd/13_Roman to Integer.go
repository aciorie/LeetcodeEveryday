package hot2nd

func romanToInt(s string) int {
	values := map[byte]int{'I': 1, 'V': 5, 'X': 10, 'L': 50, 'C': 100, 'D': 500, 'M': 1000}
	res := 0
	for i := 0; i < len(s); i++ {
		if i+1 < len(s) && values[s[i]] < values[s[i+1]] {
			res -= values[s[i]]
		} else {
			res += values[s[i]]
		}
	}
	return res
}
