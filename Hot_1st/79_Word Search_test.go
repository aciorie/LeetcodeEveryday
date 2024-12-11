package hot

import "testing"

func TestExist(t *testing.T) {
	inputs := []struct {
		board    [][]byte
		word     string
		expected bool
	}{{
		board:    [][]byte{{'A', 'B', 'C', 'E'}, {'S', 'F', 'C', 'S'}, {'A', 'D', 'E', 'E'}},
		word:     "ABCCED",
		expected: true,
	}, {
		board:    [][]byte{{'A', 'B', 'C', 'E'}, {'S', 'F', 'C', 'S'}, {'A', 'D', 'E', 'E'}},
		word:     "SEE",
		expected: true,
	}, {
		board:    [][]byte{{'A', 'B', 'C', 'E'}, {'S', 'F', 'C', 'S'}, {'A', 'D', 'E', 'E'}},
		word:     "ABCB",
		expected: false,
	},
	}
	for _, input := range inputs {
		res := Exist(input.board, input.word)
		if res != input.expected {
			t.Errorf("board:%v,word:%v,expected:%v,got:%v", input.board, input.word, input.expected, res)
		}
	}
}
