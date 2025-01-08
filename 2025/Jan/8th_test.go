package jan

import "testing"

func TestCountPrefixSuffixPairs(t *testing.T) {
	inputs := []struct {
		words []string
		want  int
	}{
		{[]string{"a", "aba", "ababa", "aa"}, 4},
		{[]string{"pa", "papa", "ma", "mama"}, 2},
		{[]string{"abab", "ab"}, 0},
	}

	for _, input := range inputs {
		res := CountPrefixSuffixPairs(input.words)
		if res != input.want {
			t.Errorf("words:%v,want:%v,got:%v", input.words, input.want, res)
		}
	}
}

func TestIsPrefixAndSuffix(t *testing.T) {
	inputs := []struct {
		a    string
		b    string
		want bool
	}{
		{"aba", "ababa", true},
		{"abc", "abcd", false},
	}

	for _, input := range inputs {
		res := IsPrefixAndSuffix(input.a, input.b)
		if res != input.want {
			t.Errorf("a:%v,b:%v,want:%v,got:%v", input.a, input.b, input.want, res)
		}
	}
}
