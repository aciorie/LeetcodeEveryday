package jan

import "testing"

func TestCanConstruct(t *testing.T) {
	inputs := []struct {
		s    string
		k    int
		want bool
	}{
		{"annabelle", 2, true},
		{"leetcode", 3, false},
		{"true", 4, true},
	}
	for _, input := range inputs {
		res := CanConstruct(input.s, input.k)
		if res != input.want {
			t.Errorf("s:%v,k:%v,want:%v,got:%v", input.s, input.k, input.want, res)
		}
	}
}
