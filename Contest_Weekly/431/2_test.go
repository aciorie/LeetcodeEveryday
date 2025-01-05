package contest

import "testing"

func TestCalculateScore(t *testing.T) {
	inputs := []struct {
		s    string
		want int64
	}{
		{"aczzx", 5},
		{"abcdef", 0},
		{"azapfwonwwcdagew", 3},
	}

	for _, input := range inputs {
		res := CalculateScore(input.s)
		if res != input.want {
			t.Errorf("s:%v,want:%v,got:%v", input.s, input.want, res)
		}
	}
}
