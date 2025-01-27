package hot

import "testing"

func TestCoinChange(t *testing.T) {
	inputs := []struct {
		coins  []int
		amount int
		want   int
	}{
		{[]int{1, 2, 5}, 11, 3},
		{[]int{2}, 3, -1},
		{[]int{1}, 0, 0},
	}

	for _, input := range inputs {
		res := CoinChange3(input.coins, input.amount)
		if res != input.want {
			t.Errorf("coins:%v,k:%v,want:%v,got:%v", input.coins, input.amount, input.want, res)
		}
	}
}
