package aug

import (
	"strconv"
	"strings"
)

func largestGoodInteger(num string) string {
	if len(num) < 3 {
		return ""
	}
	res := ""
	for i := 999; i >= 0; i = i - 111 {
		str := strconv.Itoa(i)
		if i == 0 {
			str = "000"
		}
		if strings.Contains(num, str) {
			res = str
			break
		}
	}
	return res
}
