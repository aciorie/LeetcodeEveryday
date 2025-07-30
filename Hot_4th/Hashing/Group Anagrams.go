package hashing

import "sort"

func groupAnagrams(strs []string) [][]string {
	amaMap := make(map[string][]string)
	for _, str := range strs {
		runes := []rune(str)
		sort.Slice(runes, func(i, j int) bool {
			return runes[i] < runes[j]
		})

		sortedStr := string(runes)
		amaMap[sortedStr] = append(amaMap[sortedStr], str)
	}

	res := [][]string{}
	for _, group := range amaMap {
		res = append(res, group)
	}
	return res
}
