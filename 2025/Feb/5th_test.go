package feb

import "testing"

func TestAreAlmostEqual(t *testing.T) {
	inputs := []struct {
		s1   string
		s2   string
		want bool
	}{
		{"bank", "kanb", true},
		{"attack", "defend", false},
		{"kelb", "kelb", true},
	}

	for _, input := range inputs {
		res := AreAlmostEqual2(input.s1, input.s2)
		if res != input.want {
			t.Errorf("s1:%v,s2:%v,want:%v,got:%v", input.s1, input.s2, input.want, res)
		}
	}
}
