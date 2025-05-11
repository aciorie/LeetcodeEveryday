package leetcode75

func canVisitAllRooms(rooms [][]int) bool {
	stack := []int{}
	seen := make([]bool, len(rooms))
	seen[0] = true
	stack = append(stack, 0)

	for len(stack) > 0 {
		curRoom := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		for _, nextRoom := range rooms[curRoom] {
			if !seen[nextRoom] {
				seen[nextRoom] = true
				stack = append(stack, nextRoom)
			}
		}
	}
	for _, v := range seen {
		if !v {
			return false
		}
	}
	return true
}
