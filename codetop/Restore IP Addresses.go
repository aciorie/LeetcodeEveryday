package codetop

import "strconv"

func restoreIpAddresses(s string) []string {
	var res []string

	var backtrack func(int, int, string)
	backtrack = func(left, numberOfDots int, currentCombination string) {
		if left > len(s)-1 {
			return
		}

		if numberOfDots == 1 {
			if s[left] == '0' && len(s)-left > 1 {
				return
			}

			if isValidIpAddress(s[left:]) {
				res = append(res, currentCombination+s[left:])
			}

			return
		}

		if s[left] == '0' {
			backtrack(left+1, numberOfDots-1, currentCombination+s[left:left+1]+".")
		} else {
			for i := left + 1; i < len(s); i++ {
				if isValidIpAddress(s[left:i]) {
					backtrack(i, numberOfDots-1, currentCombination+s[left:i]+".")
				}
			}
		}
	}

	backtrack(0, 4, "")

	return res
}

func isValidIpAddress(str string) bool {
	val, err := strconv.Atoi(str)

	if err != nil {
		panic(err)
	}

	return val >= 0 && val <= 255
}
