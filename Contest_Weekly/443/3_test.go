package Contest

import "testing"

func TestLongestPalindrome2(t *testing.T) {
	inputs := []struct {
		t    string
		k    string
		want int
	}{
		{"a", "a", 2},
		{"abc", "def", 1},
		{"b", "aaaa", 4},
	}

	for _, input := range inputs {
		res := LongestPalindrome2(input.t, input.k)
		if res != input.want {
			t.Errorf("t:%v,k:%v,want:%v,got:%v", input.t, input.k, input.want, res)
		}
	}
}
