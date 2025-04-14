package apr

import "testing"

func TestCountGoodTriplets(t *testing.T) {
	inputs := []struct {
		arr  []int
		a    int
		b    int
		c    int
		want int
	}{
		{[]int{3, 0, 1, 1, 9, 7}, 7, 2, 3, 4},
		{[]int{1, 1, 2, 2, 3}, 0, 0, 1, 0},
		{[]int{4, 9, 9, 8, 9, 5, 3, 7}, 1, 3, 0, 3},
	}

	for _, input := range inputs {
		res := CountGoodTriplets(input.arr, input.a, input.b, input.c)
		if res != input.want {
			t.Errorf("arr:%v,a:%v,b:%v,c:%v,want:%v,got:%v", input.arr, input.a, input.b, input.c, input.want, res)
		}
	}
}
