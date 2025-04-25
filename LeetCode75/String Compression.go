package leetcode75

import "strconv"

func compress(chars []byte) int {
	write := 0
	read := 0
	n := len(chars)

	for read < n {
		currentChar := chars[read]
		count := 0

		for read < n && chars[read] == currentChar {
			read++
			count++
		}

		chars[write] = currentChar
		write++

		if count > 1 {
			countStr := strconv.Itoa(count)
			for i := 0; i < len(countStr); i++ {
				chars[write] = countStr[i]
				write++
			}
		}
	}

	return write
}
