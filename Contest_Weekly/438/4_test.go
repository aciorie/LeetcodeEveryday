package contest

import "testing"

func TestMaxDistance(t *testing.T) {
	inputs := []struct {
		side   int
		points [][]int
		k      int
		want   int
	}{
		{2, [][]int{{0, 2}, {2, 0}, {2, 2}, {0, 0}}, 4, 2},
		{2, [][]int{{0, 0}, {1, 2}, {2, 0}, {2, 1}}, 4, 1},
		{2, [][]int{{0, 0}, {0, 1}, {0, 2}, {1, 2}, {2, 0}, {2, 2}, {2, 1}}, 5, 1},
	}

	for _, input := range inputs {
		res := MaxDistance(input.side, input.points, input.k)
		if res != input.want {
			t.Errorf("side:%v,points:%v,k:%v,want:%v,got:%v", input.side, input.points, input.k, input.want, res)
		}
	}
}
