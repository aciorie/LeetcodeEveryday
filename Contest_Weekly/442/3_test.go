package contest

import "testing"

func TestMinTime(t *testing.T) {
	inputs := []struct {
		skill []int
		mana  []int
		want  int64
	}{
		{[]int{1, 5, 2, 4}, []int{5, 1, 4, 2}, 110},
		{[]int{1, 1, 1}, []int{1, 1, 1}, 5},
		{[]int{1, 2, 3, 4}, []int{1, 2}, 21},
	}

	for _, input := range inputs {
		res := MinTime(input.skill, input.mana)
		if res != input.want {
			t.Errorf("skill:%v,mana:%v,want:%v,got:%v", input.skill, input.mana, input.want, res)
		}
	}
}
