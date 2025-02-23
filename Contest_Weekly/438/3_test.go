package contest

import "testing"

func TestHasSameDigits_2(t *testing.T) {
	inputs := []struct {
		s    string
		want bool
	}{
		{"3902", true},
		{"34789", false},
	}

	for _, input := range inputs {
		res := HasSameDigits_3(input.s)
		if res != input.want {
			t.Errorf("s:%v,want:%v,got:%v", input.s, input.want, res)
		}
	}
}
