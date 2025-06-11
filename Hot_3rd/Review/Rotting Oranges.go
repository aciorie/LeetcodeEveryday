package review

func orangesRotting(grid [][]int) int {
	rows := len(grid)
	cols := len(grid[0])
	queue := make([][]int, 0)
	freshOrangesCount := 0
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if grid[i][j] == 1 {
				freshOrangesCount++
			} else if grid[i][j] == 2 {
				queue = append(queue, []int{i, j})
			}
		}
	}

	if freshOrangesCount == 0 {
		return 0
	}

	minutes := 0
	direction := [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}
	for len(queue) > 0 && freshOrangesCount > 0 {
		curLevelSize := len(queue) // 当前这一分钟开始时，有多少个橘子是刚刚腐烂并在这一分钟内准备去感染它们的邻居

		for i := 0; i < curLevelSize; i++ {
			rottenOrange := queue[0]
			queue = queue[1:]

			row, col := rottenOrange[0], rottenOrange[1]
			// 检查上下左右四个方向的邻居
			for _, dir := range direction {
				neighborRow, neighborCol := row+dir[0], col+dir[1]

				// 检查邻居是否在网格范围内
				// 检查邻居是否是新鲜橘子 (值为 1)
				if neighborRow >= 0 && neighborRow < rows && neighborCol >= 0 && neighborCol < cols && grid[neighborRow][neighborCol] == 1 {
					grid[neighborRow][neighborCol] = 2                     // 感染新鲜橘子
					freshOrangesCount--                                    // 减少新鲜橘子的计数
					queue = append(queue, []int{neighborRow, neighborCol}) // 将新腐烂的橘子加入队列，它们将在下一分钟感染其邻居
				}
			}
		}

		if len(queue) > 0 { // 只有当队列中还有新被感染的橘子时，才算作时间流逝
			minutes++
		}
	}

	// 循环结束后，判断是否所有新鲜橘子都已腐烂
	if freshOrangesCount == 0 {
		return minutes // 所有橘子都腐烂了，返回经过的总分钟数
	} else {
		return -1 // 仍有新鲜橘子未腐烂 (被孤立)，返回 -1
	}
}
