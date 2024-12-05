package hot

import "testing"

func TestMinDistance(t *testing.T) {
	inputs := []struct {
		word1    string
		word2    string
		expected int
	}{
		{word1: "horse", word2: "ros", expected: 3},
		{word1: "intention", word2: "execution", expected: 5},
		{word1: "", word2: "a", expected: 1},
	}

	for _, input := range inputs {
		res := MinDistance(input.word1, input.word2)
		if res != input.expected {
			t.Errorf("Word1:%v,word2:%v,expect:%v,got:%v", input.word1, input.word2, input.expected, res)
		}
	}
}
