package jan

import "testing"

func TestCountPalindromicSubsequence(t *testing.T) {
	inputs := []struct {
		s    string
		want int
	}{
		{"aabca", 3},
		{"abc", 0},
		{"bbcbaba", 4},
	}

	for _, input := range inputs {
		res := CountPalindromicSubsequence(input.s)
		if res != input.want {
			t.Errorf("s:%v,want:%v,got:%v", input.s, input.want, res)
		}
	}
}
