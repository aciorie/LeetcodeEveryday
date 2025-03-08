package mar

import "testing"

func TestMinimumRecolors(t *testing.T) {
	inputs := []struct {
		blocks string
		k      int
		want   int
	}{
		{"WBBWWBBWBW", 7, 3},
		{"WBWBBBW", 2, 0},
	}

	for _, input := range inputs {
		res := MinimumRecolors(input.blocks, input.k)
		if res != input.want {
			t.Errorf("blocks:%v,k:%v,want:%v,got:%v", input.blocks, input.k, input.want, res)
		}
	}
}
