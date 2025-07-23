package jul

func maximumGain(s string, x int, y int) int {
	n := len(s)
	if n < 2 {
		return 0
	}

	totalScore := 0

	var highPriorityPair, lowPriorityPair string
	var highScore, lowScore int

	if x > y {
		highPriorityPair = "ab"
		highScore = x
		lowPriorityPair = "ba"
		lowScore = y
	} else {
		highPriorityPair = "ba"
		highScore = y
		lowPriorityPair = "ab"
		lowScore = x
	}

	sAfterHighPriority := processString(s, highPriorityPair, highScore, &totalScore)

	_ = processString(sAfterHighPriority, lowPriorityPair, lowScore, &totalScore)

	return totalScore
}

func processString(s string, targetPair string, score int, totalScore *int) string {
	stack := make([]rune, 0, len(s))

	char1 := rune(targetPair[0])
	char2 := rune(targetPair[1])

	for _, char := range s {
		if len(stack) > 0 && stack[len(stack)-1] == char1 && char == char2 {
			stack = stack[:len(stack)-1]
			*totalScore += score
		} else {
			stack = append(stack, char)
		}
	}

	return string(stack)
}
