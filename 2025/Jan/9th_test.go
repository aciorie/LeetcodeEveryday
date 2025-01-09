package jan

import "testing"

func TestPrefixCount(t *testing.T) {
	inputs := []struct {
		words []string
		pref  string
		want  int
	}{
		{[]string{"pay", "attention", "practice", "attend"}, "at", 2},
		{[]string{"leetcode", "win", "loops", "success"}, "code", 0},
	}

	for _, input := range inputs {
		res := PrefixCount(input.words, input.pref)
		if res != input.want {
			t.Errorf("words:%v,pref:%v,want:%v,got:%v", input.words, input.pref, input.want, res)
		}
	}
}
