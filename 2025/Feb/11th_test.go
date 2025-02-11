package feb

import "testing"

func TestRemoveOccurrences(t *testing.T) {
	inputs := []struct {
		s1   string
		s2   string
		want string
	}{
		{"daabcbaabcbc", "abc", "dab"},
		{"axxxxyyyyb", "xy", "ab"},
	}

	for _, input := range inputs {
		res := RemoveOccurrences(input.s1, input.s2)
		if res != input.want {
			t.Errorf("s1:%v,s2:%v,want:%v,got:%v", input.s1, input.s2, input.want, res)
		}
	}
}
