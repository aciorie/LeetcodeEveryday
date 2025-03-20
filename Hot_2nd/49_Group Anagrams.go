package hot2nd

import "sort"

func groupAnagrams(strs []string) [][]string {
	anagramMap := make(map[string][]string)

	for _, str := range strs {
		// 1. Convert string to rune slice
		runes := []rune(str)

		// 2. Sort rune slices
		sort.Slice(runes, func(i, j int) bool {
			return runes[i] < runes[j]
		})

		// 3. Convert the sorted rune slice back to a string
		sortedStr := string(runes)

		// 4. Find or create in map
		anagramMap[sortedStr] = append(anagramMap[sortedStr], str)
	}

	// 5. Extract all values ​​in the map (grouping results)
	res := [][]string{}
	for _, group := range anagramMap {
		res = append(res, group)
	}
	return res
}
