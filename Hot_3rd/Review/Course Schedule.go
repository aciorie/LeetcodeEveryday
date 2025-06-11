package review

func canFinish(numCourses int, prerequisites [][]int) bool {
	// 邻接表 adj[i] 存储所有从课程 i 出发的边，即 i 是哪些课程的先决条件
	adj := make([][]int, numCourses)
	// 入度数组 inDegree[i] 存储有多少门课程是课程 i 的先决条件
	inDegree := make([]int, numCourses)

	for _, pre := range prerequisites {
		// 从 pre[1] 到 pre[0] 有一条边：bi -> ai
		adj[pre[1]] = append(adj[pre[1]], pre[0])
		inDegree[pre[0]]++
	}

	// 将所有入度为 0 的课程（没有先决条件）加入队列
	queue := []int{}
	for i := 0; i < numCourses; i++ {
		if inDegree[i] == 0 {
			queue = append(queue, i)
		}
	}

	// BFS
	coursesFinished := 0 // 记录成功学习的课程数量
	for len(queue) > 0 {
		course := queue[0]
		queue = queue[1:]

		coursesFinished++

		// 遍历当前课程的所有“后续课程”（即以当前课程为先决条件的课程）
		for _, nextCourse := range adj[course] {
			inDegree[nextCourse]--

			if inDegree[nextCourse] == 0 {
				queue = append(queue, nextCourse)
			}
		}
	}
	return coursesFinished == numCourses
}
