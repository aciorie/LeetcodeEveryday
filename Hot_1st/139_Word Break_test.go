package hot

import "testing"

func TestWordBreak(t *testing.T) {
	inputs := []struct {
		s        string
		wordDict []string
		expected bool
	}{
		{"leetcode", []string{"leet", "code"}, true},
		{"applepenapple", []string{"apple", "pen"}, true},
		{"catsandog", []string{"cats", "dog", "sand", "and", "cat"}, false},
	}

	for _, input := range inputs {
		res := WordBreak(input.s, input.wordDict)
		if res != input.expected {
			t.Errorf("s:%v,wordDict:%v,expected:%v,got:%v", input.s, input.wordDict, input.expected, res)
		}
	}
}
