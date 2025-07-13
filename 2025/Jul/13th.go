package jul

import "sort"

func matchPlayersAndTrainers(players []int, trainers []int) int {
	if len(players) == 0 || len(trainers) == 0 {
		return 0
	}
	sort.Ints(players)
	sort.Ints(trainers)
	res := 0
	for i, j := 0, 0; i < len(players) && j < len(trainers); {
		if players[i] <= trainers[j] {
			res++
			i++
			j++
		} else {
			j++
		}
	}
	return res
}
