package jul

import "testing"

func TestFindLucky(t *testing.T) {
	inputs := []struct {
		arr  []int
		want int
	}{
		{[]int{2, 2, 3, 4}, 2},
		{[]int{1, 2, 2, 3, 3, 3}, 3},
		{[]int{2, 2, 2, 3, 3}, -1},
	}
	for _, input := range inputs {
		res := FindLucky(input.arr)
		if res != input.want {
			t.Errorf("nums:%v,want:%v,got:%v", input.arr, input.want, res)
		}
	}
}
