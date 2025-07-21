package jul

import "strings"

func makeFancyString(s string) string {
	if len(s) < 3 {
		return s
	}
	var res strings.Builder
	res.WriteString(s[:2])

	for i := 2; i < len(s); i++ {
		if s[i] == res.String()[res.Len()-2] && s[i] == res.String()[res.Len()-1] {
			continue
		} else {
			res.WriteByte(s[i])
		}
	}

	return res.String()
}
