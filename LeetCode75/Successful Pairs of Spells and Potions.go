package leetcode75

import "sort"

func successfulPairs(spells []int, potions []int, success int64) []int {
	n, m := len(spells), len(potions)
	res := make([]int, n)
	sort.Ints(potions)
	for i := 0; i < n; i++ {
		spell := spells[i]

		minIndex := sort.Search(m, func(j int) bool {
			return int64(spell)*int64(potions[j]) >= success
		})
		
		res[i] = m - minIndex
	}
	return res
}
