package leetcode75

func reverseVowels(s string) string {
	vowels := "aeiouAEIOU"
	vowelPositions := []int{}
	vowelChars := []rune{}

	for i, char := range s {
		if isVowel(char, vowels) {
			vowelPositions = append(vowelPositions, i)
			vowelChars = append(vowelChars, char)
		}
	}

	for i, j := 0, len(vowelChars)-1; i < j; i, j = i+1, j-1 {
		vowelChars[i], vowelChars[j] = vowelChars[j], vowelChars[i]
	}

	result := []rune(s)

	for i, pos := range vowelPositions {
		result[pos] = vowelChars[i]
	}

	return string(result)
}

func isVowel(char rune, vowels string) bool {
	for _, v := range vowels {
		if char == v {
			return true
		}
	}
	return false
}
