package contest

import "testing"

func TestLongestPalindromicSubsequence(t *testing.T) {
	inputs := []struct {
		s    string
		k    int
		want int
	}{
		{"abced", 2, 3},
		{"aaazzz", 4, 6},
	}

	for _, input := range inputs {
		res := LongestPalindromicSubsequence(input.s, input.k)
		if res != input.want {
			t.Errorf("s:%v,k:%v,want:%v,got:%v", input.s, input.k, input.want, res)
		}
	}
}
