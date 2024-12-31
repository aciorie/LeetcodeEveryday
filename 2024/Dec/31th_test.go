package Dec

import "testing"

func TestMincostTickets(t *testing.T) {
	inputs := []struct {
		days  []int
		costs []int
		want  int
	}{
		{[]int{1, 4, 6, 7, 8, 20}, []int{2, 7, 15}, 11},
		{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 30, 31}, []int{2, 7, 15}, 17},
	}

	for _, input := range inputs {
		res := MincostTickets(input.days, input.costs)
		if res != input.want {
			t.Errorf("days:%v,costs:%v,want:%v,got:%v", input.days, input.costs, input.want, res)
		}
	}
}
