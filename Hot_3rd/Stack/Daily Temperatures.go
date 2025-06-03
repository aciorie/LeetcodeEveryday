package stack

func dailyTemperatures(temperatures []int) []int {
	n := len(temperatures)
	res, stack := make([]int, n), make([]int, 0)

	for i := 0; i < n; i++ {
		for len(stack) > 0 && temperatures[i] > temperatures[len(stack)-1] {
			index := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			res[index] = i - index
		}
		stack = append(stack, i)
	}

	return res
}
