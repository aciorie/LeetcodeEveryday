package contest

import "testing"

func TestAnswerString(t *testing.T) {
	inputs := []struct {
		word       string
		numFriends int
		expected   string
	}{
		{"dbca", 2, "dbc"},
		{"gggg", 4, "g"},
	}

	for _, input := range inputs {
		res := AnswerString(input.word, input.numFriends)
		if res != input.expected {
			t.Errorf("word:%v,numFriends:%v,expected:%v,got:%v", input.word, input.numFriends, input.expected, res)
		}
	}
}
