package hot

import "testing"

func TestDecodeString(t *testing.T) {
	inputs := []struct {
		s    string
		want string
	}{
		{"3[a]2[bc]", "aaabcbc"},
		{"3[a2[c]]", "accaccacc"},
		{"2[abc]3[cd]ef", "abcabccdcdcdef"},
	}

	for _, input := range inputs {
		res := DecodeString2(input.s)
		if res != input.want {
			t.Errorf("s:%v,want:%v,got:%v", input.s, input.want, res)
		}
	}
}
