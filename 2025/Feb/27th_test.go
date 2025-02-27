package feb

import "testing"

func TestLenLongestFibSubseq(t *testing.T) {
	inputs := []struct {
		arr  []int
		want int
	}{
		{[]int{1, 2, 3, 4, 5, 6, 7, 8}, 5},
		{[]int{1, 3, 7, 11, 12, 14, 18}, 3},
	}

	for _, input := range inputs {
		res := LenLongestFibSubseq(input.arr)
		if res != input.want {
			t.Errorf("arr:%v,want:%v,got:%v", input.arr, input.want, res)
		}
	}
}
